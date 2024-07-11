package jsonutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitee.com/dk83/goutils/zlog"
)

func (j *JsonGo) Native() (interface{}, error) {
	if j._type == jsonMap {
		return j.mapNative()
	}
	if j._type == jsonArray {
		return j.arrayNative()
	}
	return j.v, nil
}

func (j *JsonGo) mapNative() (map[string]interface{}, error) {
	data, err := j.mapData()
	if err != nil {
		return nil, err
	}

	_re := make(map[string]interface{})
	for key, value := range data {
		_value, err := value.Native()
		if err != nil {
			return nil, err
		}
		_re[key] = _value
	}
	return _re, nil
}
func (j *JsonGo) arrayNative() ([]interface{}, error) {
	data, err := j.arrayData()
	if err != nil {
		return nil, err
	}

	_re := []interface{}{}
	for _, value := range data {
		_value, err := value.Native()
		if err != nil {
			return nil, err
		}
		_re = append(_re, _value)
	}
	return _re, nil
}

func (j *JsonGo) Byte() ([]byte, error) {
	native, err := j.Native()
	if err != nil {
		return nil, err
	}
	return json.Marshal(native)
}

func (j *JsonGo) Str() string {
	if j._type == jsonString {
		return j.v.(string)
	}
	bytes, err := j.Byte()
	if err != nil {
		zlog.Error(err)
		return ""
	}
	return string(bytes)
}
func (j *JsonGo) ValueArray() []interface{} {
	native, err := j.arrayNative()
	if err != nil {
		zlog.Error(err)
		return nil
	}
	return native
}
func (j *JsonGo) ValueMap() map[string]interface{} {
	native, err := j.mapNative()
	if err != nil {
		zlog.Error(err)
		return nil
	}
	return native
}
func (j *JsonGo) ValueStr() string {
	native, err := j.Native()
	if err != nil {
		zlog.Error(err)
		return ""
	}
	return native.(string)
}
func (j *JsonGo) ValueFloat64() (float64, error) {
	native, err := j.Native()
	if err != nil {
		return 0, err
	}
	f, ok := native.(float64)
	if !ok {
		return 0, errors.New("value is not float64")
	}
	return f, nil
}
func (j *JsonGo) ValueBool() (bool, error) {
	native, err := j.Native()
	if err != nil {
		return false, err
	}
	f, ok := native.(bool)
	if !ok {
		return false, errors.New("value is not bool")
	}
	return f, nil
}

func (j *JsonGo) As(re interface{}) (err error) {
	bytes, err := j.Byte()
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &re)
	if err != nil {
		return errors.New(fmt.Sprintf("JSON.As fail:can't as [%T] from:%s,err:%s", re, string(bytes), err.Error()))
	}
	return nil
}
