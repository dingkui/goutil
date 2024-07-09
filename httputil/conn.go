package httputil

import (
	"gitee.com/dk83/goutils/zlog"
	"io"
	"io/ioutil"
	"net/http"
)

type IReqHandler interface {
	getBaseUrl() string
	handleHead(headers map[string]string)
	handleRes(resp *http.Response, options *Options) (interface{}, error)
}

func Do(request *Request, handler IReqHandler) (interface{}, error) {
	options := request.Options
	data := request.data

	uri, err := request.GetUri(handler.getBaseUrl())
	if err != nil {
		return nil, err
	}

	readerLen, err := GetReaderLen(data)
	if err != nil {
		return nil, err
	}
	options.ContentLength(readerLen)

	headers := make(map[string]string)
	err = request.FillHeaders(headers)
	if err != nil {
		return nil, err
	}

	req := &http.Request{
		Method:     request.method,
		URL:        uri,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       uri.Host,
	}

	handler.handleHead(headers)
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	listener := options.GetProgressListener()
	var progress *progress
	if listener != nil && data != nil {
		// CRC
		request.data, progress = newTeeReader(data, nil, listener, readerLen)
		data = request.data
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
	if err != nil {
		progress.reqFailed()
		zlog.Error("[Req:%p]http error:%s\n", req, err.Error())
		return nil, err
	}
	progress.reqCompleted()

	progress.listener = nil
	if listener != nil && resp.ContentLength > 1024*1024 {
		// CRC
		resp.Body, progress = newTeeReader(resp.Body, nil, listener, resp.ContentLength)
		progress.resStarted()
	}
	defer func() {
		progress.resEnd()
	}()

	return handler.handleRes(resp, options)
}
