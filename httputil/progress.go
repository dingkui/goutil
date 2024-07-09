package httputil

import "io"

type ProgressEventType int

const (
	ReqStartedEvent ProgressEventType = 1 + iota
	ReqDataEvent
	ReqCompletedEvent
	ReqFailedEvent
	ResStartedEvent
	ResDataEvent
	ResCompletedEvent
	ResFailedEvent
)

type progressSta struct {
	consumedBytes int64
	totalBytes    int64
}

// ProgressEvent defines progress event
type ProgressEvent struct {
	progressSta
	RwBytes   int64
	EventType ProgressEventType
}

// ProgressListener listens progress change
type ProgressListener interface {
	ProgressChanged(event *ProgressEvent)
}

// ProgressListener listens progress change
type progress struct {
	listener ProgressListener
	sta      *progressSta
}

func (t ProgressEventType) mkEvent(sta *progressSta, RwBytes int64) *ProgressEvent {
	event := ProgressEvent{
		RwBytes:   RwBytes,
		EventType: t}
	event.consumedBytes = sta.consumedBytes
	event.totalBytes = sta.totalBytes
	return &event
}
func (t progress) reqStarted() {
	if t.listener != nil {
		t.sta.consumedBytes = 0
		t.listener.ProgressChanged(ReqStartedEvent.mkEvent(t.sta, 0))
	}
}
func (t progress) reqData(rwBytes int64) {
	if t.listener != nil {
		t.sta.consumedBytes += rwBytes
		t.listener.ProgressChanged(ReqDataEvent.mkEvent(t.sta, rwBytes))
	}
}
func (t progress) reqFailed() {
	if t.listener != nil {
		t.listener.ProgressChanged(ReqFailedEvent.mkEvent(t.sta, 0))
	}
}
func (t progress) reqCompleted() {
	if t.listener != nil {
		t.listener.ProgressChanged(ReqCompletedEvent.mkEvent(t.sta, 0))
	}
}

func (t progress) resStarted() {
	if t.listener != nil {
		t.sta.consumedBytes = 0
		t.listener.ProgressChanged(ResStartedEvent.mkEvent(t.sta, 0))
	}
}
func (t progress) resData(rwBytes int64) {
	if t.listener != nil {
		t.sta.consumedBytes += rwBytes
		t.listener.ProgressChanged(ResDataEvent.mkEvent(t.sta, rwBytes))
	}
}
func (t progress) resEnd() {
	if t.listener != nil {
		if t.sta.consumedBytes == t.sta.totalBytes {
			t.listener.ProgressChanged(ResCompletedEvent.mkEvent(t.sta, 0))
		} else {
			t.listener.ProgressChanged(ResFailedEvent.mkEvent(t.sta, 0))
		}
	}
}

type teeReader struct {
	reader   io.Reader
	writer   io.Writer
	progress *progress
}

func (t *teeReader) Read(p []byte) (n int, err error) {
	n, err = t.reader.Read(p)

	// Read encountered error
	if err != nil && err != io.EOF {
		t.progress.reqFailed()
	}

	if n > 0 {
		// CRC
		if t.writer != nil {
			if n, err := t.writer.Write(p[:n]); err != nil {
				return n, err
			}
		}
		// Progress
		if t.progress != nil {
			t.progress.reqData(int64(n))
		}
	}

	return
}
func (t *teeReader) Close() error {
	if rc, ok := t.reader.(io.ReadCloser); ok {
		return rc.Close()
	}
	return nil
}

// TeeReader returns a Reader that writes to w what it reads from r.
// All reads from r performed through it are matched with
// corresponding writes to w.  There is no internal buffering -
// the write must complete before the read completes.
// Any error encountered while writing is reported as a read error.
func newTeeReader(reader io.Reader, writer io.Writer, listener ProgressListener, totalBytes int64) (io.ReadCloser, *progress) {
	progress := &progress{
		listener: listener,
		sta: &progressSta{
			consumedBytes: 0,
			totalBytes:    totalBytes,
		},
	}
	return &teeReader{
		reader:   reader,
		writer:   writer,
		progress: progress,
	}, progress
}
