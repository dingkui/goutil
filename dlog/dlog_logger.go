package dlog

import (
	"fmt"
	"os"
	"sync"
)

type _logInfo struct {
	l int    //level
	t string //time
	f string //file
	m string //msg
}

type LogAppender interface {
	Close()
	Enable(level int) bool
	WriteLog(msg string, logName string) (n int, err error)
}

var (
	_leveAppender = make([]LogAppender, 0)
	_logQueue     = make(chan *_logInfo, 100)
	defLogger     = os.Stdout

	ready = false
	mu    sync.Mutex
	mu2   sync.Mutex
)

func init() {
	go func() {
		for data := range _logQueue {
			log(data)
		}
	}()
}
func log(data *_logInfo) {
	logType := getType(data.l)
	info := fmt.Sprintf("%s %s %s: %s", data.t, logType.name, data.f, data.m)
	mu.Lock()
	defer mu.Unlock()
	if ready {
		for _, w := range _leveAppender {
			if w.Enable(data.l) {
				//fmt.Printf("log %d:%s %v\n",data.l,logType.name,w)
				w.WriteLog(info, logType.name)
			}
		}
	} else {
		os.Stdout.WriteString(info + "\n")
	}
}

func AddLogger(nw LogAppender) {
	ready = false
	_leveAppender = append(_leveAppender, nw)
	ready = true
}

func ClearAppenders() {
	ready = false
	for _, w := range _leveAppender {
		w.Close()
	}
	_leveAppender = _leveAppender[:0]
}
