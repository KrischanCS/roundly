package element

import (
	"github.com/KrischanCS/htmfunc"
)

// Table creates a [table element].
//
// The table element represents data with more than one dimension, in the form of a table.
//
// [table element]: https://html.spec.whatwg.org/#the-table-element
func Table(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("table", attributes, children...)
}

// Caption creates a [caption element].
//
// The caption element represents the title of the table that is its parent, if it has a parent and that is a table
// element.
//
// [caption element]: https://html.spec.whatwg.org/#the-caption-element
func Caption(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("caption", attributes, children...)
}

// Colgroup creates a [colgroup element].
//
// The colgroup element represents a group of one or more columns in the table that is its parent, if it has a parent
// and that is a table element.
//
// [colgroup element]: https://html.spec.whatwg.org/#the-colgroup-element
func Colgroup(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("colgroup", attributes, children...)
}

// Col creates a [col element].
//
// If a col element has a parent and that is a colgroup element that itself has a parent that is a table element, then
// the col element represents one or more columns in the column group represented by that colgroup.
//
// [col element]: https://html.spec.whatwg.org/#the-col-element
func Col(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("col", attributes, children...)
}

// Tbody creates a [tbody element].
//
// The tbody element represents a block of rows that consist of a body of data for the parent table element, if the
// tbody element has a parent and it is a table.
//
// [tbody element]: https://html.spec.whatwg.org/#the-tbody-element
func Tbody(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("tbody", attributes, children...)
}

// Thead creates a [thead element].
//
// The thead element represents the block of rows that consist of the column labels (headers) and any ancillary
// non-header cells for the parent table element, if the thead element has a parent and it is a table.
//
// [thead element]: https://html.spec.whatwg.org/#the-thead-element
func Thead(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("thead", attributes, children...)
}

// Tfoot creates a [tfoot element].
//
// The tfoot element represents the block of rows that consist of the column summaries (footers) for the parent table
// element, if the tfoot element has a parent and it is a table.
//
// [tfoot element]: https://html.spec.whatwg.org/#the-tfoot-element
func Tfoot(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("tfoot", attributes, children...)
}

// Tr creates a [tr element].
//
// The tr element represents a row of cells in a table.
//
// [tr element]: https://html.spec.whatwg.org/#the-tr-element
func Tr(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("tr", attributes, children...)
}

// Td creates a [td element].
//
// The td element represents a data cell in a table.
//
// [td element]: https://html.spec.whatwg.org/#the-td-element
func Td(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("td", attributes, children...)
}

// Th creates a [th element].
//
// The th element represents a header cell in a table.
//
// [th element]: https://html.spec.whatwg.org/#the-th-element
func Th(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("th", attributes, children...)
}
