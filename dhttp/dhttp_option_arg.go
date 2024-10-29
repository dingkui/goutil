package dhttp

import (
	"bytes"
	"gitee.com/dk83/goutils/djson"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

const (
	reqPath          = "req-path"
	reqData          = "req-data"
	httpHandler      = "http-handler"
	readHandler      = "read-handler"
	progressListener = "x-progress-listener"
	responseHeader   = "x-response-header"
)

func (ops *Options) AddArg(key string, value interface{}) {
	ops.addOption(key, value, optionArg)
}
func (ops *Options) GetArgs() (map[string]interface{}, error) {
	paramsm := map[string]interface{}{}
	for _, option := range ops.options {
		if option.ot == optionArg {
			paramsm[option.key] = option.value
		}
	}
	return paramsm, nil
}
func (ops *Options) GetArg(key string) interface{} {
	for _, option := range ops.options {
		if option.ot == optionArg && option.key == key {
			return option.value
		}
	}
	return nil
}

func (ops *Options) DataStream(data io.Reader) *Options {
	if data == nil {
		return ops
	}
	ops.AddArg(reqData, data)
	ops.ContentType(ContentTypeStream)
	return ops
}
func (ops *Options) DataFile(filePath string) *Options {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 创建一个基于内存的 io.Writer，用于存储 multipart form data
	body := &bytes.Buffer{}
	// 创建一个 writer，用于写入 multipart form data
	writer := multipart.NewWriter(body)
	// 将文件添加到表单数据中
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}
	// 关闭 multipart writer
	err = writer.Close()
	if err != nil {
		panic(err)
	}
	ops.AddArg(reqData, body)
	ops.AddHeader(HTTPHeaderContentType, writer.FormDataContentType())
	ops.ContentLength(int64(body.Len()))
	return ops
}
func (ops *Options) DataJson(json *djson.JsonGo) *Options {
	if json == nil {
		return ops
	}
	b, err := json.Byte()
	if err != nil {
		panic(err)
	}
	ops.AddArg(reqData, bytes.NewReader(b))
	ops.ContentType(ContentTypeJson)
	ops.ContentLength(int64(len(b)))
	return ops
}
func (ops *Options) DataFrom(json *djson.JsonGo) *Options {
	if json == nil {
		return ops
	}
	values := url.Values{}
	m, err := json.Map()
	if err != nil {
		panic(err)
	}
	for key, value := range m {
		v := value.StrN("")
		if v != "" {
			values.Add(key, v)
		}
	}
	// 将 values 转换为字节流
	v := []byte(values.Encode())
	ops.AddArg(reqData, bytes.NewReader(v))
	ops.ContentType(ContentTypeForm)
	ops.ContentLength(int64(len(v)))
	return ops
}

func (ops *Options) GetData() (re io.Reader) {
	value := ops.GetArg(reqData)
	if value != nil {
		re = value.(io.Reader)
		readerLen := ops.GetContentLength()
		if readerLen == 0 {
			readerLen, err := GetReaderLen(re)
			if err == nil {
				ops.ContentLength(readerLen)
			}
		}
		return re
	}
	return nil
}
func (ops *Options) Path(data string) {
	ops.AddArg(reqPath, data)
}
func (ops *Options) GetPath() string {
	value := ops.GetArg(reqPath)
	if value != nil {
		return value.(string)
	}
	return ""
}
func (ops *Options) HttpHandler(data *HttpHandler) *Options {
	ops.AddArg(httpHandler, data)
	return ops
}
func (ops *Options) GetHttpHandler() *HttpHandler {
	value := ops.GetArg(httpHandler)
	if value != nil {
		return value.(*HttpHandler)
	}
	panic("no Http Handler found")
}
func (ops *Options) ReadHandler(data iReadHandler) *Options {
	ops.AddArg(readHandler, data)
	return ops
}
func (ops *Options) GetReadHandler() iReadHandler {
	value := ops.GetArg(readHandler)
	if value != nil {
		return value.(iReadHandler)
	}
	return nil
}

func (ops *Options) ResponseHeader(respHeader *http.Header) {
	ops.AddArg(responseHeader, respHeader)
}
func (ops *Options) GetResponseHeader() *http.Header {
	value := ops.GetArg(responseHeader)
	if value != nil {
		return value.(*http.Header)
	}
	return nil
}

func (ops *Options) Progress(listener ProgressListener) {
	ops.AddArg(progressListener, listener)
}
func (ops *Options) GetProgressListener() ProgressListener {
	value := ops.GetArg(progressListener)
	if value != nil {
		return value.(ProgressListener)
	}
	return nil
}
