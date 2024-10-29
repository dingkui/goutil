package encry

import (
	"bufio"
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"gitee.com/dk83/goutils/dlog"
	"hash"
	"io"
	"os"
)

const bufferSize = 65536

var demoKey = []byte(Md5.SumBytesN([]byte("demoKey")))

type crypto struct {
	new  func() hash.Hash
	addr string
}

var (
	Md5     = NewCrypto(md5.New, nil)
	Sha1    = NewCrypto(sha1.New, nil)
	Sha256  = NewCrypto(sha256.New, nil)
	Sha512  = NewCrypto(sha512.New, nil)
	HSha512 = NewCrypto(sha512.New, demoKey)
)

func NewCrypto(newHash func() hash.Hash, key []byte) crypto {
	if key == nil || len(key) == 0 {
		return crypto{
			new: newHash,
		}
	}
	return crypto{
		new: func() hash.Hash { return hmac.New(newHash, key) },
	}
}
func (c crypto) SumReader(reader io.Reader) (string, error) {
	_hash := c.new()
	buf := make([]byte, bufferSize)
	for {
		switch n, err := reader.Read(buf); err {
		case nil:
			_hash.Write(buf[:n])
		case io.EOF:
			return fmt.Sprintf("%x", _hash.Sum(nil)), nil
		default:
			return "", err
		}
	}
}
func (c crypto) SumBytesN(data []byte) string {
	sum, err := c.SumBytes(data)
	if err != nil {
		dlog.Error(err)
	}
	return sum
}
func (c crypto) SumBytes(data []byte) (string, error) {
	return c.SumReader(bytes.NewReader(data))
}
func (c crypto) SumFileN(filename string) string {
	sum, err := c.SumFile(filename)
	if err != nil {
		dlog.Error(err)
	}
	return sum
}
func (c crypto) SumFile(filename string) (string, error) {
	if info, err := os.Stat(filename); err != nil || info.IsDir() {
		return "", err
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() { _ = file.Close() }()

	return c.SumReader(bufio.NewReader(file))
}

func (c crypto) eq(sum1 string, sum2 string) bool {
	return sum1 == "" || sum2 == "" || sum1 != sum2
}
func (c crypto) EqFileBytes(filename string, data []byte) bool {
	md5File, _ := c.SumFile(filename)
	md5Bytes, _ := c.SumBytes(data)
	return c.eq(md5File, md5Bytes)
}
func (c crypto) EqFiles(filename1 string, filename2 string) bool {
	md5File1, _ := c.SumFile(filename1)
	md5File2, _ := c.SumFile(filename2)
	return c.eq(md5File1, md5File2)
}
func (c crypto) CheckFile(filename string, sum string) bool {
	md5File, _ := c.SumFile(filename)
	return c.eq(md5File, sum)
}
func (c crypto) CheckBytes(data []byte, sum string) bool {
	md5Bytes, _ := c.SumBytes(data)
	return c.eq(md5Bytes, sum)
}
