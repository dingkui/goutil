package dlog

import (
	"fmt"
	"gitee.com/dk83/goutils/native"
	"os"
	"path/filepath"
	"time"
)

//按日期划分文件
type appenderDaily struct {
	busing   bool //是否正在写入
	ready    bool //是否已经准备好
	level    int
	_day     string
	filePath string
	file     *os.File
}

func (f *appenderDaily) init(_day string) error {
	if f._day == _day {
		return nil
	}
	f.ready = false
	if f.busing {
		time.AfterFunc(time.Second*1, func() { f.init(_day) })
		return nil
	}

	filePath := fmt.Sprintf(f.filePath, _day)
	os.MkdirAll(filepath.Dir(filePath), os.ModePerm)

	if f.file != nil {
		err := f.file.Close()
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	f._day = _day
	f.file = file
	f.ready = true
	return nil
}

func (f *appenderDaily) WriteLog(s string, _ string) (int, error) {
	if !f.ready {
		Error("daily log not ready:" + s)
		return 0, nil
	}
	f.busing = true
	re, err := f.file.WriteString(s + "\n")
	if err != nil {
		panic("daily write error:" + s)
	}
	f.busing = false
	return re, err
}
func (f *appenderDaily) Enable(level int) bool {
	return f.level <= level
}
func (f *appenderDaily) Close() {
	f.ready = false
	if f.busing {
		time.AfterFunc(time.Second*1, func() { f.Close() })
		return
	}
	err := f.file.Close()
	if err != nil {
		Error("close daily log error:", err)
		return
	}
}

//每天重新设置日志文件
func reInitAppendersDaily() {
	_day := native.DateUtil.DayStr()
	for _, w := range _leveAppender {
		f, ok := w.(*appenderDaily)
		if ok {
			f.init(_day)
		}
	}

	currentTime := time.Now()
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
	delayScends := endTime.Unix() - currentTime.Unix() + 1
	if delayScends > 0 {
		time.AfterFunc(time.Second*time.Duration(delayScends), reInitAppendersDaily)
	}
}

func init() {
	go reInitAppendersDaily()
}

// AddAppenderDaily 添加每日日志文件
func AddAppenderDaily(level int, filePath string) bool {
	appenderDaily := &appenderDaily{
		filePath: filePath,
		level:    level,
	}

	day := native.DateUtil.DayStr()
	err := appenderDaily.init(day)
	if err != nil {
		Error("AddLogger failed!file is:"+filePath, day)
		return false
	}

	AddLogger(appenderDaily)
	return true
}
