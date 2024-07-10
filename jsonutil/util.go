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

func getkeys(key string) []interface{} {
	_keys := strings.Split(key, ".")
	keys := []interface{}{}
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
