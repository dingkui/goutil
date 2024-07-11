package jsonutil

import (
	"encoding/json"
	"errors"
	"gitee.com/dk83/goutils/zlog"
	"reflect"
	"strings"
)

type jsonType byte

const (
	jsonUnkonw jsonType = iota
	jsonMap
	jsonArray
	jsonString
	jsonInt
	jsonFloat
	jsonBool
)

type JsonGo struct {
	v     interface{}
	_type jsonType
}

func (j *JsonGo) mapData() (map[string]*JsonGo, error) {
	if j._type == jsonMap {
		m, ok := j.v.(map[string]*JsonGo)
		if ok {
			return m, nil
		}
	}
	return nil, errors.New("target is not a Map")
}
func (j *JsonGo) arrayData() ([]*JsonGo, error) {
	if j._type == jsonArray {
		m, ok := j.v.([]*JsonGo)
		if ok {
			return m, nil
		}
	}
	return nil, errors.New("target is not a Array")
}

func mkJsonGoByBytes(bytes []byte, re *JsonGo) error {
	if bytes == nil {
		return errors.New("input is nil when mkJsonGoByBytes")
	}
	//defer re.covert()
	if regJsonMap.Match(bytes) {
		err := json.Unmarshal(bytes, &re.v)
		if err == nil {
			re._type = jsonMap
			return nil
		}
	}
	if regJsonArray.Match(bytes) {
		err := json.Unmarshal(bytes, &re.v)
		if err == nil {
			re._type = jsonMap
			return nil
		}
	}
	return nil
}

func (j *JsonGo) covert() {
	if j == nil {
		return
	}
	if j._type == jsonMap {
		mapData, ok := j.v.(map[string]interface{})
		if !ok {
			zlog.Error("covert jsonMap error!")
			return
		}
		v := make(map[string]*JsonGo)
		for key, value := range mapData {
			v[key] = NewJsonGo(value)
		}
		j.v = v
	}
	if j._type == jsonArray {
		arrayData, ok := j.v.([]interface{})
		if !ok {
			zlog.Error("covert jsonArray error!")
			return
		}
		var v []*JsonGo
		for _, value := range arrayData {
			v = append(v, NewJsonGo(value))
		}
		j.v = v
	}
}
func NewJsonGo(data interface{}) *JsonGo {
	if data == nil {
		return nil
	}
	re := &JsonGo{}
	_dataJson, ok := data.(JsonGo)
	if ok {
		return &_dataJson
	}
	defer re.covert()
	_map, ok := data.(map[string]interface{})
	if ok {
		re.v = _map
		re._type = jsonMap
		return re
	}
	_array, ok := data.([]interface{})
	if ok {
		re.v = _array
		re._type = jsonArray
		return re
	}
	_dataByte, ok := data.([]byte)
	if ok {
		mkJsonGoByBytes(_dataByte, re)
		return re
	}
	if reflect.TypeOf(data).Kind() == reflect.Ptr {
		dataByte, err := json.Marshal(data)
		if err != nil {
			zlog.Error("NewJSON translate to byte[] fail,:", err, "data:", data)
			return nil
		}
		mkJsonGoByBytes(dataByte, re)
		return re
	}

	_dataStr, ok := data.(string)
	if ok {
		_dataByte = []byte(strings.Trim(_dataStr, " "))
		if regJsonMap.Match(_dataByte) || regJsonArray.Match(_dataByte) {
			mkJsonGoByBytes(_dataByte, re)
			return re
		}
	}
	switch data.(type) {
	case string:
		re._type = jsonString
	case int:
		re._type = jsonInt
	case float64:
		re._type = jsonFloat
	case bool:
		re._type = jsonBool
	default:
		re._type = jsonUnkonw
	}
	re.v = data
	return re
}
