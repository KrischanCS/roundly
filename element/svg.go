// Generated file. DO NOT EDIT.

package element

import (
"github.com/KrischanCS/htmfunc"
)

// Svg creates the svg element - SVG root [(More)]
//
// It belongs to the following categories:
//   - [flow]
//   - [phrasing]
//   - [embedded]
//   - [palpable]
//// It can be parent to the following elements/categories of elements:
//   - per [[SVG]]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - per [[SVG]]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [embedded]: https://html.spec.whatwg.org/dev/dom.html#embedded-content-category
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [[SVG]]: https://html.spec.whatwg.org/dev/references.html#refsSVG
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Svg(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("svg", attributes, children...)
}
