package apputil

import (
	"os"
	"strings"
)

var paras = map[string]string{}

func init() {
	for _, arg := range os.Args {
		index := strings.Index(arg, "=")
		if index > -1 {
			paras[arg[0:index]] = arg[index+1:]
		} else {
			paras[arg] = "1"
		}
	}
}

func IsPara(name string) bool {
	v, has := paras[name]
	if !has {
		return false
	}
	return v == "1" || v == "true"
}
func GetPara(name string, def string) string {
	v, has := paras[name]
	if has {
		return v
	}
	return def
}

//	"github.com/go-eden/routine"
//func GetLocal(name string) string {
//	storage := routine.NewLocalStorage()
//	get,ok := storage.Get().(map[string]string)
//	if !ok {
//		storage.Set(map[string]string{})
//		return ""
//	}
//	s,has:= get[name]
//	if has {
//		return s
//	}
//	return ""
//}
//
//func SetLocal(name string,value string) {
//	storage := routine.NewLocalStorage()
//	get,ok := storage.Get().(map[string]string)
//	if !ok {
//		storage.Set(map[string]string{
//			name:value,
//		})
//	}
//	get[name]=value
//}

func GetParaf(name string, def func() string) string {
	v, has := paras[name]
	if has {
		return v
	}
	return def()
}
