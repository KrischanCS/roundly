// Generated file. DO NOT EDIT.

package element

import (
"github.com/KrischanCS/htmfunc"
)

// Details creates the details element - Disclosure control for hiding details [(More)]
//
// It belongs to the following categories:
//   - [flow]
//   - [interactive]
//   - [palpable]
//// It can be parent to the following elements/categories of elements:
//   - [summary]
//   - [flow]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [name]
//   - [open]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [interactive]: https://html.spec.whatwg.org/dev/dom.html#interactive-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [summary]: https://html.spec.whatwg.org/dev/interactive-elements.html#the-summary-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [name]: https://html.spec.whatwg.org/dev/interactive-elements.html#attr-details-name
// [open]: https://html.spec.whatwg.org/dev/interactive-elements.html#attr-details-open
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Details(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("details", attributes, children...)
}

// Dialog creates the dialog element - Dialog box or window [(More)]
//
// It belongs to the following categories:
//   - [flow]
//// It can be parent to the following elements/categories of elements:
//   - [flow]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [open]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [open]: https://html.spec.whatwg.org/dev/interactive-elements.html#attr-dialog-open
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Dialog(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("dialog", attributes, children...)
}

// Summary creates the summary element - Caption for [details [(More)]
//
// It belongs to the following categories:
//   - none
//// It can be parent to the following elements/categories of elements:
//   - [phrasing]
//   - [heading content]
//
// If can itself be a child of the following elements/categories of elements:
//   - [details]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [details]: https://html.spec.whatwg.org/dev/interactive-elements.html#the-details-element
// [details]: https://html.spec.whatwg.org/dev/interactive-elements.html#the-details-element
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [heading content]: https://html.spec.whatwg.org/dev/dom.html#heading-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Summary(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("summary", attributes, children...)
}
