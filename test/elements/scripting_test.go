package elements

import (
	"testing"

	. "github.com/KrischanCS/roundly/html"
)

func TestScripting(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		Noscript,
		Script,
		Slot,
		Template,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}
