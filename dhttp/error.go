package dhttp

import (
	"errors"
	"gitee.com/dk83/goutils/native"
)

func Error(code int, msg interface{}, a ...interface{}) *HttpError {
	return &HttpError{
		code: code,
		err:  errors.New(native.StringUtil.Fmt(msg, a...)),
	}
}

type HttpError struct {
	code int
	err  error
}

func (e *HttpError) Error() string {
	return e.err.Error()
}
