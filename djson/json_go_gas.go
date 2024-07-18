// Package jsonutil : get and set
package djson

import (
	"fmt"
	"gitee.com/dk83/goutils/errs"
)

func (j *JsonGo) getItem(key interface{}) (*JsonGo, error) {
	err := checkKeys(key)
	if err != nil {
		return nil, err
	}
	switch key.(type) {
	case string:
		data, err := j.mapData()
		if err != nil {
			return nil, err
		}
		item, ok := data[key.(string)]
		if !ok {
			return nil, errKey.New("get [%s] from:%#q", key, data)
		}
		return item, nil
	case int:
		data, err := j.arrayData()
		if err != nil {
			return nil, err
		}
		i := key.(int)
		l := len(*data)
		if i < 0 || i >= l {
			//超出下标返回最后一个元素
			return nil, errKey.New("index is out of range :%d ,from %#q", i, data)
		}
		return (*data)[i], nil
	}
	return nil, errs.ErrSystem.New("getItem fail:[%T] %v %v", j, key, key)
}

func (j *JsonGo) Get(keys ...interface{}) (*JsonGo, error) {
	if len(keys) == 0 {
		return j, nil
	}

	keys, err := getkeys(keys)
	if err != nil {
		return nil, err
	}

	item := j
	for _, key := range keys {
		item, err = item.getItem(key)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	return item, nil
}

// setValue 设置interface中的值，只支持map和数组
func (j *JsonGo) setValue(key interface{}, val interface{}) (*JsonGo, error) {
	v, err := NewJsonGo(val)
	if err != nil {
		return nil, err
	}
	switch key.(type) {
	case int:
		data, err := j.arrayData()
		if err != nil {
			return nil, err
		}
		i := key.(int)
		//下标超出则在数组中追加元素
		l := len(*data)

		if i < -1 || i > l {
			str, _ := j.Str("")
			return nil, errTarget.New("setValue fail:%s[%d] is invalid.index value in(-1,%d) means append item,otherwise replace the item", str, i, l)
		}
		if i == -1 || i == l {
			j.v = append(*data, v)
			return v, nil
		}
		(*data)[i] = v
		return v, nil
	case string:
		data, err := j.mapData()
		if err != nil {
			return nil, err
		}
		data[key.(string)] = v
		return v, nil
	}
	return nil, errs.ErrSystem.New(fmt.Sprintf("setItem type is not support %T", j))
}

// setValue 设置interface中的值，只支持map和数组
func (j *JsonGo) removeValue(key interface{}) error {
	switch key.(type) {
	case int:
		data, err := j.arrayData()
		if err != nil {
			return err
		}
		_data := *data
		i := key.(int)
		//下标超出则在数组中追加元素
		l := len(_data)

		if i < -1 || i > l-1 || l == 0 {
			str, _ := j.Str("")
			return errTarget.New("removeValue fail:%s[%d] is invalid.index value in(-1) means remove last item", str, i, l)
		}
		if i == -1 {
			j.v = _data[:l-1]
			return nil
		}
		j.v = append(_data[:i], _data[i+1:]...)
		return nil
	case string:
		data, err := j.mapData()
		if err != nil {
			return err
		}
		delete(data, key.(string))
		return nil
	}
	return errs.ErrSystem.New(fmt.Sprintf("removeValue type is not support %T", key))
}

// Remove Set 设置interface中的值，支持多级设置，支持map,数组和json字符串
func (j *JsonGo) Remove(keys ...interface{}) error {
	if len(keys) == 0 {
		return errs.ErrValidate.New("remove key is empty")
	}

	keys, err := getkeys(keys)
	if err != nil {
		return err
	}
	l := len(keys)
	if l == 1 {
		return j.removeValue(keys[0])
	} else {
		_item, err := j.Get(keys[:l-1]...)
		if err != nil {
			return err
		}
		return _item.removeValue(keys[l-1])
	}
}

// Set 设置interface中的值，支持多级设置，支持map,数组和json字符串
func (j *JsonGo) Set(val interface{}, keys ...interface{}) error {
	if len(keys) == 0 {
		return j.from(val)
	}

	keys, err := getkeys(keys)
	if err != nil {
		return err
	}

	item := j

	for indx, key := range keys {
		_item, err := item.getItem(key)
		if err != nil {
			//遇到不存在的key,创建对象
			if errKey.Is(err) {
				addVas := val
				//遇到不存在的key,创建对象
				if indx < len(keys)-1 {
					addVas = getAddVal(keys[indx+1])
				}
				_item, err = item.setValue(key, addVas)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
		item = _item
	}

	return item.from(val)
}

func getAddVal(key interface{}) interface{} {
	err := checkKeys(key)
	if err != nil {
		panic(err)
	}

	switch key.(type) {
	case int:
		return make([]interface{}, 0)
	case string:
		return make(map[string]interface{})
	}
	panic(errs.ErrSystem.New("key type error:%T", key))
}
