package jsonutil

import (
	"bytes"
	"encoding/json"
	"errors"
	"gitee.com/dk83_admin/goutil/utils/apputil"
	"gitee.com/dk83_admin/goutil/utils/fileutil"
	"gitee.com/dk83_admin/goutil/utils/zlog"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Read_json_file(file string) map[string]interface{} {
	b, e := ioutil.ReadFile(file)
	if e != nil {
		return nil
	}
	rs := make(map[string]interface{})
	e = json.Unmarshal(b, &rs)
	return rs
}

func Readdata_json_file(file string) []byte {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return data
	}
	return nil
}

func Write_json_file(file string, data map[string]interface{}) {
	b, e := json.Marshal(data)
	if e != nil {
		zlog.ErrorLn(e)
		panic(e)
		return
	}

	os.MkdirAll(filepath.Dir(file), os.ModePerm)
	fileutil.WriteAndSyncFile(file, format(b), os.ModePerm)
	//ioutil.WriteFile(file,b,os.ModePerm)
}

func format(input []byte) []byte {
	return formatJson(input, apputil.IsPara("formatJson"))
}

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

func Get_bytes(data map[string]interface{}) []byte {
	b, e := json.Marshal(data)
	if e != nil {
		return b
	}
	return nil
}

func Data2Json(data interface{}) string {
	b, e := json.Marshal(data)
	if e == nil {
		return string(format(b))
	}
	return ""
}

func GetItem(data interface{}, key ...interface{}) interface{} {
	if len(key) == 0 || data == nil {
		return data
	}

	_dataStr, ok := data.(string)
	if ok {
		if strings.Index(_dataStr, "{") == 0 {
			rs := make(map[string]interface{})
			err := json.Unmarshal([]byte(_dataStr), &rs)
			if err != nil {
				return nil
			}
			return GetItem(rs, key...)
		} else if strings.Index(_dataStr, "[") == 0 {
			var rs []interface{}
			err := json.Unmarshal([]byte(_dataStr), &rs)
			if err != nil {
				return nil
			}
			return GetItem(rs, key...)
		}
	}

	_dataMap, ok := data.(map[string]interface{})
	if ok {
		_key, ok := key[0].(string)
		if !ok {
			zlog.Error("key type is error %T", key[0])
			return nil
		}

		_value, has := _dataMap[_key]
		if has && len(key) > 1 {
			return GetItem(_value, key[1:]...)
		}
		return _value
	}

	_dataArray, ok := data.([]interface{})
	if ok {
		_key, ok := key[0].(int)
		if !ok {
			zlog.Error("key type is error %T", key[0])
			return nil
		}

		if _key >= len(_dataArray) {
			zlog.Error("index out of range:[%d] with length %d", _key, len(_dataArray))
			return nil
		}

		_value := _dataArray[_key]
		if len(key) > 1 {
			return GetItem(_value, key[1:]...)
		}
		return _value
	}
	return nil
}

func GetArray(data interface{}, key ...interface{}) []interface{} {
	item := GetItem(data, key...)
	if item == nil {
		return nil
	}
	i, ok := item.([]interface{})
	if ok {
		return i
	}
	return nil
}
func GetMap(data interface{}, key ...interface{}) map[string]interface{} {
	item := GetItem(data, key...)
	if item == nil {
		return nil
	}
	m, ok := item.(map[string]interface{})
	if ok {
		return m
	}
	return nil
}

func GetString(data interface{}, key ...interface{}) (string, error) {
	item := GetItem(data, key...)
	if item == nil {
		return "", errors.New("not exsits")
	}
	return item.(string), nil
}
func GetNum(data interface{}, key ...interface{}) (float64, error) {
	item := GetItem(data, key...)
	if item == nil {
		return 0, errors.New("not exsits")
	}
	return item.(float64), nil
}
func GetBool(data interface{}, key ...interface{}) (bool, error) {
	item := GetItem(data, key...)
	if item == nil {
		return false, errors.New("not exsits")
	}
	return item.(bool), nil
}
