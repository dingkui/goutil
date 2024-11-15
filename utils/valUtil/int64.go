package valUtil

import (
	"github.com/dingkui/goutil/consts"
	"github.com/dingkui/goutil/errs"
	"strconv"
)

func ToInt64(data interface{}) (int64, error) {
	if data == nil {
		return consts.EmptyInt64, errs.ErrTargetType.New("value is nil")
	}
	switch t := data.(type) {
	case int64:
		return t, nil
	case consts.IfToInt64:
		return t.ToInt64()
	case int:
		return int64(t), nil
	case int8:
		return int64(t), nil
	case int16:
		return int64(t), nil
	case int32:
		return int64(t), nil
	case uint:
		return int64(t), nil
	case uint8:
		return int64(t), nil
	case uint16:
		return int64(t), nil
	case uint32:
		return int64(t), nil
	case uint64:
		return int64(t), nil
	case float64:
		v := int64(t)
		if float64(v) == t {
			return v, nil
		}
		return consts.EmptyInt64, errs.ErrTargetType.New("float64 to int64,losing precision:%v -> %v", data, v)
	case float32:
		v := int64(t)
		if float32(v) == t {
			return v, nil
		}
		return consts.EmptyInt64, errs.ErrTargetType.New("float32 to int64,losing precision:%v -> %v", data, v)
	case string:
		num, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			return consts.EmptyInt64, err
		}
		return num, nil
	}
	return consts.EmptyInt64, errs.ErrTargetType.New("value is not Int")
}
func Int64(val interface{}, def ...int64) (int64, error) {
	re, err := ToInt64(val)
	if err != nil {
		if len(def) > 0 {
			return def[0], err
		}
		return consts.EmptyInt64, err
	}
	return re, err
}
