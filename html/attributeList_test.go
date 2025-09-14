package html

import (
	"fmt"

	"github.com/KrischanCS/roundly"
)

func ExampleAttributes() {
	el := Div(Attributes(Id("greeting"), Class("g1", "g2")))
	fmt.Println(el.String())
	// Output: <div id="greeting" class="g1 g2"></div>
}

func ExampleAttributes_none() {
	el := Div(Attributes())
	fmt.Println(el.String())
	// Output: <div></div>
}

func ExampleAttributes_nilSlice() {
	var attrs []roundly.Attribute
	el := Div(Attributes(attrs...))
	fmt.Println(el.String())
	// Output: <div></div>
}

func ExampleAttributes_emptyNonNil() {
	empty := make([]roundly.Attribute, 0)
	el := Div(Attributes(empty...))
	fmt.Println(el.String())
	// Output: <div></div>
}

func ExampleAttributes_single() {
	el := Div(Attributes(Id("main")))
	fmt.Println(el.String())
	// Output: <div id="main"></div>
}

//nolint:errcheck
func ExampleAttributes_pretty_multiline() {
	// Pretty rendering with 3 attributes triggers joinWithLineBrakes and multi-line formatting
	el := Div(Attributes(Id("main"), TitleAttribute("Hello"), TranslateYes()))

	w := roundly.NewWriter()
	_ = el.RenderElementWithOptions(w, &roundly.RenderOptions{Pretty: true})

	fmt.Println(w.String())
	// Output:
	//
	// <div
	// 	id="main"
	// 	title="Hello"
	// 	translate="yes"
	// >
	// </div>
}
