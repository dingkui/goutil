package jsonutil

import (
	"gitee.com/dk83/goutils/errs"
	"regexp"
)

type jsonType byte

const (
	jsonUnkonw jsonType = iota
	jsonMap
	jsonArray
	jsonString
	jsonInt
	jsonFloat
	jsonBool
)

var (
	regJsonMap, _   = regexp.Compile("^\\{.*\\}$")
	regJsonArray, _ = regexp.Compile("^\\[.*\\]$")

	errNewJsonGo  = errs.Err(9000, "errNewJsonGo")  //json 构建时错误
	errJsonType   = errs.Err(9001, "errJsonType")   //json数据类型错误
	errTarget     = errs.Err(9002, "errTarget")     //目标错误
	errKey        = errs.Err(9003, "errKey")        //目标错误
	errTargetType = errs.Err(9004, "errTargetType") //目标类型错误
)
