package encry

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"os"
	"strings"
)

type h byte

const EncryUtil h = iota

//md5加密
func (x h) Md5(src string) string {
	m := md5.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

//Sha256加密
func (x h) Sha256(src string) string {
	m := sha256.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
func (x h) Base64Decode(str string) string {
	reader := strings.NewReader(str)
	decoder := base64.NewDecoder(base64.RawStdEncoding, reader)
	// 以流式解码
	buf := make([]byte, 1024)
	// 保存解码后的数据
	dst := ""
	for {
		n, err := decoder.Read(buf)
		dst += string(buf[:n])
		if n == 0 || err != nil {
			break
		}
	}
	return dst
}

//返回一个32位md5加密后的字符串
func (x h) GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//返回一个16位md5加密后的字符串
func (x h) Get16MD5Encode(data string) string {
	return x.GetMD5Encode(data)[8:24]
}

//strByXOR 异或运算方法
func (x h) strByXOR(message string, keywords string) string {
	messageLen := len(message)
	keywordsLen := len(keywords)

	result := ""

	for i := 0; i < messageLen; i++ {
		result += string(message[i] ^ keywords[i%keywordsLen])
	}
	return result
}

// byteByXOR 异或运算方法
func (x h) ByteByXOR(message, keywords []byte) []byte {
	messageLen := len(message)
	keywordsLen := len(keywords)

	var result []byte

	for i := 0; i < messageLen; i++ {
		result = append(result, message[i]^keywords[i%keywordsLen])
	}
	return result
}

//XOREncode 对srcPath文件内容进行异域运算(加密)
func (x h) XOREncode(srcPath, tempPath, prefix string) error {
	b, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	f, err := os.OpenFile(tempPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	data := []byte(prefix)
	content := append(data, x.ByteByXOR(b, []byte{0xc6, 0x89, 0xb7, 0xba, 0xcc})...)
	f.Write(content)
	f.Close()

	return nil
}

// XORDecode 对srcPath文件内容进行异或运算(解密)
func (x h) XORDecode(srcPath, prefix string) error {
	b, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return err
	}
	content := string(b)
	if strings.Contains(content, prefix) {
		data := b[len(prefix):]
		message := x.ByteByXOR(data, []byte{0xc6, 0x89, 0xb7, 0xba, 0xcc})
		err = ioutil.WriteFile(srcPath, message, os.ModePerm)

		if err != nil {
			return err
		}
	}

	return nil
}
