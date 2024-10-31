package dateUtil

import (
	"time"
)

type dateLayout string

var (
	DateTimeM dateLayout = "2006/01/02 15:04:05.000"
	DateTime  dateLayout = "2006/01/02 15:04:05"
	DateTime2 dateLayout = "20060102_150405"
	Minute    dateLayout = "20060102_1504"
	Hour      dateLayout = "20060102_15"
	Day       dateLayout = "20060102"
)

func Layout(layout string) dateLayout {
	return dateLayout(layout)
}
func Now() time.Time {
	return time.Now()
}
func Unix(ti int64) time.Time {
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
