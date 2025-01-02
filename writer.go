package htmfunc

import (
	"bytes"
	"io"
)

func NewWriter(initialLength int) *bytes.Buffer {
	return bytes.NewBuffer(make([]byte, 0, initialLength))
}

type Writer interface {
	io.Writer
	io.ByteWriter
	io.StringWriter
}

// TODO this needs to go somewhere else

type TextDirection string

const (
	LeftToRight TextDirection = "ltr"
	RightToLeft TextDirection = "rtl"
)
