package httputil

import (
	"strconv"
)

// ResponseContentType is an option to set response-content-type param
func (ops *Options) ResponseContentType(value string) {
	ops.AddParam("response-content-type", value)
}

// ResponseContentLanguage is an option to set response-content-language param
func (ops *Options) ResponseContentLanguage(value string) {
	ops.AddParam("response-content-language", value)
}

// ResponseExpires is an option to set response-expires param
func (ops *Options) ResponseExpires(value string) {
	ops.AddParam("response-expires", value)
}

// ResponseCacheControl is an option to set response-cache-control param
func (ops *Options) ResponseCacheControl(value string) {
	ops.AddParam("response-cache-control", value)
}

// ResponseContentDisposition is an option to set response-content-disposition param
func (ops *Options) ResponseContentDisposition(value string) {
	ops.AddParam("response-content-disposition", value)
}

// ResponseContentEncoding is an option to set response-content-encoding param
func (ops *Options) ResponseContentEncoding(value string) {
	ops.AddParam("response-content-encoding", value)
}

// Process is an option to set x-oss-process param
func (ops *Options) Process(value string) {
	ops.AddParam("x-oss-process", value)
}

// TrafficLimitParam is a option to set x-oss-traffic-limit
func (ops *Options) TrafficLimitParam(value int64) {
	ops.AddParam("x-oss-traffic-limit", strconv.FormatInt(value, 10))
}
