package jsonutil

import (
	"errors"
	"fmt"
	"gitee.com/dk83/goutils/zlog"
)

//设置interface中的值，只支持map和数组
func (j *JsonGo) setItem(key interface{}, val interface{}) error {
	switch key.(type) {
	case int:
		data, err := j.arrayData()
		if err != nil {
			return err
		}
		data[key.(int)] = val
	case string:
		data, err := j.mapData()
		if err != nil {
			return err
		}
		data[key.(string)] = val
	default:
		return errors.New(fmt.Sprintf("setItem type is not sopport %T", j))
	}
	return nil
}

//设置interface中的值，支持多级设置，支持map,数组和json字符串
func (j *JsonGo) SetItem(key string, val interface{}) bool {
	return j.SetItemByKeys(val, getkeys(key)...)
}

//设置interface中的值，支持多级设置，支持map,数组和json字符串
func (j *JsonGo) SetItemByKeys(val interface{}, keys ...interface{}) bool {
	err := checkKeys(keys...)
	if err != nil {
		zlog.Error(err)
		return false
	}
	if len(keys) == 1 {
		err := j.setItem(keys[0], val)
		if err != nil {
			zlog.Error(err)
			return false
		}
	} else if len(keys) > 1 {
		item := j.GetJsonByKeys(keys[:len(keys)-2])
		err := item.check()
		if err != nil {
			zlog.Error(err)
			return false
		}
		err = item.setItem(keys[len(keys)-1], val)
		if err != nil {
			zlog.Error(err)
			return false
		}
	}
	return true
}
