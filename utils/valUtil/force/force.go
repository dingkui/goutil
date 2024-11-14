package force

import (
	"github.com/dingkui/goutil/utils/valUtil"
)

func Bool(val interface{}, def ...bool) bool {
	re, _ := valUtil.Bool(val, def...)
	return re
}
func Bytes(val interface{}, def ...[]byte) []byte {
	re, _ := valUtil.Bytes(val, def...)
	return re
}
func Float64(val interface{}, def ...float64) float64 {
	re, _ := valUtil.Float64(val, def...)
	return re
}
func Int(val interface{}, def ...int) int {
	re, _ := valUtil.Int(val, def...)
	return re
}
func Int64(val interface{}, def ...int64) int64 {
	re, _ := valUtil.Int64(val, def...)
	return re
}
func Str(val interface{}, def ...string) string {
	re, _ := valUtil.Str(val, def...)
	return re
}
