package attribute

import (
    "github.com/ch-schulz/htmfunc"
    "strconv"
)

// Start creates the start attribute - Starting value of the list
//
// It can be applied to the following elements:
//   - [ol]
//
// Value constraints: [Valid integer]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [ol]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-ol-start
// [Valid integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Start(start int) htmfunc.Attribute {
    return htmfunc.WriteAttribute("start", strconv.Itoa(start))
}

// TabIndex creates the tabindex attribute - Whether the element is focusable and sequentially focusable, andthe relative order of the element for the purposes of sequential focus navigation
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Value constraints: [Valid integer]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/interaction.html#attr-tabindex
// [Valid integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func TabIndex(tabIndex int) htmfunc.Attribute {
    return htmfunc.WriteAttribute("tabindex", strconv.Itoa(tabIndex))
}

// Value_Li creates the value attribute - Ordinal value of the list item
//
// It can be applied to the following elements:
//   - [li]
//
// Value constraints: [Valid integer]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [li]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-li-value
// [Valid integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Value_Li(value int) htmfunc.Attribute {
    return htmfunc.WriteAttribute("value", strconv.Itoa(value))
}
