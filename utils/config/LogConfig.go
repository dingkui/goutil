package config

import (
	"codeup.aliyun.com/dk/go/goutil/utils/jsonutil"
	"encoding/json"
)

type LogConfig struct {
	level float64
	root  string
}

var LogConf = &LogConfig{}

func InitLog() {
	configMap := GetConfigMap("logger")
	if configMap != nil {
		bytes := jsonutil.Get_bytes(configMap)
		json.Unmarshal(bytes, LogConf)
	}
}
