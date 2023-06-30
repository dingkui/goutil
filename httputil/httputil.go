package httputil

import (
	"bytes"
	"fmt"
	"gitee.com/dk83/goutils/fileutil"
	"gitee.com/dk83/goutils/jsonutil"
	"gitee.com/dk83/goutils/logutil"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func delfaultCheckStatus(url string, status int) error {
	if status == http.StatusNotFound || status == http.StatusForbidden {
		logutil.Error("无效请求:url=%s,status：%d", url, status)
		return fmt.Errorf("非法请求")
	}

	if status != http.StatusOK {
		logutil.Error("服务器异常:url=%s,status：%d", url, status)
		return fmt.Errorf("服务器异常")
	}
	return nil
}
func DoHttp(method string, url string, bodys []byte, headers map[string]string, checkStatus func(string, int) error) ([]byte, error) {
	//把[]byte 转成实现了read接口的Reader结构体
	var body io.Reader
	if bodys != nil {
		body = bytes.NewReader(bodys)
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logutil.Error("网络故障-1:url=%s,%s", url, err.Error())
		return nil, fmt.Errorf("网络故障-1")
	}
	if headers != nil && len(headers) > 0 {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logutil.Error("网络故障-2:url=%s,%s", url, err.Error())
		return nil, fmt.Errorf("网络故障-2")
	}
	defer resp.Body.Close()

	if checkStatus == nil {
		checkStatus = delfaultCheckStatus
	}
	err = checkStatus(url, resp.StatusCode)
	if err != nil {
		return nil, err
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logutil.Error("网络故障-3:url=%s,%s", url, err.Error())
		return nil, fmt.Errorf("网络故障-3")
	}
	return resBody, err
}
func GET(url string, headers map[string]string, checkStatus func(string, int) error) ([]byte, error) {
	return DoHttp("GET", url, nil, headers, checkStatus)
}
func POST(url string, param []byte, headers map[string]string, checkStatus func(string, int) error) ([]byte, error) {
	return DoHttp("POST", url, nil, headers, checkStatus)
}
func POSTAsJson(url string, param []byte, headers map[string]string) (*jsonutil.JSON, error) {
	resBody, err := POST(url, param, headers, nil)
	if err != nil {
		return nil, err
	}
	return jsonutil.MkJSON(resBody), nil
}
func Get(url string) ([]byte, error) {
	return GET(url, nil, nil)
}
func GetAsJson(url string) (*jsonutil.JSON, error) {
	resBody, err := GET(url, nil, nil)
	if err != nil {
		return nil, err
	}
	return jsonutil.MkJSON(resBody), nil
}
func GETAsJson(url string, headers map[string]string) (*jsonutil.JSON, error) {
	resBody, err := GET(url, headers, nil)
	if err != nil {
		return nil, err
	}
	return jsonutil.MkJSON(resBody), nil
}

func Get2File(dataFile, url string) (result bool) {
	result = false
	err := os.MkdirAll(filepath.Dir(dataFile), os.ModePerm)
	if err != nil {
		logutil.Error("保存文件失败：", err.Error())
		return
	}
	defer func() {
		if !result {
			os.Remove(dataFile)
		}
	}()
	respBody, err := GET(url, nil, nil)
	if err != nil {
		return
	}
	err = fileutil.WriteAndSyncFile(dataFile, respBody, os.ModePerm)
	if err != nil {
		logutil.Error("保存文件失败：%s dataFile=%s", err.Error(), dataFile)
		return
	}
	result = true
	return
}
func Get2Folder(folder, url string) (result bool) {
	fileName := filepath.Base(url)
	dataFile := filepath.Join(folder, fileName)
	return Get2File(dataFile, url)
}
