package dateutil

import (
	"gitee.com/dk83/goutils/zlog"
	"os"
	"time"
)

const (
	TfDateTime  string = "2006/01/02 15:04:05"
	TfTimeStr   string = "20060102_150405"
	TfMinuteStr string = "20060102_1504"
	TfHourStr   string = "20060102_15"
	TfDayStr    string = "20060102"
)

//获取文件修改时间 返回unix时间戳
func GetFileModTime(path string) (time.Time, bool) {
	f, err := os.Open(path)
	if err != nil {
		return time.Now(), false
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		zlog.Error("stat fileinfo error")
		return time.Now(), false
	}
	return fi.ModTime(), true
}
func DateTimeStr() string {
	return time.Now().Format(TfDateTime)
}
func TimeStr() string {
	return time.Now().Format(TfTimeStr)
}
func MinuteStr() string {
	return time.Now().Format(TfMinuteStr)
}
func HourStr() string {
	return time.Now().Format(TfHourStr)
}
func DayStr() string {
	return time.Now().Format(TfDayStr)
}
