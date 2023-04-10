package config

import (
	"encoding/json"
	"gitee.com/dk83_admin/goutil/utils/jsonutil"
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
