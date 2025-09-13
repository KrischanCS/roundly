// Generated file. DO NOT EDIT.

package html

import (
	"github.com/KrischanCS/roundly"
)

// Math creates the math element - MathML root [(More)]
//
// It belongs to the following categories:
//   - [flow] [phrasing] [embedded] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - per [[MATHML]]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - per [[MATHML]]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/https://w3c.github.io/mathml-core/#the-top-level-math-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [embedded]: https://html.spec.whatwg.org/dev/dom.html#embedded-content-category
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [[MATHML]]: https://html.spec.whatwg.org/dev/references.html#refsMATHML
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Math(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("math", attributes, children...)
}
