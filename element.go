package htmfunc

import "io"

type Element interface {
	RenderHtml(w Writer) error
}

type Attribute interface {
	RenderAttribute(w Writer) error
}

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
