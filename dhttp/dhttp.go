package dhttp

import (
	"fmt"
	"gitee.com/dk83/goutils/dlog"
	"io"
	"io/ioutil"
	"net/http"
)

type HTTPMethod string

var (
	_methods   = make(map[string]bool)
	HTTPGet    = HM("GET")
	HTTPPut    = HM("PUT")
	HTTPHead   = HM("HEAD")
	HTTPPost   = HM("POST")
	HTTPDelete = HM("DELETE")
)

type IReqHandler interface {
	GetBaseUrl() string
	HandleReq(options *Options)
	HandleRes(resp *http.Response, options *Options) (interface{}, error)
}

func HM(method string) HTTPMethod {
	_, exists := _methods[method]
	if exists {
		panic(fmt.Sprintf("Make new HTTPMehods %s has error", method))
	}
	return HTTPMethod(method)
}

func (ops HTTPMethod) Do(options *Options) (interface{}, error) {
	handler := options.GetReqHandler()
	//添加自定义headers,默认options等
	handler.HandleReq(options)

	data := options.GetData()
	headers := options.GetHeaders()

	uri, err := options.GetUri(handler.GetBaseUrl())
	if err != nil {
		return nil, err
	}

	req := &http.Request{
		Method:     string(ops),
		URL:        uri,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       uri.Host,
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	listener := options.GetProgressListener()
	var progress *progress
	if listener != nil && data != nil {
		// CRC
		data, progress = newTeeReader(data, nil, listener, options.GetContentLength())
	}
	if data != nil {
		// HTTP body
		rc, ok := data.(io.ReadCloser)
		if !ok {
			rc = ioutil.NopCloser(data)
		}
		req.Body = rc
	}
	progress.reqStarted()

	// Transfer started
	resp, err := http.DefaultClient.Do(req)
	progress.reqEnd()
	if err != nil {
		dlog.Error("[Req:%p]http error:%s\n", req, err.Error())
		return nil, err
	}

	if listener != nil && resp.ContentLength > 1024*1024 {
		// CRC
		resBody, reProgress := newTeeReader(resp.Body, nil, listener, resp.ContentLength)
		resp.Body = io.NopCloser(resBody)
		reProgress.resStarted()
		defer func() {
			reProgress.resEnd()
			resBody.Close()
		}()
	}

	return handler.HandleRes(resp, options)
}
