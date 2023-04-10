package config

import (
	"encoding/json"
	"fmt"
	"gitee.com/dk83/goutils/dateutil"
	"gitee.com/dk83/goutils/jsonutil"
	"gitee.com/dk83/goutils/logutil"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type localFileLog struct {
	Level  int
	Root   string
	day    string
	_debug *log.Logger
	_error *log.Logger
	locker sync.Mutex
}

func (this *localFileLog) Log(level int, depthPre int, msg string) {
	levelStr := "debug"
	log := this._debug
	switch level {
	case 1:
		levelStr = "info"
		break
	case 2:
		levelStr = "warn"
		break
	case 3:
		levelStr = "error"
		log = this._error
	}
	log.Output(3+depthPre, fmt.Sprintf("%s %s", levelStr, msg))
}

func (this *localFileLog) LogAble(level int) bool {
	return this.Level <= level
}

func (this *localFileLog) Init() {
	if this.day == dateutil.DayStr() {
		return
	}
	this.locker.Lock()
	defer func() {
		this.locker.Unlock()
	}()
	if this.day == dateutil.DayStr() {
		return
	}
	this.day = dateutil.DayStr()
	debugLogFile := filepath.Join(this.Root, "server."+this.day+".log")
	errLogFile := filepath.Join(this.Root, "server-error."+this.day+".log")
	os.MkdirAll(filepath.Dir(debugLogFile), os.ModePerm)

	f, err := os.OpenFile(debugLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	e, err := os.OpenFile(errLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		f.Close()
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
			logutil.Info("log change last")
			f.Close()
			e.Close()
			this.Init()
			logutil.Info("log change first")
		})
	}
}

func InitLog() {
	configMap := GetConfigMap("Logger")
	_type, _ := jsonutil.GetString(configMap, "Type")
	if _type == "localFileLog" {
		var logConf = &localFileLog{}
		bytes := jsonutil.Get_bytes(configMap)
		err := json.Unmarshal(bytes, logConf)
		if err != nil {
			fmt.Print(err)
		}
		logutil.InitLog(logConf)
	}
}
