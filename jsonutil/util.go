package jsonutil

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	regJsonMap, _   = regexp.Compile("^\\{.*\\}$")
	regJsonArray, _ = regexp.Compile("^\\[.*\\]$")
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
	var keys []interface{}
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
func getInt(str string) (int, error) {
	return strconv.Atoi(str)
}
