package valUtil

import (
	"gitee.com/dk83/goutils/dlog"
	"strconv"
)

type interfaceFloat64 interface{ ToFloat64() (float64, error) }

func ToFloat64(data interface{}) (float64, error) {
	if data == nil {
		return Emputy_float64, errTargetType.New("value is nil")
	}
	switch t := data.(type) {
	case float64:
		return t, nil
	case interfaceFloat64:
		return t.ToFloat64()
	case float32:
		return float64(t), nil
	case int:
		return float64(t), nil
	case int8:
		return float64(t), nil
	case int16:
		return float64(t), nil
	case int32:
		return float64(t), nil
	case int64:
		return float64(t), nil
	case uint:
		return float64(t), nil
	case uint8:
		return float64(t), nil
	case uint16:
		return float64(t), nil
	case uint32:
		return float64(t), nil
	case uint64:
		return float64(t), nil
	case string:
		num, err := strconv.ParseFloat(t, 64)
		if err != nil {
			return Emputy_float64, err
		}
		return num, nil
	}

	return Emputy_float64, errTargetType.New("value is not Float")
}
func Float64(val interface{}, def ...float64) (float64, error) {
	re, err := ToFloat64(val)
	if err != nil {
		if len(def) > 0 {
			return def[0], err
		}
		return Emputy_float64, err
	}
	return re, err
}
func Float64N(val interface{}, def ...float64) float64 {
	re, err := Float64(val, def...)
	if err != nil {
		dlog.ErrorCaller(err)
	}
	return re
}
