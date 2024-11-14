package valUtil

import (
	"github.com/dingkui/goutil/consts"
	"github.com/dingkui/goutil/utils/stringUtil"
)

func Str(val interface{}, def ...string) (string, error) {
	re, err := stringUtil.ToStr(val)
	if err != nil {
		if len(def) > 0 {
			return def[0], err
		}
		return consts.EmptyStr, err
	}
	return re, err
}
