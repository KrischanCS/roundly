package htmfunc

import "github.com/valyala/bytebufferpool"

type Element interface {
	RenderHTML(w Writer) error
}

type Attribute func(w Writer) error
type Value func(w Writer) error

type Writer = *bytebufferpool.ByteBuffer

type ListItem = Element

type TextDirection string

const (
	LeftToRight TextDirection = "ltr"
	RightToLeft TextDirection = "rtl"
)
