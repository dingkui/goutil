package valUtil

import (
	"encoding/json"
	"fmt"
	"gitee.com/dk83/goutils/dlog"
)

type interfaceStr interface{ ToStr() (string, error) }

func ToStr(data interface{}) (string, error) {
	switch t := data.(type) {
	case string:
		return t, nil
	case interfaceStr:
		return t.ToStr()
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float64, float32:
		return fmt.Sprintf("%v", t), nil
	case []byte:
		return string(t), nil
	}

	b, e := json.Marshal(data)
	if e != nil {
		return Emputy_str, e
	}
	return string(b), nil
}
func Str(val interface{}, def ...string) (string, error) {
	re, err := ToStr(val)
	if err != nil {
		if len(def) > 0 {
			return def[0], err
		}
		return Emputy_str, err
	}
	return re, err
}
func StrN(val interface{}, def ...string) string {
	re, err := Str(val, def...)
	if err != nil {
		dlog.Warn(err)
	}
	return re
}
