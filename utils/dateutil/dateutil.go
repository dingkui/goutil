package dateutil

import (
	"gitee.com/dk83_admin/goutil/utils/zlog"
	"os"
	"time"
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
	return time.Now().Format("2006/01/02 15:04:05")
}
func TimeStr() string {
	return time.Now().Format("20060102_150405")
}
func MinuteStr() string {
	return time.Now().Format("20060102_1504")
}
func HourStr() string {
	return time.Now().Format("20060102_15")
}
func DayStr() string {
	return time.Now().Format("20060102")
}
