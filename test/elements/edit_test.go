package elements

import (
	"testing"

	. "github.com/KrischanCS/roundly/html"
)

func TestEdit(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		Ins,
		Del,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}
