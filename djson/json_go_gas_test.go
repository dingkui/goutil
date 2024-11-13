// Package jsonutil : get and set
package djson_test

import (
	"github.com/dingkui/goutil/djson"
	"github.com/dingkui/goutil/dlog"
	"testing"
)

func TestArrayGas2(t *testing.T) {
	jsonGo := djson.NewJsonMap()
	dlog.Info(jsonGo.Set("xxValue", -1))
	dlog.Info(jsonGo.Set("xxValue2", -1))
	dlog.Info(jsonGo.Set("xxValue3", -1))
	check(t, "jsonGo:%s", jsonGo, `{}`)
}
func TestArrayGas3(t *testing.T) {
	jsonGo := djson.NewJsonArray()
	jsonGo.Set("xxValue", -1)
	jsonGo.Set("xxValue2", -1)
	check(t, "jsonGo:%s", jsonGo, `["xxValue2","xxValue"]`)
	jsonGo.Set("xxValue3", -1)
	check(t, "jsonGo:%s", jsonGo, `["xxValue3","xxValue2","xxValue"]`)
	jsonGo.Set("xxValue4", -2)
	check(t, "jsonGo:%s", jsonGo, `["xxValue3","xxValue2","xxValue","xxValue4"]`)
	jsonGo.Remove(-2)
	check(t, "jsonGo:%s", jsonGo, `["xxValue3","xxValue2","xxValue"]`)
	jsonGo.Remove(-1)
	check(t, "jsonGo:%s", jsonGo, `["xxValue2","xxValue"]`)
	dlog.Info(jsonGo.Remove(9))
	check(t, "jsonGo:%s", jsonGo, `["xxValue2","xxValue"]`)
}
