package test

//
//import (
//	"errors"
//	"fmt"
//)
//
//type keyType struct {
//	getItem func(jsonGo *JsonGo, key interface{}) (*JsonGo, error)
//	setItem func(j *JsonGo, key interface{}, v *JsonGo) error
//}
//var (
//	keyInt  = &keyType{}
//	keyString = &keyType{}
//	keyErr = &keyType{}
//)
//
//func init() {
//	keyErr.setItem = func(j *JsonGo, key interface{}, v *JsonGo) error {
//		return errors.New(fmt.Sprintf("setItem type is not support with key:%v[%T]", key,key))
//	}
//	keyErr.getItem = func(j *JsonGo, key interface{}) (*JsonGo, error) {
//		return nil, errors.New(fmt.Sprintf("getItem fail with key:%v[%T] jsonValue:%v[%T]", key, key, j.v, j.v))
//	}
//
//	keyInt.setItem = func(j *JsonGo, key interface{}, v *JsonGo) error {
//		data, err := j.arrayData()
//		if err != nil {
//			return err
//		}
//		i := key.(int)
//		if i < -1 || i >= len(data) {
//			return errors.New(fmt.Sprintf("setItem index is out of range %d", i))
//		}
//		//-1表示往数组中追加元素
//		if i == -1 {
//			data = append(data, v)
//			return nil
//		}
//		data[i] = v
//		return nil
//	}
//	keyInt.getItem = func(j *JsonGo, key interface{}) (*JsonGo, error) {
//		data, err := j.mapData()
//		if err != nil {
//			return nil, err
//		}
//		item, ok := data[key.(string)]
//		if !ok {
//			return nil, errors.New(fmt.Sprintf("can't get [%s] from:%v", key, j))
//		}
//		return item, nil
//	}
//
//	keyString.setItem = func(j *JsonGo, key interface{}, v *JsonGo) error {
//		s := key.(string)
//		if s == "" {
//			j.from(v)
//		} else if j._type == jsonMap {
//			data, err := j.mapData()
//			if err != nil {
//				return err
//			}
//			data[s] = v
//		} else if j._type == jsonArray {
//			_key, err := getInt(s)
//			if err != nil {
//				return err
//			}
//			j.setValue(_key, v)
//		} else {
//			return errors.New(fmt.Sprintf("setItem type key is not support:%T key:%s", j, s))
//		}
//		return nil
//	}
//	keyString.getItem = func(j *JsonGo, key interface{}) (*JsonGo, error) {
//		data, err := j.arrayData()
//		if err != nil {
//			return nil, err
//		}
//		i := key.(int)
//		if i < 0 || i >= len(data) {
//			return nil, errors.New(fmt.Sprintf("getItem index is out of range %d", i))
//		}
//		return data[key.(int)], nil
//	}
//
//}
//
//func getkeyType(key interface{}) *keyType {
//	switch key.(type) {
//	case string:
//		return keyString
//	case int:
//		return keyInt
//	}
//	return keyErr
//}
