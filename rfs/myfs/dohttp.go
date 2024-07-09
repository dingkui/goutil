package myfs

import (
	"fmt"
	oss "gitee.com/dk83/goutils/rfs/alioss"
	"hash"
	"io"
	"net/url"
	"strconv"
	"strings"
)

func (bucket MyFs) Do(method, path string, urlParams map[string]interface{}, data io.Reader, options []oss.Option) (*oss.Response, error) {
	headers := make(map[string]string)
	err := oss.HandleOptions(headers, options)
	if err != nil {
		return nil, err
	}

	urlPath := fmt.Sprintf("%s/%s", bucket.Config.Endpoint, path)
	return DefaultConn.do(method, urlPath, urlParams, headers, data)
}

func (bucket MyFs) PostToMyfs(path string, params map[string]interface{}, options []oss.Option) (*oss.Response, error) {
	options = append(options,
		oss.ContentType("application/x-www-form-urlencoded"),
		oss.SetHeader("sts_token", bucket.Config.SecurityToken),
	)
	formData := url.Values{}
	for k, v := range params {
		if v != nil {
			formData.Add(k, v.(string))
		}
	}
	data := strings.NewReader(formData.Encode())
	resp, err := bucket.Do("POST", path, nil, data, options)
	return resp, err
}

func (bucket MyFs) DoGetObject(request *oss.GetObjectRequest, options []oss.Option) (*oss.GetObjectResult, error) {
	params, _ := oss.GetRawParams(options)
	resp, err := bucket.Do("GET", request.ObjectKey, params, nil, options)
	if err != nil {
		return nil, err
	}

	result := &oss.GetObjectResult{
		Response: resp,
	}

	// CRC
	var crcCalc hash.Hash64

	// Progress
	listener := oss.GetProgressListener(options)

	contentLen, _ := strconv.ParseInt(resp.Headers.Get(oss.HTTPHeaderContentLength), 10, 64)
	resp.Body = oss.TeeReader(resp.Body, crcCalc, contentLen, listener, nil)

	return result, nil
}
