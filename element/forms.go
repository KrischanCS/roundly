// Generated file. DO NOT EDIT.

package element

import (
"github.com/KrischanCS/htmfunc"
)

// Form creates the form element - User-submittable form [(More)]
//
// It belongs to the following categories:
//   - [flow]
//   - [palpable]
//// It can be parent to the following elements/categories of elements:
//   - [flow]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [accept-charset]
//   - [action]
//   - [autocomplete]
//   - [enctype]
//   - [method]
//   - [name]
//   - [novalidate]
//   - [rel]
//   - [target]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [accept-charset]: https://html.spec.whatwg.org/dev/forms.html#attr-form-accept-charset
// [action]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-action
// [autocomplete]: https://html.spec.whatwg.org/dev/forms.html#attr-form-autocomplete
// [enctype]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-enctype
// [method]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-method
// [name]: https://html.spec.whatwg.org/dev/forms.html#attr-form-name
// [novalidate]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-novalidate
// [rel]: https://html.spec.whatwg.org/dev/forms.html#attr-form-rel
// [target]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-target
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Form(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("form", attributes, children...)
}

// Label creates the label element - Caption for a form control [(More)]
//
// It belongs to the following categories:
//   - [flow]
//   - [phrasing]
//   - [interactive]
//   - [palpable]
//// It can be parent to the following elements/categories of elements:
//   - [phrasing]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [for]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [interactive]: https://html.spec.whatwg.org/dev/dom.html#interactive-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [for]: https://html.spec.whatwg.org/dev/forms.html#attr-label-for
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Label(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("label", attributes, children...)
}
