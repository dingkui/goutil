package dhttp

import (
	"fmt"
	"gitee.com/dk83/goutils/djson"
	"net/url"
	"strings"
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

func (ops *Options) GetUri(base string) (*url.URL, error) {
	path := ops.GetPath()
	params := ops.GetUrlParams()
	urlParams := GetURLParams(params)
	if strings.LastIndex(base, "/") == len(base)-1 {
		base = base[:len(base)-1]
	}
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
