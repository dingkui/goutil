package dlog

import (
	"fmt"
)

type consoleAppender struct {
	level  int
	colors map[string]string
}

func (f *consoleAppender) Enable(level int) bool {
	return f.level <= level
}
func (f *consoleAppender) WriteLog(s string, name string) (int, error) {
	if name == "" {
		return fmt.Println(s)
	}
	s2, has := f.colors[name]
	if has {
		return fmt.Println("\033["+s2+"m"+s+"\033[0m")
	}
	return fmt.Println(s)
}
func (f *consoleAppender) Close() {
}
func (f *consoleAppender) Color(colors map[string]string) {
	f.colors = colors
}

func AddAppenderConsole(level int) *consoleAppender {
	appender := &consoleAppender{level, map[string]string{
		"DEBUG":  "37",
		"INFO":  "36",
		"WARN":  "33",
		"ERROR": "31",
		"TEST":  "32",
	}}
	AddLogger(appender)
	return appender
}
