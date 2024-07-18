package djson

import "gitee.com/dk83/goutils/dlog"

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
func (j *JsonGo) Byte(keys ...interface{}) ([]byte, error) {
	native, err := j.Native(keys...)
	if err != nil {
		return nil, err
	}
	return Byte(native)
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
	return Str(def, native)
}
func (j *JsonGo) StrN(def string, keys ...interface{}) string {
	t, err := j.Str(def, keys...)
	if err != nil {
		dlog.Warn(err)
	}
	return t
}
func (j *JsonGo) Int(def int64, keys ...interface{}) (int64, error) {
	t, err := j.Get(keys...)
	if err != nil {
		return def, err
	}
	native, err := t.Native()
	if err != nil {
		return def, err
	}
	return Int(def, native)
}
func (j *JsonGo) IntN(def int64, keys ...interface{}) int64 {
	t, err := j.Int(def, keys...)
	if err != nil {
		dlog.Warn(err)
	}
	return t
}
func (j *JsonGo) Float(def float64, keys ...interface{}) (float64, error) {
	native, err := j.Native(keys...)
	if err != nil {
		return def, err
	}
	return Float(def, native)
}
func (j *JsonGo) FloatN(def float64, keys ...interface{}) float64 {
	t, err := j.Float(def, keys...)
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
	return Bool(def, native)
}

func (j *JsonGo) BoolN(def bool, keys ...interface{}) bool {
	t, err := j.Bool(def, keys...)
	if err != nil {
		dlog.Warn(err)
	}
	return t
}
