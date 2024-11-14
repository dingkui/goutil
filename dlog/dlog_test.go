package dlog_test

import (
	"fmt"
	"github.com/dingkui/goutil/dlog"
	"github.com/dingkui/goutil/utils/runtimeUtil"
	"testing"
	"time"
)

func init() {
	dlog.AddAppenderConsole(0)
	dlog.AddAppenderDaily(1, "./logs/test.%s.log")
	dlog.AddAppenderDaily(2, "./logs/test.err.%s.log")
	//dlog.AddAppenderRemote(2, "https://xxx.chan3d.com/log/upload", nil)
}

func TestColors(t *testing.T) {
	printColor := func (colorCode string, s string,colorName string) {
		fmt.Printf("\033[%sm%s \033[0m %s： \\033[%sm \n", colorCode,  s, colorName, colorCode)
	}

	s := "Hello, World!"
	printColor("30", s,"黑色") //黑色: \033[30m
	printColor("31", s,"红色") //红色: \033[31m
	printColor("32", s,"绿色") //绿色: \033[32m
	printColor("33", s,"黄色") //黄色: \033[33m
	printColor("34", s,"蓝色") //蓝色: \033[34m
	printColor("35", s,"洋红") //洋红色: \033[35m
	printColor("36", s,"青色") //青色: \033[36m
	printColor("37", s,"白色") //白色: \033[37m
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
