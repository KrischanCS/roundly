package example

import (
	"fmt"

	"github.com/KrischanCS/htmfunc"
	. "github.com/KrischanCS/htmfunc/attribute"
	. "github.com/KrischanCS/htmfunc/element"
	. "github.com/KrischanCS/htmfunc/text"
)

//nolint:lll
func ExampleTable() {
	w := htmfunc.NewWriter(256)

	element := Table(
		Class("test"),
		Caption(nil, Text("Test")),
		Thead(nil,
			Tr(nil,
				Th(nil, Text("Test")),
				Th(nil, Text("Test")),
			),
		),
		Tbody(nil,
			Tr(nil,
				Th(ScopeRow(), Text("Test")),
				Td(nil, Text("Test")),
			),
		),
	)

	err := element.RenderElement(w)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output:
	// <table class="test"><caption>Test</caption><thead><tr><th>Test</th><th>Test</th></tr></thead><tbody><tr><th scope="row">Test</th><td>Test</td></tr></tbody></table>
}
