package jsonutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitee.com/dk83/goutils/zlog"
	"reflect"
	"strings"
)

type jsonType byte

const (
	jsonSimple jsonType = iota
	jsonMap
	jsonArray
)

type JsonGo struct {
	v     interface{}
	_type jsonType
}

func (j *JsonGo) check() error {
	if j == nil || j.v == nil || j._type == jsonArray {
		return errors.New("target is not a JsonGo")
	}
	return nil
}

func (j *JsonGo) Byte() ([]byte, error) {
	native, err := j.Native()
	if err != nil {
		zlog.Error(err)
		return nil, err
	}
	return json.Marshal(native)
}

func (j *JsonGo) Str() string {
	bytes, err := j.Byte()
	if err != nil {
		zlog.Error(err)
		return ""
	}
	return string(bytes)
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

func unmarshalJsonGo(bytes []byte, re *JsonGo) error {
	if bytes == nil {
		return errors.New("input is nil where Unmarshal")
	}
	defer re.covert()
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

func NewJsonGo(data interface{}) *JsonGo {
	if data == nil {
		return nil
	}
	re := &JsonGo{}
	defer re.covert()
	_dataJson, ok := data.(JsonGo)
	if ok {
		return &_dataJson
	}
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
		unmarshalJsonGo(_dataByte, re)
		return re
	}
	if reflect.TypeOf(data).Kind() == reflect.Ptr {
		dataByte, err := json.Marshal(data)
		if err != nil {
			zlog.Error("NewJSON translate to byte[] fail,:", err, "data:", data)
			return nil
		}
		unmarshalJsonGo(dataByte, re)
		return re
	}

	_dataStr, ok := data.(string)
	if ok {
		_dataByte = []byte(strings.Trim(_dataStr, " "))
		if regJsonMap.Match(_dataByte) || regJsonArray.Match(_dataByte) {
			unmarshalJsonGo(_dataByte, re)
			return re
		}
	}
	return re
}
