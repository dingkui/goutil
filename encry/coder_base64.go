package encry

import (
	"encoding/base64"
	"io"
)

var Base64 = NewCoder(func() Coder { return _coderBase64 }, nil)

type coderBase64 struct{}

var _coderBase64 = coderBase64{}

func (coderBase64) Encode(reader io.Reader, writer io.Writer, _ []byte) error {
	//dst := base64.StdEncoding.EncodeToString(src)
	encoder := base64.NewEncoder(base64.StdEncoding, writer)
	defer encoder.Close()
	// 以流式解码
	buf := make([]byte, 1024)
	// 保存解码后的数据
	for {
		n, err := reader.Read(buf)
		encoder.Write(buf[:n])
		if n == 0 || err != nil {
			break
		}
	}
	return nil
}
func (coderBase64) Decode(reader io.Reader, writer io.Writer, _ []byte) error {
	decoder := base64.NewDecoder(base64.StdEncoding, reader)
	// 以流式解码
	buf := make([]byte, 1024)
	// 保存解码后的数据
	for {
		n, err := decoder.Read(buf)
		writer.Write(buf[:n])
		if n == 0 || err != nil {
			break
		}
	}
	return nil
}
