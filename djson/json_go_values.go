package djson

import (
	"github.com/dingkui/goutil/dlog"
	"github.com/dingkui/goutil/utils/valUtil"
)

// Native 获取原生数据
func (j *JsonGo) Native(keys ...interface{}) (interface{}, error) {
	t, err := j.Get(keys...)
	if err != nil {
		return nil, err
	}

	if t._type == jsonMap {
		return t.NativeMap()
	}
	if t._type == jsonArray {
		return t.NativeArray()
	}
	return t.v, nil
}

// NativeN 获取原生数据
func (j *JsonGo) NativeN(keys ...interface{}) interface{} {
	t, err := j.Native(keys...)
	if err != nil {
		dlog.Warn(err)
	}
	return t
}

func (j *JsonGo) NativeArray(keys ...interface{}) ([]interface{}, error) {
	t, err := j.Get(keys...)
	if err != nil {
		return nil, err
	}

	data, err := t.arrayData()
	if err != nil {
		return nil, err
	}

	_re := make([]interface{}, 0)

	for _, value := range *data {
		_value, err := value.Native()
		if err != nil {
			return nil, err
		}
		_re = append(_re, _value)
	}
	return _re, nil
}
func (j *JsonGo) NativeMap(keys ...interface{}) (map[string]interface{}, error) {
	t, err := j.Get(keys...)
	if err != nil {
		return nil, err
	}

	data, err := t.mapData()
	if err != nil {
		return nil, err
	}

	_re := make(map[string]interface{})
	for key, value := range data {
		_value, err := value.Native()
		if err != nil {
			return nil, err
		}
		_re[key] = _value
	}
	return _re, nil
}

func (j *JsonGo) Array(keys ...interface{}) (*[]*JsonGo, error) {
	t, err := j.Get(keys...)
	if err != nil {
		return nil, err
	}

	return t.arrayData()
}
func (j *JsonGo) Map(keys ...interface{}) (map[string]*JsonGo, error) {
	t, err := j.Get(keys...)
	if err != nil {
		return nil, err
	}
	return t.mapData()
}
func (j *JsonGo) Bytes(keys ...interface{}) ([]byte, error) {
	native, err := j.Native(keys...)
	if err != nil {
		return nil, err
	}
	return valUtil.Bytes(native)
}

func (j *JsonGo) Str(def string, keys ...interface{}) (string, error) {
	t, err := j.Get(keys...)
	if err != nil {
		return def, err
	}
	native, err := t.Native()
	if err != nil {
		return def, err
	}
	return valUtil.Str(native, def)
}
func (j *JsonGo) StrN(def string, keys ...interface{}) string {
	t, err := j.Str(def, keys...)
	if err != nil {
		dlog.Warn(err)
	}
	return t
}
func (j *JsonGo) Int64(def int64, keys ...interface{}) (int64, error) {
	t, err := j.Get(keys...)
	if err != nil {
		return def, err
	}
	native, err := t.Native()
	if err != nil {
		return def, err
	}
	return valUtil.Int64(native, def)
}
func (j *JsonGo) Int64N(def int64, keys ...interface{}) int64 {
	t, err := j.Int64(def, keys...)
	if err != nil {
		dlog.Warn(err)
	}
	return t
}
func (j *JsonGo) Int(def int, keys ...interface{}) (int, error) {
	native, err := j.Native(keys...)
	if err != nil {
		return def, err
	}
	return valUtil.Int(native, def)
}
func (j *JsonGo) IntN(def int, keys ...interface{}) int {
	t, err := j.Int(def, keys...)
	if err != nil {
		dlog.Warn(err)
	}
	return t
}
func (j *JsonGo) Float64(def float64, keys ...interface{}) (float64, error) {
	native, err := j.Native(keys...)
	if err != nil {
		return def, err
	}
	return valUtil.Float64(native, def)
}
func (j *JsonGo) Float64N(def float64, keys ...interface{}) float64 {
	t, err := j.Float64(def, keys...)
	if err != nil {
		dlog.Warn(err)
	}
	return t
}
func (j *JsonGo) Bool(def bool, keys ...interface{}) (bool, error) {
	native, err := j.Native(keys...)
	if err != nil {
		return false, err
	}
	return valUtil.Bool(native, def)
}

func (j *JsonGo) BoolN(def bool, keys ...interface{}) bool {
	t, err := j.Bool(def, keys...)
	if err != nil {
		dlog.Warn(err)
	}
	return t
}

func (j *JsonGo) ToStr() (string, error) {
	return j.Str(valUtil.Emputy_str)
}
func (j *JsonGo) ToInt64() (int64, error) {
	return j.Int64(valUtil.Emputy_int64)
}
func (j *JsonGo) ToInt() (int, error) {
	return j.Int(valUtil.Emputy_int)
}
func (j *JsonGo) ToBool() (bool, error) {
	return j.Bool(valUtil.Emputy_bool)
}
func (j *JsonGo) ToBytes() ([]byte, error) {
	return j.Bytes()
}
func (j *JsonGo) ToFloat64() (float64, error) {
	return j.Float64(valUtil.Emputy_float64)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (j *JsonGo) UnmarshalJSON(data []byte) error {
	return j.ReNew(data)
}

// MarshalJSON implements the json.Marshaler interface.
func (j *JsonGo) MarshalJSON() ([]byte, error) {
	return j.ToBytes()
}
