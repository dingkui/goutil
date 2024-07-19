package encry

import (
	"bufio"
	"gitee.com/dk83/goutils/errs"
	"io"
)

var Xor = NewCoder(func() Coder { return _xor }, demoKey)

type xor struct{}

var _xor = &xor{}

func (xor) Encode(reader io.Reader, writer io.Writer, pass []byte) error {
	return doXor(reader, writer, pass)
}
func (xor) Decode(reader io.Reader, writer io.Writer, pass []byte) error {
	return doXor(reader, writer, pass)
}

// byteByXOR 异或运算方法
func doXor(reader io.Reader, writer io.Writer, pass []byte) error {
	if pass == nil || len(pass) == 0 {
		_, err := io.Copy(writer, reader)
		return err
	}
	newReader := bufio.NewReader(reader)

	buf := make([]byte, len(pass))
	for {
		switch n, err := newReader.Read(buf); err {
		case nil:
			for i := 0; i < n; i++ {
				buf[i] ^= pass[i]
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
