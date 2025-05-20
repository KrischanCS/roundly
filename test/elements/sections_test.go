package elements

import (
	"testing"

	"github.com/KrischanCS/htmfunc/element"
)

func TestSections(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		element.Body,
		element.Article,
		element.Section,
		element.Nav,
		element.Aside,
		element.H1,
		element.H2,
		element.H3,
		element.H4,
		element.H5,
		element.H6,
		element.Hgroup,
		element.Header,
		element.Footer,
		element.Address,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}
