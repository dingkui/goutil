package test

import (
	"bytes"
	"encoding/json"
	"github.com/dingkui/goutil/dlog"
	"io/ioutil"
	"net/http"
)

type CacheUser struct {
	UserId string  `json:"id"`
	Login  string  `json:"login"`
	Token  string  `json:"token"`
	Level  float64 `json:"level"`
	Mac    string  `json:"mac"`
	Info   []byte  `json:"info"`
}

var VERSION = "3.0"
var LogVer = "3.0"
var UserCache = &CacheUser{
	Token:  "",
	UserId: "",
}
var ApiHost = ""

func remoteLogger(level string, msg string, caller string) {
	//把[]byte 转成实现了read接口的Reader结构体
	body := map[string]string{}
	body["msg"] = msg
	body["level"] = level
	body["caller"] = caller
	b, e := json.Marshal(body)
	if e != nil {
		return
	}

	req, err := http.NewRequest("POST", ApiHost+"/s/log", bytes.NewReader(b))
	if err != nil {
		dlog.ERROR.LogLocal("日志网络故障-101:" + err.Error())
		return
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("cid", UserCache.Mac)
	req.Header.Add("ver", LogVer)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		dlog.ERROR.LogLocal("日志网络故障-102：" + err.Error())
		return
	}
	if resp.StatusCode != 200 {
		dlog.WARN.LogLocal("日志网络故障-104：%d", resp.StatusCode)
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		dlog.ERROR.LogLocal("日志网络故障-103：" + err.Error())
		return
	}
}

func getRemoteLogger(level int, loggerLevel int) func(level string, msg string, caller string) {
	if level < loggerLevel {
		return nil
	}
	return remoteLogger
}
