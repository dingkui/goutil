package httputil

import (
	"errors"
	"gitee.com/dk83/goutils/stringutil"
)

func Error(code int, a ...interface{}) *HttpError {
	return &HttpError{
		code: code,
		err:  errors.New(stringutil.Fmt(a...)),
	}
}

type HttpError struct {
	code int
	err  error
}

func (e *HttpError) Error() string {
	return e.err.Error()
}
