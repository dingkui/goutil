package valUtil

import (
	"github.com/dingkui/goutil/dlog"
	"strings"
)

type interfaceBool interface{ ToBool() (bool, error) }

func ToBool(data interface{}) (bool, error) {
	if data == nil {
		return Emputy_bool, errTargetType.New("value is nil")
	}
	switch t := data.(type) {
	case bool:
		return t, nil
	case interfaceBool:
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
	return Emputy_bool, errTargetType.New("value is not Bool")
}
func Bool(val interface{}, def ...bool) (bool, error) {
	re, err := ToBool(val)
	if err != nil {
		if len(def) > 0 {
			return def[0], err
		}
		return Emputy_bool, err
	}
	return re, err
}
func BoolN(val interface{}, def ...bool) bool {
	re, err := Bool(val, def...)
	if err != nil {
		dlog.Warn(err)
	}
	return re
}
