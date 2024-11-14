package base_test

import (
	"github.com/dingkui/goutil/dhttp"
	"github.com/dingkui/goutil/djson"
	"github.com/dingkui/goutil/dlog"
	"github.com/dingkui/goutil/encry"
	"github.com/dingkui/goutil/errs"
	"github.com/dingkui/goutil/utils/apputil"
	"github.com/dingkui/goutil/utils/confUtil"
	"github.com/dingkui/goutil/utils/dateUtil"
	"github.com/dingkui/goutil/utils/fileUtil"
	"github.com/dingkui/goutil/utils/idUtil"
	"github.com/dingkui/goutil/utils/mathUtil"
	"github.com/dingkui/goutil/utils/runtimeUtil"
	"github.com/dingkui/goutil/utils/stringUtil"
	"github.com/dingkui/goutil/utils/valUtil"
	"github.com/dingkui/goutil/utils/valUtil/forceVal"
	"testing"
)
func TestAll(t *testing.T) {
	dlog.Info(apputil.Para("xx", "1"))
	dlog.Info(dhttp.Client{})
	dlog.Info(djson.NewJsonMap())
	dlog.Info(encry.Md5)
	dlog.Info(errs.ErrSystem.New("ErrSystem"))
	dlog.Info(dateUtil.DateTime.FormatNow())
	dlog.Info(fileUtil.Exists("d:/"))
	dlog.Info(idUtil.ID16(32))
	dlog.Info(mathUtil.Round(12.1, 3))
	dlog.Info(runtimeUtil.GetCaller(1))
	dlog.Info(stringUtil.Fmt("%d", 1))
	dlog.Info(valUtil.Str(1))
	dlog.Info(valUtil.Int64("", 999))
	data := map[string]interface{}{"a": 2}
	dlog.Info(confUtil.NewConf("d:/xx/xx2.json", &data, true))
	dlog.Info(data)
	dlog.Info(forceVal.Int64("64"))
	dlog.Info(string([]byte{0xc6, 0x89, 0xb7, 0xba, 0xcc}))
}
