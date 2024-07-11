package jsonutil

//import (
//	"encoding/json"
//	"errors"
//	"gitee.com/dk83/goutils/zlog"
//	"reflect"
//	"strings"
//)
//
//type JSONArray struct {
//	JsonGo
//}
//type JSONMap struct {
//	JsonGo
//}
//type JSONObject struct {
//	JSONMap
//	JSONArray
//}
//
//func (j *JSONArray) CheckArray() error {
//	if j == nil || j.v == nil || j._type != jsonArray {
//		return errors.New("target is not a Array")
//	}
//	return nil
//}
//func (j *JSONArray) Array() *JSONArray {
//	err := j.CheckArray()
//	if err != nil {
//		zlog.Error(err)
//		return &JSONArray{}
//	}
//
//	var v interface{} = j
//	m, _ := v.(*JSONArray)
//	return m
//}
//
//
//func (j *JSONMap) CheckMap() error {
//	if j == nil || j.v == nil || j._type != jsonMap {
//		return errors.New("target is not a Map")
//	}
//	return nil
//}
//func (j *JSONMap) Map() *JSONMap {
//	err := j.CheckMap()
//	if err != nil {
//		zlog.Error(err)
//		return &JSONMap{}
//	}
//
//	var v interface{} = j
//	m, _ := v.(*JSONMap)
//	return m
//}
//
//func Unmarshal(bytes []byte, re *JSONObject) error {
//	if bytes == nil {
//		return errors.New("input is nil where Unmarshal")
//	}
//	defer re.covert()
//	if regJsonMap.Match(bytes) {
//		err := json.Unmarshal(bytes, &re.JSONMap.v)
//		if err == nil {
//			re._type = jsonMap
//			return nil
//		}
//	}
//	if regJsonArray.Match(bytes) {
//		err := json.Unmarshal(bytes, &re.JSONArray.v)
//		if err == nil {
//			re._type = jsonMap
//			return nil
//		}
//	}
//	return nil
//}
//
//func NewJsonObject(data interface{}) *JSONObject {
//	if data == nil {
//		return nil
//	}
//	re := &JSONObject{}
//	defer re.covert()
//	_dataJson, ok := data.(JSONObject)
//	if ok {
//		return &_dataJson
//	}
//	_map, ok := data.(map[string]interface{})
//	if ok {
//		re.JSONMap.v = _map
//		re._type = jsonMap
//		return re
//	}
//	_array, ok := data.([]interface{})
//	if ok {
//		re.JSONArray.v = _array
//		re._type = jsonArray
//		return re
//	}
//	_dataByte, ok := data.([]byte)
//	if ok {
//		Unmarshal(_dataByte, re)
//		return re
//	}
//	if reflect.TypeOf(data).Kind() == reflect.Ptr {
//		dataByte, err := json.Marshal(data)
//		if err != nil {
//			zlog.Error("NewJSON translate to byte[] fail,:", err, "data:", data)
//			return nil
//		}
//		Unmarshal(dataByte, re)
//		return re
//	}
//
//	_dataStr, ok := data.(string)
//	if ok {
//		_dataByte = []byte(strings.Trim(_dataStr, " "))
//		if regJsonMap.Match(_dataByte) || regJsonArray.Match(_dataByte) {
//			Unmarshal(_dataByte, re)
//			return re
//		}
//	}
//	return re
//}
