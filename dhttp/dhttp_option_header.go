package dhttp

import (
	"fmt"
	"gitee.com/dk83/goutils/djson"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	HTTPHeaderAcceptEncoding     = "Accept-Encoding"
	HTTPHeaderAuthorization      = "Authorization"
	HTTPHeaderCacheControl       = "Cache-Control"
	HTTPHeaderContentDisposition = "Content-Disposition"
	HTTPHeaderContentEncoding    = "Content-Encoding"
	HTTPHeaderContentLength      = "Content-Length"
	HTTPHeaderContentMD5         = "Content-MD5"
	HTTPHeaderContentType        = "Content-Type"
	HTTPHeaderContentLanguage    = "Content-Language"
	HTTPHeaderDate               = "Date"
	HTTPHeaderEtag               = "ETag"
	HTTPHeaderExpires            = "Expires"
	HTTPHeaderHost               = "Host"
	HTTPHeaderLastModified       = "Last-Modified"
	HTTPHeaderRange              = "Range"
	HTTPHeaderLocation           = "Location"
	HTTPHeaderOrigin             = "Origin"
	HTTPHeaderServer             = "Server"
	HTTPHeaderUserAgent          = "User-Agent"
	HTTPHeaderIfModifiedSince    = "If-Modified-Since"
	HTTPHeaderIfUnmodifiedSince  = "If-Unmodified-Since"
	HTTPHeaderIfMatch            = "If-Match"
	HTTPHeaderIfNoneMatch        = "If-None-Match"
	HTTPHeaderACReqMethod        = "Access-Control-Request-Method"
	HTTPHeaderACReqHeaders       = "Access-Control-Request-Headers"
)

type ContentType string

const (
	ContentTypeForm   ContentType = "application/x-www-form-urlencoded;charset=utf-8"
	ContentTypeStream ContentType = "application/octet-stream"
	ContentTypeJson   ContentType = "application/json;charset=utf-8"
	ContentTypeXml    ContentType = "application/xml"
	ContentTypeText   ContentType = "text/plain"
	ContentTypeHtml   ContentType = "text/html"
	ContentTypeJpeg   ContentType = "image/jpeg"
	ContentTypeMpeg   ContentType = "audio/mpeg"
)

func (ops *Options) AddHeaders(headers *djson.JsonGo) {
	if headers != nil {
		m, _ := headers.Map()
		for key, jsonGo := range m {
			ops.AddHeader(key, jsonGo.StrN(""))
		}
	}
}
func (ops *Options) AddHeader(key string, value string) {
	ops.addOption(key, value, optionHeader)
}
func (ops *Options) GetHeaders() map[string]string {
	headers := make(map[string]string)
	for _, option := range ops.options {
		if option.ot == optionHeader {
			headers[option.key] = option.value.(string)
		}
	}
	return headers
}
func (ops *Options) GetHeader(key string) string {
	for _, option := range ops.options {
		if option.ot == optionArg && option.key == key {
			return option.value.(string)
		}
	}
	return ""
}

// ContentType is an option to set Content-Type header
func (ops *Options) ContentType(value ContentType) {
	ops.AddHeader(HTTPHeaderContentType, fmt.Sprintf("%s", value))
}

// ContentLength is an option to set Content-Length header
func (ops *Options) ContentLength(length int64) {
	ops.AddHeader(HTTPHeaderContentLength, strconv.FormatInt(length, 10))
}
func (ops *Options) GetContentLength() int64 {
	length := ops.GetHeader(HTTPHeaderContentLength)
	if length != "" {
		return 0
	}
	num, _ := strconv.ParseInt(length, 10, 64)
	return num
}

// CacheControl is an option to set Cache-Control header
func (ops *Options) CacheControl(value string) {
	ops.AddHeader(HTTPHeaderCacheControl, value)
}

// ContentDisposition is an option to set Content-Disposition header
func (ops *Options) ContentDisposition(value string) {
	ops.AddHeader(HTTPHeaderContentDisposition, value)
}

// ContentEncoding is an option to set Content-Encoding header
func (ops *Options) ContentEncoding(value string) {
	ops.AddHeader(HTTPHeaderContentEncoding, value)
}

// ContentLanguage is an option to set Content-Language header
func (ops *Options) ContentLanguage(value string) {
	ops.AddHeader(HTTPHeaderContentLanguage, value)
}

// ContentMD5 is an option to set Content-MD5 header
func (ops *Options) ContentMD5(value string) {
	ops.AddHeader(HTTPHeaderContentMD5, value)
}

// Expires is an option to set Expires header
func (ops *Options) Expires(t time.Time) {
	ops.AddHeader(HTTPHeaderExpires, t.Format(http.TimeFormat))
}

// Range is an option to set Range header, [start, end]
func (ops *Options) Range(start, end int64) {
	ops.AddHeader(HTTPHeaderRange, fmt.Sprintf("bytes=%d-%d", start, end))
}

// NormalizedRange is an option to set Range header, such as 1024-2048 or 1024- or -2048
func (ops *Options) NormalizedRange(nr string) {
	ops.AddHeader(HTTPHeaderRange, fmt.Sprintf("bytes=%s", strings.TrimSpace(nr)))
}

// AcceptEncoding is an option to set Accept-Encoding header
func (ops *Options) AcceptEncoding(value string) {
	ops.AddHeader(HTTPHeaderAcceptEncoding, value)
}

// IfModifiedSince is an option to set If-Modified-Since header
func (ops *Options) IfModifiedSince(t time.Time) {
	ops.AddHeader(HTTPHeaderIfModifiedSince, t.Format(http.TimeFormat))
}

// IfUnmodifiedSince is an option to set If-Unmodified-Since header
func (ops *Options) IfUnmodifiedSince(t time.Time) {
	ops.AddHeader(HTTPHeaderIfUnmodifiedSince, t.Format(http.TimeFormat))
}

// IfMatch is an option to set If-Match header
func (ops *Options) IfMatch(value string) {
	ops.AddHeader(HTTPHeaderIfMatch, value)
}

// IfNoneMatch is an option to set IfNoneMatch header
func (ops *Options) IfNoneMatch(value string) {
	ops.AddHeader(HTTPHeaderIfNoneMatch, value)
}

// Origin is an option to set Origin header
func (ops *Options) Origin(value string) {
	ops.AddHeader(HTTPHeaderOrigin, value)
}

// ACReqMethod is an option to set Access-Control-Request-Method header
func (ops *Options) ACReqMethod(value string) {
	ops.AddHeader(HTTPHeaderACReqMethod, value)
}

// ACReqHeaders is an option to set Access-Control-Request-Headers header
func (ops *Options) ACReqHeaders(value string) {
	ops.AddHeader(HTTPHeaderACReqHeaders, value)
}

// UserAgentHeader is an option to set HTTPHeaderUserAgent
func (ops *Options) UserAgentHeader(ua string) {
	ops.AddHeader(HTTPHeaderUserAgent, ua)
}
