package hfcomp

import (
	"fmt"

	"github.com/KrischanCS/roundly"
)

//nolint:gosmopolitan,errcheck
func ExampleRubyText() {
	rb := RubyText([]RubySegment{
		{"漢", "kan"},
		{"字", "ji"},
	})

	w := roundly.NewWriter()
	_ = rb.RenderElementWithOptions(w, &roundly.RenderOptions{Pretty: true})

	fmt.Println(w.String())

	// Output:
	//
	// <ruby>
	// 	漢
	// 	<rp>
	// 		(
	// 	</rp>
	// 	<rt>
	// 		漢
	// 	</rt>
	// 	<rp>
	// 		)
	// 	</rp>
	// 	字
	// 	<rp>
	// 		(
	// 	</rp>
	// 	<rt>
	// 		字
	// 	</rt>
	// 	<rp>
	// 		)
	// 	</rp>
	// </ruby>
}
