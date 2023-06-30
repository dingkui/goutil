package httputil

import (
	"bytes"
	"fmt"
	"gitee.com/dk83/goutils/jsonutil"
	"gitee.com/dk83/goutils/logutil"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type HTTP struct {
	Method      string
	Headers     map[string]string
	CheckStatus func(string, int) error
	PreReq      func(*http.Request) error
	AfterRes    func(*http.Response) error
}
type Para struct {
	Url         string
	UrlParas    jsonutil.JSON
	Body        jsonutil.JSON
	Headers     map[string]string
	CheckStatus func(string, int) error
	PreReq      func(*http.Request) error
	AfterRes    func(*http.Response) error
}

func (p *Para) IsBodyNil() bool {
	return p == nil || p.Body.IsNil()
}
func (p *Para) IsUrlPparsNil() bool {
	return p == nil || p.UrlParas.IsNil()
}
func (h *HTTP) SendAsStr(para Para) (string, error) {
	re, err := h.Send(para)
	if err != nil {
		return "", err
	}
	return string(re), nil
}
func (h *HTTP) SendAsJson(para Para) (*jsonutil.JSON, error) {
	re, err := h.Send(para)
	if err != nil {
		return nil, err
	}
	return jsonutil.MkJSON(re), nil
}
func (h *HTTP) Send(para Para) ([]byte, error) {
	//1.准备body
	var body io.Reader
	if !para.IsBodyNil() {
		body = bytes.NewReader(para.Body.Bytes())
	}
	_url := para.Url
	//2.准备url
	if !para.IsUrlPparsNil() {
		re, err := para.UrlParas.Map()
		if err != nil {
			logutil.Error("http请求参数错误:", err.Error())
			return nil, fmt.Errorf("http请求参数错误")
		}
		params := url.Values{}
		for key, value := range re {
			params.Set(key, jsonutil.AsStr(value))
		}
		if strings.Index(_url, "?") > -1 {
			_url += "&" + params.Encode()
		} else {
			_url += "?" + params.Encode()
		}
	}
	//3.req生成
	req, err := http.NewRequest(h.Method, _url, body)
	if err != nil {
		logutil.Error("网络故障-1:url=%s,%s", _url, err.Error())
		return nil, fmt.Errorf("网络故障-1")
	}

	//4.请求前处理
	if h.Headers != nil && len(h.Headers) > 0 {
		for key, value := range h.Headers {
			req.Header.Set(key, value)
		}
	}
	if para.Headers != nil && len(para.Headers) > 0 {
		for key, value := range para.Headers {
			req.Header.Set(key, value)
		}
	}

	preReq := para.PreReq
	if preReq == nil {
		preReq = h.PreReq
	}
	if preReq != nil {
		err = preReq(req)
		if err != nil {
			return nil, err
		}
	}

	//5.执行网络请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logutil.Error("网络故障-2:url=%s,%s", _url, err.Error())
		return nil, fmt.Errorf("网络故障-2")
	}
	defer resp.Body.Close()

	//6.返回信息处理
	checkStatus := para.CheckStatus
	if checkStatus == nil {
		checkStatus = h.CheckStatus
	}
	if checkStatus == nil {
		checkStatus = delfaultCheckStatus
	}
	err = checkStatus(_url, resp.StatusCode)
	if err != nil {
		return nil, err
	}
	afterRes := para.AfterRes
	if afterRes == nil {
		afterRes = h.AfterRes
	}
	if afterRes != nil {
		err = afterRes(resp)
		if err != nil {
			return nil, err
		}
	}

	//8.读取返回信息
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logutil.Error("网络故障-3:url=%s,%s", _url, err.Error())
		return nil, fmt.Errorf("网络故障-3")
	}
	return resBody, err
}
