package el

import (
	"github.com/ch-schulz/htmfunc"
	"github.com/ch-schulz/htmfunc/attr"
)

// Canvas creates a [canvas element].
//
// The canvas element provides scripts with a resolution-dependent bitmap canvas, which can be used for rendering
// graphs, game graphics, art, or other visual images on the fly.
//
// [canvas element]: https://html.spec.whatwg.org/#the-canvas-element
func Canvas(attributes attr.Ls, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("canvas", attributes, children...)
}
