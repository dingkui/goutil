package apputil

import (
	"fmt"
	"gitee.com/dk83/goutils/djson"
	"gitee.com/dk83/goutils/dlog"
	"path/filepath"
)

var (
	Conf             = &djson.JsonFile{}
	RemoteLogHeaders = make(map[string]string)
)

func InitConfig(configFile string) bool {
	if configFile == "" {
		configFile = GetPara("configFile", "config/conf.json")
	}
	//读取配置文件
	err := djson.ReadFile(configFile, Conf)
	if err != nil {
		panic(err)
	}
	//初始化日志
	return InitLog(Conf.JsonGo)
}

func InitLog(conf *djson.JsonGo) bool {
	logRoot := conf.StrN("../logs", "@logger.logRoot")
	logFiles, err := Conf.Array("@logger.appender")
	if err != nil {
		fmt.Sprintf("logConfig is not set: @logger.appender")
		return false
	}
	for _, log := range *logFiles {
		logType := log.StrN("", "logType")
		level := log.IntN(1, "level")
		switch logType {
		case "console":
			dlog.AddAppenderConsole(level)
		case "remote":
			logHost := log.StrN("", "host")
			if logHost != "" {
				addRemote := dlog.AddAppenderRemote(level, logHost, &RemoteLogHeaders)
				if !addRemote {
					fmt.Sprintf("add remote log appender failed: %s", logHost)
					return false
				}
			}
		case "daily":
			name := log.StrN("debug", "name")
			file := log.StrN("server%s."+name+".log", "file")
			addDaily := dlog.AddAppenderDaily(level, filepath.Join(logRoot, file))
			if !addDaily {
				fmt.Sprintf("add addDaily log appender failed: %s", file)
				return false
			}
		}
	}
	return true
}
