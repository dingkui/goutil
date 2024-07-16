package native

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"
)

type f byte

const FileUtil f = iota

//获取文件修改时间 返回unix时间戳
func (x f) GetFileModTime(path string) (time.Time, error) {
	f, err := os.Open(path)
	if err != nil {
		return time.Now(), err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return time.Now(), err
	}
	return fi.ModTime(), nil
}

func (x f) WriteAndSyncFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err == nil {
		err = f.Sync()
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

func (x f) Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (x f) IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

func (x f) MakeUnique(file_path string) string {
	idx := 2
	root := filepath.Dir(file_path)
	ext := filepath.Ext(file_path)
	filename := filepath.Base(file_path)
	root = filepath.Join(root, strings.TrimSuffix(filename, ext))
	//root += strings.TrimSuffix(filename,ext)

	for x.Exists(file_path) {
		file_path = fmt.Sprintf("%s-%d%s", root, idx, ext)
		idx += 1
	}

	return file_path
}

func (x f) Splitext(file_path string) (name string, ext string) {
	ext = filepath.Ext(file_path)
	name = strings.TrimSuffix(file_path, ext)
	return name, ext
}

func (x f) CopyFile(src, dest string) error {
	if dest == src {
		return nil
	}
	os.MkdirAll(filepath.Dir(dest), os.ModePerm)
	f, err := os.Open(src)
	if err != nil {
		return err
	}

	defer f.Close()

	f2, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f2.Close()

	io.Copy(f2, f)
	return nil
}

/**
 * 拷贝文件夹,同时拷贝文件夹中的文件
 * @param srcPath  		需要拷贝的文件夹路径: D:/test
 * @param destPath		拷贝到的位置: D:/backup/
 */
func (x f) CopyDir(srcPath string, destPath string) error {
	//检测目录正确性
	if srcInfo, err := os.Stat(srcPath); err != nil {
		return err
	} else {
		if !srcInfo.IsDir() {
			e := errors.New("srcPath不是一个正确的目录！")
			return e
		}
	}

	err := os.MkdirAll(destPath, os.ModePerm)
	if err != nil {
		return err
	}
	//加上拷贝时间:不用可以去掉
	//destPath = destPath + "_" + time.Now().Format("20060102150405")

	err = filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			path := strings.Replace(path, "/", "\\", -1)
			destNewPath := strings.Replace(path, srcPath, destPath, -1)
			//zlog.Debug("复制文件:" + path + " 到 " + destNewPath)
			x.CopyFile(path, destNewPath)
		}
		return nil
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
	return err
}

//检测文件夹路径时候存在
func (x f) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (x f) RenameFile(oldFile, newFile string) error {
	if x.Exists(oldFile) {
		err := os.Rename(oldFile, newFile)
		return err
	}
	return nil
}

func (x f) Read(path string) ([]byte, error) {
	by, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}
	return by, nil
}

func (x f) GetFileCreateTime(path string) int64 {
	osType := runtime.GOOS
	fileInfo, _ := os.Stat(path)
	if osType == "windows" {
		wFileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
		tNanSeconds := wFileSys.CreationTime.Nanoseconds() /// 返回的是纳秒
		tSec := tNanSeconds / 1e9                          ///秒
		return tSec
	}
	return time.Now().Unix()
}