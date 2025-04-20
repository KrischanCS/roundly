package element

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/htmfunc"
	attr "github.com/KrischanCS/htmfunc/attribute"
)

func TestCanvas(t *testing.T) {
	t.Parallel()

	w := htmfunc.NewWriter(256)

	element := Canvas(attr.Attributes(attr.Class("test", "other")))

	err := element.RenderElement(w)

	want := `<canvas class="test other"></canvas>`
	got := w.String()

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
