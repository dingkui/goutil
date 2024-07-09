package fileutil

import (
	"errors"
	"fmt"
	"gitee.com/dk83/goutils/zlog"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func WriteAndSyncFile(filename string, data []byte, perm os.FileMode) error {
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

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

func MakeUnique(file_path string) string {
	idx := 2
	root := filepath.Dir(file_path)
	ext := filepath.Ext(file_path)
	filename := filepath.Base(file_path)
	root = filepath.Join(root, strings.TrimSuffix(filename, ext))
	//root += strings.TrimSuffix(filename,ext)

	for Exists(file_path) {
		file_path = fmt.Sprintf("%s-%d%s", root, idx, ext)
		idx += 1
	}

	return file_path
}

func Splitext(file_path string) (name string, ext string) {
	ext = filepath.Ext(file_path)
	name = strings.TrimSuffix(file_path, ext)
	return name, ext
}

func CopyFile(src, dest string) bool {
	if dest == src {
		return true
	}
	os.MkdirAll(filepath.Dir(dest), os.ModePerm)
	f, err := os.Open(src)
	if err != nil {
		zlog.Error(err)
		return false
	}

	defer f.Close()

	f2, err := os.Create(dest)
	if err != nil {
		zlog.Error(err)
		return false
	}
	defer f2.Close()

	io.Copy(f2, f)
	return true
}

/**
 * 拷贝文件夹,同时拷贝文件夹中的文件
 * @param srcPath  		需要拷贝的文件夹路径: D:/test
 * @param destPath		拷贝到的位置: D:/backup/
 */
func CopyDir(srcPath string, destPath string) error {
	//检测目录正确性
	if srcInfo, err := os.Stat(srcPath); err != nil {
		zlog.Error(err)
		return err
	} else {
		if !srcInfo.IsDir() {
			e := errors.New("srcPath不是一个正确的目录！")
			zlog.Error(e)
			return e
		}
	}

	err := os.MkdirAll(destPath, os.ModePerm)
	if err != nil {
		zlog.Error(err)
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
			CopyFile(path, destNewPath)
		}
		return nil
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
	return err
}

//检测文件夹路径时候存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func RenameFile(oldFile, newFile string) error {
	if Exists(oldFile) {
		err := os.Rename(oldFile, newFile)
		return err
	}
	return nil
}

func Read(path string) []byte {
	by, e := ioutil.ReadFile(path)
	if e != nil {
		zlog.Error(e)
	}
	return by
}

func GetFileCreateTime(path string) int64 {
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
