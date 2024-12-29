package generated

import (
    "github.com/ch-schulz/htmfunc"
    "github.com/ch-schulz/htmfunc/attribute"
)

// Type creates the type attribute - Type of form control
//
// It can be applied to the following elements:
//   - [[input]]
//
// Value constraints: [input type keyword]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-type
// [input type keyword]: https://html.spec.whatwg.org/dev/input.html#attr-input-type
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Type(typeV string) htmfunc.AttributeRenderer {
    return attribute.Attribute("type", typeV)
}
