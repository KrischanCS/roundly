// Generated file. DO NOT EDIT.

package element

import (
	"github.com/KrischanCS/roundly"
)

// Area creates the area element - Hyperlink or dead area on an image map [(More)]
//
// It belongs to the following categories:
// [flow] [phrasing]
//
// It is a void element and cannot contain any child elements
//
// If can itself be a child of the following elements/categories of elements:
// [phrasing]
//
// The following attributes can be added to this element:
// [globals] [alt] [coords] [shape] [href] [target] [download] [ping] [rel] [referrerpolicy]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/image-maps.html#the-area-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [alt]: https://html.spec.whatwg.org/dev/image-maps.html#attr-area-alt
// [coords]: https://html.spec.whatwg.org/dev/image-maps.html#attr-area-coords
// [shape]: https://html.spec.whatwg.org/dev/image-maps.html#attr-area-shape
// [href]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-href
// [target]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-target
// [download]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-download
// [ping]: https://html.spec.whatwg.org/dev/links.html#ping
// [rel]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-rel
// [referrerpolicy]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-referrerpolicy
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Area(attributes roundly.Attribute) roundly.Element {
    return roundly.WriteVoidElement("area", attributes)
}

// Map creates the map element - Image map [(More)]
//
// It belongs to the following categories:
// [flow] [phrasing] [palpable]
//
// It can be parent to the following elements/categories of elements:
// [transparent] [area]
//
// If can itself be a child of the following elements/categories of elements:
// [phrasing]
//
// The following attributes can be added to this element:
// [globals] [name]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/image-maps.html#the-map-element
// [Image map]: https://html.spec.whatwg.org/dev/image-maps.html#image-map
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [transparent]: https://html.spec.whatwg.org/dev/dom.html#transparent
// [area]: https://html.spec.whatwg.org/dev/image-maps.html#the-area-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [name]: https://html.spec.whatwg.org/dev/image-maps.html#attr-map-name
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Map(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("map", attributes, children...)
}
