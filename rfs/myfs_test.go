package rfs

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	oss "gitee.com/dk83/goutils/rfs/alioss"
	"gitee.com/dk83/goutils/rfs/myfs"
	"testing"
)

var bucket = myfs.Init(&myfs.Config{
	Endpoint:        "http://192.168.1.2:3050",
	AccessKeyID:     "accessKeyID",
	AccessKeySecret: "accessKeySecret",
	SecurityToken:   "securityToken",
})

func TestPutObjectFromFile(t *testing.T) {
	objectKey := "test/50.jpg"
	filePath := "d:/50.jpg"
	err := bucket.PutObjectFromFile(objectKey, filePath)
	if err != nil {
		fmt.Println(err)
	}
}
func TestPutObjectFromFile2(t *testing.T) {
	objectKey := "test/lightmap-rgbm0.png"
	filePath := "d:/lightmap-rgbm0.png"
	err := bucket.PutObjectFromFile(objectKey, filePath)
	if err != nil {
		fmt.Println(err)
	}
}
func TestPutObject(t *testing.T) {
	objectKey := "test/test.json"
	data := make(map[string]interface{}, 0)
	data["test"] = "123"
	data["test2"] = 123
	datas, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = bucket.PutObject(objectKey, bytes.NewReader(datas))
	if err != nil {
		fmt.Println(err)
	}
}

func TestCopyObject(t *testing.T) {
	objectKey := "test/test.json"
	destObjectKey := "test/test2.json"
	re, err := bucket.CopyObject(objectKey, destObjectKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(re)
}

func TestCopyObject2(t *testing.T) {
	objectKey := "test"
	destObjectKey := "test2"
	re, err := bucket.CopyObject(objectKey, destObjectKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(re)
}

func TestListObjects(t *testing.T) {
	lsRes, err := bucket.ListObjects(oss.Marker(""), oss.MaxKeys(800), oss.Prefix("test"))
	if err != nil {
		fmt.Println(err)
	}
	for _, object := range lsRes.Objects {
		fmt.Println(object.Key, object.Size, object.LastModified)
	}
}

func TestGetObjectToFile(t *testing.T) {
	objectKey := "test/test.json"
	filePath := "d:/test.json"
	err := bucket.GetObjectToFile(objectKey, filePath)
	if err != nil {
		fmt.Println(err)
	}
}

func TestProcessObject(t *testing.T) {
	objectKey := "test/lightmap-rgbm0.png"

	targetImageName := fmt.Sprintf("%s", objectKey)
	style := fmt.Sprintf("image/format,webp")
	process := fmt.Sprintf("%s|sys/saveas,o_%v", style, base64.URLEncoding.EncodeToString([]byte(targetImageName+".webp")))

	re, err := bucket.ProcessObject(objectKey, process)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(re)
}

//func doPostFile(url string, params map[string]string, headers map[string]string, fileToUpload *os.File) (string, error) {
//	body := &bytes.Buffer{}
//	writer := multipart.NewWriter(body)
//
//	// 添加文件
//	fileWriter, err := writer.CreateFormFile("file", fileToUpload.Name())
//	if err != nil {
//		return "", err
//	}
//	_, err = io.Copy(fileWriter, fileToUpload)
//	if err != nil {
//		return "", err
//	}
//
//	// 设置其他可能的表单字段
//	for key, val := range params {
//		_ = writer.WriteField(key, val)
//	}
//
//	// 关闭writer以完成multipart/form-data的编写
//	err = writer.Close()
//	if err != nil {
//		return "", err
//	}
//
//	// 创建请求
//	req, err := http.NewRequest("POST", url, body)
//	if err != nil {
//		return "", err
//	}
//	req.Header.Set("Content-Type", writer.FormDataContentType())
//
//	// 添加自定义头
//	for key, value := range headers {
//		req.Header.Set(key, value)
//	}
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	// 检查响应状态码
//	if resp.StatusCode != http.StatusOK {
//		log.Printf("File upload failed with status: %s", resp.Status)
//		return "", fmt.Errorf("upload failed with status: %s", resp.Status)
//	}
//
//	// 读取并返回响应体
//	respBody, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return "", err
//	}
//	return string(respBody), nil
//}
//
//
//
//
//
//
//// handleBody handles request body
//func handleBody(req *http.Request, body io.Reader) {
//	reader := body
//	readerLen, err := oss.GetReaderLen(reader)
//	if err == nil {
//		req.ContentLength = readerLen
//	}
//	req.Header.Set(oss.HTTPHeaderContentLength, strconv.FormatInt(req.ContentLength, 10))
//
//	// HTTP body
//	rc, ok := reader.(io.ReadCloser)
//	if !ok && reader != nil {
//		rc = ioutil.NopCloser(reader)
//	}
//	req.Body = rc
//}
//
//
//func doPost(u string, headerMap map[string]string, param map[string]string) (string, error) {
//	// 创建HTTP客户端
//	client := &http.Client{}
//
//	// 构建请求
//	req, err := http.NewRequest("POST", u, nil)
//	if err != nil {
//		return "", err
//	}
//
//	// 设置请求头
//	for k, v := range headerMap {
//		req.Header.Set(k, v)
//	}
//
//	// 添加表单数据
//	var requestBody bytes.Buffer
//	if param != nil {
//		formData := url.Values{}
//		for key, value := range param {
//			formData.Add(key, value)
//		}
//		req.Body = ioutil.NopCloser(&requestBody)
//		req.ContentLength = int64(requestBody.Len())
//		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//		_, err = requestBody.WriteString(formData.Encode())
//		if err != nil {
//			return "", err
//		}
//	}
//
//	// 发送请求
//	resp, err := client.Do(req)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	// 读取响应体
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return "", err
//	}
//
//	return string(body), nil
//}
