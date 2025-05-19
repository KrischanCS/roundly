package element

import (
	"fmt"
	"testing"

	"github.com/KrischanCS/htmfunc"
	. "github.com/KrischanCS/htmfunc/attribute"
	. "github.com/KrischanCS/htmfunc/text"
)

func TestTable(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		Table,
		Caption,
		Colgroup,
		Tbody,
		Thead,
		Tfoot,
		Tr,
		Td,
		Th,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}

func ExampleTable() {
	w := htmfunc.NewWriter(256)

	element := Table(
		Attributes(Class("test")),
		Caption(nil, Text("Test")),
		Tbody(nil,
			Tr(nil,
				Td(nil, Text("Test")),
				Td(nil, Text("Test")),
			),
		),
	)

	_ = element.RenderElement(w)

	fmt.Println(w.String())

	// Output:
	// <table class="test"><caption>Test</caption><tbody><tr><td>Test</td><td>Test</td></tr></tbody></table>
}
