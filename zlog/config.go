package zlog

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type ILoggerGetter interface {
	getLocalLogger(level int) *log.Logger
	getRemoteLogger(level int) func(level string, msg string, caller string)
}

var _logGetter ILoggerGetter

func InitLog(logger ILoggerGetter) *ILoggerGetter {
	_logGetter = logger
	return &_logGetter
}

type defaultLogGetter struct {
	localLevel  int
	remoteLevel int
	local       *localFileLog
}
type localFileLog struct {
	root    string
	logName string
	_day    string
	_debug  *log.Logger
	_error  *log.Logger
	_locker sync.Mutex
}

func NewDefaultLogGette(local int, remote int, logRoot string, logName string) ILoggerGetter {
	return &defaultLogGetter{
		localLevel:  local,
		remoteLevel: remote,
		local: &localFileLog{
			root:    logRoot,
			logName: logName,
		},
	}
}

func (this *defaultLogGetter) getLocalLogger(level int) *log.Logger {
	if this.local == nil || level < this.localLevel {
		return nil
	}
	return this.local.getLocalLogger(level)
}
func (this *defaultLogGetter) getRemoteLogger(level int) func(level string, msg string, caller string) {
	if level < this.remoteLevel {
		return nil
	}
	return nil
}

func (this *localFileLog) getLocalLogger(level int) *log.Logger {
	if this._debug == nil {
		this.initLocal()
	}
	if level < 3 {
		return this._debug
	}
	return this._error
}

func (this *localFileLog) initLocal() {
	if this._day == dayStr() {
		return
	}
	this._locker.Lock()
	defer func() {
		this._locker.Unlock()
	}()
	if this._day == dayStr() {
		return
	}
	this._day = dayStr()
	if this.root == "" {
		this.root = filepath.Join(RootDir, "logs")
	}
	if this.logName == "" {
		this.logName = "server%s.%s.log"
	}

	debugLogFile := filepath.Join(this.root, fmt.Sprintf(this.logName, "", this._day))
	errLogFile := filepath.Join(this.root, fmt.Sprintf(this.logName, "-error", this._day))
	os.MkdirAll(filepath.Dir(debugLogFile), os.ModePerm)

	f, err := os.OpenFile(debugLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		f.Close()
		return
	}
	e, err := os.OpenFile(errLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		e.Close()
		return
	}

	// 组合一下即可，os.Stdout代表标准输出流
	this._debug = log.New(io.MultiWriter(os.Stdout, f), "", log.Ldate|log.Ltime|log.Lshortfile)
	this._error = log.New(io.MultiWriter(os.Stderr, e), "", log.Ldate|log.Ltime|log.Lshortfile)

	currentTime := time.Now()
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())

	delayScends := int64(endTime.Unix() - currentTime.Unix() + 1)
	if delayScends > 0 {
		time.AfterFunc(time.Second*time.Duration(delayScends), func() {
			Info("log change last")
			f.Close()
			e.Close()
			this.initLocal()
			Info("log change first")
		})
	}
}
