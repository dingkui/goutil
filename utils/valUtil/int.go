package valUtil

import (
	"github.com/dingkui/goutil/consts"
	"github.com/dingkui/goutil/errs"
	"strconv"
)

func ToInt(data interface{}) (int, error) {
	if data == nil {
		return consts.EmptyInt, errs.ErrTargetType.New("value is nil")
	}
	switch t := data.(type) {
	case int:
		return t, nil
	case consts.IfToInt:
		return t.ToInt()
	case int64:
		return int(t), nil
	case int8:
		return int(t), nil
	case int16:
		return int(t), nil
	case int32:
		return int(t), nil
	case uint64:
		return int(t), nil
	case uint8:
		return int(t), nil
	case uint16:
		return int(t), nil
	case uint32:
		return int(t), nil
	case uint:
		return int(t), nil
	case float64:
		v := int(t)
		if float64(v) == t {
			return v, nil
		}
		return consts.EmptyInt, errs.ErrTargetType.New("float64 to int,losing precision:%v -> %v", data, v)
	case float32:
		v := int(t)
		if float32(v) == t {
			return v, nil
		}
		return consts.EmptyInt, errs.ErrTargetType.New("float32 to int,losing precision:%v -> %v", data, v)
	case string:
		num, err := strconv.Atoi(t)
		if err != nil {
			return consts.EmptyInt, errs.ErrTargetType.New("string to int err.",err)
		}
		return num, nil
	}
	return consts.EmptyInt, errs.ErrTargetType.New("value is not Int")
}
func Int(val interface{}, def ...int) (int, error) {
	re, err := ToInt(val)
	if err != nil {
		if len(def) > 0 {
			return def[0], err
		}
		return consts.EmptyInt, err
	}
	return re, err
}
