package dlog

import (
	"fmt"
	"os"
	"syscall"
)

var Stderr = os.NewFile(uintptr(syscall.Stderr), "/dev/stderr99")

type consoleAppender struct {
	level int
}

func (f *consoleAppender) Enable(level int) bool {
	return f.level <= level
}
func (f *consoleAppender) WriteLog(s string, _ string) (int, error) {
	return fmt.Print(s)
}

func AddAppenderConsole(level int) {
	AddLogger(&consoleAppender{level})
}
