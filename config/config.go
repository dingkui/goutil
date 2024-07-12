package config

import (
	"gitee.com/dk83/goutils/apputil"
	"gitee.com/dk83/goutils/jsonutil"
	"gitee.com/dk83/goutils/zlog"
)

var (
	Conf = &jsonutil.JsonFile{}
)

func InitConfig(configFile string) {
	if configFile == "" {
		configFile = apputil.GetPara("configFile", "config/conf.json")
	}
	//读取配置文件
	err := jsonutil.ReadFile(configFile, Conf)
	if err != nil {
		panic(err)
	}

	//初始化日志
	zlog.InitLog(zlog.NewDefaultLogGette(
		Conf.IntN(0, "@logger.localLevel"),
		Conf.IntN(0, "@logger.remoteLevel"),
		Conf.StrN("../logs", "@logger.logRoot"),
		Conf.StrN("server%s.%s.log", "@logger.logName")))
}
