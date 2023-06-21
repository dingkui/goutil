package config

import (
	"gitee.com/dk83/goutils/apputil"
	"gitee.com/dk83/goutils/jsonutil"
	"os"
	"strings"
)

var (
	confFile = apputil.GetPara("configFile", "config/conf.json")
	conf     map[string]interface{}
)

func init() {
	conf = jsonutil.Read_json_file(confFile)
	if conf == nil {
		conf = make(map[string]interface{})
	}
	for _, arg := range os.Args {
		index := strings.Index(arg, "=")
		if index > -1 {
			conf[arg[0:index]] = arg[index+1:]
		} else {
			conf[arg] = "1"
		}
	}
}
func SetConf(key string, val interface{}) {
	jsonutil.SetItem(conf, val, getKeys(key)...)
}
func SaveConf() {
	jsonutil.Write_formatjson_file(confFile, conf)
}
func GetConfigMap(key string) map[string]interface{} {
	item_map := jsonutil.GetMap(conf, key)
	if item_map != nil {
		return item_map
	}
	return jsonutil.GetMap(conf, getKeys(key)...)
}
func getKeys(key string) []interface{} {
	split := strings.Split(key, ".")
	var keys []interface{}
	for _, s := range split {
		keys = append(keys, s)
	}
	return keys
}

func GetConfigString(key string) string {
	str, err := jsonutil.GetString(conf, key)
	if err != nil {
		return str
	}
	str, err = jsonutil.GetString(conf, getKeys(key)...)
	return str
}
func GetConfigNum(key string) float64 {
	num, err := jsonutil.GetNum(conf, key)
	if err != nil {
		return num
	}
	num, err = jsonutil.GetNum(conf, getKeys(key)...)
	return num
}
func GetConfigBool(key string) bool {
	num, err := jsonutil.GetBool(conf, key)
	if err != nil {
		return num
	}
	num, err = jsonutil.GetBool(conf, getKeys(key)...)
	return num
}
