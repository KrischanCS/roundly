// Generated file. DO NOT EDIT.

package html

import (
	"github.com/KrischanCS/roundly"
)

// Noscript creates the noscript element - Fallback content for script [(More)]
//
// It belongs to the following categories:
//   - [metadata] [flow] [phrasing] [select element inner content elements] [optgroup element inner content elements]
//
// It can be parent to the following elements/categories of elements:
//   - varies
//
// If can itself be a child of the following elements/categories of elements:
//   - [head] [phrasing]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/scripting.html#the-noscript-element
// [metadata]: https://html.spec.whatwg.org/dev/dom.html#metadata-content-2
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [select element inner content elements]: https://html.spec.whatwg.org/dev/dom.html#select-element-inner-content-elements-2
// [optgroup element inner content elements]: https://html.spec.whatwg.org/dev/dom.html#optgroup-element-inner-content-elements-2
// [head]: https://html.spec.whatwg.org/dev/semantics.html#the-head-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Noscript(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("noscript", attributes, children...)
}

// Script creates the script element - Embedded script [(More)]
//
// It belongs to the following categories:
//   - [metadata] [flow] [phrasing] [script-supporting]
//
// It can be parent to the following elements/categories of elements:
//   - script, data, or script documentation
//
// If can itself be a child of the following elements/categories of elements:
//   - [head] [phrasing] [script-supporting]
//
// The following attributes can be added to this element:
//   - [globals] [src] [type] [nomodule] [async] [defer] [crossorigin] [integrity] [referrerpolicy] [blocking] [fetchpriority]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/scripting.html#the-script-element
// [metadata]: https://html.spec.whatwg.org/dev/dom.html#metadata-content-2
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [script-supporting]: https://html.spec.whatwg.org/dev/dom.html#script-supporting-elements-2
// [head]: https://html.spec.whatwg.org/dev/semantics.html#the-head-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [src]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-src
// [type]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-type
// [nomodule]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-nomodule
// [async]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-async
// [defer]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-defer
// [crossorigin]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-crossorigin
// [integrity]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-integrity
// [referrerpolicy]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-referrerpolicy
// [blocking]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-blocking
// [fetchpriority]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-fetchpriority
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Script(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("script", attributes, children...)
}

// Slot creates the slot element - Shadow tree slot [(More)]
//
// It belongs to the following categories:
//   - [flow] [phrasing]
//
// It can be parent to the following elements/categories of elements:
//   - [transparent]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals] [name]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/scripting.html#the-slot-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [transparent]: https://html.spec.whatwg.org/dev/dom.html#transparent
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [name]: https://html.spec.whatwg.org/dev/scripting.html#attr-slot-name
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Slot(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("slot", attributes, children...)
}

// Template creates the template element - Template [(More)]
//
// It belongs to the following categories:
//   - [metadata] [flow] [phrasing] [script-supporting]
//
// It can be parent to the following elements/categories of elements:
//   - empty
//
// If can itself be a child of the following elements/categories of elements:
//   - [metadata] [phrasing] [script-supporting] [colgroup]
//
// The following attributes can be added to this element:
//   - [globals] [shadowrootmode] [shadowrootdelegatesfocus] [shadowrootclonable] [shadowrootserializable] [shadowrootcustomelementregistry]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/scripting.html#the-template-element
// [metadata]: https://html.spec.whatwg.org/dev/dom.html#metadata-content-2
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [script-supporting]: https://html.spec.whatwg.org/dev/dom.html#script-supporting-elements-2
// [colgroup]: https://html.spec.whatwg.org/dev/tables.html#the-colgroup-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [shadowrootmode]: https://html.spec.whatwg.org/dev/scripting.html#attr-template-shadowrootmode
// [shadowrootdelegatesfocus]: https://html.spec.whatwg.org/dev/scripting.html#attr-template-shadowrootdelegatesfocus
// [shadowrootclonable]: https://html.spec.whatwg.org/dev/scripting.html#attr-template-shadowrootclonable
// [shadowrootserializable]: https://html.spec.whatwg.org/dev/scripting.html#attr-template-shadowrootserializable
// [shadowrootcustomelementregistry]: https://html.spec.whatwg.org/dev/scripting.html#attr-template-shadowrootcustomelementregistry
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Template(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("template", attributes, children...)
}
