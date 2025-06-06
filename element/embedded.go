// Generated file. DO NOT EDIT.

package element

import (
"github.com/KrischanCS/roundly"
)

// Img creates the img element - Image [(More)]
//
// It belongs to the following categories:
// [flow] [phrasing] [embedded] [interactive] [form-associated] [palpable]
//
// It is a void element and cannot contain any child elements
//
// If can itself be a child of the following elements/categories of elements:
// [phrasing] [picture]
//
// The following attributes can be added to this element:
// [globals] [alt] [src] [srcset] [sizes] [crossorigin] [usemap] [ismap] [width] [height] [referrerpolicy] [decoding] [loading] [fetchpriority]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/embedded-content.html#the-img-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [embedded]: https://html.spec.whatwg.org/dev/dom.html#embedded-content-category
// [interactive]: https://html.spec.whatwg.org/dev/dom.html#interactive-content-2
// [form-associated]: https://html.spec.whatwg.org/dev/forms.html#form-associated-element
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [picture]: https://html.spec.whatwg.org/dev/embedded-content.html#the-picture-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [alt]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-alt
// [src]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-src
// [srcset]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-srcset
// [sizes]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-sizes
// [crossorigin]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-crossorigin
// [usemap]: https://html.spec.whatwg.org/dev/image-maps.html#attr-hyperlink-usemap
// [ismap]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-ismap
// [width]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [height]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [referrerpolicy]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-referrerpolicy
// [decoding]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-decoding
// [loading]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-loading
// [fetchpriority]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-fetchpriority
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Img(attributes roundly.Attribute) roundly.Element {
    return roundly.WriteVoidElement("img", attributes)
}

// Picture creates the picture element - Image [(More)]
//
// It belongs to the following categories:
// [flow] [phrasing] [embedded] [palpable]
//
// It can be parent to the following elements/categories of elements:
// [source] one [img] [script-supporting elements]
//
// If can itself be a child of the following elements/categories of elements:
// [phrasing]
//
// The following attributes can be added to this element:
// [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/embedded-content.html#the-picture-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [embedded]: https://html.spec.whatwg.org/dev/dom.html#embedded-content-category
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [source]: https://html.spec.whatwg.org/dev/embedded-content.html#the-source-element
// [img]: https://html.spec.whatwg.org/dev/embedded-content.html#the-img-element
// [script-supporting elements]: https://html.spec.whatwg.org/dev/dom.html#script-supporting-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Picture(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("picture", attributes, children...)
}

// Source creates the source element - Image source for [img] or media source for [video] or [audio [(More)]
//
// It belongs to the following categories:
// none
//
// It is a void element and cannot contain any child elements
//
// If can itself be a child of the following elements/categories of elements:
// [picture] [video] [audio]
//
// The following attributes can be added to this element:
// [globals] [type] [media] [src] [srcset] [sizes] [width] [height]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/embedded-content.html#the-source-element
// [img]: https://html.spec.whatwg.org/dev/embedded-content.html#the-img-element
// [video]: https://html.spec.whatwg.org/dev/media.html#the-video-element
// [audio]: https://html.spec.whatwg.org/dev/media.html#the-audio-element
// [picture]: https://html.spec.whatwg.org/dev/embedded-content.html#the-picture-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [type]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-source-type
// [media]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-source-media
// [src]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-source-src
// [srcset]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-source-srcset
// [sizes]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-source-sizes
// [width]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [height]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Source(attributes roundly.Attribute) roundly.Element {
    return roundly.WriteVoidElement("source", attributes)
}
