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
	fmt.Println(getCaller(0)) // 注意换行符的位置
	fmt.Println(getCaller(1)) // 注意换行符的位置
	fmt.Println(getCaller(2)) // 注意换行符的位置
}
func TestLog(t *testing.T) {
	Debug("113: %s", "Debug") // 注意换行符的位置
	Info("113: %s", "Info")   // 注意换行符的位置
	Warn("113: %s", "Warn")   // 注意换行符的位置
	Error("113: %s", "Error") // 注意换行符的位置
}
func TestTrace(t *testing.T) {
	testTrace() // 注意换行符的位置
}
func testTrace() {
	//ErrorStack("ErrorStack:","short")
	ErrorStackShrot("ErrorStack:", "short")
	ErrorStackTrace(3, true, "ErrorStackTrace:", "short")
	//ErrorStackTrace(20,3,false,"ErrorStackTrace:","full")
	//ErrorStack(0,20,3,false,"113:","xxx")
}
