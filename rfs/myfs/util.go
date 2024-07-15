package myfs

import (
	"encoding/json"
	"gitee.com/dk83/goutils/dlog"
	oss "gitee.com/dk83/goutils/rfs/alioss"
	"io"
	"io/ioutil"
	"net/http"
)

func jsonUnmarshal(body io.Reader, v interface{}) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
func changeFsResultToListObjectsResult(result *ListObjectsResult, re *oss.ListObjectsResult) error {
	re.IsTruncated = false
	for _, item := range result.Data {
		re.Objects = append(re.Objects, oss.ObjectProperties{
			Key:          item.Key,
			Size:         item.Size,
			LastModified: item.LastModified,
		})
	}
	return nil
}

func serviceErr(body []byte, statusCode int, requestID string) (oss.ServiceError, error) {
	var jsonRe ObjectsResult
	var storageErr oss.ServiceError
	dlog.Error("body:%s", string(body))
	err := json.Unmarshal(body, &jsonRe)

	if err == nil {
		storageErr.StatusCode = jsonRe.Code
		storageErr.RequestID = requestID
		storageErr.RawMessage = jsonRe.ErrorMsg
		return storageErr, err
	}

	//if err := xml.Unmarshal(body, &storageErr); err != nil {
	//	return storageErr, err
	//}

	storageErr.StatusCode = statusCode
	storageErr.RequestID = requestID
	storageErr.RawMessage = string(body)
	return storageErr, nil
}

func readResponseBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)
	if err == io.EOF {
		err = nil
	}
	return out, err
}
