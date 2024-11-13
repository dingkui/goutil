package encry

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/dingkui/goutil/errs"
	"io"
	"time"
)

var PXor = NewCoder(func() Coder { return _pxor }, demoKey)
var PXor2 = NewCoder(func() Coder { return _pxor2 }, demoKey)

type pxor struct {
	pass   []byte
	base64 bool
}

var _pxor = &pxor{nil, true}
var _pxor2 = &pxor{demoKey, false}

func (t pxor) Encode(reader io.Reader, writer io.Writer, pass []byte) error {
	out := writer
	if t.base64 {
		out = bytes.NewBuffer([]byte{})
	}

	randKey := []byte(RandKey(4))
	pasSum, _ := Md5.SumBytes(append(append(pass, randKey...), t.pass...))
	md5Pass := []byte(pasSum)
	out.Write(append(randKey, md5Pass[:4]...))
	//dlog.Debug("randKey:%s checkPass:%s pasSum:%s", string(randKey), string(md5Pass[:4]), pasSum)
	err := doXorPlus(reader, out, md5Pass)
	if err != nil {
		return err
	}

	if t.base64 {
		return coderBase64{}.Encode(out.(io.Reader), writer, nil)
	} else {
		return nil
	}
}
func (t pxor) Decode(reader io.Reader, writer io.Writer, pass []byte) error {
	if t.base64 {
		in := bytes.NewBuffer([]byte{})
		err := coderBase64{}.Decode(reader, in, nil)
		if err != nil {
			return err
		}
		reader = in
	}

	randKey := make([]byte, 8)
	read, err := reader.Read(randKey)
	if err != nil || read != 8 {
		return errs.ErrValidate.New(err, "读取失败")
	}
	pasSum, _ := Md5.SumBytes(append(append(pass, randKey[:4]...), t.pass...))
	//dlog.Debug("randKey:%s checkPass:%s pasSum:%s", string(randKey[:4]), string(randKey[4:]), pasSum)
	md5Pass := []byte(pasSum)
	if string(md5Pass[:4]) != string(randKey[4:]) {
		return errs.ErrValidate.New("密码错误")
	}
	return doXorPlus(reader, writer, md5Pass)
}
func RandKey(len int) string {
	bytes, _ := Md5.SumBytes([]byte(fmt.Sprint(time.Now().UnixNano())))
	if len > 32 {
		len = 32
	}
	return bytes[:len]
}

// byteByXOR 异或运算方法
func doXorPlus(reader io.Reader, writer io.Writer, pass []byte) error {
	lenPass := len(pass)
	if pass == nil || lenPass == 0 {
		_, err := io.Copy(writer, reader)
		return err
	}
	newReader := bufio.NewReader(reader)

	buf := make([]byte, lenPass)
	for {
		switch n, err := newReader.Read(buf); err {
		case nil:
			for i := 0; i < n; i++ {
				buf[i] ^= pass[i]
				if n < lenPass {
					for j := i; j < lenPass; j++ {
						buf[i] ^= pass[j]
					}
				}
			}
			_, err := writer.Write(buf[:n])
			if err != nil {
				return err
			}
		case io.EOF:
			return nil
		default:
			return errs.ErrRuntime.New("read error")
		}
	}
}
