// Generated file. DO NOT EDIT.

package element

import (
"github.com/KrischanCS/htmfunc"
)

// Caption creates the caption element - Table caption [(More)]
//
// It belongs to the following categories:
//   - none
//// It can be parent to the following elements/categories of elements:
//   - [flow]
//
// If can itself be a child of the following elements/categories of elements:
//   - [table]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [table]: https://html.spec.whatwg.org/dev/tables.html#the-table-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Caption(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("caption", attributes, children...)
}

// Col creates the col element - Table column [(More)]
//
// It belongs to the following categories:
//   - none
//
// It is a void element and cannot contain any child elements
//
// If can itself be a child of the following elements/categories of elements:
//   - [colgroup]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [span]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [colgroup]: https://html.spec.whatwg.org/dev/tables.html#the-colgroup-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [span]: https://html.spec.whatwg.org/dev/tables.html#attr-col-span
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Col(attributes htmfunc.Attribute) htmfunc.Element {
    return htmfunc.WriteVoidElement("col", attributes)
}

// Colgroup creates the colgroup element - Group of columns in a table [(More)]
//
// It belongs to the following categories:
//   - none
//// It can be parent to the following elements/categories of elements:
//   - [col]
//   - [template]
//
// If can itself be a child of the following elements/categories of elements:
//   - [table]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [span]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [table]: https://html.spec.whatwg.org/dev/tables.html#the-table-element
// [col]: https://html.spec.whatwg.org/dev/tables.html#the-col-element
// [template]: https://html.spec.whatwg.org/dev/scripting.html#the-template-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [span]: https://html.spec.whatwg.org/dev/tables.html#attr-colgroup-span
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Colgroup(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("colgroup", attributes, children...)
}

// Table creates the table element - Table [(More)]
//
// It belongs to the following categories:
//   - [flow]
//   - [palpable]
//// It can be parent to the following elements/categories of elements:
//   - [caption]
//   - [colgroup]
//   - [thead]
//   - [tbody]
//   - [tfoot]
//   - [tr]
//   - [script-supporting elements]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [caption]: https://html.spec.whatwg.org/dev/tables.html#the-caption-element
// [colgroup]: https://html.spec.whatwg.org/dev/tables.html#the-colgroup-element
// [thead]: https://html.spec.whatwg.org/dev/tables.html#the-thead-element
// [tbody]: https://html.spec.whatwg.org/dev/tables.html#the-tbody-element
// [tfoot]: https://html.spec.whatwg.org/dev/tables.html#the-tfoot-element
// [tr]: https://html.spec.whatwg.org/dev/tables.html#the-tr-element
// [script-supporting elements]: https://html.spec.whatwg.org/dev/dom.html#script-supporting-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Table(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("table", attributes, children...)
}

// Tbody creates the tbody element - Group of rows in a table [(More)]
//
// It belongs to the following categories:
//   - none
//// It can be parent to the following elements/categories of elements:
//   - [tr]
//   - [script-supporting elements]
//
// If can itself be a child of the following elements/categories of elements:
//   - [table]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [table]: https://html.spec.whatwg.org/dev/tables.html#the-table-element
// [tr]: https://html.spec.whatwg.org/dev/tables.html#the-tr-element
// [script-supporting elements]: https://html.spec.whatwg.org/dev/dom.html#script-supporting-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Tbody(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("tbody", attributes, children...)
}

// Td creates the td element - Table cell [(More)]
//
// It belongs to the following categories:
//   - none
//// It can be parent to the following elements/categories of elements:
//   - [flow]
//
// If can itself be a child of the following elements/categories of elements:
//   - [tr]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [colspan]
//   - [rowspan]
//   - [headers]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [tr]: https://html.spec.whatwg.org/dev/tables.html#the-tr-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [colspan]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-colspan
// [rowspan]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-rowspan
// [headers]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-headers
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Td(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("td", attributes, children...)
}

// Tfoot creates the tfoot element - Group of footer rows in a table [(More)]
//
// It belongs to the following categories:
//   - none
//// It can be parent to the following elements/categories of elements:
//   - [tr]
//   - [script-supporting elements]
//
// If can itself be a child of the following elements/categories of elements:
//   - [table]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [table]: https://html.spec.whatwg.org/dev/tables.html#the-table-element
// [tr]: https://html.spec.whatwg.org/dev/tables.html#the-tr-element
// [script-supporting elements]: https://html.spec.whatwg.org/dev/dom.html#script-supporting-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Tfoot(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("tfoot", attributes, children...)
}

// Th creates the th element - Table header cell [(More)]
//
// It belongs to the following categories:
//   - [interactive]
//// It can be parent to the following elements/categories of elements:
//   - [flow]
//
// If can itself be a child of the following elements/categories of elements:
//   - [tr]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [colspan]
//   - [rowspan]
//   - [headers]
//   - [scope]
//   - [abbr]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [interactive]: https://html.spec.whatwg.org/dev/dom.html#interactive-content-2
// [tr]: https://html.spec.whatwg.org/dev/tables.html#the-tr-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [colspan]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-colspan
// [rowspan]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-rowspan
// [headers]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-headers
// [scope]: https://html.spec.whatwg.org/dev/tables.html#attr-th-scope
// [abbr]: https://html.spec.whatwg.org/dev/tables.html#attr-th-abbr
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Th(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("th", attributes, children...)
}

// Thead creates the thead element - Group of heading rows in a table [(More)]
//
// It belongs to the following categories:
//   - none
//// It can be parent to the following elements/categories of elements:
//   - [tr]
//   - [script-supporting elements]
//
// If can itself be a child of the following elements/categories of elements:
//   - [table]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [table]: https://html.spec.whatwg.org/dev/tables.html#the-table-element
// [tr]: https://html.spec.whatwg.org/dev/tables.html#the-tr-element
// [script-supporting elements]: https://html.spec.whatwg.org/dev/dom.html#script-supporting-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Thead(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("thead", attributes, children...)
}

// Tr creates the tr element - Table row [(More)]
//
// It belongs to the following categories:
//   - none
//// It can be parent to the following elements/categories of elements:
//   - [th]
//   - [td]
//   - [script-supporting elements]
//
// If can itself be a child of the following elements/categories of elements:
//   - [table]
//   - [thead]
//   - [tbody]
//   - [tfoot]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [table]: https://html.spec.whatwg.org/dev/tables.html#the-table-element
// [thead]: https://html.spec.whatwg.org/dev/tables.html#the-thead-element
// [tbody]: https://html.spec.whatwg.org/dev/tables.html#the-tbody-element
// [tfoot]: https://html.spec.whatwg.org/dev/tables.html#the-tfoot-element
// [th]: https://html.spec.whatwg.org/dev/tables.html#the-th-element
// [td]: https://html.spec.whatwg.org/dev/tables.html#the-td-element
// [script-supporting elements]: https://html.spec.whatwg.org/dev/dom.html#script-supporting-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Tr(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("tr", attributes, children...)
}
