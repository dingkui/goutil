package webutil

import (
	"net/http"
)

func SetCookie(name string, value string, w *http.ResponseWriter) {
	cookie := http.Cookie{
		Name:  name,
		Value: value,
	}
	http.SetCookie(*w, &cookie)
}
func GetCookie(r *http.Request, name string) string {
	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie.Value
		}
	}
	return ""
}
