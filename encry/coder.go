package encry

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
)

type Coder interface {
	Encode(io.Reader, io.Writer, []byte) error
	Decode(io.Reader, io.Writer, []byte) error
}

type coder struct {
	get         func() Coder
	defaultPass []byte
}

func NewCoder(getCoder func() Coder, defaultPass []byte) coder {
	return coder{
		get:         getCoder,
		defaultPass: defaultPass,
	}
}

func (c coder) EncodeByPass(reader io.Reader, writer io.Writer, pass []byte) error {
	return c.get().Encode(reader, writer, pass)
}
func (c coder) Encode(reader io.Reader, writer io.Writer) error {
	return c.EncodeByPass(reader, writer, c.defaultPass)
}
func (c coder) EncodeBytesByPass(data []byte, pass []byte) (*bytes.Buffer, error) {
	read := bytes.NewBuffer(data)
	out := bytes.NewBuffer([]byte{})
	err := c.EncodeByPass(read, out, pass)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c coder) EncodeBytes(data []byte) (*bytes.Buffer, error) {
	return c.EncodeBytesByPass(data, c.defaultPass)
}
func (c coder) EncodeToFile(filePathInput string, encodedFile string, pass []byte) error {
	if info, err := os.Stat(filePathInput); err != nil || info.IsDir() {
		return err
	}
	fileInput, err := os.Open(filePathInput)
	if err != nil {
		return err
	}
	defer fileInput.Close()

	os.MkdirAll(filepath.Dir(encodedFile), os.ModePerm)
	fileOut, err := os.Create(encodedFile)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	return c.EncodeByPass(fileInput, fileOut, pass)
}

func (c coder) DecodeByPass(reader io.Reader, writer io.Writer, pass []byte) error {
	return c.get().Decode(reader, writer, pass)
}
func (c coder) Decode(reader io.Reader, writer io.Writer) error {
	return c.EncodeByPass(reader, writer, c.defaultPass)
}
func (c coder) DecodeBytesByPass(data []byte, pass []byte) (*bytes.Buffer, error) {
	read := bytes.NewBuffer(data)
	out := bytes.NewBuffer([]byte{})
	err := c.DecodeByPass(read, out, pass)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c coder) DecodeBytes(data []byte) (*bytes.Buffer, error) {
	return c.DecodeBytesByPass(data, c.defaultPass)
}
func (c coder) DecodeToFile(encodedFile string, filePathOutput string, pass []byte) error {
	if info, err := os.Stat(encodedFile); err != nil || info.IsDir() {
		return err
	}
	fileInput, err := os.Open(encodedFile)
	if err != nil {
		return err
	}
	defer fileInput.Close()

	os.MkdirAll(filepath.Dir(filePathOutput), os.ModePerm)
	fileOut, err := os.Create(filePathOutput)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	return c.DecodeByPass(fileInput, fileOut, pass)
}
