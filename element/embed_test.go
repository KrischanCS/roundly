package element

import (
	"testing"
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
		Source,
		Svg,
		Track,
		Video,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}
