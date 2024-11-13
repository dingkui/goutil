package dhttp

import (
	"github.com/dingkui/goutil/djson"
)

func (ops *Options) AddUrlParams(urlParam *djson.JsonGo) {
	if urlParam != nil {
		m, _ := urlParam.Map()
		for key, jsonGo := range m {
			ops.AddUrlParam(key, jsonGo.NativeN())
		}
	}
}
func (ops *Options) AddUrlParam(key string, value interface{}) {
	ops.addOption(key, value, optionUrlParam)
}
func (ops *Options) GetUrlParams() map[string]interface{} {
	paramsm := map[string]interface{}{}
	for _, option := range ops.options {
		if option.ot == optionUrlParam {
			paramsm[option.key] = option.value
		}
	}
	return paramsm
}
func (ops *Options) GetUrlParam(key string) interface{} {
	for _, option := range ops.options {
		if option.ot == optionUrlParam && option.key == key {
			return option.value
		}
	}
	return ""
}
