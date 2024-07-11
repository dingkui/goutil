package zlog

import (
	"gitee.com/dk83/goutils/stringutil"
	"runtime/debug"
	"strings"
)

func Debug(v1 interface{}, v ...interface{}) {
	DEBUG.Log(v1, v...)
}
func Info(v1 interface{}, v ...interface{}) {
	INFO.Log(v1, v...)
}
func Warn(v1 interface{}, v ...interface{}) {
	WARN.Log(v1, v...)
}
func Error(v1 interface{}, v ...interface{}) {
	ERROR.Log(v1, v...)
}
func WarnCaller(v1 interface{}, v ...interface{}) {
	WARN.LogCaller(1, v1, v...)
}
func ErrorCaller(v1 interface{}, v ...interface{}) {
	ERROR.LogCaller(1, v1, v...)
}
func ErrorStack(v1 interface{}, v ...interface{}) {
	stack := getStack(string(debug.Stack()), 5, 5, true)
	errorStack(stack, v1, v...)
}
func ErrorStackShrot(v1 interface{}, v ...interface{}) {
	stack := getStack(string(debug.Stack()), 5, 3, true)
	errorStack(stack, v1, v...)
}
func ErrorStackTrace(calldepth int, isShort bool, v1 interface{}, v ...interface{}) {
	stack := getStack(string(debug.Stack()), 5, calldepth, isShort)
	errorStack(stack, v1, v...)
}
func errorStack(stack string, v1 interface{}, v ...interface{}) {
	msg := stringutil.Fmt(v1, v...)
	ERROR.LogCaller(1, "%s stacks:\n%s", msg, stack)
}
func getStack(stackInfo string, start int, calldepth int, isShort bool) string {
	stacks := strings.Split(stackInfo, "\n")
	_len := len(stacks)
	if _len > start {
		begin := start
		if begin < 5 {
			begin = 5
		}
		if calldepth < 1 {
			calldepth = 1
		}
		end := begin + calldepth*2
		if end < begin {
			end = _len - 1
		}
		stacks = stacks[begin:end]
	}
	_len = len(stacks)
	if _len > 30 {
		stacks = stacks[:_len-30]
	}

	if isShort {
		var shortInfo []string
		for i := 0; i < len(stacks); i++ {
			s := stacks[i]
			index := strings.LastIndex(s, "/")
			if index > -1 {
				stacks[i] = s[index+1:]
			}
			index = strings.Index(s, ".go")
			simpleLen := len(shortInfo)
			if index > -1 && simpleLen > 0 {
				shortInfo[simpleLen-1] = "  -> " + shortInfo[simpleLen-1] + " " + stacks[i]
			} else {
				shortInfo = append(shortInfo, stacks[i])
			}
		}
		stacks = shortInfo
	}
	return strings.Join(stacks, "\n")
}
