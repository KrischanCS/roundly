// Generated file. DO NOT EDIT.

package element

import (
"github.com/KrischanCS/htmfunc"
)

// Base creates the base element - Base URL and default target [navigable] for [hyperlinks] and [forms [(More)]
//
// It belongs to the following categories:
//   - [metadata]
//
// It is a void element and cannot contain any child elements
//
// If can itself be a child of the following elements/categories of elements:
//   - [head]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [href]
//   - [target]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [navigable]: https://html.spec.whatwg.org/dev/document-sequences.html#navigable
// [hyperlinks]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-target
// [forms]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-target
// [metadata]: https://html.spec.whatwg.org/dev/dom.html#metadata-content-2
// [head]: https://html.spec.whatwg.org/dev/semantics.html#the-head-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [href]: https://html.spec.whatwg.org/dev/semantics.html#attr-base-href
// [target]: https://html.spec.whatwg.org/dev/semantics.html#attr-base-target
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Base(attributes htmfunc.Attribute) htmfunc.Element {
    return htmfunc.WriteVoidElement("base", attributes)
}

// Head creates the head element - Container for document metadata [(More)]
//
// It belongs to the following categories:
//   - none
//// It can be parent to the following elements/categories of elements:
//   - [metadata content]
//
// If can itself be a child of the following elements/categories of elements:
//   - [html]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [html]: https://html.spec.whatwg.org/dev/semantics.html#the-html-element
// [metadata content]: https://html.spec.whatwg.org/dev/dom.html#metadata-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Head(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {return htmfunc.WriteElement("head", attributes, children...)
}

// Html creates the html element - Root element [(More)]
//
// It belongs to the following categories:
//   - none
//// It can be parent to the following elements/categories of elements:
//   - [head]
//   - [body]
//
// If can itself be a child of the following elements/categories of elements:
//   - none
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [head]: https://html.spec.whatwg.org/dev/semantics.html#the-head-element
// [body]: https://html.spec.whatwg.org/dev/sections.html#the-body-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Html(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {return htmfunc.WriteElement("html", attributes, children...)
}

// Link creates the link element - Link metadata [(More)]
//
// It belongs to the following categories:
//   - [metadata]
//   - [flow]
//   - [phrasing]
//
// It is a void element and cannot contain any child elements
//
// If can itself be a child of the following elements/categories of elements:
//   - [head]
//   - [noscript]
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [href]
//   - [crossorigin]
//   - [rel]
//   - [as]
//   - [media]
//   - [hreflang]
//   - [type]
//   - [sizes]
//   - [imagesrcset]
//   - [imagesizes]
//   - [referrerpolicy]
//   - [integrity]
//   - [blocking]
//   - [color]
//   - [disabled]
//   - [fetchpriority]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [metadata]: https://html.spec.whatwg.org/dev/dom.html#metadata-content-2
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [head]: https://html.spec.whatwg.org/dev/semantics.html#the-head-element
// [noscript]: https://html.spec.whatwg.org/dev/scripting.html#the-noscript-element
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [href]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-href
// [crossorigin]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-crossorigin
// [rel]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [as]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-as
// [media]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-media
// [hreflang]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-hreflang
// [type]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-type
// [sizes]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-sizes
// [imagesrcset]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-imagesrcset
// [imagesizes]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-imagesizes
// [referrerpolicy]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-referrerpolicy
// [integrity]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-integrity
// [blocking]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-blocking
// [color]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-color
// [disabled]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-disabled
// [fetchpriority]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-fetchpriority
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Link(attributes htmfunc.Attribute) htmfunc.Element {
    return htmfunc.WriteVoidElement("link", attributes)
}

// Meta creates the meta element - Text metadata [(More)]
//
// It belongs to the following categories:
//   - [metadata]
//   - [flow]
//   - [phrasing]
//
// It is a void element and cannot contain any child elements
//
// If can itself be a child of the following elements/categories of elements:
//   - [head]
//   - [noscript]
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [name]
//   - [http-equiv]
//   - [content]
//   - [charset]
//   - [media]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [metadata]: https://html.spec.whatwg.org/dev/dom.html#metadata-content-2
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [head]: https://html.spec.whatwg.org/dev/semantics.html#the-head-element
// [noscript]: https://html.spec.whatwg.org/dev/scripting.html#the-noscript-element
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [name]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-name
// [http-equiv]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-http-equiv
// [content]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-content
// [charset]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-charset
// [media]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-media
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Meta(attributes htmfunc.Attribute) htmfunc.Element {
    return htmfunc.WriteVoidElement("meta", attributes)
}

// Style creates the style element - Embedded styling information [(More)]
//
// It belongs to the following categories:
//   - [metadata]
//// It can be parent to the following elements/categories of elements:
//   - text
//
// If can itself be a child of the following elements/categories of elements:
//   - [head]
//   - [noscript]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [media]
//   - [blocking]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [metadata]: https://html.spec.whatwg.org/dev/dom.html#metadata-content-2
// [head]: https://html.spec.whatwg.org/dev/semantics.html#the-head-element
// [noscript]: https://html.spec.whatwg.org/dev/scripting.html#the-noscript-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [media]: https://html.spec.whatwg.org/dev/semantics.html#attr-style-media
// [blocking]: https://html.spec.whatwg.org/dev/semantics.html#attr-style-blocking
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Style(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {return htmfunc.WriteElement("style", attributes, children...)
}

// Title creates the title element - Document title [(More)]
//
// It belongs to the following categories:
//   - [metadata]
//// It can be parent to the following elements/categories of elements:
//   - [text]
//
// If can itself be a child of the following elements/categories of elements:
//   - [head]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [metadata]: https://html.spec.whatwg.org/dev/dom.html#metadata-content-2
// [head]: https://html.spec.whatwg.org/dev/semantics.html#the-head-element
// [text]: https://html.spec.whatwg.org/dev/dom.html#text-content
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Title(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {return htmfunc.WriteElement("title", attributes, children...)
}
