package jsonutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitee.com/dk83/goutils/zlog"
	"strconv"
	"strings"
)

func formatJson(input []byte, format bool) []byte {
	if !format {
		return input
	}
	var out bytes.Buffer
	json.Indent(&out, input, "", "  ")
	return out.Bytes()
}

func CopyMapVal(target map[string]interface{}, source map[string]interface{}, keys ...string) {
	for _, key := range keys {
		value, has := source[key]
		if has {
			target[key] = value
		}
	}
}

func Str(def string, data interface{}) (string, error) {
	switch data.(type) {
	case *JsonFile:
		return data.(*JsonFile).Str(def)
	case *JsonGo:
		return data.(*JsonGo).Str(def)
	case string:
		return data.(string), nil
	case int:
		return strconv.Itoa(data.(int)), nil
	case []byte:
		return string(data.([]byte)), nil
	}

	b, e := json.Marshal(data)
	if e != nil {
		return def, e
	}

	return string(b), nil
}

func StrN(def string, data interface{}) string {
	str, err := Str(def, data)
	if err != nil {
		zlog.Error(err)
	}
	return str
}
func Byte(data interface{}) ([]byte, error) {
	if data == nil {
		return nil, errTargetType.New("value is not Bytes")
	}
	switch data.(type) {
	case *JsonGo:
		return data.(*JsonGo).Byte()
	case *JsonFile:
		return data.(*JsonFile).Byte()
	case string:
		return []byte(data.(string)), nil
	}

	return json.Marshal(data)
}
func Bool(def bool, data interface{}) (bool, error) {

	if data == nil {
		return def, errTargetType.New("value is not Bool")
	}

	if data == nil {
		return false, nil
	}
	switch data.(type) {
	case bool:
		return data.(bool), nil
	case string:
		str := strings.ToLower(data.(string))
		return str != "false" && str != "0" && str != "", nil
	case int:
		i := data.(int)
		return i != 0, nil
	case *JsonGo:
		return data.(*JsonGo).Bool(def)
	case *JsonFile:
		return data.(*JsonFile).Bool(def)
	}

	return def, errTargetType.New("value is not Bool")
}
func Int(def int, data interface{}) (int, error) {
	if data == nil {
		return def, errTargetType.New("value is not Int")
	}
	switch data.(type) {
	case int:
		i := data.(int)
		return i, nil
	case string:
		num, err := strconv.Atoi(data.(string))
		if err != nil {
			return def, err
		}
		return num, nil
	case *JsonGo:
		return data.(*JsonGo).Int(def)
	case *JsonFile:
		return data.(*JsonFile).Int(def)
	}

	return def, errTargetType.New("value is not Int")
}
func Float(def float64, data interface{}) (float64, error) {
	if data == nil {
		return def, errTargetType.New("value is not Float")
	}
	switch data.(type) {
	case float64:
		i := data.(float64)
		return i, nil
	case int:
		num, err := strconv.ParseFloat(fmt.Sprintf("%d", data), 64)
		if err != nil {
			return def, err
		}
		return num, nil
	case string:
		num, err := strconv.ParseFloat(data.(string), 64)
		if err != nil {
			return def, err
		}
		return num, nil
	case *JsonGo:
		return data.(*JsonGo).Float(def)
	case *JsonFile:
		return data.(*JsonFile).Float(def)
	}

	return def, errTargetType.New("value is not Float")
}
