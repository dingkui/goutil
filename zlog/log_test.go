package zlog_test

import (
	"fmt"
	"gitee.com/dk83/goutils/zlog"
	"testing"
)

func init() {
	zlog.InitLog(zlog.NewDefaultLogGette(0, 0, "../logs", ""))
}

func TestGetCaller(t *testing.T) {
	fmt.Println(zlog.GetCaller(0))
	fmt.Println(zlog.GetCaller(1))
	fmt.Println(zlog.GetCaller(2))
}
func TestLog(t *testing.T) {
	zlog.Debug("113: %s", "Debug")
	zlog.Info("113: %s", "Info")
	zlog.Warn("113: %s", "Warn")
	zlog.Error("113: %s", "Error")
}
func TestTrace(t *testing.T) {
	testTrace()
}
func testTrace1() {
	zlog.ErrorStack("ErrorStack:", "short")
	zlog.ErrorStackTrace(3, true, "ErrorStackTrace:", "short")
}
func testTrace() {
	testTrace1()
}
