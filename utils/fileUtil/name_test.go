package fileUtil_test

import (
	"gitee.com/dk83/goutils/dlog"
	"gitee.com/dk83/goutils/utils/fileUtil"
	"testing"
)

func TestFileName(t *testing.T) {
	set := fileUtil.NewFileNameSet()
	dlog.Info(set.Unique_file_name("x/te st.json"))
	dlog.Info(set.Unique_file_name("t est.json"))
	dlog.Info(set.Unique_file_name("tes t.json"))
	dlog.Info(set.Unique_file_name("tes t2.json"))
	dlog.Info(set.Unique_file_name("2tes t2.json"))
	dlog.Info(set.Unique_file_name("2tes t2.json"))
	dlog.Info(set.Unique_file_name("2tes t2.json"))
}