package elements

import (
	"testing"

	. "github.com/KrischanCS/roundly/html"
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
