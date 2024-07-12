package fileutil

import (
	"gitee.com/dk83/goutils/zlog"
	"testing"
)

func TestFileName(t *testing.T) {
	set := NewFileNameSet()
	zlog.Info(set.Unique_file_name("x/te st.json"))
	zlog.Info(set.Unique_file_name("t est.json"))
	zlog.Info(set.Unique_file_name("tes t.json"))
	zlog.Info(set.Unique_file_name("tes t2.json"))
	zlog.Info(set.Unique_file_name("2tes t2.json"))
	zlog.Info(set.Unique_file_name("2tes t2.json"))
	zlog.Info(set.Unique_file_name("2tes t2.json"))
}
