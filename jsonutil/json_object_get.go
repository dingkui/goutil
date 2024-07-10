package jsonutil

import (
	"gitee.com/dk83/goutils/zlog"
)

func (j *JSONObject) GetJSONObjectByKeys(keys ...interface{}) *JSONObject {
	errorRe := &JSONObject{}
	item, err := j.GetItemByKeys(keys...)
	if err != nil {
		zlog.Error(err)
		return errorRe
	}
	re, ok := item.(JSONObject)
	if !ok {
		zlog.Error("GetJson error item:%v", item)
		return errorRe
	}
	return &re
}
func (j *JSONObject) GetJSONArrayByKeys(key ...interface{}) *JSONArray {
	return j.GetJSONObjectByKeys(key...).Array()
}
func (j *JSONObject) GetJSONMapByKeys(key ...interface{}) *JSONMap {
	return j.GetJSONObjectByKeys(key...).Map()
}

func (j *JSONObject) GetJSONObject(key string) *JSONObject {
	return j.GetJSONObjectByKeys(getkeys(key)...)
}
func (j *JSONObject) GetJSONArray(key string) *JSONArray {
	return j.GetJSONArrayByKeys(getkeys(key)...)
}
func (j *JSONObject) GetJSONMap(key string) *JSONMap {
	return j.GetJSONMapByKeys(getkeys(key)...)
}
