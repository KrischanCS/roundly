package elements

import (
	"testing"

	"github.com/KrischanCS/roundly/element"
)

func TestEmbed(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		element.Audio,
		element.Iframe,
		element.Map,
		element.Math,
		element.Object,
		element.Picture,
		element.Svg,
		element.Video,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}
