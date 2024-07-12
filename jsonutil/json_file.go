package jsonutil

import (
	"gitee.com/dk83/goutils/apputil"
	"gitee.com/dk83/goutils/errs"
	"gitee.com/dk83/goutils/fileutil"
	"gitee.com/dk83/goutils/zlog"
	"io/ioutil"
	"os"
	"path/filepath"
)

type JsonFile struct {
	*JsonGo
	_file string
}

func NewJsonFile(file string, data interface{}) (*JsonFile, error) {
	if data == nil {
		data = make(map[string]interface{})
	}
	jsonGo, err := NewJsonGo(data)
	if err != nil {
		return nil, err
	}
	j := &JsonFile{
		_file: file,
	}
	j.JsonGo = jsonGo
	return j, nil
}
func ReadFile(file string, j *JsonFile) error {
	j._file = file
	err := j.ReLoad()
	if err != nil {
		return err
	}
	return nil
}
func (j *JsonFile) ReLoad() error {
	if j._file == "" {
		panic("JsonFile file is nil!")
	}

	b, e := ioutil.ReadFile(j._file)
	if e != nil {
		return errs.ErrEnv.New(e, "JsonFile load error!")
	}
	jsonGo, e := NewJsonGo(b)
	if e != nil {
		return e
	}
	*j.JsonGo = *jsonGo
	return nil
}
func (j *JsonFile) save(format bool) bool {
	if j._file == "" {
		zlog.Error("Save json to file fail, JsonFile is not Read completed!", j._file)
		panic(errs.ErrSystem.New("JsonFile is not Read completed!"))
		return false
	}
	bytes, err := j.Byte()
	if err != nil {
		zlog.Error(err, j._file)
		panic(err)
		return false
	}
	os.MkdirAll(filepath.Dir(j._file), os.ModePerm)
	fileutil.WriteAndSyncFile(j._file, formatJson(bytes, format), os.ModePerm)
	return true
}
func (j *JsonFile) Save() bool {
	return j.save(apputil.IsPara("formatJson"))
}
func (j *JsonFile) SaveFormat() bool {
	return j.save(true)
}
func (j *JsonFile) SaveUnFormat() bool {
	return j.save(false)
}
