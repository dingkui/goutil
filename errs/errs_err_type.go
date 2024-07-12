package errs

import (
	"fmt"
	"gitee.com/dk83/goutils/stringutil"
	"gitee.com/dk83/goutils/zlog"
)

type ErrType struct {
	code int
	msg  string
	addr string
}

func (e *ErrType) Trace(err error, msg interface{}, a ...interface{}) error {
	trace := []*errInfo{{t: e,
		m: stringutil.Fmt(msg, a...),
		a: zlog.GetCaller(2),
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
			_msg = stringutil.Fmt(a[0], a[1:]...)
		}

		trace := []*errInfo{{t: e,
			m: _msg,
			a: zlog.GetCaller(2),
		}}
		return &Error{trace: append(trace, _e.trace...)}
	} else {
		return &Error{trace: []*errInfo{{t: e,
			m: stringutil.Fmt(msg, a...),
			a: zlog.GetCaller(2),
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
