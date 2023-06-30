package dhttp

import (
	"net/http"
)

type BaseHttpHandler struct {
	base string
}

func (h *BaseHttpHandler) GetBaseUrl() string {
	return h.base
}

func (h *BaseHttpHandler) HandleReq(_ *Options) {

}
func (h *BaseHttpHandler) HandleRes(resp *http.Response, options *Options) (interface{}, error) {
	return nil, nil
}
