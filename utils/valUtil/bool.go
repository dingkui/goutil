package valUtil

import (
	"github.com/dingkui/goutil/consts"
	"github.com/dingkui/goutil/errs"
	"strings"
)

func ToBool(data interface{}) (bool, error) {
	if data == nil {
		return consts.EmptyBool, errs.ErrTargetType.New("value is nil")
	}
	switch t := data.(type) {
	case bool:
		return t, nil
	case consts.IfToBool:
		return t.ToBool()
	case string:
		str := strings.ToLower(t)
		return str != "false" && str != "0" && str != "", nil
	case int64:
		return t != 0, nil
	case int:
		return t != 0, nil
	case int8:
		return t != 0, nil
	case int16:
		return t != 0, nil
	case int32:
		return t != 0, nil
	case uint:
		return t != 0, nil
	case uint8:
		return t != 0, nil
	case uint16:
		return t != 0, nil
	case uint32:
		return t != 0, nil
	case uint64:
		return t != 0, nil
	case float64:
		return t != 0, nil
	case float32:
		return t != 0, nil
	}
	return consts.EmptyBool, errs.ErrTargetType.New("value is not Bool")
}
func Bool(val interface{}, def ...bool) (bool, error) {
	re, err := ToBool(val)
	if err != nil {
		if len(def) > 0 {
			return def[0], err
		}
		return consts.EmptyBool, err
	}
	return re, err
}
