package rfs

import (
	oss "gitee.com/dk83/goutils/rfs/alioss"
	"io"
)

type IBucket interface {
	ListObjects(options ...oss.Option) (oss.ListObjectsResult, error)
	GetObjectToFile(objectKey, filePath string, options ...oss.Option) error
	PutObjectFromFile(objectKey, filePath string, options ...oss.Option) error
	PutObject(objectKey string, reader io.Reader, options ...oss.Option) error
	CopyObject(srcObjectKey, destObjectKey string, options ...oss.Option) (oss.CopyObjectResult, error)
	ProcessObject(objectKey string, process string, options ...oss.Option) (oss.ProcessObjectResult, error)
}
