package djson

import (
	"bytes"
	"encoding/json"
)

func formatJson(input []byte, format bool) []byte {
	if !format {
		return input
	}
	var out bytes.Buffer
	json.Indent(&out, input, "", "  ")
	return out.Bytes()
}

func CopyMapVal(target map[string]interface{}, source map[string]interface{}, keys ...string) {
	for _, key := range keys {
		value, has := source[key]
		if has {
			target[key] = value
		}
	}
}

//
//func Str(def string, data interface{}) (string, error) {
//	switch t := data.(type) {
//	case *JsonFile:
//		return t.Str(def)
//	case *JsonGo:
//		return t.Str(def)
//	case string:
//		return t, nil
//	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float64, float32:
//		return fmt.Sprintf("%v", t), nil
//	case []byte:
//		return string(t), nil
//	}
//
//	b, e := json.Marshal(data)
//	if e != nil {
//		return def, e
//	}
//
//	return string(b), nil
//}
//
//func StrN(def string, data interface{}) string {
//	str, err := Str(def, data)
//	if err != nil {
//		dlog.ErrorCaller(err)
//	}
//	return str
//}
//func Byte(data interface{}) ([]byte, error) {
//	if data == nil {
//		return nil, errTargetType.New("value is not Bytes")
//	}
//	switch t := data.(type) {
//	case *JsonGo:
//		return t.Byte()
//	case *JsonFile:
//		return t.Byte()
//	case string:
//		return []byte(t), nil
//	}
//
//	return json.Marshal(data)
//}
//func BoolN(def bool, data interface{}) bool {
//	re, err := Bool(def, data)
//	if err != nil {
//		dlog.ErrorCaller(err)
//	}
//	return re
//}
//func Bool(def bool, data interface{}) (bool, error) {
//	if data == nil {
//		return def, errTargetType.New("value is not Bool")
//	}
//
//	if data == nil {
//		return false, nil
//	}
//	switch t := data.(type) {
//	case bool:
//		return t, nil
//	case string:
//		str := strings.ToLower(t)
//		return str != "false" && str != "0" && str != "", nil
//	case int64:
//		return t != 0, nil
//	case int:
//		return t != 0, nil
//	case int8:
//		return t != 0, nil
//	case int16:
//		return t != 0, nil
//	case int32:
//		return t != 0, nil
//	case uint:
//		return t != 0, nil
//	case uint8:
//		return t != 0, nil
//	case uint16:
//		return t != 0, nil
//	case uint32:
//		return t != 0, nil
//	case uint64:
//		return t != 0, nil
//	case float64:
//		return t != 0, nil
//	case float32:
//		return t != 0, nil
//	case *JsonGo:
//		return t.Bool(def)
//	case *JsonFile:
//		return t.Bool(def)
//	}
//	return def, errTargetType.New("value is not Bool")
//}
//
//func IntN(def int64, data interface{}) int64 {
//	re, err := Int(def, data)
//	if err != nil {
//		dlog.ErrorCaller(err)
//	}
//	return re
//}
//func Int(def int64, data interface{}) (int64, error) {
//	if data == nil {
//		return def, errTargetType.New("value is not Int")
//	}
//	switch t := data.(type) {
//	case int64:
//		return t, nil
//	case int:
//		return int64(t), nil
//	case int8:
//		return int64(t), nil
//	case int16:
//		return int64(t), nil
//	case int32:
//		return int64(t), nil
//	case uint:
//		return int64(t), nil
//	case uint8:
//		return int64(t), nil
//	case uint16:
//		return int64(t), nil
//	case uint32:
//		return int64(t), nil
//	case uint64:
//		return int64(t), nil
//	case float64:
//		v := int64(t)
//		if float64(v) == t {
//			return v, nil
//		}
//		return def, errTargetType.New("float64 to int64,losing precision:%v -> %v", data, v)
//	case float32:
//		v := int64(t)
//		if float32(v) == t {
//			return v, nil
//		}
//		return def, errTargetType.New("float32 to int64,losing precision:%v -> %v", data, v)
//	case string:
//		num, err := strconv.ParseInt(t, 10, 64)
//		if err != nil {
//			return def, err
//		}
//		return num, nil
//	case *JsonGo:
//		return t.Int(def)
//	case *JsonFile:
//		return t.Int(def)
//	}
//	return def, errTargetType.New("value is not Int")
//}
//
//func FloatN(def float64, data interface{}) float64 {
//	re, err := Float(def, data)
//	if err != nil {
//		dlog.ErrorCaller(err)
//	}
//	return re
//}
//func Float(def float64, data interface{}) (float64, error) {
//	if data == nil {
//		return def, errTargetType.New("value is not Float")
//	}
//	switch t := data.(type) {
//	case float64:
//		return t, nil
//	case float32:
//		return float64(t), nil
//	case int:
//		return float64(t), nil
//	case int8:
//		return float64(t), nil
//	case int16:
//		return float64(t), nil
//	case int32:
//		return float64(t), nil
//	case int64:
//		return float64(t), nil
//	case uint:
//		return float64(t), nil
//	case uint8:
//		return float64(t), nil
//	case uint16:
//		return float64(t), nil
//	case uint32:
//		return float64(t), nil
//	case uint64:
//		return float64(t), nil
//	case string:
//		num, err := strconv.ParseFloat(t, 64)
//		if err != nil {
//			return def, err
//		}
//		return num, nil
//	case *JsonGo:
//		return t.Float(def)
//	case *JsonFile:
//		return t.Float(def)
//	}
//
//	return def, errTargetType.New("value is not Float")
//}
