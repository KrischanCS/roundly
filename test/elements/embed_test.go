package elements

import (
	"testing"

	. "github.com/KrischanCS/roundly/html"
)

func TestEmbed(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		Audio,
		Iframe,
		Map,
		Math,
		Object,
		Picture,
		Svg,
		Video,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}
