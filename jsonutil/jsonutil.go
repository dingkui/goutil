package jsonutil

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gitee.com/dk83/goutils/apputil"
	"gitee.com/dk83/goutils/fileutil"
	"gitee.com/dk83/goutils/logutil"
	"io/ioutil"
	"os"
	"path/filepath"
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

func Write_formatjson_file(file string, data map[string]interface{}) {
	b, e := json.Marshal(data)
	if e != nil {
		logutil.ErrorLn(e)
		panic(e)
		return
	}

	os.MkdirAll(filepath.Dir(file), os.ModePerm)
	fileutil.WriteAndSyncFile(file, formatJson(b, true), os.ModePerm)
	//ioutil.WriteFile(file,b,os.ModePerm)
}
func Write_json_file(file string, data map[string]interface{}) {
	b, e := json.Marshal(data)
	if e != nil {
		logutil.ErrorLn(e)
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
	if e == nil {
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

//取得interface中的值，只支持map和数组
func getItem(data interface{}, key interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("data is nil")
	}

	_dataMap, ok := data.(map[string]interface{})
	if ok {
		_key, ok := key.(string)
		if !ok {
			return nil, errors.New(fmt.Sprintf("map key type is error %T", key))
		}
		return _dataMap[_key], nil
	}
	_dataArray, ok := data.([]interface{})
	if ok {
		_key, ok := key.(int)
		if !ok {
			return nil, errors.New(fmt.Sprintf("array key type is error %T", key))
		}
		if _key >= len(_dataArray) {
			return nil, errors.New(fmt.Sprintf("index out of range:[%d] with length %d", _key, len(_dataArray)))
		}
		return _dataArray[_key], nil
	}
	return nil, errors.New(fmt.Sprintf("setItem type is not sopport %T", data))
}

//设置interface中的值，只支持map和数组
func setItem(data interface{}, val interface{}, key interface{}) error {
	if data == nil {
		return errors.New("data is nil")
	}
	_dataMap, ok := data.(map[string]interface{})
	if ok {
		_key, ok := key.(string)
		if !ok {
			return errors.New(fmt.Sprintf("map key type is error %T", key))
		}
		_dataMap[_key] = val
		return nil
	}
	_dataArray, ok := data.([]interface{})
	if ok {
		_key, ok := key.(int)
		if !ok {
			return errors.New(fmt.Sprintf("array key type is error %T", key))
		}
		if _key >= len(_dataArray) {
			return errors.New(fmt.Sprintf("index out of range:[%d] with length %d", _key, len(_dataArray)))
		}
		_dataArray[_key] = val
		return nil
	}
	return errors.New(fmt.Sprintf("setItem type is not sopport %T", data))
}

//设置interface中的值，只支持map和数组
func checkKeys(keys ...interface{}) error {
	for _, key := range keys {
		_, isStr := key.(string)
		_, isInt := key.(int)
		if !isStr && !isInt {
			return errors.New(fmt.Sprintf("key type is not sopport %T", key))
		}
	}
	return nil
}

//设置interface中的值，支持多级设置，支持map,数组和json字符串
func SetItem(data interface{}, val interface{}, keys ...interface{}) interface{} {
	err := checkKeys(keys...)
	if err != nil {
		logutil.ErrorLn(err)
		return nil
	}
	var result interface{}
	_dataStr, isString := data.(string)
	if isString {
		err := json.Unmarshal([]byte(_dataStr), &data)
		if err != nil {
			logutil.ErrorLn(err)
			return nil
		}
	}
	if len(keys) == 1 {
		err := setItem(data, val, keys[0])
		if err != nil {
			logutil.ErrorLn(err)
			return nil
		}
		result = data
	} else if len(keys) > 1 {
		item, err := getItem(data, keys[0])
		if err != nil {
			logutil.ErrorLn(err)
			return nil
		}
		if item == nil {
			_, ok := keys[1].(string)
			if ok {
				item = make(map[string]interface{})
			} else {
				len, ok := keys[1].(int)
				if ok {
					item = make([]interface{}, len+1)
				}
			}
		}
		result = SetItem(item, val, keys[1:]...)
	}
	if isString {
		return Data2Json(result)
	}
	return result
}

//取得interface中的值，支持map，数组和json字符串
func GetItem(data interface{}, keys ...interface{}) interface{} {
	err := checkKeys(keys...)
	if err != nil {
		logutil.ErrorLn(err)
		return nil
	}
	if len(keys) == 0 || data == nil {
		return data
	}
	_dataStr, ok := data.(string)
	if ok {
		err := json.Unmarshal([]byte(_dataStr), &data)
		if err != nil {
			logutil.ErrorLn(err)
			return nil
		}
	}
	item, err := getItem(data, keys[0])
	if err != nil {
		logutil.ErrorLn(err)
		return nil
	}
	return GetItem(item, keys[1:]...)
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
