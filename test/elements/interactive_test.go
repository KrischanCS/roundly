package elements

import (
	"testing"

	. "github.com/KrischanCS/roundly/html"
)

func TestInteractive(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		Details,
		Summary,
		Dialog,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}
