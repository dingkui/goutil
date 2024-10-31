package fileUtil_test

import (
	"gitee.com/dk83/goutils/dlog"
	"gitee.com/dk83/goutils/utils/fileUtil"
	"os"
	"testing"
)

func TestFileList(t *testing.T) {
	fileUtil.LL("d:/test", func(path string, info os.FileInfo) error {
		dlog.Info("%s %s %d", path, info.Name(), info.Size())
		return nil
	})
}
