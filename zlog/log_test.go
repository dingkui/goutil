package zlog

import (
	"fmt"
	"testing"
)

var (
	xx = InitLog(&defaultLogGetter{
		local: &localFileLog{
			root: "../logs",
		},
	})
)

func TestGetCaller(t *testing.T) {
	fmt.Println(getCaller(0))
	fmt.Println(getCaller(1))
	fmt.Println(getCaller(2))
}
func TestLog(t *testing.T) {
	Debug("113: %s", "Debug")
	Info("113: %s", "Info")
	Warn("113: %s", "Warn")
	Error("113: %s", "Error")
}
func TestTrace(t *testing.T) {
	testTrace()
}
func testTrace() {
	//ErrorStack("ErrorStack:","short")
	ErrorStackShrot("ErrorStack:", "short")
	ErrorStackTrace(3, true, "ErrorStackTrace:", "short")
	//ErrorStackTrace(20,3,false,"ErrorStackTrace:","full")
	//ErrorStack(0,20,3,false,"113:","xxx")
}
