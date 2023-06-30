package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitee.com/dk83/goutils/fileutil"
	"gitee.com/dk83/goutils/logutil"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func delfaultCheckSstatus(url string, status int) error {
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
func doHttp(method string, url string, bodys []byte, headers map[string]string, checkStatus func(string, int) error) ([]byte, error) {
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
		checkStatus = delfaultCheckSstatus
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
	return doHttp("GET", url, nil, headers, checkStatus)
}
func POST(url string, param []byte, headers map[string]string, checkStatus func(string, int) error) ([]byte, error) {
	return doHttp("POST", url, nil, headers, checkStatus)
}
func POSTAsStr(url string, param []byte, headers map[string]string) (string, error) {
	resBody, err := POST(url, param, headers, nil)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}
func POSTAsMap(url string, param []byte, headers map[string]string) (map[string]interface{}, error) {
	resBody, err := POST(url, param, headers, nil)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		logutil.Error("网络故障-4:url=%s,%s", url, err.Error())
		return nil, fmt.Errorf("网络故障-4")
	}
	return res, nil
}
func Get(url string) (string, error) {
	resBody, err := GET(url, nil, nil)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}
func GetAsStr(url string) (string, error) {
	resBody, err := GET(url, nil, nil)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}
func GetAsByte(url string) ([]byte, error) {
	return GET(url, nil, nil)
}
func GETAsStr(url string, headers map[string]string) (string, error) {
	resBody, err := GET(url, headers, nil)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}
func GETAsMap(url string, headers map[string]string) (map[string]interface{}, error) {
	resBody, err := GET(url, headers, nil)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		logutil.Error("网络故障-4:url=%s,%s", url, err.Error())
		return nil, fmt.Errorf("网络故障-4")
	}
	return res, nil
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
