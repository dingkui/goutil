package dateUtil_test

import (
	"fmt"
	"gitee.com/dk83/goutils/dlog"
	"gitee.com/dk83/goutils/utils/dateUtil"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	duration := 24 * time.Hour * 365 * 82
	add := time.Now().Add(duration)
	second := add.Unix()
	second1 := time.Now().Unix()
	dlog.Info("TestTime:", second, second1, second-second1, add)
	dlog.Info("TestTime2:", fmt.Sprintf("%x", second))
	dlog.Info("TestTime2:", fmt.Sprintf("%x", second1))
}
func TestTime1(t *testing.T) {
	dlog.Info("FormatNow:", dateUtil.Day.FormatNow())
	dlog.Info("FormatNow:", dateUtil.DateTimeM.FormatNow())
}
