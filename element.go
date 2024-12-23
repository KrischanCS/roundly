package htmfunc

import "io"

type Element interface {
	RenderHTML(w Writer) error
}

type Attribute func(w Writer) error
type Value func(w Writer) error

type Writer interface {
	io.Writer
	io.ByteWriter
	io.StringWriter
}

type ListItem = Element

type TextDirection string

const (
	LeftToRight TextDirection = "ltr"
	RightToLeft TextDirection = "rtl"
)
