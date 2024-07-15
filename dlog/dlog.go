package dlog

func Debug(v1 interface{}, v ...interface{}) {
	DEBUG.Log(v1, v...)
}
func DebugStack(v1 interface{}, v ...interface{}) {
	DEBUG.Stack(3, true, v1, v...)
}

func Info(v1 interface{}, v ...interface{}) {
	INFO.Log(v1, v...)
}
func InfoStack(v1 interface{}, v ...interface{}) {
	INFO.Stack(3, true, v1, v...)
}

func Warn(v1 interface{}, v ...interface{}) {
	WARN.Log(v1, v...)
}
func WarnCaller(v1 interface{}, v ...interface{}) {
	WARN.LogCaller(1, v1, v...)
}
func WarnStack(v1 interface{}, v ...interface{}) {
	WARN.Stack(3, true, v1, v...)
}

func Error(v1 interface{}, v ...interface{}) {
	ERROR.Log(v1, v...)
}
func ErrorCaller(v1 interface{}, v ...interface{}) {
	ERROR.LogCaller(1, v1, v...)
}
func ErrorStack(v1 interface{}, v ...interface{}) {
	ERROR.Stack(3, true, v1, v...)
}
func ErrorStackTrace(calldepth int, isShort bool, v1 interface{}, v ...interface{}) {
	ERROR.Stack(calldepth, isShort, v1, v...)
}

func Recover() interface{} {
	if r := recover(); r != nil {
		ErrorStackTrace(3, true, r)
		return r
	}
	return nil
}
