package elements

import (
	"testing"

	"github.com/KrischanCS/htmfunc/element"
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
