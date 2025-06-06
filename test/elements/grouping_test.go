package elements

import (
	"testing"

	"github.com/KrischanCS/roundly/element"
)

func TestGrouping(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		element.Blockquote,
		element.Dd,
		element.Div,
		element.Dl,
		element.Dt,
		element.Figcaption,
		element.Figure,
		element.Li,
		element.Main,
		element.Menu,
		element.Ol,
		element.P,
		element.Pre,
		element.Search,
		element.Ul,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}
