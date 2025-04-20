// Generated file. DO NOT EDIT.

package attribute

import (
	"github.com/KrischanCS/htmfunc"
)

// Type_Input creates the type attribute - Type of form control
//
// It can be applied to the following elements:
//   - [input]
//
// Value constraints: [input type keyword]. Must be one of the following:
//   - hidden
//   - text
//   - search
//   - tel
//   - url
//   - email
//   - password
//   - date
//   - month
//   - week
//   - time
//   - datetime
//   - number
//   - range
//   - color
//   - checkbox
//   - radio
//   - file
//   - submit
//   - image
//   - reset
//   - button
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-type
// [input type keyword]: https://html.spec.whatwg.org/dev/input.html#attr-input-type
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Type_Input(typeV string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("type", typeV)
}
