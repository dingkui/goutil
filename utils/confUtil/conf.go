package confUtil

import (
	"encoding/json"
	"github.com/dingkui/goutil/errs"
	"github.com/dingkui/goutil/utils/fileUtil"
	"github.com/dingkui/goutil/utils/stringUtil"
	"github.com/dingkui/goutil/utils/valUtil"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Conf struct {
	file   string
	data   interface{}
	format bool
}

func NewConf(file string, data interface{}, format bool) *Conf {
	if file == "" {
		panic("conf file is empty!")
	}
	if data == nil {
		panic("conf data is nil!")
	}
	cf := &Conf{
		file:   file,
		data:   data,
		format: format,
	}
	var err error
	if !fileUtil.Exists(file) {
		err = cf.SaveFormat(format)
	} else {
		err = cf.ReLoad()
	}
	if err != nil {
		panic(err)
	}
	return cf
}
func (j *Conf) ReLoad() error {
	b, e := ioutil.ReadFile(j.file)
	if e != nil {
		return errs.ErrEnv.New(e, "conf load error:%s", j.file)
	}
	e = json.Unmarshal(b, j.data)
	if e != nil {
		return errs.ErrTargetType.New(e, "parse error:%s → [%T]", string(b), j.data)
	}
	return nil
}
func (j *Conf) Save() error {
	return j.SaveFormat(j.format)
}
func (j *Conf) SaveFormat(format bool) error {
	bytes, err := valUtil.Bytes(j.data)
	if err != nil {
		return errs.ErrSystem.New(err, "conf save error:%s", j.file)
	}
	os.MkdirAll(filepath.Dir(j.file), os.ModePerm)
	fileUtil.WriteAndSyncFile(j.file, stringUtil.FormatJson(bytes, format), os.ModePerm)
	return nil
}
