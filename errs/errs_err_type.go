package errs

import (
	"fmt"
	"gitee.com/dk83/goutils/native"
)

type ErrType struct {
	code int
	msg  string
	addr string
}

func (e *ErrType) Trace(err error, msg interface{}, a ...interface{}) error {
	trace := []*errInfo{{t: e,
		m: native.StringUtil.Fmt(msg, a...),
		a: native.RuntimeUtil.GetCaller(2),
	}}
	_e, ok := err.(*Error)
	if ok {
		trace = append(trace, _e.trace...)
	}
	return &Error{trace: trace}
}
func (e *ErrType) New(msg interface{}, a ...interface{}) error {
	_e, ok := msg.(*Error)
	if ok {
		_msg := e.msg
		if len(_e.trace) > 0 {
			_msg = native.StringUtil.Fmt(a[0], a[1:]...)
		}

		trace := []*errInfo{{t: e,
			m: _msg,
			a: native.RuntimeUtil.GetCaller(2),
		}}
		return &Error{trace: append(trace, _e.trace...)}
	} else {
		return &Error{trace: []*errInfo{{t: e,
			m: native.StringUtil.Fmt(msg, a...),
			a: native.RuntimeUtil.GetCaller(2),
		}}}
	}
}
func (e *ErrType) Is(err error) bool {
	_e, ok := err.(*Error)
	if !ok {
		return false
	}
	for _, i2 := range _e.trace {
		if i2.t == e {
			return true
		}
	}
	return false
}
func (e *ErrType) Msg(err error) string {
	_e, ok := err.(*Error)
	if !ok {
		return ""
	}
	for _, i2 := range _e.trace {
		if i2.t == e {
			return fmt.Sprintf("[%d:%s]:%s", e.code, e.msg, i2.m)
		}
	}
	return ""
}
func (e *ErrType) Error() string {
	return fmt.Sprintf("%d:%s", e.code, e.msg)
}
