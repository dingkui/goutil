package native

import (
	"fmt"
	"strings"
)

type s byte

const StringUtil s = iota

func (x s) Fmt(s interface{}, v ...interface{}) string {
	if len(v) == 0 {
		return fmt.Sprint(s)
	}
	message, ok := s.(string)
	if !ok || strings.Index(message, "%") == -1 {
		return fmt.Sprint(append([]interface{}{s}, v...)...)
	}
	return fmt.Sprintf(message, v...)
}

func (x s) InStringArray(list []string, find string) bool {
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
