package errs

import (
	"fmt"
	"strings"
)

type errInfo struct {
	t *ErrType
	m string
	a string
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
		stackMsg[i] = fmt.Sprintf("  %s [%d:Err%s] %s", v.a, v.t.code, v.t.msg, v.m)
	}
	return "trace info:\n" + strings.Join(stackMsg, "\n")
}
