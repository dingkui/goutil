package forceVal

import (
	"github.com/dingkui/goutil/dlog"
	"github.com/dingkui/goutil/utils/valUtil"
)

func Bool(val interface{}, def ...bool) bool {
	re, err := valUtil.Bool(val, def...)
	if err != nil {
		dlog.Warn(err)
	}
	return re
}
func Bytes(val interface{}, def ...[]byte) []byte {
	re, err := valUtil.Bytes(val, def...)
	if err != nil {
		dlog.Warn(err)
	}
	return re
}
func Float64(val interface{}, def ...float64) float64 {
	re, err := valUtil.Float64(val, def...)
	if err != nil {
		dlog.Warn(err)
	}
	return re
}
func Int(val interface{}, def ...int) int {
	re, err := valUtil.Int(val, def...)
	if err != nil {
		dlog.Warn(err)
	}
	return re
}
func Int64(val interface{}, def ...int64) int64 {
	re, err := valUtil.Int64(val, def...)
	if err != nil {
		dlog.Warn(err)
	}
	return re
}
func Str(val interface{}, def ...string) string {
	re, err := valUtil.Str(val, def...)
	if err != nil {
		dlog.Warn(err)
	}
	return re
}
