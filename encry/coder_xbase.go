package encry

import (
	"bytes"
	"io"
)

var Xbase = NewCoder(func() Coder { return xbase{} }, demoKey)

type xbase struct{}

func (xbase) Encode(reader io.Reader, writer io.Writer, pass []byte) error {
	base := coderBase64{}
	out := bytes.NewBuffer([]byte{})
	err := doXor(reader, out, pass)
	if err != nil {
		return err
	}
	return base.Encode(out, writer, nil)
}
func (xbase) Decode(reader io.Reader, writer io.Writer, pass []byte) error {
	base := coderBase64{}
	out := bytes.NewBuffer([]byte{})
	err := base.Decode(reader, out, nil)
	if err != nil {
		return err
	}
	return doXor(out, writer, pass)
}
