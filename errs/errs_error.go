package errs

import (
	"fmt"
	"strings"
)

type errInfo struct {
	t *ErrType
	m string  //错误信息
	a string  //错误发生位置
	d interface{}
}
type Error struct {
	trace []*errInfo
}

func (e *Error) Error() string {
	if len(e.trace) == 1 {
		v := e.trace[0]
		return fmt.Sprintf("%s [%d:Err%s] %s", v.a, v.t.code, v.t.msg, v.m)
	}
	stackMsg := make([]string, len(e.trace))
	for i, v := range e.trace {
		stackMsg[i] = v.msg()
	}
	return "trace info:\n" + strings.Join(stackMsg, "\n")
}
func (e *Error) Is(etype *ErrType) (*errInfo, bool) {
	for _, v := range e.trace {
		if v.t == etype{
			return v,true
		}
	}
	return nil, false
}
func (e *Error) Msg() string {
	return e.trace[0].m
}
func (e *Error) Data() interface{} {
	return e.trace[0].d
}
func (e *Error) MsgWithType() string {
	ei := e.trace[0]
	return fmt.Sprintf("[%d:%s]%s", ei.t.code, ei.t.msg, ei.m)
}

func (ei *errInfo) Msg() string {
	return ei.m
}
func (ei *errInfo) Data() interface{} {
	return ei.d
}
func (ei *errInfo) MsgWithType() string {
	return fmt.Sprintf("[%d:%s]%s", ei.t.code, ei.t.msg, ei.m)
}
func (ei *errInfo) msg() string {
	return fmt.Sprintf("  %s [%d:Err%s] %s", ei.a, ei.t.code, ei.t.msg, ei.m)
}
