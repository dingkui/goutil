package fileUtil_test

import (
	"github.com/dingkui/goutil/dlog"
	"github.com/dingkui/goutil/utils/fileUtil"
	"os"
	"testing"
)

func TestFileList(t *testing.T) {
	fileUtil.LL("d:/test", func(path string, info os.FileInfo) error {
		dlog.Info("%s %s %d", path, info.Name(), info.Size())
		return nil
	})
}
