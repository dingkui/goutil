package jsonutil

import (
	"errors"
	"fmt"
	"gitee.com/dk83/goutils/zlog"
)

func (j *JsonGo) getItem(key interface{}) (*JsonGo, error) {
	err := checkKeys(key)
	if err != nil {
		return nil, err
	}

	switch key.(type) {
	case string:
		data, err := j.mapData()
		if err != nil {
			return nil, err
		}
		item, ok := data[key.(string)]
		if !ok {
			return nil, errors.New(fmt.Sprintf("can't get [%s] from:%v", key, j))
		}
		return item, nil
	case int:
		data, err := j.arrayData()
		if err != nil {
			return nil, err
		}
		i := key.(int)
		if i < 0 || i >= len(data) {
			return nil, errors.New(fmt.Sprintf("getItem index is out of range %d", i))
		}
		return data[key.(int)], nil
	}
	return nil, errors.New(fmt.Sprintf("getItem fail:[%T] %v %v", j, key, key))
}

func (j *JsonGo) AtKeys(keys ...interface{}) (*JsonGo, error) {
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
	return item.AtKeys(keys[1:]...)
}
func (j *JsonGo) At(key string) *JsonGo {
	value, err := j.AtKeys(getkeys(key)...)
	if err != nil {
		zlog.Error(err)
		return &JsonGo{}
	}
	return value
}
