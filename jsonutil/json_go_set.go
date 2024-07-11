package jsonutil

import (
	"errors"
	"fmt"
	"gitee.com/dk83/goutils/zlog"
)

//设置interface中的值，只支持map和数组
func (j *JsonGo) SetValue(key interface{}, val interface{}) error {
	v := NewJsonGo(val)
	switch key.(type) {
	case int:
		data, err := j.arrayData()
		if err != nil {
			return err
		}
		i := key.(int)
		if i < -1 || i >= len(data) {
			return errors.New(fmt.Sprintf("setItem index is out of range %d", i))
		}
		if i == -1 {
			data = append(data, v)
			return nil
		}
		data[i] = v
	case string:
		s := key.(string)

		if j._type == jsonMap {
			data, err := j.mapData()
			if err != nil {
				return err
			}
			data[s] = v
		} else if j._type == jsonArray {
			_key, err := getInt(s)
			if err != nil {
				return err
			}
			j.SetValue(_key, v)
		} else {
			j.v = v
		}
		return nil
	}
	return errors.New(fmt.Sprintf("setItem type is not sopport %T", j))
}

//设置interface中的值，支持多级设置，支持map,数组和json字符串
func (j *JsonGo) Set(key string, val interface{}) bool {
	return j.SetByKeys(val, getkeys(key)...)
}

//设置interface中的值，支持多级设置，支持map,数组和json字符串
func (j *JsonGo) SetByKeys(val interface{}, keys ...interface{}) bool {
	err := checkKeys(keys...)
	if err != nil {
		zlog.Error(err)
		return false
	}
	if len(keys) == 1 {
		err := j.SetValue(keys[0], val)
		if err != nil {
			zlog.Error(err)
			return false
		}
	} else if len(keys) > 1 {
		item, err := j.AtKeys(keys[:len(keys)-1]...)
		if err != nil {
			zlog.Error(err)
			return false
		}
		err = item.SetValue(keys[len(keys)-1], val)
		if err != nil {
			zlog.Error(err)
			return false
		}
	}
	return true
}
