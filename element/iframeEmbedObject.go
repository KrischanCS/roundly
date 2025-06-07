// Generated file. DO NOT EDIT.

package element

import (
	"github.com/KrischanCS/roundly"
)

// Embed creates the embed element - Plugin [(More)]
//
// It belongs to the following categories:
//   - [flow] [phrasing] [embedded] [interactive] [palpable]
//
// It is a void element and cannot contain any child elements
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals] [src] [type] [width] [height] any
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#the-embed-element
// [Plugin]: https://html.spec.whatwg.org/dev/infrastructure.html#plugin
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [embedded]: https://html.spec.whatwg.org/dev/dom.html#embedded-content-category
// [interactive]: https://html.spec.whatwg.org/dev/dom.html#interactive-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [src]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-embed-src
// [type]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-embed-type
// [width]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [height]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Embed(attributes roundly.Attribute) roundly.Element {
    return roundly.WriteVoidElement("embed", attributes)
}

// Iframe creates the iframe element - Child navigable [(More)]
//
// It belongs to the following categories:
//   - [flow] [phrasing] [embedded] [interactive] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - empty
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals] [src] [srcdoc] [name] [sandbox] [allow] [allowfullscreen] [width] [height] [referrerpolicy] [loading]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#the-iframe-element
// [Child navigable]: https://html.spec.whatwg.org/dev/document-sequences.html#child-navigable
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [embedded]: https://html.spec.whatwg.org/dev/dom.html#embedded-content-category
// [interactive]: https://html.spec.whatwg.org/dev/dom.html#interactive-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [src]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-src
// [srcdoc]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-srcdoc
// [name]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-name
// [sandbox]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-sandbox
// [allow]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-allow
// [allowfullscreen]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-allowfullscreen
// [width]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [height]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [referrerpolicy]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-referrerpolicy
// [loading]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-loading
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Iframe(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("iframe", attributes, children...)
}

// Object creates the object element - Image, [child navigable], or [plugin [(More)]
//
// It belongs to the following categories:
//   - [flow] [phrasing] [embedded] [interactive] [listed] [form-associated] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [transparent]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals] [data] [type] [name] [form] [width] [height]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#the-object-element
// [child navigable]: https://html.spec.whatwg.org/dev/document-sequences.html#child-navigable
// [plugin]: https://html.spec.whatwg.org/dev/infrastructure.html#plugin
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [embedded]: https://html.spec.whatwg.org/dev/dom.html#embedded-content-category
// [interactive]: https://html.spec.whatwg.org/dev/dom.html#interactive-content-2
// [listed]: https://html.spec.whatwg.org/dev/forms.html#category-listed
// [form-associated]: https://html.spec.whatwg.org/dev/forms.html#form-associated-element
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [transparent]: https://html.spec.whatwg.org/dev/dom.html#transparent
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [data]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-object-data
// [type]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-object-type
// [name]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-object-name
// [form]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [width]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [height]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Object(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("object", attributes, children...)
}
