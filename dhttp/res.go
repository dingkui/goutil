package dhttp

import (
	"io"
	"net/http"
)

// Response defines HTTP response from OSS
type Response struct {
	StatusCode int
	Headers    http.Header
	Body       io.ReadCloser
}

func (r *Response) Read(p []byte) (n int, err error) {
	return r.Body.Read(p)
}

// Close close http reponse body
func (r *Response) Close() error {
	return r.Body.Close()
}
