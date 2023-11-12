package htmfunc

import (
	"io"
)

type Component func(w Writer) error

type Attribute func(w Writer) error

type Writer interface {
	io.Writer
	io.ByteWriter
	io.StringWriter
}

type ListItem = Component

type TextDirection string

const (
	LeftToRight TextDirection = "ltr"
	RightToLeft TextDirection = "rtl"
)
