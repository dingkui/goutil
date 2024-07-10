package jsonutil

import (
	"errors"
	"fmt"
	"gitee.com/dk83/goutils/zlog"
)

func (j *JsonGo) mapData() (map[string]interface{}, error) {
	if j._type == jsonMap {
		m, ok := j.v.(map[string]interface{})
		if ok {
			return m, nil
		}
	}
	return nil, errors.New("target is not a Map")
}
func (j *JsonGo) arrayData() ([]interface{}, error) {
	if j._type == jsonArray {
		m, ok := j.v.([]interface{})
		if ok {
			return m, nil
		}
	}
	return nil, errors.New("target is not a Array")
}

func (j *JsonGo) covert() {
	if j == nil {
		return
	}
	if j._type == jsonMap {
		mapData, err := j.mapData()
		if err != nil {
			zlog.Error(err)
			return
		}
		v := make(map[string]interface{})
		for key, value := range mapData {
			v[key] = value
			_, ok := value.(JsonGo)
			if !ok {
				json := NewJSON(value)
				if json.IsNil() {
					v[key] = value
				}
			}
		}
		j.v = v
	}
	if j._type == jsonArray {
		arrayData, err := j.arrayData()
		if err != nil {
			zlog.Error(err)
			return
		}
		v := append([]interface{}{}, arrayData...)
		for index, value := range v {
			_, ok := value.(JsonGo)
			if !ok {
				v[index] = NewJSON(value)
			}
		}
		j.v = v
	}
}

func (j *JsonGo) getItem(key interface{}) (interface{}, error) {
	err := checkKeys(key)
	if err != nil {
		return nil, err
	}
	err = j.check()
	if err != nil {
		return nil, err
	}

	var item interface{}
	switch key.(type) {
	case string:
		data, err := j.mapData()
		if err != nil {
			return nil, err
		}
		_item, ok := data[key.(string)]
		if !ok {
			return nil, errors.New(fmt.Sprintf("can't get [%s] from:%v+", key, j))
		}
		item = _item
	case int:
		data, err := j.arrayData()
		if err != nil {
			return nil, err
		}
		item = data[key.(int)]
	default:
		return nil, errors.New(fmt.Sprintln("getItem fail:[%T] %v+ %v+", key, key, j))
	}
	return item, nil
}

func (j *JsonGo) Native() (interface{}, error) {
	if j._type == jsonMap {
		return j.mapNative()
	}
	if j._type == jsonArray {
		return j.arrayNative()
	}
	return nil, errors.New("unReachAble")
}

func (j *JsonGo) mapNative() (map[string]interface{}, error) {
	data, err := j.mapData()
	if err != nil {
		return nil, err
	}

	_re := make(map[string]interface{})
	for key, value := range data {
		obj, ok := value.(JSONObject)
		if ok {
			_re[key], err = obj.Native()
			if err != nil {
				return nil, err
			}
		} else {
			_re[key] = value
		}
	}
	return _re, nil
}
func (j *JsonGo) arrayNative() ([]interface{}, error) {
	data, err := j.arrayData()
	if err != nil {
		return nil, err
	}

	_re := append([]interface{}{}, data...)
	for key, value := range _re {
		obj, ok := value.(JSONObject)
		if ok {
			_re[key], err = obj.Native()
			if err != nil {
				return nil, err
			}
		} else {
			_re[key] = value
		}
	}
	return _re, nil
}
