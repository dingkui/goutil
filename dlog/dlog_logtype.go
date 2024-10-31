package dlog

import (
	"fmt"
	"gitee.com/dk83/goutils/utils/dateUtil"
	"gitee.com/dk83/goutils/utils/runtimeUtil"
	"gitee.com/dk83/goutils/utils/stringUtil"
	"runtime/debug"
	"strings"
)

type logType struct {
	level int
	name  string
	sync  bool // 是否同步执行
	addr  string
}

var (
	DEBUG = AddType(0, "DEBUG", true)
	INFO  = AddType(1, "INFO", true)
	WARN  = AddType(2, "WARN", true)
	ERROR = AddType(3, "ERROR", true)
	TEST  = AddType(4, "TEST", true)

	_logTypes = make(map[int]*logType)
)

func AddType(level int, name string, sync bool) *logType {
	_lt, exists := _logTypes[level]
	if exists {
		panic(fmt.Sprintf(
			"Make new logType [%d:%s] has error,the logType [%d:%s] is maked in %s",
			level, name, level, _lt.name, _lt.addr))
	}
	e := &logType{
		level: level,
		name:  name,
		sync:  sync,
		addr:  runtimeUtil.GetCaller(2),
	}
	_logTypes[level] = e
	return e
}
func getType(level int) *logType {
	_lt, exists := _logTypes[level]
	if !exists {
		panic(fmt.Sprintf("logType [%d] not exists!", level))
	}
	return _lt
}

func (t logType) log(depth int, v1 interface{}, v ...interface{}) {
	data := &_logInfo{
		t: dateUtil.DateTimeM.FormatNow(),
		l: t.level,
		f: runtimeUtil.GetCaller(3 + depth),
		m: stringUtil.Fmt(v1, v...),
	}
	if t.sync {
		log(data)
	} else {
		_logQueue <- data
	}
}

func (t logType) Log(v1 interface{}, v ...interface{}) {
	t.log(1, v1, v...)
}
func (t logType) LogCaller(caller int, v1 interface{}, v ...interface{}) {
	t.log(1+caller, v1, v...)
}
func (t logType) LogLocal(v1 interface{}, v ...interface{}) {
	t.log(1, v1, v...)
}
func (t logType) Stack(calldepth int, isShort bool, v1 interface{}, v ...interface{}) {
	stack := getStack(string(debug.Stack()), 7, calldepth, isShort)
	msg := stringUtil.Fmt(v1, v...)
	t.log(1, "%s stacks:\n%s", msg, stack)
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
