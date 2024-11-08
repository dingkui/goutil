package apputil

import (
	"os"
	"strconv"
	"strings"
)

var paras = map[string]string{}

func init() {
	for _, arg := range os.Args {
		index := strings.Index(arg, "=")
		if index > -1 {
			paras[arg[0:index]] = arg[index+1:]
		} else {
			paras[arg] = "1"
		}
	}
}

func IsPara(name string) bool {
	v, has := paras[name]
	if !has {
		return false
	}
	return v == "1" || v == "true"
}
func GetPara(name string, def string) string {
	v, has := paras[name]
	if has {
		return v
	}
	return def
}
func GetParaInt(name string, def int) int {
	v, has := paras[name]
	if has {
		num, err := strconv.Atoi(v)
		if err != nil {
			return def
		}
		return num
	}
	return def
}
func GetParaInt64(name string, def int64) int64 {
	v, has := paras[name]
	if has {
		num, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return def
		}
		return num
	}
	return def
}

func GetParaf(name string, def func() string) string {
	v, has := paras[name]
	if has {
		return v
	}
	return def()
}
