package stringUtil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dingkui/goutil/consts"
	"strings"
)
func ToStr(data interface{}) (string, error) {
	switch t := data.(type) {
	case string:
		return t, nil
	case consts.IfToStr:
		return t.ToStr()
	case error:
		return t.Error(),nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float64, float32:
		return fmt.Sprintf("%v", t), nil
	case []byte:
		return string(t), nil
	}

	b, e := json.Marshal(data)
	if e != nil {
		return consts.EmptyStr, e
	}
	return string(b), nil
}
func Fmt(s interface{}, v ...interface{}) string {
	if len(v) == 0 {
		str, _ := ToStr(s)
		return str
	}
	message, ok := s.(string)
	if !ok || strings.Index(message, "%") == -1 {
		return fmt.Sprint(append([]interface{}{s}, v...)...)
	}
	return fmt.Sprintf(message, v...)
}
func FormatJson(input []byte, format bool) []byte {
	if !format {
		return input
	}
	var out bytes.Buffer
	json.Indent(&out, input, "", "  ")
	return out.Bytes()
}
func InStringArray(list []string, find string) bool {
	if list == nil || len(list) == 0 {
		return false
	}

	for _, value := range list {
		if value == find {
			return true
		}
	}

	return false
}
