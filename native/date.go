package native

import (
	"time"
)

type dateLayout string
type dateUtil struct {
	DateTimeM dateLayout
	DateTime  dateLayout
	DateTime2 dateLayout
	Minute    dateLayout
	Hour      dateLayout
	Day       dateLayout
}

var DateUtil = dateUtil{
	DateTimeM: "2006/01/02 15:04:05.000",
	DateTime:  "2006/01/02 15:04:05",
	DateTime2: "20060102_150405",
	Minute:    "20060102_1504",
	Hour:      "20060102_15",
	Day:       "20060102",
}

func (t dateUtil) Layout(layout string) dateLayout {
	return dateLayout(layout)
}
func (t dateUtil) Now() time.Time {
	return time.Now()
}
func (t dateUtil) Unix(ti int64) time.Time {
	return time.Unix(ti, 0)
}

func (t dateLayout) Pause(ti string) (time.Time, error) {
	return time.Parse(string(t), ti)
}
func (t dateLayout) FormatNow() string {
	return time.Now().Format(string(t))
}
func (t dateLayout) Format(ti time.Time) string {
	return ti.Format(string(t))
}
func (t dateLayout) FormatUnix(ti int64) string {
	return time.Unix(ti, 0).Format(string(t))
}
func (t dateLayout) FormatStr(ti string, layoutSource dateLayout) string {
	parse, err := layoutSource.Pause(ti)
	if err != nil {
		return ""
	}
	return t.Format(parse)
}
