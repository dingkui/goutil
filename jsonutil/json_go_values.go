package jsonutil

import (
	"encoding/json"
	"gitee.com/dk83/goutils/zlog"
)

// Native 获取原生数据
func (j *JsonGo) Native(keys ...interface{}) (interface{}, error) {
	t, err := j.Get(keys...)
	if err != nil {
		return nil, err
	}

	if t._type == jsonMap {
		return t.NativeMap()
	}
	if t._type == jsonArray {
		return t.NativeArray()
	}
	return t.v, nil
}

func (j *JsonGo) NativeArray(keys ...interface{}) ([]interface{}, error) {
	t, err := j.Get(keys...)
	if err != nil {
		return nil, err
	}

	data, err := t.arrayData()
	if err != nil {
		return nil, err
	}

	var _re []interface{}

	for _, value := range *data {
		_value, err := value.Native()
		if err != nil {
			return nil, err
		}
		_re = append(_re, _value)
	}
	return _re, nil
}
func (j *JsonGo) NativeMap(keys ...interface{}) (map[string]interface{}, error) {
	t, err := j.Get(keys...)
	if err != nil {
		return nil, err
	}

	data, err := t.mapData()
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

func (j *JsonGo) Byte(keys ...interface{}) ([]byte, error) {
	native, err := j.Native(keys...)
	if err != nil {
		return nil, err
	}
	return json.Marshal(native)
}

func (j *JsonGo) Str(keys ...interface{}) string {
	t, err := j.Get(keys...)
	if err != nil {
		zlog.Warn(err)
		return ""
	}

	if t._type == jsonString {
		return t.v.(string)
	}
	bytes, err := t.Byte()
	if err != nil {
		zlog.Error(err)
		return ""
	}
	return string(bytes)
}
func (j *JsonGo) Float(keys ...interface{}) (float64, error) {
	native, err := j.Native(keys...)
	if err != nil {
		return 0, err
	}
	f, ok := native.(float64)
	if !ok {
		return 0, errTargetType.New("value is not float64")
	}
	return f, nil
}
func (j *JsonGo) Bool(keys ...interface{}) (bool, error) {
	native, err := j.Native(keys...)
	if err != nil {
		return false, err
	}
	f, ok := native.(bool)
	if !ok {
		return false, errTargetType.New("value is not bool")
	}
	return f, nil
}

func (j *JsonGo) As(re interface{}, keys ...interface{}) (err error) {
	bytes, err := j.Byte(keys...)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &re)
	if err != nil {
		return errTargetType.New("JSON.As fail:can't as [%T] from:%s,err:%s", re, string(bytes), err.Error())
	}
	return nil
}
