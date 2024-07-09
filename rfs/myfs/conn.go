package myfs

import (
	"bytes"
	"fmt"
	oss "gitee.com/dk83/goutils/rfs/alioss"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Conn struct {
}

var DefaultConn = &Conn{}

func (conn Conn) GetUri(urlPath string, params map[string]interface{}) (*url.URL, error) {
	urlParams := conn.GetURLParams(params)
	if urlParams != "" {
		urlPath = fmt.Sprintf("%s?%s", urlPath, urlParams)
	}
	return url.ParseRequestURI(urlPath)
}

// Do sends request and returns the response
func (conn Conn) do(method, urlPath string, params map[string]interface{}, headers map[string]string, data io.Reader) (*oss.Response, error) {
	uri, err := conn.GetUri(urlPath, params)
	if err != nil {
		return nil, err
	}

	date := time.Now().UTC().Format(http.TimeFormat)
	headers[oss.HTTPHeaderDate] = date
	headers[oss.HTTPHeaderHost] = uri.Host

	resp, err := conn.DoRequest(method, uri, headers, data)
	if err != nil {
		return nil, err
	}

	return conn.handleResponse(resp)
}

func (conn Conn) GetURLParams(params map[string]interface{}) string {
	// Sort
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Serialize
	var buf bytes.Buffer
	for _, k := range keys {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(url.QueryEscape(k))
		if params[k] != nil {
			buf.WriteString("=" + strings.Replace(url.QueryEscape(params[k].(string)), "+", "%20", -1))
		}
	}
	return buf.String()
}

func (conn Conn) DoRequest(method string, uri *url.URL, headers map[string]string, data io.Reader) (*http.Response, error) {
	method = strings.ToUpper(method)
	req := &http.Request{
		Method:     method,
		URL:        uri,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       uri.Host,
	}
	conn.handleBody(req, data)

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	return http.DefaultClient.Do(req)
}

// handleBody handles request body
func (conn Conn) handleBody(req *http.Request, body io.Reader) {
	reader := body
	readerLen, err := oss.GetReaderLen(reader)
	if err == nil {
		req.ContentLength = readerLen
	}
	req.Header.Set(oss.HTTPHeaderContentLength, strconv.FormatInt(req.ContentLength, 10))

	// HTTP body
	rc, ok := reader.(io.ReadCloser)
	if !ok && reader != nil {
		rc = ioutil.NopCloser(reader)
	}
	req.Body = rc
}

// handleResponse handles response
func (conn Conn) handleResponse(resp *http.Response) (*oss.Response, error) {
	var cliCRC uint64
	var srvCRC uint64

	statusCode := resp.StatusCode
	if statusCode >= 400 && statusCode <= 505 {
		// 4xx and 5xx indicate that the operation has error occurred
		var respBody []byte
		respBody, err := readResponseBody(resp)
		if err != nil {
			return nil, err
		}

		if len(respBody) == 0 {
			err = oss.ServiceError{
				StatusCode: statusCode,
				RequestID:  resp.Header.Get(oss.HTTPHeaderOssRequestID),
			}
		} else {
			// Response contains storage service error object, unmarshal
			srvErr, errIn := serviceErr(respBody, resp.StatusCode, resp.Header.Get(oss.HTTPHeaderOssRequestID))
			if errIn != nil { // error unmarshaling the error response
				err = fmt.Errorf("oss: service returned invalid response body, status = %s, RequestId = %s", resp.Status, resp.Header.Get(oss.HTTPHeaderOssRequestID))
			} else {
				err = srvErr
			}
		}

		return &oss.Response{
			StatusCode: resp.StatusCode,
			Headers:    resp.Header,
			Body:       ioutil.NopCloser(bytes.NewReader(respBody)), // restore the body
		}, err
	} else if statusCode >= 300 && statusCode <= 307 {
		// OSS use 3xx, but response has no body
		err := fmt.Errorf("oss: service returned %d,%s", resp.StatusCode, resp.Status)
		return &oss.Response{
			StatusCode: resp.StatusCode,
			Headers:    resp.Header,
			Body:       resp.Body,
		}, err
	}

	// 2xx, successful
	return &oss.Response{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       resp.Body,
		ClientCRC:  cliCRC,
		ServerCRC:  srvCRC,
	}, nil
}
