package zlog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

var RootDir = getCurrentDirectory()
var _logs = filepath.Join(RootDir, "logs")

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err)
	}
	join := filepath.Join(dir, "..")
	return strings.Replace(join, "\\", "/", -1)
}
func dayStr() string {
	return time.Now().Format("20060102")
}
func getCaller(calldepth int) string {
	_, file, line, ok := runtime.Caller(calldepth) // 回溯两层，拿到写日志的调用方的函数信息
	if !ok {
		return ""
	}
	index := strings.LastIndex(file, "/")
	if index > -1 {
		return fmt.Sprintf("%s:%d", file[index+1:], line)
	}
	return fmt.Sprintf("%s:%d", file[26:], line)
}
func funcThatPanics() {
	log.Printf(string(debug.Stack()))
}
