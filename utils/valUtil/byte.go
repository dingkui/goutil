package valUtil

import (
	"encoding/json"
	"gitee.com/dk83/goutils/dlog"
)

type interfaceBytes interface{ ToBytes() ([]byte, error) }

func ToBytes(data interface{}) ([]byte, error) {
	if data == nil {
		return nil, errTargetType.New("value is nil")
	}
	switch t := data.(type) {
	case []byte:
		return t, nil
	case interfaceBytes:
		return t.ToBytes()
	case string:
		return []byte(t), nil
	}

	return json.Marshal(data)
}
func Bytes(val interface{}, def ...[]byte) ([]byte, error) {
	re, err := ToBytes(val)
	if err != nil {
		if len(def) > 0 {
			return def[0], err
		}
		return nil, err
	}
	return re, err
}
func BytesN(val interface{}, def ...[]byte) []byte {
	re, err := Bytes(val, def...)
	if err != nil {
		dlog.Warn(err)
	}
	return re
}
