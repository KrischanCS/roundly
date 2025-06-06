package roundly

import (
	"bytes"
	"io"
)

// NewWriter creates a bytes.Buffer with a default initial capacity. It can be used as a [Writer].
func NewWriter() *bytes.Buffer {
	return bytes.NewBuffer(make([]byte, 0, 256))
}

// NewWriterSize creates a bytes.Buffer the given initial size. It can be used as a [Writer].
func NewWriterSize(size int) *bytes.Buffer {
	return bytes.NewBuffer(make([]byte, 0, size))
}

// Writer is the interface required by the [Element] and [Attribute]
// to render themselves.
type Writer interface {
	io.Writer
	io.ByteWriter
	io.StringWriter
}
