package stringutil

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Fmt(v ...interface{}) string {
	msg := ""
	if len(v) > 0 {
		message, ok := v[0].(string)
		if ok && strings.Index(message, "%") > -1 {
			msg = fmt.Sprintf(message, v[1:]...)
		} else {
			msg = fmt.Sprintln(v...)
		}
	}
	return msg
}

//返回一个32位md5加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//返回一个16位md5加密后的字符串
func Get16MD5Encode(data string) string {
	return GetMD5Encode(data)[8:24]
}

//strByXOR 异或运算方法
func strByXOR(message string, keywords string) string {
	messageLen := len(message)
	keywordsLen := len(keywords)

	result := ""

	for i := 0; i < messageLen; i++ {
		result += string(message[i] ^ keywords[i%keywordsLen])
	}
	return result
}

// byteByXOR 异或运算方法
func ByteByXOR(message, keywords []byte) []byte {
	messageLen := len(message)
	keywordsLen := len(keywords)

	var result []byte

	for i := 0; i < messageLen; i++ {
		result = append(result, message[i]^keywords[i%keywordsLen])
	}
	return result
}

//XOREncode 对srcPath文件内容进行异域运算(加密)
func XOREncode(srcPath, tempPath, prefix string) error {
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
	content := append(data, ByteByXOR(b, []byte{0xc6, 0x89, 0xb7, 0xba, 0xcc})...)
	f.Write(content)
	f.Close()

	return nil
}

// XORDecode 对srcPath文件内容进行异或运算(解密)
func XORDecode(srcPath, prefix string) error {
	b, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return err
	}
	content := string(b)
	if strings.Contains(content, prefix) {
		data := b[len(prefix):]
		message := ByteByXOR(data, []byte{0xc6, 0x89, 0xb7, 0xba, 0xcc})
		err = ioutil.WriteFile(srcPath, message, os.ModePerm)

		if err != nil {
			return err
		}
	}

	return nil
}

func InStringArray(list []string, find string) bool {
	if list == nil || len(list) == 0 {
		return false
	}

	for _, value := range list {
		if value == find {
			return true
		}
	}

	return false
}
