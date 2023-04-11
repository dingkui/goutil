package logutil

import (
	"fmt"
	"strings"
)

type LogInterface interface {
	Log(level int, depthPre int, msg string)
}

var (
	Logger LogInterface
)

func Debug(message string, v ...interface{}) {
	Logger.Log(0, 0, _msg(message, v...))
}
func DebugLn(v ...interface{}) {
	Logger.Log(0, 0, fmt.Sprintln(v...))
}
func Info(message string, v ...interface{}) {
	Logger.Log(1, 0, _msg(message, v...))
}
func InfoLn(v ...interface{}) {
	Logger.Log(1, 0, fmt.Sprintln(v...))
}
func Warn(message string, v ...interface{}) {
	Logger.Log(2, 0, _msg(message, v...))
}
func WarnLn(v ...interface{}) {
	Logger.Log(2, 0, fmt.Sprintln(v...))
}
func Error(message string, v ...interface{}) {
	Logger.Log(3, 0, _msg(message, v...))
}
func ErrorLn(v ...interface{}) {
	Logger.Log(3, 0, fmt.Sprintln(v...))
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
