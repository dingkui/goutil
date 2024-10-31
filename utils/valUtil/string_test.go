package valUtil_test

import (
	"gitee.com/dk83/goutils/djson"
	"gitee.com/dk83/goutils/dlog"
	"gitee.com/dk83/goutils/utils/valUtil"
	"testing"
)

func TestStrJson(t *testing.T) {
	json, _ := djson.NewJsonGo(make(map[string]interface{}))
	json.Set("123", "@xx.x")
	dlog.Info(valUtil.Str(json))
	json1, _ := djson.NewJsonGo("123")
	dlog.Info(valUtil.Str(json1))
	json2, _ := djson.NewJsonFile("", make(map[string]interface{}))
	json2.Set("123", "@xx.x22")
	dlog.Info(valUtil.Str(json2))
}
