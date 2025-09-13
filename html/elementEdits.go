// Generated file. DO NOT EDIT.

package html

import (
	"github.com/KrischanCS/roundly"
)

// Del creates the del element - A removal from the document [(More)]
//
// It belongs to the following categories:
//   - [flow] [phrasing] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [transparent]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals] [cite] [datetime]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/edits.html#the-del-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [transparent]: https://html.spec.whatwg.org/dev/dom.html#transparent
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [cite]: https://html.spec.whatwg.org/dev/edits.html#attr-mod-cite
// [datetime]: https://html.spec.whatwg.org/dev/edits.html#attr-mod-datetime
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Del(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("del", attributes, children...)
}

// Ins creates the ins element - An addition to the document [(More)]
//
// It belongs to the following categories:
//   - [flow] [phrasing] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [transparent]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals] [cite] [datetime]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/edits.html#the-ins-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [transparent]: https://html.spec.whatwg.org/dev/dom.html#transparent
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [cite]: https://html.spec.whatwg.org/dev/edits.html#attr-mod-cite
// [datetime]: https://html.spec.whatwg.org/dev/edits.html#attr-mod-datetime
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Ins(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("ins", attributes, children...)
}
