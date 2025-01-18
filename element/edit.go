package element

import (
	"github.com/KrischanCS/htmfunc"
)

// Del creates a [del element].
//
// The del element represents a removal from the document.
//
// del elements should not cross implied paragraph boundaries.
//
// [del element]: https://html.spec.whatwg.org/#the-del-element
func Del(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("del", attributes, children...)
}

// Ins creates a [ins element].
//
// The ins element represents an addition to the document.
//
// ins elements should not cross implied paragraph boundaries.
//
// [ins element]: https://html.spec.whatwg.org/#the-ins-element
func Ins(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("ins", attributes, children...)
}
