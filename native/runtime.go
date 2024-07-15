package native

import (
	"fmt"
	"runtime"
	"strings"
)

type r byte

const RuntimeUtil r = iota

func (r) GetCaller(calldepth int) string {
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
