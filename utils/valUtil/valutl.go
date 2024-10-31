package valUtil

import (
	"gitee.com/dk83/goutils/errs"
)

const (
	Emputy_str     string  = ""
	Emputy_int64   int64   = iota
	Emputy_int     int     = iota
	Emputy_bool    bool    = false
	Emputy_float64 float64 = iota
)

var errTargetType = errs.Err(9005, "errTargetType") //目标类型错误
