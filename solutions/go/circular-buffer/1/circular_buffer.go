package circular

import (
	"errors"
)

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.

type Buffer struct {
	data     []byte
	size     int
	readPos  int
	writePos int
	isFull   bool
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		data: make([]byte, size),
		size: size,
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.readPos == b.writePos && !b.isFull {
		return 0, errors.New("Buffer is Empty")
	}

	// Read the byte
	value := b.data[b.readPos]
	b.readPos = (b.readPos + 1) % b.size
	b.isFull = false

	return value, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.isFull {
		return errors.New("Buffer is Full")
	}
	b.data[b.writePos] = c
	b.writePos = (b.writePos + 1) % b.size

	if b.writePos == b.readPos {
		b.isFull = true
	}
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	b.data[b.writePos] = c
	b.writePos = (b.writePos + 1) % b.size

	if b.isFull {
		b.readPos = (b.readPos + 1) % b.size
	}

	if b.writePos == b.readPos {
		b.isFull = true
	}
}

func (b *Buffer) Reset() {
	b.readPos = 0
	b.writePos = 0
	b.isFull = false
}
