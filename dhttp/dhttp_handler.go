package dhttp

import (
	"fmt"
	"gitee.com/dk83/goutils/djson"
	"gitee.com/dk83/goutils/errs"
	"io"
	"io/ioutil"
	"net/url"
	"strings"
)

type iUrlHandler interface {
	GetUri(ops *Options) (*url.URL, error)
}
type iReqHandler interface {
	HandleReq(options *Options)
}
type iReadHandler interface {
	// HandleRead 读取body数据，可根据resp的返回类型进行自定义，如文件下载等
	// 一般，返回值内容比较大，做特殊处理，如保存成文件，或者其他操作。为避免内存占用过大 resp.Body中的byte数组为空
	HandleRead(body io.Reader, resp *Response) error
}
type iResHandler interface {
	// HandleRes 可根据resp的返回类型进行自定义返回值
	HandleRes(resp *Response) (interface{}, int, error)
	HandleResAsJson(resp *Response) (*djson.JsonGo, int, error)
	HandleResAsStr(resp *Response) (string, int, error)
}

type HttpHandler struct {
	url iUrlHandler
	req iReqHandler
	res iResHandler
}

func (t *HttpHandler) handles() (iUrlHandler, iReqHandler, iResHandler) {
	return t.url, t.req, t.res
}

type DefaultUrlHandler struct{ Base string }

func (t *DefaultUrlHandler) GetUri(ops *Options) (*url.URL, error) {
	path := ops.GetPath()
	params := ops.GetUrlParams()
	urlParams := GetURLParams(params)
	base := t.Base
	isHttpPath := strings.Index(path, "http") == 0
	if isHttpPath {
		base = ""
	} else {
		if len(base) > 0 && base[len(base)-1:] == "/" {
			base = base[:len(base)-1]
		}
		if len(path) == 0 || path[:1] != "/" {
			path = "/" + path
		}
	}

	if urlParams != "" {
		if strings.Index(path, "?") > -1 {
			path += "&"
		} else {
			path += "?"
		}
	}
	return url.ParseRequestURI(fmt.Sprintf("%s%s%s", base, path, urlParams))
}

type DefaultReqHandler struct{}

func (t *DefaultReqHandler) HandleReq(_ *Options) {}

type DefaultReadHandler struct{}

func (t *DefaultReadHandler) HandleRead(body io.Reader, res *Response) error {
	out, err := ioutil.ReadAll(body)
	if err == io.EOF {
		err = nil
	}
	res.Body = out
	return err
}

type DefaultResHandler struct{}

func (t *DefaultResHandler) HandleResAsStr(res *Response) (string, int, error) {
	return string(res.Body), res.StatusCode, nil
}
func (t *DefaultResHandler) HandleResAsJson(res *Response) (*djson.JsonGo, int, error) {
	code := res.StatusCode
	respBody := res.Body
	if len(respBody) == 0 {
		return nil, code, errs.ErrHttp.New("returned invalid response body, status=%s,path=%s", res.StatusCode, res.reqPath)
	}
	jsonGo, err := djson.NewJsonGo(respBody)
	return jsonGo, code, err
}
func (t *DefaultResHandler) HandleRes(res *Response) (interface{}, int, error) {
	code := res.StatusCode
	resContextType := res.Header.Get(HTTPHeaderContentType)
	s := strings.Split(resContextType, ";")[0]
	switch s {
	case string(ContentTypeXml):
	case string(ContentTypeHtml):
	case string(ContentTypeText):
		return t.HandleResAsStr(res)
	case strings.Split(string(ContentTypeJson), ";")[0]:
		return t.HandleResAsJson(res)
	}
	return res, code, nil
}
