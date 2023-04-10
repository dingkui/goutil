package logutil

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type printLog struct {
	out    *log.Logger
	err    *log.Logger
	locker sync.Mutex
}

func (this *printLog) Log(level int, depthPre int, msg string) {
	levelStr := "debug"
	log := this.out
	switch level {
	case 1:
		levelStr = "info"
		break
	case 2:
		levelStr = "warn"
		break
	case 3:
		levelStr = "error"
		log = this.err
	}
	log.Output(3+depthPre, fmt.Sprintf("%s %s", levelStr, msg))
}

func (this *printLog) LogAble(level int) bool {
	if this.out != nil {
		return true
	}
	this.locker.Lock()
	defer func() {
		this.locker.Unlock()
	}()
	if this.out != nil {
		return true
	}
	this.out = log.New(io.MultiWriter(os.Stdout), "", log.Ldate|log.Ltime|log.Lshortfile)
	this.err = log.New(io.MultiWriter(os.Stderr), "", log.Ldate|log.Ltime|log.Lshortfile)
	return true
}

func (this *printLog) Init() {}
func init() {
	InitLog(&printLog{})
}
