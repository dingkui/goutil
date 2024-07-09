package httputil

import (
	"fmt"
	"io"
	"net/url"
	"strings"
)

type HTTPMethod string

const (
	// HTTPGet HTTP GET
	HTTPGet HTTPMethod = "GET"
	// HTTPPut HTTP PUT
	HTTPPut HTTPMethod = "PUT"
	// HTTPHead HTTP HEAD
	HTTPHead HTTPMethod = "HEAD"
	// HTTPPost HTTP POST
	HTTPPost HTTPMethod = "POST"
	// HTTPDelete HTTP DELETE
	HTTPDelete HTTPMethod = "DELETE"
)

type Request struct {
	method  string
	path    string
	data    io.Reader
	Options *Options
}

func (ops HTTPMethod) NewRequest(path string, data io.Reader) *Request {
	return &Request{
		method: fmt.Sprintf("%s", ops),
		path:   path,
		data:   data,
	}
}
func (ops Request) FillHeaders(headers map[string]string) error {
	return ops.Options.FillHeaders(headers)
}

func (req Request) GetRawParams() (map[string]interface{}, error) {
	return req.Options.GetRawParams()
}

func (req Request) GetUri(base string) (*url.URL, error) {
	params, err := req.Options.GetUrlParams()
	if err != nil {
		return nil, err
	}
	urlParams := GetURLParams(params)
	if strings.LastIndex(base, "/") == len(base)-1 {
		base = base[:len(base)-1]
	}
	path := req.path
	if strings.Index(path, "/") == 0 {
		path = path[1:]
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

func LimitReadCloser(r io.Reader, n int64) io.Reader {
	var lc LimitedReadCloser
	lc.R = r
	lc.N = n
	return &lc
}

// LimitedRC support Close()
type LimitedReadCloser struct {
	io.LimitedReader
}

func (lc *LimitedReadCloser) Close() error {
	if closer, ok := lc.R.(io.ReadCloser); ok {
		return closer.Close()
	}
	return nil
}
