package dhttp

import (
	"net/http"
)

const (
	deleteObjectsQuiet = "delete-objects-quiet"
	routineNum         = "x-routine-num"
	checkpointConfig   = "x-cp-config"
	initCRC64          = "init-crc64"
	progressListener   = "x-progress-listener"
	storageClass       = "storage-class"
	responseHeader     = "x-response-header"
	redundancyType     = "redundancy-type"
)

type DataRedundancyType string

const (
	// RedundancyLRS Local redundancy, default value
	RedundancyLRS DataRedundancyType = "LRS"
	// RedundancyZRS Same city redundancy
	RedundancyZRS DataRedundancyType = "ZRS"
)

// StorageClassType bucket storage type
type StorageClassType string

const (
	// StorageStandard standard
	StorageStandard StorageClassType = "Standard"
	// StorageIA infrequent access
	StorageIA StorageClassType = "IA"
	// StorageArchive archive
	StorageArchive StorageClassType = "Archive"
	// StorageColdArchive cold archive
	StorageColdArchive StorageClassType = "ColdArchive"
)

// Checkpoint configuration
type cpConfig struct {
	IsEnable bool
	FilePath string
	DirPath  string
}

// DeleteObjectsQuiet false:DeleteObjects in verbose mode; true:DeleteObjects in quite mode. Default is false.
func (ops *Options) DeleteObjectsQuiet(isQuiet bool) {
	ops.AddArg(deleteObjectsQuiet, isQuiet)
}

// StorageClass bucket storage class
func (ops *Options) StorageClass(value StorageClassType) {
	ops.AddArg(storageClass, value)
}

// RedundancyType bucket data redundancy type
func (ops *Options) RedundancyType(value DataRedundancyType) {
	ops.AddArg(redundancyType, value)
}

// Checkpoint sets the isEnable flag and checkpoint file path for DownloadFile/UploadFile.
func (ops *Options) Checkpoint(isEnable bool, filePath string) {
	ops.AddArg(checkpointConfig, &cpConfig{IsEnable: isEnable, FilePath: filePath})
}

// CheckpointDir sets the isEnable flag and checkpoint dir path for DownloadFile/UploadFile.
func (ops *Options) CheckpointDir(isEnable bool, dirPath string) {
	ops.AddArg(checkpointConfig, &cpConfig{IsEnable: isEnable, DirPath: dirPath})
}

// Routines DownloadFile/UploadFile routine count
func (ops *Options) Routines(n int) {
	ops.AddArg(routineNum, n)
}

// InitCRC Init AppendObject CRC
func (ops *Options) InitCRC(initCRC uint64) {
	ops.AddArg(initCRC64, initCRC)
}

// GetResponseHeader for get response http header
func (ops *Options) GetResponseHeader(respHeader *http.Header) {
	ops.AddArg(responseHeader, respHeader)
}

// Progress set progress listener
func (ops *Options) Progress(listener ProgressListener) {
	ops.AddArg(progressListener, listener)
}

// Progress set progress listener
func (ops *Options) GetProgressListener() ProgressListener {
	value := ops.GetArg(progressListener)
	if value != nil {
		return value.(ProgressListener)
	}
	return nil
}
