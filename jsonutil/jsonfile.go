package jsonutil

import (
	"gitee.com/dk83/goutils/fileutil"
	"gitee.com/dk83/goutils/zlog"
	"io/ioutil"
	"os"
	"path/filepath"
)

type JSONFile struct {
	JsonGo
	_file string
}

func (j *JSONFile) NewFile(file string) {
	newJSON := NewJsonGo(make(map[string]interface{}))
	if newJSON != nil {
		j.JsonGo = *newJSON
		j._file = file
	}
}
func (j *JSONFile) Read(file string) {
	b, e := ioutil.ReadFile(file)
	if e != nil {
		zlog.Error("read file error:", e)
		return
	}
	newJSON := NewJsonGo(b)
	if newJSON != nil {
		j.JsonGo = *newJSON
		j._file = file
	}
}

func (j *JSONFile) Save() bool {
	if j._file == "" {
		zlog.Error("Save json to file fail, JSONFile is not Read completed!", j._file)
		return false
	}
	bytes, err := j.Byte()
	if err != nil {
		zlog.Error(err, j._file)
		return false
	}
	os.MkdirAll(filepath.Dir(j._file), os.ModePerm)
	fileutil.WriteAndSyncFile(j._file, formatJson(bytes, true), os.ModePerm)
	return true
}
