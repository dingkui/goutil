package jsonutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitee.com/dk83/goutils/logutil"
)

type JSON struct {
	_data []byte
}

func MkJSON(data interface{}) *JSON {
	if data == nil {
		return nil
	}
	_dataJson, ok := data.(JSON)
	if ok {
		return &_dataJson
	}
	_dataByte, ok := data.([]byte)
	if ok {
		return &JSON{
			_data: _dataByte,
		}
	}
	_dataStr, ok := data.(string)
	if ok {
		return &JSON{
			_data: []byte(_dataStr),
		}
	}
	_dataByte, e := json.Marshal(data)
	if e == nil {
		return &JSON{
			_data: _dataByte,
		}
	}
	return nil
}

func (j *JSON) Array() (re []interface{}, err error) {
	if j == nil || j._data == nil {
		return nil, errors.New(fmt.Sprintf("JSON.Array fail:data is nil"))
	}
	err = json.Unmarshal(j._data, &re)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("JSON.Array fail: can't as [%T] from:%s,err:%s", re, j.Str(), err.Error()))
	}
	return re, nil
}
func (j *JSON) Map() (re map[string]interface{}, err error) {
	if j == nil || j._data == nil {
		return nil, errors.New(fmt.Sprintf("JSON.Map fail:data is nil"))
	}
	err = json.Unmarshal(j._data, &re)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("JSON.Map fail:can't as [%T] from:%s,err:%s", re, j.Str(), err.Error()))
	}
	return re, nil
}
func (j *JSON) As(re interface{}) (err error) {
	if j == nil || j._data == nil {
		return errors.New(fmt.Sprintf("JSON.As fail:data is nil"))
	}
	err = json.Unmarshal(j._data, &re)
	if err != nil {
		return errors.New(fmt.Sprintf("JSON.As fail:can't as [%T] from:%s,err:%s", re, j.Str(), err.Error()))
	}
	return nil
}
func (j *JSON) Str() string {
	if j == nil || j._data == nil {
		return ""
	}
	return string(j._data)
}
func (j *JSON) Bytes() []byte {
	return j._data
}

func (j *JSON) getItem(key interface{}) (interface{}, error) {
	err := checkKeys(key)
	if err != nil {
		logutil.Error(err)
		return nil, err
	}
	_key, ok := key.(string)
	if ok {
		asMap, err := j.Map()
		if err != nil {
			return nil, err
		}
		item, ok := asMap[_key]
		if !ok {
			return nil, errors.New(fmt.Sprintf("can't get [%s] from:%s", _key, j.Str()))
		}
		return item, nil
	}
	_keyInt, ok := key.(int)
	if ok {
		asArray, err := j.Array()
		if err != nil {
			return nil, err
		}
		if _keyInt < 0 || _keyInt > len(asArray)-1 {
			return nil, errors.New(fmt.Sprintf("can't get [%d] from:%s", _keyInt, j.Str()))
		}
		return asArray[_keyInt], nil
	}
	return nil, errors.New(fmt.Sprintln("getJson fail:", _key, j.Str()))
}

func (j *JSON) GetJson(keys ...interface{}) *JSON {
	return MkJSON(j.GetItem(keys...))
}
func (j *JSON) GetItem(keys ...interface{}) interface{} {
	err := checkKeys(keys...)
	if err != nil {
		logutil.Error(err)
		return nil
	}
	if len(keys) == 0 {
		return j
	}
	item, err := j.getItem(keys[0])
	if err != nil {
		logutil.Error(err)
		return nil
	}
	if len(keys) == 1 {
		return item
	}
	return MkJSON(item).GetItem(keys[1:]...)
}

func (j *JSON) GetArray(key ...interface{}) []interface{} {
	item := j.GetJson(key...)
	if item == nil {
		return nil
	}
	array, _ := item.Array()
	return array
}
func (j *JSON) GetMap(key ...interface{}) map[string]interface{} {
	item := j.GetJson(key...)
	if item == nil {
		return nil
	}
	array, _ := item.Map()
	return array
}
func (j *JSON) GetStr(key ...interface{}) string {
	item := j.GetJson(key...)
	if item == nil {
		return ""
	}
	return item.Str()
}
func (j *JSON) GetNum(key ...interface{}) (float64, error) {
	item := j.GetItem(key...)
	if item == nil {
		return 0, errors.New("not exsits")
	}
	return item.(float64), nil
}
func (j *JSON) GetBool(key ...interface{}) (bool, error) {
	item := j.GetItem(key...)
	if item == nil {
		return false, errors.New("not exsits")
	}
	return item.(bool), nil
}
