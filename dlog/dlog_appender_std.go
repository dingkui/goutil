package dlog

import (
	"fmt"
)

type consoleAppender struct {
	level int
}

func (f *consoleAppender) Enable(level int) bool {
	return f.level <= level
}
func (f *consoleAppender) WriteLog(s string, _ string) (int, error) {
	return fmt.Print(s)
}
func (f *consoleAppender) Close() {
}

func AddAppenderConsole(level int) {
	AddLogger(&consoleAppender{level})
}
