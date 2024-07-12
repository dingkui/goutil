package zlog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
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
func GetCaller(calldepth int) string {
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

func getStack(stackInfo string, start int, calldepth int, isShort bool) string {
	stacks := strings.Split(stackInfo, "\n")
	_len := len(stacks)
	begin := start
	if begin < 5 {
		begin = 5
	}
	if begin >= _len {
		begin = _len - 1
	}
	end := _len
	if calldepth < 0 {
		end = _len + calldepth*2
	} else if calldepth > 0 {
		end = begin + calldepth*2
	}
	if end < begin {
		end = begin + 2
	}
	if end >= _len {
		end = _len
	}
	stacks = stacks[begin:end]
	_len = len(stacks)
	if _len > 30 {
		stacks = stacks[:_len-30]
	}

	if isShort {
		var shortInfo []string
		for i := 0; i < len(stacks); i++ {
			s := stacks[i]
			index := strings.Index(s, ".go:")
			if index == -1 {
				continue
			}
			//index = strings.LastIndex(s, "/")
			//if index > -1 {
			//	stacks[i] = s[index+1:]
			//}
			shortInfo = append(shortInfo, stacks[i])
		}
		stacks = shortInfo
	}
	return strings.Join(stacks, "\n")
}
