package elements

import (
	"testing"

	"github.com/KrischanCS/roundly/element"
)

func TestInteractive(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		element.Details,
		element.Summary,
		element.Dialog,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}
