package elements

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/htmfunc"
	. "github.com/KrischanCS/htmfunc/attribute"
	element2 "github.com/KrischanCS/htmfunc/element"
)

func TestCanvas(t *testing.T) {
	t.Parallel()

	w := htmfunc.NewWriter(256)

	element := element2.Canvas(Attributes(Class("test", "other")))

	err := element.RenderElement(w)

	want := `<canvas class="test other"></canvas>`
	got := w.String()

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
