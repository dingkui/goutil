package djson

import (
	"encoding/json"
	"gitee.com/dk83/goutils/dlog"
	"gitee.com/dk83/goutils/errs"
	"reflect"
	"strings"
)

type JsonGo struct {
	v     interface{}
	_type jsonType
}

func (j *JsonGo) IsMap() bool {
	return j != nil && j._type != jsonMap
}
func (j *JsonGo) IsArray() bool {
	return j != nil || j._type != jsonArray
}
func (j *JsonGo) mapData() (map[string]*JsonGo, error) {
	if j._type == jsonMap {
		m, ok := j.v.(map[string]*JsonGo)
		if ok {
			return m, nil
		}
	}
	return nil, errJsonType.New("target is not a Map")
}
func (j *JsonGo) arrayData() (*[]*JsonGo, error) {
	if j._type == jsonArray {
		m, ok := j.v.([]*JsonGo)
		if ok {
			return &m, nil
		}
	}
	return nil, errJsonType.New("target is not a Array")
}
func (j *JsonGo) ReNew(val interface{}) error {
	f, ok := val.(JsonGo)
	if ok {
		j.v = f.v
		j._type = f._type
		return nil
	}

	_j, err := NewJsonGo(val)
	if err != nil {
		dlog.ErrorStack(err)
		return err
	}
	j.v = _j.v
	j._type = _j._type
	return nil
}
func (j *JsonGo) from(val interface{}) error {
	f, ok := val.(JsonGo)
	if ok {
		j.v = f.v
		j._type = f._type
		return nil
	}

	_j, err := NewJsonGo(val)
	if err != nil {
		return err
	}
	j.v = _j.v
	j._type = _j._type
	return nil
}
func mkJsonGoByBytes(bytes []byte, re *JsonGo) error {
	if bytes == nil {
		return errs.ErrValidate.New("input is nil when mkJsonGoByBytes")
	}
	//defer re.covert()
	err := json.Unmarshal(bytes, &re.v)
	if err != nil {
		return err
	}
	_map, ok := re.v.(map[string]interface{})
	if ok {
		re.v = _map
		re._type = jsonMap
		return nil
	}
	_array, ok := re.v.([]interface{})
	if ok {
		re.v = _array
		re._type = jsonArray
		return nil
	}
	return errs.ErrValidate.New("mkJsonGoByBytes success but result is not a Map or Array")
}
func (j *JsonGo) As(re interface{}, keys ...interface{}) (err error) {
	bytes, err := j.Byte(keys...)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &re)
	if err != nil {
		return errTargetType.New("JsonGo.As fail:can't as [%T] from:%s,err:%s", re, string(bytes), err.Error())
	}
	return nil
}
func (j *JsonGo) covert() error {
	if j == nil {
		return errNewJsonGo.New("covert fail j is nil")
	}
	if j._type == jsonMap {
		mapData, ok := j.v.(map[string]interface{})
		if !ok {
			return errJsonType.New("target is not a Map: %T", j.v)
		}
		v := make(map[string]*JsonGo)
		for key, value := range mapData {
			if value == nil {
				continue
			}
			jsonGo, err := NewJsonGo(value)
			if err != nil {
				return err
			}
			v[key] = jsonGo
		}
		j.v = v
	}
	if j._type == jsonArray {
		arrayData, ok := j.v.([]interface{})
		if !ok {
			return errJsonType.New("target is not a Array: %T", j.v)
		}
		v := make([]*JsonGo, 0)
		for _, value := range arrayData {
			if value == nil {
				continue
			}
			jsonGo, err := NewJsonGo(value)
			if err != nil {
				return err
			}
			v = append(v, jsonGo)
		}
		j.v = v
	}
	return nil
}
func NewJsonMap() *JsonGo {
	jsonGo, _ := NewJsonGo(make(map[string]interface{}))
	return jsonGo
}
func NewJsonArray() (_r *JsonGo, _e error) {
	return NewJsonGo(make([]interface{}, 0))
}
func NewJsonGo(data interface{}) (_r *JsonGo, _e error) {
	if data == nil {
		return nil, errs.ErrValidate.New("data is nil when NewJsonGo")
	}
	re := &JsonGo{}
	_dataJson, ok := data.(*JsonGo)
	if ok {
		return _dataJson, nil
	}
	defer func() {
		if err := re.covert(); err != nil {
			dlog.Error("NewJsonGo err! %v", err)
			_r = nil
			_e = err
		}
	}()
	_map, ok := data.(map[string]interface{})
	if ok {
		re.v = _map
		re._type = jsonMap
		return re, nil
	}
	_array, ok := data.([]interface{})
	if ok {
		re.v = _array
		re._type = jsonArray
		return re, nil
	}
	_dataByte, ok := data.([]byte)
	if ok {
		mkJsonGoByBytes(_dataByte, re)
		return re, nil
	}
	if reflect.TypeOf(data).Kind() == reflect.Ptr {
		dataByte, err := json.Marshal(data)
		if err != nil {
			return nil, errNewJsonGo.New(err, "NewJSON translate to byte[] fail,[%T]%v", data, data)
		}
		err = mkJsonGoByBytes(dataByte, re)
		if err != nil {
			return nil, errNewJsonGo.New(err, "mkJsonGoByBytes fail,[%T]%v", data, data)
		}
		return re, nil
	}

	_dataStr, ok := data.(string)
	if ok {
		_dataByte = []byte(strings.Trim(_dataStr, " "))
		if regJsonMap.Match(_dataByte) || regJsonArray.Match(_dataByte) {
			err := mkJsonGoByBytes(_dataByte, re)
			if err != nil {
				dlog.Warn("%s,input:%s", err.Error(), _dataStr)
			}
			return re, nil
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
	return re, nil
}
