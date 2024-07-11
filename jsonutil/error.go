package jsonutil

import (
	"errors"
	"gitee.com/dk83/goutils/stringutil"
)

type errType byte

const (
	errUnReached  errType = iota //不应该到达的错误
	errValid                     //校验错误
	errJsonType                  //json数据类型错误
	errTarget                    //目标错误
	errTargetType                //目标类型错误
)

func (e errType) New(msg interface{}, a ...interface{}) error {
	return &JsonError{
		errT: e,
		err:  errors.New(stringutil.Fmt(msg, a...)),
	}
}

func (e errType) Is(err error) bool {
	jErr, ok := err.(*JsonError)
	return ok && jErr.errT == e
}

type JsonError struct {
	errT errType
	err  error
}

func (e *JsonError) Error() string {
	return e.err.Error()
}
