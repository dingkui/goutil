package zlog

import (
	"fmt"
	"gitee.com/dk83/goutils/stringutil"
	"log"
	"os"
	"runtime/debug"
)

type logType struct {
	level int
	name  string
}

var (
	TEST  = &logType{level: 0, name: "test"}
	DEBUG = &logType{level: 0, name: "debug"}
	INFO  = &logType{level: 1, name: "info"}
	WARN  = &logType{level: 2, name: "warn"}
	ERROR = &logType{level: 3, name: "error"}

	basecalldepth = 3
)

var _stdoutLogger *log.Logger

func (t *logType) log(localDepth int, remoteDepth int, v1 interface{}, v ...interface{}) {
	var local *log.Logger = nil
	var remote func(level string, msg string, caller string) = nil
	if localDepth > -1 {
		//localDepth>=0 表示需要本地日志
		if _logGetter != nil {
			local = _logGetter.getLocalLogger(t.level)
		}
		if local == nil {
			if _stdoutLogger == nil {
				_stdoutLogger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
				_stdoutLogger.Output(1, "---- log is not init,use default stdout -----")
			}
			local = _stdoutLogger
		}
	}
	if remoteDepth > -1 && _logGetter != nil {
		remote = _logGetter.getRemoteLogger(t.level)
	}
	if local == nil && remote == nil {
		return
	}
	msg := stringutil.Fmt(v1, v...)
	if local != nil {
		local.Output(basecalldepth+localDepth, fmt.Sprintf("%s %s", t.name, msg))
	}
	if remote != nil {
		caller := GetCaller(basecalldepth + remoteDepth)
		go remote(t.name, msg, caller)
	}
}

func (t *logType) Log(v1 interface{}, v ...interface{}) {
	t.log(1, 1, v1, v...)
}
func (t *logType) LogCaller(caller int, v1 interface{}, v ...interface{}) {
	t.log(1+caller, 1+caller, v1, v...)
}
func (t *logType) LogLocal(v1 interface{}, v ...interface{}) {
	t.log(1, -1, v1, v...)
}
func (t *logType) LogRemote(v1 interface{}, v ...interface{}) {
	t.log(-1, 1, v1, v...)
}
func (t *logType) Stack(calldepth int, isShort bool, v1 interface{}, v ...interface{}) {
	stack := getStack(string(debug.Stack()), 7, calldepth, isShort)
	msg := stringutil.Fmt(v1, v...)
	t.log(1, 1, "%s stacks:\n%s", msg, stack)
}
