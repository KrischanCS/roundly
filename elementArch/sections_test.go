package elementArch

import (
	"testing"
)

func TestSections(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		Body,
		Article,
		Section,
		Nav,
		Aside,
		H1,
		H2,
		H3,
		H4,
		H5,
		H6,
		Hgroup,
		Header,
		Footer,
		Address,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}
