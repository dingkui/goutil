package djson

import (
	"github.com/dingkui/goutil/errs"
	"strconv"
	"strings"
)

func getkeys(keys []interface{}) ([]interface{}, error) {
	if len(keys) == 1 {
		s, ok := keys[0].(string)
		if ok && strings.Index(s, "@") == 0 {
			keys = _getkeys(s)
		}
	}
	_changeKeys(&keys)

	err := checkKeys(keys...)
	if err != nil {
		return nil, err
	}
	return keys, nil
}
func _getkeys(key string) []interface{} {
	keys := make([]interface{}, 0)
	if key == "" {
		return keys
	}
	_keys := strings.Split(key[1:], ".")
	for _, k := range _keys {
		num, err := strconv.Atoi(k)
		if err != nil {
			keys = append(keys, k)
		} else {
			keys = append(keys, num)
		}
	}
	return keys
}
func _changeKeys(keys *[]interface{}) *[]interface{} {
	for i, k := range *keys {
		if str, ok := k.(string); ok {
			num, err := strconv.Atoi(str)
			if err == nil {
				(*keys)[i] = num
			}
		}
	}
	return keys
}

//设置interface中的值，只支持map和数组
func checkKeys(keys ...interface{}) error {
	for _, key := range keys {
		kStr, isStr := key.(string)
		if isStr && kStr == "" {
			return errs.ErrValidate.New("key is string but empty!")
		}
		kInt, isInt := key.(int)
		if isStr && kInt < -1 {
			return errs.ErrValidate.New("key is int but not effective!:%d", kInt)
		}
		if !isStr && !isInt {
			return errs.ErrValidate.New("key type is not support %T", key)
		}
	}
	return nil
}
