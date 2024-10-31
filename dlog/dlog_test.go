package dlog_test

import (
	"fmt"
	"gitee.com/dk83/goutils/dlog"
	"gitee.com/dk83/goutils/utils/runtimeUtil"
	"testing"
	"time"
)

func init() {
	dlog.AddAppenderConsole(0)
	dlog.AddAppenderDaily(1, "./logs/test.%s.log")
	dlog.AddAppenderDaily(2, "./logs/test.err.%s.log")
	dlog.AddAppenderRemote(2, "https://xxx.chan3d.com/log/upload", nil)
}

func TestGetCaller(t *testing.T) {
	fmt.Println(runtimeUtil.GetCaller(0))
	fmt.Println(runtimeUtil.GetCaller(1))
	fmt.Println(runtimeUtil.GetCaller(2))
}
func TestLog(t *testing.T) {
	//dlog.Debug("113: %s", "Debug")
	//dlog.Info("113: %s", "Info")
	//dlog.Warn("113: %s", "Warn")
	dlog.Error("113: %s", "Error1111111")
	time.Sleep(time.Second * 10)
}
func TestTrace(t *testing.T) {
	testTrace()
}
func TestClearAppenders(t *testing.T) {
	dlog.ClearAppenders()
}
func testTrace1() {
	dlog.ErrorStack("ErrorStack:", "short")
	dlog.ErrorStackTrace(3, true, "ErrorStackTrace:", "short")
}
func testTrace() {
	testTrace1()
}
