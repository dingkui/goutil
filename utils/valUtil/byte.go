package valUtil

import (
	"encoding/json"
	"github.com/dingkui/goutil/consts"
	"github.com/dingkui/goutil/errs"
)

func ToBytes(data interface{}) ([]byte, error) {
	if data == nil {
		return nil, errs.ErrTargetType.New("value is nil")
	}
	switch t := data.(type) {
	case []byte:
		return t, nil
	case consts.IfToBytes:
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
