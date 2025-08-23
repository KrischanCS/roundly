// Generated file. DO NOT EDIT.

package element

import (
	"github.com/KrischanCS/roundly"
)

// Blockquote creates the blockquote element - A section quoted from another source [(More)]
//
// It belongs to the following categories:
//   - [flow] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [flow]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow]
//
// The following attributes can be added to this element:
//   - [globals] [cite]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/grouping-content.html#the-blockquote-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [cite]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-blockquote-cite
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Blockquote(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("blockquote", attributes, children...)
}

// Dd creates the dd element - Content for corresponding [dt] element(s) [(More)]
//
// It belongs to the following categories:
//   - none
//
// It can be parent to the following elements/categories of elements:
//   - [flow]
//
// If can itself be a child of the following elements/categories of elements:
//   - [dl] [div]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/grouping-content.html#the-dd-element
// [dt]: https://html.spec.whatwg.org/dev/grouping-content.html#the-dt-element
// [dl]: https://html.spec.whatwg.org/dev/grouping-content.html#the-dl-element
// [div]: https://html.spec.whatwg.org/dev/grouping-content.html#the-div-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Dd(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("dd", attributes, children...)
}

// Div creates the div element - Generic flow container, or container for name-value groups in [dl] elements [(More)]
//
// It belongs to the following categories:
//   - [flow] [palpable] [select element inner content elements] [optgroup element inner content elements] [option element inner content elements]
//
// It can be parent to the following elements/categories of elements:
//   - [flow][select element inner content elements] [optgroup element inner content elements] [option element inner content elements]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow] [dl] [select element inner content elements] [optgroup element inner content elements] [option element inner content elements]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/grouping-content.html#the-div-element
// [dl]: https://html.spec.whatwg.org/dev/grouping-content.html#the-dl-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [select element inner content elements]: https://html.spec.whatwg.org/dev/dom.html#select-element-inner-content-elements-2
// [optgroup element inner content elements]: https://html.spec.whatwg.org/dev/dom.html#optgroup-element-inner-content-elements-2
// [option element inner content elements]: https://html.spec.whatwg.org/dev/dom.html#option-element-inner-content-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Div(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("div", attributes, children...)
}

// Dl creates the dl element - Association list consisting of zero or more name-value groups [(More)]
//
// It belongs to the following categories:
//   - [flow] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [dt] [dd] [div] [script-supporting elements]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/grouping-content.html#the-dl-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [dt]: https://html.spec.whatwg.org/dev/grouping-content.html#the-dt-element
// [dd]: https://html.spec.whatwg.org/dev/grouping-content.html#the-dd-element
// [div]: https://html.spec.whatwg.org/dev/grouping-content.html#the-div-element
// [script-supporting elements]: https://html.spec.whatwg.org/dev/dom.html#script-supporting-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Dl(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("dl", attributes, children...)
}

// Dt creates the dt element - Legend for corresponding [dd] element(s) [(More)]
//
// It belongs to the following categories:
//   - none
//
// It can be parent to the following elements/categories of elements:
//   - [flow]
//
// If can itself be a child of the following elements/categories of elements:
//   - [dl] [div]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/grouping-content.html#the-dt-element
// [dd]: https://html.spec.whatwg.org/dev/grouping-content.html#the-dd-element
// [dl]: https://html.spec.whatwg.org/dev/grouping-content.html#the-dl-element
// [div]: https://html.spec.whatwg.org/dev/grouping-content.html#the-div-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Dt(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("dt", attributes, children...)
}

// Figcaption creates the figcaption element - Caption for [figure [(More)]
//
// It belongs to the following categories:
//   - none
//
// It can be parent to the following elements/categories of elements:
//   - [flow]
//
// If can itself be a child of the following elements/categories of elements:
//   - [figure]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/grouping-content.html#the-figcaption-element
// [figure]: https://html.spec.whatwg.org/dev/grouping-content.html#the-figure-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Figcaption(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("figcaption", attributes, children...)
}

