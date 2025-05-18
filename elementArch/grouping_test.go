package elementArch

import (
	"testing"
)

func TestGrouping(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		Blockquote,
		Dd,
		Div,
		Dl,
		Dt,
		Figcaption,
		Figure,
		Hr,
		Li,
		Main,
		Menu,
		Ol,
		P,
		Pre,
		Search,
		Ul,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}
