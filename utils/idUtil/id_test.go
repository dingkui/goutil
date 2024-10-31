package idUtil_test

import (
	"gitee.com/dk83/goutils/dlog"
	"gitee.com/dk83/goutils/utils/idUtil"
	"testing"
)

func TestId(t *testing.T) {
	dlog.Info("TestTime:", idUtil.Rand64.Rand(32))
	dlog.Info("TestTime:", idUtil.Rand62.Rand(8))
	dlog.Info("TestTime:", idUtil.Rand32.Rand(32))
	dlog.Info("TestTime:", idUtil.Rand16.Rand(32))
	dlog.Info("TestTime:", idUtil.Rand10.Rand(32))
	dlog.Info("TestTime:", idUtil.Rand8.Rand(32))
	//dlog.Info("TestTime:", native.idUtil.ID62(32))
	//dlog.Info("TestTime:", native.idUtil.ID62(8))
	//dlog.Info("TestTime:", native.idUtil.NUM(32))
}
func TestRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		idUtil.Rand62.Rand(32)
	}
	for i := 0; i < 10000; i++ {
		idUtil.NUM(32)
	}
}