// Figure creates the figure element - Figure with optional caption [(More)]
//
// It belongs to the following categories:
//   - [flow] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [figcaption] [flow]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/grouping-content.html#the-figure-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [figcaption]: https://html.spec.whatwg.org/dev/grouping-content.html#the-figcaption-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Figure(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("figure", attributes, children...)
}

// Hr creates the hr element - Thematic break [(More)]
//
// It belongs to the following categories:
//   - [flow] [select element inner content elements]
//
// It is a void element and cannot contain any child elements
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow] [select element inner content elements]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/grouping-content.html#the-hr-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [select element inner content elements]: https://html.spec.whatwg.org/dev/dom.html#select-element-inner-content-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Hr(attributes roundly.Attribute) roundly.Element {
    return roundly.WriteVoidElement("hr", attributes)
}

// Li creates the li element - List item [(More)]
//
// It belongs to the following categories:
//   - none
//
// It can be parent to the following elements/categories of elements:
//   - [flow]
//
// If can itself be a child of the following elements/categories of elements:
//   - [ol] [ul] [menu]
//
// The following attributes can be added to this element:
//   - [globals] [value]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/grouping-content.html#the-li-element
// [ol]: https://html.spec.whatwg.org/dev/grouping-content.html#the-ol-element
// [ul]: https://html.spec.whatwg.org/dev/grouping-content.html#the-ul-element
// [menu]: https://html.spec.whatwg.org/dev/grouping-content.html#the-menu-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [value]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-li-value
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Li(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("li", attributes, children...)
}

// Main creates the main element - Container for the dominant contents of the document [(More)]
//
// It belongs to the following categories:
//   - [flow] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [flow]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/grouping-content.html#the-main-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Main(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("main", attributes, children...)
}

// Menu creates the menu element - Menu of commands [(More)]
//
// It belongs to the following categories:
//   - [flow] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [li] [script-supporting elements]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/grouping-content.html#the-menu-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [li]: https://html.spec.whatwg.org/dev/grouping-content.html#the-li-element
// [script-supporting elements]: https://html.spec.whatwg.org/dev/dom.html#script-supporting-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Menu(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("menu", attributes, children...)
}

// Ol creates the ol element - Ordered list [(More)]
//
// It belongs to the following categories:
//   - [flow] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [li] [script-supporting elements]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow]
//
// The following attributes can be added to this element:
//   - [globals] [reversed] [start] [type]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/grouping-content.html#the-ol-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [li]: https://html.spec.whatwg.org/dev/grouping-content.html#the-li-element
// [script-supporting elements]: https://html.spec.whatwg.org/dev/dom.html#script-supporting-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [reversed]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-ol-reversed
// [start]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-ol-start
// [type]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-ol-type
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Ol(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("ol", attributes, children...)
}

// P creates the p element - Paragraph [(More)]
//
// It belongs to the following categories:
//   - [flow] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [phrasing]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/grouping-content.html#the-p-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func P(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("p", attributes, children...)
}

// Pre creates the pre element - Block of preformatted text [(More)]
//
// It belongs to the following categories:
//   - [flow] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [phrasing]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/grouping-content.html#the-pre-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Pre(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("pre", attributes, children...)
}

// Search creates the search element - Container for search controls [(More)]
//
// It belongs to the following categories:
//   - [flow] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [flow]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/grouping-content.html#the-search-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Search(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("search", attributes, children...)
}

// Ul creates the ul element - List [(More)]
//
// It belongs to the following categories:
//   - [flow] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [li] [script-supporting elements]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/grouping-content.html#the-ul-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [li]: https://html.spec.whatwg.org/dev/grouping-content.html#the-li-element
// [script-supporting elements]: https://html.spec.whatwg.org/dev/dom.html#script-supporting-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Ul(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("ul", attributes, children...)
}
