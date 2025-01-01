package htmfunc

import "io"

type Writer interface {
	io.Writer
	io.ByteWriter
	io.StringWriter
}

type TextDirection string

const (
	LeftToRight TextDirection = "ltr"
	RightToLeft TextDirection = "rtl"
)
