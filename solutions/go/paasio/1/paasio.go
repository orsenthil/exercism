package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.

type readCounter struct {
	reader    io.Reader
	bytesRead int64
	reads     int
	mu        sync.Mutex
}

type writeCounter struct {
	writer     io.Writer
	bytesWrite int64
	writes     int
	mu         sync.Mutex
}

type readWriteCounter struct {
	*readCounter
	*writeCounter
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		readCounter:  &readCounter{reader: readwriter},
		writeCounter: &writeCounter{writer: readwriter},
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	n, err := rc.reader.Read(p)

	rc.mu.Lock()
	rc.reads++
	rc.bytesRead += int64(n)
	rc.mu.Unlock()

	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return rc.bytesRead, rc.reads
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.writer.Write(p)

	wc.mu.Lock()
	wc.bytesWrite += int64(n)
	wc.writes++
	wc.mu.Unlock()

	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	return wc.bytesWrite, wc.writes
}
