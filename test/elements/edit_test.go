package elements

import (
	"testing"

	"github.com/KrischanCS/roundly/element"
)

func TestEdit(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		element.Ins,
		element.Del,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}
