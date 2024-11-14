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
func Para(name string, def ...string) string {
	v, found := paras[name]
	if !found && len(def) > 0 {
		return def[0]
	}
	return v
}
func ParaF(name string, def func() string) string {
	v, found := paras[name]
	if found {
		return v
	}
	return def()
}
func ParaIs(name string) bool {
	v, found := paras[name]
	if !found {
		return false
	}
	return v == "1" || v == "true"
}
func ParaInt(name string, def int) int {
	num, err := strconv.Atoi(Para(name))
	if err != nil {
		return def
	}
	return num
}
func ParaInt64(name string, def int64) int64 {
	num, err := strconv.ParseInt(Para(name), 10, 64)
	if err != nil {
		return def
	}
	return num
}
