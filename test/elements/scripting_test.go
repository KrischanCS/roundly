package elements

import (
	"testing"

	"github.com/KrischanCS/roundly/element"
)

func TestScripting(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		element.Noscript,
		element.Script,
		element.Slot,
		element.Template,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}
