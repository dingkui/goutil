package jsonutil

import (
	"gitee.com/dk83/goutils/fileutil"
	"gitee.com/dk83/goutils/zlog"
	"io/ioutil"
	"os"
	"path/filepath"
)

type JSONFile struct {
	*JsonGo
	_file string
}

func NewJsonFile(file string, json *JsonGo) *JSONFile {
	if json == nil {
		json = NewJsonGo(make(map[string]interface{}))
	}
	j := &JSONFile{
		JsonGo: json,
		_file:  file,
	}
	j._file = file
	return j
}

func (j *JSONFile) Reload() {
	j.Read(j._file)
}
func (j *JSONFile) Read(file string) {
	b, e := ioutil.ReadFile(file)
	if e != nil {
		zlog.Error("read file error:", e)
		return
	}
	newJSON := NewJsonGo(b)
	if newJSON != nil {
		*j.JsonGo = *newJSON
		j._file = file
	}
}
func (j *JSONFile) save(format bool) bool {
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
	fileutil.WriteAndSyncFile(j._file, formatJson(bytes, format), os.ModePerm)
	return true
}
func (j *JSONFile) Save() bool {
	return j.save(false)
}
func (j *JSONFile) SaveFormat() bool {
	return j.save(true)
}
