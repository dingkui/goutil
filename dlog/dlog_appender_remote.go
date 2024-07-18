package dlog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	doRemoteErr = AddType(9, "RemoteError", false)
)

type remoteAppender struct {
	level  int
	url    string
	header *map[string]string
}

func (f *remoteAppender) Enable(level int) bool {
	return f.level <= level && level != 9
}
func (f *remoteAppender) WriteLog(s string, logName string) (int, error) {
	go f.postLog(logName, s)
	return 0, nil
}
func (f *remoteAppender) Close() {
}
func errRemote(v1 interface{}, v ...interface{}) {
	doRemoteErr.Stack(3, true, v1, v...)
}
func (f *remoteAppender) postLog(level string, msg string) bool {
	if f.url == "" {
		return false
	}

	defer func() {
		if r := recover(); r != nil {
			errRemote(r)
			return
		}
	}()
	//把[]byte 转成实现了read接口的Reader结构体
	body := map[string]string{}
	body["msg"] = msg
	body["level"] = level
	b, e := json.Marshal(body)
	if e != nil {
		return false
	}
	req, err := http.NewRequest("POST", f.url, bytes.NewReader(b))
	if err != nil {
		errRemote("日志网络故障-101:%v", err)
		return false
	}

	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	if f.header != nil {
		for k, v := range *f.header {
			req.Header.Add(k, v)
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if resp == nil {
		errRemote("日志网络故障-102:%v", err)
		return false
	}
	defer resp.Body.Close()
	if err != nil {
		errRemote("日志网络故障-102:%v", err)
		return false
	}
	if resp.StatusCode != 200 {
		errRemote("日志网络故障-104:%v", resp)
		return false
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		errRemote(fmt.Sprintf("日志网络故障-103：%d,%v", resp.StatusCode, err))
		return false
	}
	return true
}

func AddAppenderRemote(level int, url string, header *map[string]string) bool {
	if url == "" {
		return false
	}
	remoteAppender := &remoteAppender{level, url, header}

	re := remoteAppender.postLog("debug", "init remote appender sucess!")
	if !re {
		return false
	}

	AddLogger(remoteAppender)
	return true
}
