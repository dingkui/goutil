package myfs

import (
	"encoding/base64"
	"errors"
	oss "gitee.com/dk83/goutils/rfs/alioss"
	"io"
	"os"
	"strings"
	"time"
)

// 派生结构体，通过嵌入 Bucket 来继承其属性和方法
type MyFs struct {
	oss.Bucket // 嵌入 Base 结构体
	Config     *Config
}

// Config defines oss configuration
type Config struct {
	Endpoint        string // OSS endpoint
	AccessKeyID     string // AccessId
	AccessKeySecret string // AccessKey
	SecurityToken   string // AccessKey
}

var myFs *MyFs

func Init(config *Config) *MyFs {
	if myFs == nil {
		myFs = &MyFs{
			Config: config,
		}
	} else {
		myFs.Config = config
	}
	return myFs
}

type ObjectsResult struct {
	Code     int         `json:"code"`
	Rid      string      `json:"rid"`
	Msg      string      `json:"msg"`
	ErrorMsg string      `json:"errorMsg"`
	Data     interface{} `json:"data"`
}

type ListObjectsResult struct {
	ObjectsResult
	Data []ObjectProperties `json:"data"` // Max keys to return
}

type CopyObjectResult struct {
	ObjectsResult
	Data int64 `json:"data"` // Max keys to return
}

type ObjectProperties struct {
	Key          string    `json:"key"`          // Object key
	Size         int64     `json:"size"`         // Object size
	LastModified time.Time `json:"lastModified"` // Object last modified time
}

func (bucket MyFs) ListObjects(options ...oss.Option) (oss.ListObjectsResult, error) {
	var out oss.ListObjectsResult
	var fsOut ListObjectsResult

	params, err := oss.GetRawParams(options)
	if err != nil {
		return out, err
	}
	params["dirPath"] = params["prefix"]

	resp, err := bucket.PostToMyfs("listObject", params, options)
	if err != nil {
		return out, err
	}
	defer resp.Body.Close()
	err = jsonUnmarshal(resp.Body, &fsOut)
	if err != nil {
		return out, err
	}
	if fsOut.Code != 200 {
		return out, errors.New(fsOut.ErrorMsg)
	}

	err = changeFsResultToListObjectsResult(&fsOut, &out)
	return out, err
}
func (bucket MyFs) PutObjectFromFile(objectKey, filePath string, options ...oss.Option) error {
	fd, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fd.Close()

	return bucket.PutObject(objectKey, fd, options...)
}
func (bucket MyFs) PutObject(objectKey string, reader io.Reader, options ...oss.Option) error {
	options = append(options, oss.ContentType("application/octet-stream"))
	base64Fp := base64.StdEncoding.EncodeToString([]byte(objectKey))
	options = append(options, oss.SetHeader("my-filepath", base64Fp))

	resp, err := bucket.Do("POST", "uploadStream", nil, reader, options)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var fsOut ObjectsResult
	err = jsonUnmarshal(resp.Body, &fsOut)
	if fsOut.Code != 200 {
		return errors.New(fsOut.ErrorMsg)
	}

	return err
}
func (bucket MyFs) CopyObject(srcObjectKey, destObjectKey string, options ...oss.Option) (oss.CopyObjectResult, error) {
	var out oss.CopyObjectResult

	options = append(options, oss.ContentType("application/x-www-form-urlencoded"))
	params := map[string]interface{}{}
	params["oldDirPath"] = srcObjectKey
	params["newDirPath"] = destObjectKey

	resp, err := bucket.PostToMyfs("copyObject", params, options)
	if err != nil {
		return out, err
	}
	defer resp.Body.Close()

	var fsOut ObjectsResult
	err = jsonUnmarshal(resp.Body, &fsOut)
	if fsOut.Code != 200 {
		return out, errors.New(fsOut.ErrorMsg)
	}

	return out, err
}
func (bucket MyFs) ProcessObject(objectKey string, process string, options ...oss.Option) (oss.ProcessObjectResult, error) {
	var out oss.ProcessObjectResult
	params, _ := oss.GetRawParams(options)
	processSplit := strings.Split(process, "|")
	xProcess := processSplit[0]
	xSaves := processSplit[1]
	processSavePrifix := "sys/saveas,o_"
	if strings.Index(xProcess, "image/") != 0 || strings.Index(xSaves, processSavePrifix) != 0 {
		return out, errors.New("process or save format error")
	}

	xSavesPath := xSaves[len(processSavePrifix):]
	if xSavesPath == "" {
		return out, errors.New("process or save format error")
	}

	params["x-oss-process"] = xProcess
	params["dirPath"] = objectKey
	params["newDirPath"] = xSavesPath
	resp, err := bucket.PostToMyfs("processObject", params, options)
	if err != nil {
		return out, err
	}
	defer resp.Body.Close()

	var fsOut ObjectsResult
	err = jsonUnmarshal(resp.Body, &fsOut)
	if fsOut.Code != 200 {
		return out, errors.New(fsOut.ErrorMsg)
	}

	//err = jsonUnmarshal(resp.Body, &out)
	return out, err
}

func (bucket MyFs) GetObjectToFile(objectKey, filePath string, options ...oss.Option) error {
	tempFilePath := filePath + oss.TempFileSuffix

	// Calls the API to actually download the object. Returns the result instance.
	result, err := bucket.DoGetObject(&oss.GetObjectRequest{objectKey}, options)
	if err != nil {
		return err
	}
	defer result.Response.Close()

	// If the local file does not exist, create a new one. If it exists, overwrite it.
	fd, err := os.OpenFile(tempFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, oss.FilePermMode)
	if err != nil {
		return err
	}

	// Copy the data to the local file path.
	_, err = io.Copy(fd, result.Response.Body)
	fd.Close()
	if err != nil {
		return err
	}
	return os.Rename(tempFilePath, filePath)
}
