package roundly

import (
	"bufio"
	"bytes"
	"io"
	"sync"
)

type pooledBuffer struct {
	*bufio.Writer
}

// NewWriter creates a bytes.Buffer with a default initial capacity. It can be used as a [Writer].
func NewWriter() *bytes.Buffer {
	return bytes.NewBuffer(make([]byte, 0, 4096))
}

// placeholder is a buffer never actually written to. bufio.NewWriter can not be called without providing
// a target writer, but the target writer is not available in pool.New, so this is used.
//
// When returning the buffer to the pool, it's also put via Reset, so that the buffer doesn't hold a
// reference to the given writer, so that the gc can release it.
var placeholder = bytes.NewBuffer(make([]byte, 0, 0))

var bufPool = sync.Pool{
	New: func() interface{} {
		return pooledBuffer{bufio.NewWriter(placeholder)}
	},
}

// WrapWriter takes an io.Writer and returns a roundly.Writer forwarding to it. If w implements
// roundly.Writer, it's directly returned.
//
// The writers are managed in a sync.Pool, a writer obtained using this method must return it via
// PutWriter after use.
func WrapWriter(w io.Writer) Writer {
	if w, ok := w.(Writer); ok {
		return w
	}

	buf := bufPool.Get().(pooledBuffer)
	buf.Reset(w)
	return buf
}

// PutWriter puts w a pool of buffers, which can be obtained by WrapWriter.
//
// If WrapWriter returned the original writer without wrapping it, the given writer will also not
// be put to the pool.
//
// The writer is resetted before it's put to the pool, so it's not holding a reference to the
// wrapped writer, so that the gc can release it.
func PutWriter(w Writer) error {
	b, ok := w.(pooledBuffer)
	if !ok {
		return nil
	}

	err := b.Flush()

	b.Reset(placeholder)
	bufPool.Put(w)

	return err
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
