package elements

import (
	"fmt"
	"testing"

	"github.com/KrischanCS/htmfunc"
	. "github.com/KrischanCS/htmfunc/attribute"
	"github.com/KrischanCS/htmfunc/element"
	. "github.com/KrischanCS/htmfunc/text"
)

func TestTable(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		element.Table,
		element.Caption,
		element.Colgroup,
		element.Tbody,
		element.Thead,
		element.Tfoot,
		element.Tr,
		element.Td,
		element.Th,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}

//nolint:lll
func ExampleTable() {
	w := htmfunc.NewWriter(256)

	element := element.Table(
		Attributes(Class("test")),
		element.Caption(nil, Text("Test")),
		element.Thead(nil,
			element.Tr(nil,
				element.Th(nil, Text("Test"))),
			element.Th(nil, Text("Test")),
		),
		element.Tbody(nil,
			element.Tr(nil,
				element.Td(nil, Text("Test")),
				element.Td(nil, Text("Test")),
			),
		),
	)

	_ = element.RenderElement(w)

	fmt.Println(w.String())

	// Output:
	// <table class="test"><caption>Test</caption><thead><tr><th>Test</th></tr><th>Test</th></thead><tbody><tr><td>Test</td><td>Test</td></tr></tbody></table>
}
