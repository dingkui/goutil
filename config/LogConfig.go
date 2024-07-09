package config

import (
	"encoding/json"
	"fmt"
	"gitee.com/dk83/goutils/jsonutil"
	"gitee.com/dk83/goutils/zlog"
)

func InitLog() {
	configMap := GetConfigMap("Logger")
	_type, _ := jsonutil.GetString(configMap, "Type")
	if _type == "localFileLog" {
		var logConf = zlog.NewDefaultLogGette(0, 0, "../logs")
		bytes := jsonutil.Get_bytes(configMap)
		err := json.Unmarshal(bytes, logConf)
		if err != nil {
			fmt.Print(err)
		}
		zlog.InitLog(logConf)
	}
}
