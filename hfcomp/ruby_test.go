package hfcomp

import (
	"fmt"

	"github.com/yosssi/gohtml"

	"github.com/KrischanCS/htmfunc"
)

//nolint:gosmopolitan,errcheck
func ExampleRubyText() {
	rb := RubyText([]RubySegment{
		{"漢", "kan"},
		{"字", "ji"},
	})

	w := htmfunc.NewWriter()
	_ = rb.RenderElement(w)

	formatted := gohtml.Format(w.String())

	fmt.Println(formatted)

	// Output:
	// <ruby>
	//   漢
	//   <rp>
	//     (
	//   </rp>
	//   <rt>
	//     kan
	//   </rt>
	//   <rp>
	//     )
	//   </rp>
	//   字
	//   <rp>
	//     (
	//   </rp>
	//   <rt>
	//     ji
	//   </rt>
	//   <rp>
	//     )
	//   </rp>
	// </ruby>
}
