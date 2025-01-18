//nolint:nolintlint,lll
package attribute

import (
	"github.com/KrischanCS/htmfunc"
	"strconv"
)

// Coords creates the coords attribute - Coordinates for the shape to be created in an [image map]
//
// It can be applied to the following elements:
//   - [area]
//
// Value constraints: [Valid list of floating-point numbers] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [area]: https://html.spec.whatwg.org/dev/image-maps.html#attr-area-coords
// [image map]: https://html.spec.whatwg.org/dev/image-maps.html#image-map
// [Valid list of floating-point numbers]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-list-of-floating-point-numbers
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Coords(coords ...float64) htmfunc.Attribute {
	floats := make([]string, len(coords))
	for i, val := range coords {
		floats[i] = strconv.FormatFloat(val, 'f', -1, 64)
	}

	return htmfunc.WriteMultiValueAttribute("coords", ',', floats...)
}
