package httputil

import (
	"bytes"
	"codeup.aliyun.com/dk/go/goutil/utils/zlog"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

type tsdbrelayHTTPTransport struct {
	Transport http.RoundTripper
}

// RoundTrip sets a predefined agent in the request and then forwards it to the
// default RountTrip implementation.
func (this *tsdbrelayHTTPTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Safari/537.36")
	return this.Transport.RoundTrip(r)
}

func PostForNode(url string, param map[string]interface{}) (map[string]interface{}, error) {
	b, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("参数格式化失败")
	}
	re, err, _ := Post(url, b, nil, nil)
	return re, err
}

func Post(url string, param []byte, headers map[string]string, checkRes func(resBody []byte) (map[string]interface{}, error, bool)) (map[string]interface{}, error, bool) {
	//把[]byte 转成实现了read接口的Reader结构体
	var body io.Reader
	if param != nil {
		body = bytes.NewReader(param)
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		zlog.Error("网络故障-1:url=%s,%s", url, err.Error())
		return nil, fmt.Errorf("网络故障-1"), true
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	if headers != nil && len(headers) > 0 {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		zlog.Error("网络故障-2:url=%s,%s", url, err.Error())
		return nil, fmt.Errorf("网络故障-2"), true
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zlog.Error("网络故障-3:url=%s,%s", url, err.Error())
		return nil, fmt.Errorf("网络故障-3"), true
	}

	if checkRes != nil {
		res, err, b := checkRes(resBody)
		if !b {
			zlog.Warn("网络故障-4:url=%s,%s", url, err.Error())
		}
		return res, err, b
	}

	res := make(map[string]interface{})
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		zlog.Error("网络故障-4:url=%s,%s", url, err.Error())
		return nil, fmt.Errorf("网络故障-4"), true
	}
	return res, err, true
}

func Download2Dir(dir, urlPath string) (string, error) {
	i, err := url.Parse(urlPath)
	if err != nil {
		return "", err
	}
	dstfile := i.Path

	dataFile := filepath.Join(dir, filepath.Base(dstfile))

	tryNum := 0

	for ; tryNum < 10; tryNum++ {
		if downloadImageOnce(dataFile, urlPath) {
			return dataFile, nil
		}
	}

	return "", fmt.Errorf("下载失败")
}

func downloadImageOnce(dataFile, url string) (result bool) {
	result = false

	rootDir := filepath.Dir(dataFile)
	os.MkdirAll(rootDir, os.ModePerm)

	defer func() {
		if !result {
			os.Remove(dataFile)
		}

	}()
	_, err := os.Stat(dataFile)
	if err == nil {
		//ctx.SendFile(imgFile,ctx.Params().Get("imgId"))
		result = true
		return
	}

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		zlog.ErrorLn(err)

		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
		result = true
		return
	}

	if resp.StatusCode != http.StatusOK {
		zlog.ErrorLn(err, resp.StatusCode, url)
		return
	}

	clen := resp.Header.Get("Content-Length")
	len, _ := strconv.Atoi(clen)

	f, err := os.Create(dataFile)
	if err != nil {
		zlog.ErrorLn(err)
		return
	}
	defer f.Close()
	sz, err := io.Copy(f, resp.Body)
	if err != nil {
		zlog.ErrorLn(err)
		return
	}

	if sz != int64(len) && len > 0 {
		zlog.ErrorLn(sz, len)
		return
	}

	result = true

	return
}

func MustDownload(fileUrl, dir string) (string, error) {
	data, err := MustGet(fileUrl)
	if err != nil {
		zlog.ErrorLn(err, fileUrl)
		return "", err
	}
	i, _ := url.Parse(fileUrl)
	dstfile := i.Path
	os.MkdirAll(filepath.Dir(dir+dstfile), os.ModePerm)
	ioutil.WriteFile(dir+dstfile, data, os.ModePerm)

	return dir + dstfile, nil
}

func MustGet(url string) ([]byte, error) {
	for {
		resp, err := http.DefaultClient.Get(url)
		if err != nil {
			zlog.ErrorLn(err, url)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
			zlog.ErrorLn("404:", url)
			return nil, fmt.Errorf("404")
		}

		if resp.StatusCode != http.StatusOK {
			zlog.ErrorLn(err, url)
			continue
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			zlog.ErrorLn(err, url)
			continue
		}
		return bytes, err
	}
}
