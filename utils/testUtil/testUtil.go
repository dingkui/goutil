package testUtil

import (
	"github.com/dingkui/goutil/dlog"
	"github.com/dingkui/goutil/utils/stringUtil"
	"github.com/dingkui/goutil/utils/valUtil"
	"testing"
)

const (
	UnCheck = "unCheck"
	FailDef = "failDef"
)
func init() {
	dlog.AddAppenderConsole(0)
}

func Check(t *testing.T, fmtx string, got interface{}, want string) {
	DoTest(t,func()(ok bool,info string){
		_got, e := valUtil.Str(got, FailDef)
		if e != nil {
			return false, stringUtil.Fmt(e)
		}
		if want != UnCheck && _got != want {
			return false,stringUtil.Fmt("got %s; want %s", _got, want)
		}
		return true,stringUtil.Fmt(fmtx, _got)
	})
}
func DoTest(t *testing.T,fn func()(ok bool,info string)) {
	ok, info := fn()
	if !ok {
		dlog.ERROR.LogCaller(1,"NG: %s", info)
		t.Fail()
		return
	}
	dlog.TEST.LogCaller(1,"OK: %s",info)
}
