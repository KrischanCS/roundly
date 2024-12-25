package htmfunc

import "io"

type ElementRenderer interface {
	RenderElement(w Writer) error
}

type AttributeRenderer interface {
	RenderAttribute(w Writer) error
}

type ValueRenderer interface {
	RenderValue(w Writer) error
}

type Writer interface {
	io.Writer
	io.ByteWriter
	io.StringWriter
}

type ListItem = ElementRenderer

type TextDirection string

const (
	LeftToRight TextDirection = "ltr"
	RightToLeft TextDirection = "rtl"
)
