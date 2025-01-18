//nolint:nolintlint,lll
package attribute

import (
	"github.com/KrischanCS/htmfunc"
	"strconv"
)

// Cols creates the cols attribute - Maximum number of characters per line
//
// It can be applied to the following elements:
//   - [textarea]
//
// Value constraints: [Valid non-negative integer] greater than zero
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [textarea]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-cols
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Cols(cols uint) htmfunc.Attribute {
	return htmfunc.WriteAttribute("cols", strconv.FormatUint(uint64(cols), 10))
}

// ColSpan creates the colspan attribute - Number of columns that the cell is to span
//
// It can be applied to the following elements:
//   - [td]
//   - [th]
//
// Value constraints: [Valid non-negative integer] greater than zero
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [td]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-colspan
// [th]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-colspan
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ColSpan(colSpan uint) htmfunc.Attribute {
	return htmfunc.WriteAttribute("colspan", strconv.FormatUint(uint64(colSpan), 10))
}

// Height creates the height attribute - Vertical dimension
//
// It can be applied to the following elements:
//   - [canvas]
//   - [embed]
//   - [iframe]
//   - [img]
//   - [input]
//   - [object]
//   - [source] (in [picture])
//   - [video]
//
// Value constraints: [Valid non-negative integer]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [canvas]: https://html.spec.whatwg.org/dev/canvas.html#attr-canvas-height
// [embed]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [iframe]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [img]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [input]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [object]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [source]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [picture]: https://html.spec.whatwg.org/dev/embedded-content.html#the-picture-element
// [video]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Height(height uint) htmfunc.Attribute {
	return htmfunc.WriteAttribute("height", strconv.FormatUint(uint64(height), 10))
}

// MaxLength creates the maxlength attribute - Maximum [length] of value
//
// It can be applied to the following elements:
//   - [input]
//   - [textarea]
//
// Value constraints: [Valid non-negative integer]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-maxlength
// [textarea]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-maxlength
// [length]: https://html.spec.whatwg.org/dev/https://infra.spec.whatwg.org/#string-length
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func MaxLength(maxLength uint) htmfunc.Attribute {
	return htmfunc.WriteAttribute("maxlength", strconv.FormatUint(uint64(maxLength), 10))
}

// MinLength creates the minlength attribute - Minimum [length] of value
//
// It can be applied to the following elements:
//   - [input]
//   - [textarea]
//
// Value constraints: [Valid non-negative integer]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-minlength
// [textarea]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-minlength
// [length]: https://html.spec.whatwg.org/dev/https://infra.spec.whatwg.org/#string-length
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func MinLength(minLength uint) htmfunc.Attribute {
	return htmfunc.WriteAttribute("minlength", strconv.FormatUint(uint64(minLength), 10))
}

// Rows creates the rows attribute - Number of lines to show
//
// It can be applied to the following elements:
//   - [textarea]
//
// Value constraints: [Valid non-negative integer] greater than zero
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [textarea]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-rows
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Rows(rows uint) htmfunc.Attribute {
	return htmfunc.WriteAttribute("rows", strconv.FormatUint(uint64(rows), 10))
}

// RowSpan creates the rowspan attribute - Number of rows that the cell is to span
//
// It can be applied to the following elements:
//   - [td]
//   - [th]
//
// Value constraints: [Valid non-negative integer]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [td]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-rowspan
// [th]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-rowspan
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func RowSpan(rowSpan uint) htmfunc.Attribute {
	return htmfunc.WriteAttribute("rowspan", strconv.FormatUint(uint64(rowSpan), 10))
}

// Size creates the size attribute - Size of the control
//
// It can be applied to the following elements:
//   - [input]
//   - [select]
//
// Value constraints: [Valid non-negative integer] greater than zero
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-size
// [select]: https://html.spec.whatwg.org/dev/form-elements.html#attr-select-size
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Size(size uint) htmfunc.Attribute {
	return htmfunc.WriteAttribute("size", strconv.FormatUint(uint64(size), 10))
}

// Span creates the span attribute - Number of columns spanned by the element
//
// It can be applied to the following elements:
//   - [col]
//   - [colgroup]
//
// Value constraints: [Valid non-negative integer] greater than zero
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [col]: https://html.spec.whatwg.org/dev/tables.html#attr-col-span
// [colgroup]: https://html.spec.whatwg.org/dev/tables.html#attr-colgroup-span
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Span(span uint) htmfunc.Attribute {
	return htmfunc.WriteAttribute("span", strconv.FormatUint(uint64(span), 10))
}

// Width creates the width attribute - Horizontal dimension
//
// It can be applied to the following elements:
//   - [canvas]
//   - [embed]
//   - [iframe]
//   - [img]
//   - [input]
//   - [object]
//   - [source] (in [picture])
//   - [video]
//
// Value constraints: [Valid non-negative integer]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [canvas]: https://html.spec.whatwg.org/dev/canvas.html#attr-canvas-width
// [embed]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [iframe]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [img]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [input]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [object]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [source]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [picture]: https://html.spec.whatwg.org/dev/embedded-content.html#the-picture-element
// [video]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Width(width uint) htmfunc.Attribute {
	return htmfunc.WriteAttribute("width", strconv.FormatUint(uint64(width), 10))
}
