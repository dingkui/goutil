package native_test

import (
	"gitee.com/dk83/goutils/dlog"
	"gitee.com/dk83/goutils/native"
	"testing"
)

func TestId(t *testing.T) {
	dlog.Info("TestTime:", native.IdUtil.Rand64.Rand(32))
	dlog.Info("TestTime:", native.IdUtil.Rand62.Rand(8))
	dlog.Info("TestTime:", native.IdUtil.Rand32.Rand(32))
	dlog.Info("TestTime:", native.IdUtil.Rand16.Rand(32))
	dlog.Info("TestTime:", native.IdUtil.Rand10.Rand(32))
	dlog.Info("TestTime:", native.IdUtil.Rand8.Rand(32))
	//dlog.Info("TestTime:", native.IdUtil.ID62(32))
	//dlog.Info("TestTime:", native.IdUtil.ID62(8))
	//dlog.Info("TestTime:", native.IdUtil.NUM(32))
}
func TestRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		native.IdUtil.Rand62.Rand(32)
	}
	for i := 0; i < 10000; i++ {
		native.IdUtil.NUM(32)
	}
}
