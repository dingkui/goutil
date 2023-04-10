package config

import (
	"codeup.aliyun.com/dk/go/goutil/utils/jsonutil"
	"os"
	"strings"
)

var (
	confFile string = "config/conf.json"
	conf     map[string]interface{}
)

func init() {
	conf = jsonutil.Read_json_file(confFile)
	for _, arg := range os.Args {
		index := strings.Index(arg, "=")
		if index > -1 {
			conf[arg[0:index]] = arg[index+1:]
		} else {
			conf[arg] = "1"
		}
	}
}
func GetConfigMap(key string) map[string]interface{} {
	item_map := jsonutil.GetMap(conf, key)
	if item_map != nil {
		return item_map
	}
	return jsonutil.GetMap(conf, strings.Split(key, "."))
}

func GetConfigString(key string) string {
	str, err := jsonutil.GetString(conf, key)
	if err != nil {
		return str
	}
	str, err = jsonutil.GetString(conf, strings.Split(key, "."))
	return str
}
func GetConfigNum(key string) float64 {
	num, err := jsonutil.GetNum(conf, key)
	if err != nil {
		return num
	}
	num, err = jsonutil.GetNum(conf, strings.Split(key, "."))
	return num
}
func GetConfigBool(key string) bool {
	num, err := jsonutil.GetBool(conf, key)
	if err != nil {
		return num
	}
	num, err = jsonutil.GetBool(conf, strings.Split(key, "."))
	return num
}
