package dhttp

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

func TestMultipart(t *testing.T) {
	// 定义文件路径
	filePath := "/path/to/your/file.txt"

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

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", "https://example.com/upload", body)
	if err != nil {
		panic(err)
	}

	// 设置 Content-Type header
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送 HTTP 请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 处理响应
	// ...
}
