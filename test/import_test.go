package test

import (
	"testing"

	. "github.com/KrischanCS/htmfunc/attribute"
	. "github.com/KrischanCS/htmfunc/element"
	. "github.com/KrischanCS/htmfunc/flow"
	. "github.com/KrischanCS/htmfunc/text"
)

func TestShouldCompileWhenAllSubpackagesAreDotImported(t *testing.T) {
	// This test is only to ensure that all subpackages can be dot-imported at
	// once. There are some name collisions between attributes and elements, so
	// this makes sure all collisions are handled.
	t.Parallel()

	// Using a function from each package to avoid the unused imports being removed.
	_ = []any{
		Html,      // Function from element package
		Class,     // Function from attribute package
		Text,      // Function from text package
		RangeIter, // Function from flow package
	}
}
