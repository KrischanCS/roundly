// Generated file. DO NOT EDIT.

package element

import (
"github.com/KrischanCS/htmfunc"
)

// Canvas creates the canvas element - Scriptable bitmap canvas [(More)]
//
// It belongs to the following categories:
//   - [flow]
//   - [phrasing]
//   - [embedded]
//   - [palpable]
//// It can be parent to the following elements/categories of elements:
//   - [transparent]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [width]
//   - [height]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [embedded]: https://html.spec.whatwg.org/dev/dom.html#embedded-content-category
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [transparent]: https://html.spec.whatwg.org/dev/dom.html#transparent
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [width]: https://html.spec.whatwg.org/dev/canvas.html#attr-canvas-width
// [height]: https://html.spec.whatwg.org/dev/canvas.html#attr-canvas-height
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Canvas(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("canvas", attributes, children...)
}
