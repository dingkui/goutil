package errs

import (
	"fmt"
	"gitee.com/dk83/goutils/native"
)

func Err(code int, msg string) *ErrType {
	err, exists := _errs[code]
	if exists {
		panic(fmt.Sprintf(
			"Make new ErrType [%d:%s] has error,the ErrType [%d:%s] is maked in %s",
			code, msg, code, err.msg, err.addr))
	}
	e := &ErrType{
		code: code,
		msg:  msg,
		addr: native.RuntimeUtil.GetCaller(2),
	}
	_errs[code] = e
	return e
}
