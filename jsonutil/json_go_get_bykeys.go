package jsonutil

import (
	"errors"
	"fmt"
	"gitee.com/dk83/goutils/zlog"
)

func (j *JsonGo) GetItemByKeys(keys ...interface{}) (interface{}, error) {
	err := checkKeys(keys...)
	if err != nil {
		return nil, err
	}
	if len(keys) == 0 {
		return j, nil
	}

	item, err := j.getItem(keys[0])
	if err != nil {
		return nil, err
	}

	if len(keys) == 1 {
		return item, nil
	}
	last, ok := item.(JsonGo)
	if ok {
		return last.GetItemByKeys(keys[1:]...)
	}
	return nil, errors.New(fmt.Sprintf("GetItem error obj:%v keys:%v", item, keys))
}

func (j *JsonGo) GetJsonByKeys(keys ...interface{}) *JsonGo {
	errorRe := &JsonGo{}
	item, err := j.GetItemByKeys(keys...)
	if err != nil {
		zlog.Error(err)
		return errorRe
	}
	re, ok := item.(JsonGo)
	if !ok {
		zlog.Error("GetJson error item:%v", item)
		return errorRe
	}
	return &re
}

func (j *JsonGo) GetArrayByKeys(key ...interface{}) []interface{} {
	item := j.GetJsonByKeys(key...)
	native, err := item.arrayNative()
	if err != nil {
		zlog.Error(err)
		return nil
	}
	return native
}
func (j *JsonGo) GetMapByKeys(key ...interface{}) map[string]interface{} {
	item := j.GetJsonByKeys(key...)
	native, err := item.mapNative()
	if err != nil {
		zlog.Error(err)
		return nil
	}
	return native
}

func (j *JsonGo) GetArrayIntByKeys(key ...interface{}) []int {
	array := j.GetJsonByKeys(key...)
	native, err := array.Native()
	if err != nil {
		zlog.Error(err)
		return nil
	}
	return native.([]int)
}
func (j *JsonGo) GetStrByKeys(key ...interface{}) string {
	item, err := j.GetItemByKeys(key...)
	if err != nil {
		zlog.Error(err)
		return ""
	}
	return item.(string)
}
func (j *JsonGo) GetNumByKeys(key ...interface{}) (float64, error) {
	item, err := j.GetItemByKeys(key...)
	if err != nil {
		return 0, err
	}
	return item.(float64), nil
}
func (j *JsonGo) GetBoolByKeys(key ...interface{}) (bool, error) {
	item, err := j.GetItemByKeys(key...)
	if err == nil {
		return false, err
	}
	return item.(bool), nil
}
