package dhttp

import (
	"gitee.com/dk83/goutils/djson"
	"gitee.com/dk83/goutils/dlog"
	"gitee.com/dk83/goutils/errs"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type HTTPMethod string

var (
	Get    HTTPMethod = "GET"
	Put    HTTPMethod = "PUT"
	Head   HTTPMethod = "HEAD"
	Post   HTTPMethod = "POST"
	Delete HTTPMethod = "DELETE"
)

// Response defines HTTP
type Response struct {
	StatusCode    int
	Header        http.Header
	Body          []byte
	ContentLength int64
	resHandler    iResHandler
	reqPath       string
}

func (ops *Response) HandleResAsStr() (string, int, error) {
	handler := ops.resHandler
	if handler != nil {
		return handler.HandleResAsStr(ops)
	}
	return string(ops.Body), ops.StatusCode, nil
}
func (ops *Response) HandleResAsJson() (*djson.JsonGo, int, error) {
	handler := ops.resHandler
	if handler != nil {
		return handler.HandleResAsJson(ops)
	}
	code := ops.StatusCode
	respBody := ops.Body
	if len(respBody) == 0 {
		return nil, code, errs.ErrHttp.New("returned invalid response body, status=%s,path=%s", ops.StatusCode, ops.reqPath)
	}
	jsonGo, err := djson.NewJsonGo(respBody)
	return jsonGo, code, err
}
func (ops *Response) HandleRes() (interface{}, int, error) {
	handler := ops.resHandler
	if handler != nil {
		return handler.HandleRes(ops)
	}
	code := ops.StatusCode
	resContextType := ops.Header.Get(HTTPHeaderContentType)
	s := strings.Split(resContextType, ";")[0]
	switch s {
	case string(ContentTypeXml):
	case string(ContentTypeHtml):
	case string(ContentTypeText):
		return ops.HandleResAsStr()
	case strings.Split(string(ContentTypeJson), ";")[0]:
		return ops.HandleResAsJson()
	}
	return ops, code, nil
}

func (ops HTTPMethod) Do(options *Options) (*Response, error) {
	urlHandler, reqHandler, resHandler := options.GetHttpHandler().handles()
	//添加自定义headers,默认options等
	if reqHandler != nil {
		reqHandler.HandleReq(options)
	}
	uri, err := urlHandler.GetUri(options)
	if err != nil {
		return nil, err
	}
	//data 取得headers 顺序不能反
	data := options.GetData()
	headers := options.GetHeaders()

	listener := options.GetProgressListener()
	resp, err := ops.DoReq(uri, headers, data, listener)
	if err != nil {
		return nil, err
	}

	if listener != nil && resp.ContentLength > 1024*1024 {
		// CRC
		resBody, reProgress := newTeeReader(resp.Body, nil, listener, resp.ContentLength)
		resp.Body = ioutil.NopCloser(resBody)
		reProgress.resStarted()
		defer func() {
			reProgress.resEnd()
			resBody.Close()
		}()
	} else {
		resBody := resp.Body
		resp.Body = ioutil.NopCloser(resBody)
		defer func() {
			resBody.Close()
		}()
	}
	response := &Response{
		StatusCode:    resp.StatusCode,
		Header:        resp.Header,
		ContentLength: resp.ContentLength,
		resHandler:    resHandler,
		reqPath:       options.GetPath(),
	}
	handler := options.GetReadHandler()
	if handler == nil {
		handler = &DefaultReadHandler{}
	}
	err = handler.HandleRead(resp.Body, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (ops HTTPMethod) DoReq(uri *url.URL, headers map[string]string, body io.Reader, listener ProgressListener) (*http.Response, error) {
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

	contentLength, ok := headers[HTTPHeaderContentLength]
	if listener != nil && body != nil && ok {
		teeData, progress := newTeeReader(body, nil, listener, djson.IntN(0, contentLength))
		body = teeData
		// Transfer started
		progress.reqStarted()
		defer progress.reqEnd()
	}
	if body != nil {
		// HTTP body
		rc, ok := body.(io.ReadCloser)
		if !ok {
			rc = ioutil.NopCloser(body)
		}
		req.Body = rc
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		dlog.Error("[Req:%p]http error:%s\n", req, err.Error())
	}
	return response, err
}
