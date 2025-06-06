package elements

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/attribute"
	element2 "github.com/KrischanCS/roundly/element"
)

func TestCanvas(t *testing.T) {
	t.Parallel()

	w := roundly.NewWriter()

	element := element2.Canvas(Attributes(Class("test", "other")))

	err := element.RenderElement(w)

	want := `<canvas class="test other"></canvas>`
	got := w.String()

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
