// Generated file. DO NOT EDIT.

package element

import (
	"github.com/KrischanCS/roundly"
)

// Canvas creates the canvas element - Scriptable bitmap canvas [(More)]
//
// It belongs to the following categories:
// [flow] [phrasing] [embedded] [palpable]
//
// It can be parent to the following elements/categories of elements:
// [transparent]
//
// If can itself be a child of the following elements/categories of elements:
// [phrasing]
//
// The following attributes can be added to this element:
// [globals] [width] [height]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/canvas.html#the-canvas-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [embedded]: https://html.spec.whatwg.org/dev/dom.html#embedded-content-category
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [transparent]: https://html.spec.whatwg.org/dev/dom.html#transparent
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [width]: https://html.spec.whatwg.org/dev/canvas.html#attr-canvas-width
// [height]: https://html.spec.whatwg.org/dev/canvas.html#attr-canvas-height
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Canvas(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("canvas", attributes, children...)
}
