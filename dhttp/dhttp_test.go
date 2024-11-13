package dhttp_test

import (
	"bytes"
	"fmt"
	"github.com/dingkui/goutil/dhttp"
	"github.com/dingkui/goutil/djson"
	"github.com/dingkui/goutil/dlog"
	"github.com/dingkui/goutil/utils/valUtil"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
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

var client = dhttp.NewDefaultClient("http://127.0.0.1:8888")

func TestClient1(t *testing.T) {
	jsonGo, _ := djson.NewJsonGo(map[string]interface{}{
		"user": "admin",
		"pass": "1234",
	})
	jsonGo.Set("admin", "user")
	jsonGo.Set("1234", "pass")
	options := &dhttp.Options{}
	options.ReadHandler(&TestReadHandler{})
	res, err := client.SendForm("POST", "/api/ma/auth", jsonGo, options)
	if err != nil {
		dlog.Info(err)
		return
	}
	json, i, err := res.HandleResAsStr()
	dlog.Info(valUtil.StrN(json), i, err)
}
func TestClient2(t *testing.T) {
	jsonGo, _ := djson.NewJsonGo(make(map[string]interface{}))
	jsonGo.Set("admin", "user")
	jsonGo.Set("1234", "pass")
	form, err := client.SendJson("POST", "/api/ma/auth", jsonGo)
	dlog.Info(err)
	dlog.Info(valUtil.StrN(form))
}

func TestHttp(t *testing.T) {
	// 定义请求的URL
	urlx := "http://127.0.0.1:8888/api/ma/auth"

	// 定义请求的参数
	data := url.Values{}
	data.Add("user", "admin")
	data.Add("email", "john@example.com")

	// 将参数转换为字符串
	payload := data.Encode()

	// 创建一个POST请求
	req, err := http.NewRequest("POST", urlx, bytes.NewBufferString(payload))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 打印响应
	fmt.Println("响应状态:", resp.Status)
	fmt.Println("响应体:", string(body))
}

type TestReadHandler struct{}

func (t *TestReadHandler) HandleRead(body io.Reader, res *dhttp.Response) error {
	out, err := ioutil.ReadAll(body)
	if err == io.EOF {
		err = nil
	}
	dlog.Info("do special things with out:%s", string(out))
	res.Body = out
	return err
}
