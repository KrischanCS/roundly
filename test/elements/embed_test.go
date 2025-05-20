package elements

import (
	"testing"

	"github.com/KrischanCS/htmfunc/element"
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
		element.Source,
		element.Svg,
		element.Track,
		element.Video,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}
