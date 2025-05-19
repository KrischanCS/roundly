// Generated file. DO NOT EDIT.

package attribute

import (
	"github.com/KrischanCS/htmfunc"
)

// AccessKey creates the accesskey attribute - Keyboard shortcut to activate or focus element
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Value constraints: [Ordered set of unique space-separated tokens], none of which are [identical to] another, each consisting of one code point in length
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/interaction.html#the-accesskey-attribute
// [Ordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#ordered-set-of-unique-space-separated-tokens
// [identical to]: https://html.spec.whatwg.org/dev/https://infra.spec.whatwg.org/#string-is
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func AccessKey(accessKey ...string) htmfunc.Attribute {
	return htmfunc.WriteMultiValueAttribute("accesskey", ' ', accessKey...)
}

// Blocking creates the blocking attribute - Whether the element is [potentially render-blocking]
//
// It can be applied to the following elements:
//   - [link]
//   - [script]
//   - [style]
//
// Value constraints: [Unordered set of unique space-separated tokens] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-blocking
// [script]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-blocking
// [style]: https://html.spec.whatwg.org/dev/semantics.html#attr-style-blocking
// [potentially render-blocking]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#potentially-render-blocking
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Blocking(blocking ...string) htmfunc.Attribute {
	return htmfunc.WriteMultiValueAttribute("blocking", ' ', blocking...)
}

// Class creates the class attribute - Classes to which the element belongs
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Value constraints: [Set of space-separated tokens]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#classes
// [Set of space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#set-of-space-separated-tokens
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Class(class ...string) htmfunc.Attribute {
	return htmfunc.WriteMultiValueAttribute("class", ' ', class...)
}

// For creates the for attribute - Specifies controls from which the output was calculated
//
// It can be applied to the following elements:
//   - [output]
//
// Value constraints: [Unordered set of unique space-separated tokens] consisting of IDs (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [output]: https://html.spec.whatwg.org/dev/form-elements.html#attr-output-for
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func For(forV ...string) htmfunc.Attribute {
	return htmfunc.WriteMultiValueAttribute("for", ' ', forV...)
}

// Headers creates the headers attribute - The header cells for this cell
//
// It can be applied to the following elements:
//   - [td]
//   - [th]
//
// Value constraints: [Unordered set of unique space-separated tokens] consisting of IDs (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [td]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-headers
// [th]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-headers
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Headers(headers ...string) htmfunc.Attribute {
	return htmfunc.WriteMultiValueAttribute("headers", ' ', headers...)
}

// ItemProp creates the itemprop attribute - [Property names] of a microdata item
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Value constraints: [Unordered set of unique space-separated tokens] consisting of [valid absolute URLs], [defined property names], or text (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/microdata.html#names:-the-itemprop-attribute
// [Property names]: https://html.spec.whatwg.org/dev/microdata.html#property-names
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [valid absolute URLs]: https://html.spec.whatwg.org/dev/https://url.spec.whatwg.org/#syntax-url-absolute
// [defined property names]: https://html.spec.whatwg.org/dev/microdata.html#defined-property-name
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ItemProp(itemProp ...string) htmfunc.Attribute {
	return htmfunc.WriteMultiValueAttribute("itemprop", ' ', itemProp...)
}

// ItemRef creates the itemref attribute - [Referenced] elements
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Value constraints: [Unordered set of unique space-separated tokens] consisting of IDs (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/microdata.html#attr-itemref
// [Referenced]: https://html.spec.whatwg.org/dev/dom.html#referenced
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ItemRef(itemRef ...string) htmfunc.Attribute {
	return htmfunc.WriteMultiValueAttribute("itemref", ' ', itemRef...)
}

// ItemType creates the itemtype attribute - [Item types] of a microdata item
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Value constraints: [Unordered set of unique space-separated tokens] consisting of [valid absolute URLs] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/microdata.html#attr-itemtype
// [Item types]: https://html.spec.whatwg.org/dev/microdata.html#item-types
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [valid absolute URLs]: https://html.spec.whatwg.org/dev/https://url.spec.whatwg.org/#syntax-url-absolute
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ItemType(itemType ...string) htmfunc.Attribute {
	return htmfunc.WriteMultiValueAttribute("itemtype", ' ', itemType...)
}

// Ping creates the ping attribute - [URLs] to ping
//
// It can be applied to the following elements:
//   - [a]
//   - [area]
//
// Value constraints: [Set of space-separated tokens] consisting of [valid non-empty URLs]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#ping
// [area]: https://html.spec.whatwg.org/dev/links.html#ping
// [URLs]: https://html.spec.whatwg.org/dev/https://url.spec.whatwg.org/#concept-url
// [Set of space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#set-of-space-separated-tokens
// [valid non-empty URLs]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Ping(ping ...string) htmfunc.Attribute {
	return htmfunc.WriteMultiValueAttribute("ping", ' ', ping...)
}

// Rel creates the rel attribute - Relationship between the location in the document containing the [hyperlink] and the destination resource
//
// It can be applied to the following elements:
//   - [a]
//   - [area]
//
// Value constraints: [Unordered set of unique space-separated tokens] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-rel
// [area]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-rel
// [hyperlink]: https://html.spec.whatwg.org/dev/links.html#hyperlink
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Rel(rel ...string) htmfunc.Attribute {
	return htmfunc.WriteMultiValueAttribute("rel", ' ', rel...)
}

// Rel creates the rel attribute - Relationship between the document containing the [hyperlink] and the destination resource
//
// It can be applied to the following elements:
//   - [link]
//
// Value constraints: [Unordered set of unique space-separated tokens] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [hyperlink]: https://html.spec.whatwg.org/dev/links.html#hyperlink
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Rel(rel ...string) htmfunc.Attribute {
	return htmfunc.WriteMultiValueAttribute("rel", ' ', rel...)
}

// SandBox creates the sandbox attribute - Security rules for nested content
//
// It can be applied to the following elements:
//   - [iframe]
//
// Value constraints: [Unordered set of unique space-separated tokens], [ASCII case-insensitive], consisting of"allow-downloads", "allow-forms", "allow-modals", "allow-orientation-lock", "allow-pointer-lock", "allow-popups", "allow-popups-to-escape-sandbox", "allow-presentation", "allow-same-origin", "allow-scripts", "allow-top-navigation", "allow-top-navigation-by-user-activation", "allow-top-navigation-to-custom-protocols"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-sandbox
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [ASCII case-insensitive]: https://html.spec.whatwg.org/dev/https://infra.spec.whatwg.org/#ascii-case-insensitive
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func SandBox(sandBox ...string) htmfunc.Attribute {
	return htmfunc.WriteMultiValueAttribute("sandbox", ' ', sandBox...)
}

// Sizes creates the sizes attribute - Sizes of the icons (for [rel]="[icon]")
//
// It can be applied to the following elements:
//   - [link]
//
// Value constraints: [Unordered set of unique space-separated tokens], [ASCII case-insensitive], consisting of sizes (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-sizes
// [rel]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [icon]: https://html.spec.whatwg.org/dev/links.html#rel-icon
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [ASCII case-insensitive]: https://html.spec.whatwg.org/dev/https://infra.spec.whatwg.org/#ascii-case-insensitive
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Sizes(sizes ...string) htmfunc.Attribute {
	return htmfunc.WriteMultiValueAttribute("sizes", ' ', sizes...)
}
