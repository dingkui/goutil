package dhttp

import (
	"bytes"
	"gitee.com/dk83/goutils/djson"
	"gitee.com/dk83/goutils/dlog"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func defCheckStatus(url string, status int) *HttpError {
	if status == http.StatusNotFound || status == http.StatusForbidden {
		dlog.Error("无效请求:url=%s,status：%d", url, status)
		return Error(status, "非法请求")
	}

	if status != http.StatusOK {
		dlog.Error("服务器异常:url=%s,status：%d", url, status)
		return Error(status, "服务器异常")
	}
	return nil
}

type HTTP struct {
	Method      string
	Headers     map[string]string
	CheckStatus func(string, int) *HttpError
	PreReq      func(*http.Request) *HttpError
	AfterRes    func(*http.Response) *HttpError
}
type Para struct {
	Url         string
	UrlParas    djson.JsonGo
	Body        djson.JsonGo
	Headers     map[string]string
	CheckStatus func(string, int) *HttpError
	PreReq      func(*http.Request) *HttpError
	AfterRes    func(*http.Response) *HttpError
}

func (p *Para) IsBodyNil() bool {
	return p == nil || !p.Body.IsMap()
}
func (p *Para) IsUrlPparsNil() bool {
	return p == nil || !p.UrlParas.IsMap()
}
func (h *HTTP) SendAsStr(para Para) (string, *HttpError) {
	re, err := h.Send(para)
	if err != nil {
		return "", err
	}
	return string(re), nil
}
func (h *HTTP) SendAsJson(para Para) (*djson.JsonGo, *HttpError) {
	re, err := h.Send(para)
	if err != nil {
		return nil, err
	}
	return djson.NewJsonGo(re), nil
}
func (h *HTTP) Send(para Para) ([]byte, *HttpError) {
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
			dlog.Error("http请求参数错误:", err.Error())
			return nil, Error(1, "http请求参数错误")
		}
		params := url.Values{}
		for key, value := range re {
			params.Set(key, djson.AsStr(value))
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
		dlog.Error("网络故障-1:url=%s,%s", _url, err.Error())
		return nil, Error(1, "网络故障-1")
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
		httpError := preReq(req)
		if httpError != nil {
			return nil, httpError
		}
	}

	//5.执行网络请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		dlog.Error("网络故障-2:url=%s,%s", _url, err.Error())
		return nil, Error(2, "网络故障-2")
	}
	defer resp.Body.Close()

	//6.返回信息处理
	checkStatus := para.CheckStatus
	if checkStatus == nil {
		checkStatus = h.CheckStatus
	}
	if checkStatus == nil {
		checkStatus = defCheckStatus
	}
	httpError := checkStatus(_url, resp.StatusCode)
	if err != nil {
		return nil, httpError
	}
	afterRes := para.AfterRes
	if afterRes == nil {
		afterRes = h.AfterRes
	}
	if afterRes != nil {
		httpError = afterRes(resp)
		if err != nil {
			return nil, httpError
		}
	}

	//8.读取返回信息
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		dlog.Error("网络故障-3:url=%s,%s", _url, err.Error())
		return nil, Error(3, "网络故障-3")
	}
	return resBody, httpError
}
