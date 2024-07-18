package djson

import "gitee.com/dk83/goutils/dlog"

type DefValue struct {
	V interface{}
}

//type IntValue int64
//type StrValue string
//type FloatValue float64
//type BoolValue bool

var (
	N0        = DefValue{int64(0)}
	StrEmpty  = DefValue{""}
	BoolFalse = DefValue{false}
	BoolTrue  = DefValue{true}
)

func (v DefValue) IntN(val interface{}) int64 {
	re, err := Int(v.V.(int64), val)
	if err != nil {
		dlog.ErrorCaller(err)
	}
	return re
}
func (v DefValue) Int(val interface{}) (int64, error) {
	return Int(v.V.(int64), val)
}
func (v DefValue) StrN(val interface{}) string {
	return StrN(v.V.(string), val)
}
func (v DefValue) Str(val interface{}) (string, error) {
	return Str(v.V.(string), val)
}
func (v DefValue) ToStr() string {
	return StrEmpty.StrN(v.V)
}
func (v DefValue) FloatN(val interface{}) float64 {
	return FloatN(v.V.(float64), val)
}
func (v DefValue) Float(val interface{}) (float64, error) {
	return Float(v.V.(float64), val)
}
func (v DefValue) BoolN(val interface{}) bool {
	return BoolN(v.V.(bool), val)
}
func (v DefValue) Bool(val interface{}) (bool, error) {
	return Bool(v.V.(bool), val)
}
