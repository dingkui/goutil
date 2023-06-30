package logutil

import (
	"fmt"
	"gitee.com/dk83/goutils/stringutil"
	"strings"
)

type LogInterface interface {
	Log(level int, depthPre int, msg string)
	LogAble(level int) bool
	Init()
}

var (
	_log LogInterface
)

func Debug(v ...interface{}) {
	Log(0, 1, v...)
}
func Info(v ...interface{}) {
	Log(1, 1, v...)
}
func Warn(v ...interface{}) {
	Log(2, 1, v...)
}
func Error(v ...interface{}) {
	Log(3, 1, v...)
}
func Log(level int, depthPre int, v ...interface{}) {
	if !_log.LogAble(level) {
		return
	}
	_log.Log(level, depthPre, stringutil.Fmt(v...))
}
func _msg(message string, v ...interface{}) string {
	msg := message
	if len(v) > 0 {

		if strings.Index(message, "%") == -1 {
			msgs := make([]interface{}, 1)
			msgs[0] = message
			for _, i2 := range v {
				msgs = append(msgs, i2)
			}
			msg = fmt.Sprintln(msgs...)
		} else {
			msg = fmt.Sprintf(message, v...)
		}
	}
	return msg
}

func InitLog(logger LogInterface) {
	_log = logger
	_log.Init()
}
