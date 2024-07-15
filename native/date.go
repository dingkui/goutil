package native

import (
	"time"
)

type date struct {
	TfDateTimeM string
	TfDateTime  string
	TfTimeStr   string
	TfMinuteStr string
	TfHourStr   string
	TfDayStr    string
}

var DateUtil = date{
	TfDateTimeM: "2006/01/02 15:04:05.000",
	TfDateTime:  "2006/01/02 15:04:05",
	TfTimeStr:   "20060102_150405",
	TfMinuteStr: "20060102_1504",
	TfHourStr:   "20060102_15",
	TfDayStr:    "20060102",
}

func (t date) DateTimeStrM() string {
	return time.Now().Format(t.TfDateTimeM)
}
func (t date) DateTimeStr() string {
	return time.Now().Format(t.TfDateTime)
}
func (t date) TimeStr() string {
	return time.Now().Format(t.TfTimeStr)
}
func (t date) MinuteStr() string {
	return time.Now().Format(t.TfMinuteStr)
}
func (t date) HourStr() string {
	return time.Now().Format(t.TfHourStr)
}
func (t date) DayStr() string {
	return time.Now().Format(t.TfDayStr)
}
