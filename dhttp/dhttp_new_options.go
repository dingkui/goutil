package dhttp

import (
	"gitee.com/dk83/goutils/djson"
	"io"
)

func NewOpiton(path string, handler IReqHandler) *Options {
	if handler == nil {
		panic("handler is nil")
	}
	options := &Options{}
	options.Path(path)
	options.ReqHandler(handler)
	return options
}
func NewStramOpiton(path string, handler IReqHandler, body io.Reader) *Options {
	return NewOpiton(path, handler).DataStram(body)
}
func NewFormOpiton(path string, handler IReqHandler, data *djson.JsonGo) *Options {
	return NewOpiton(path, handler).DataFrom(data)
}
func NewJsonOpiton(path string, handler IReqHandler, data *djson.JsonGo) *Options {
	return NewOpiton(path, handler).DataJson(data)
}
