package apputil

import (
	"strconv"
	"syscall"
)

func Env(name string, def ...string) string {
	v, found := syscall.Getenv(name)
	if !found && len(def) > 0 {
		return def[0]
	}
	return v
}
func EnvF(name string, def func() string) string {
	v, found := syscall.Getenv(name)
	if found {
		return v
	}
	return def()
}
func EnvIs(name string) bool {
	v, found := syscall.Getenv(name)
	if !found {
		return false
	}
	return v == "1" || v == "true"
}
func EnvInt(name string, def int) int {
	num, err := strconv.Atoi(Env(name))
	if err != nil {
		return def
	}
	return num
}
func EnvInt64(name string, def int64) int64 {
	num, err := strconv.ParseInt(Env(name), 10, 64)
	if err != nil {
		return def
	}
	return num
}
