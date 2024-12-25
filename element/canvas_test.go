package element

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ch-schulz/htmfunc"
	attr "github.com/ch-schulz/htmfunc/attribute"
)

func TestCanvas(t *testing.T) {
	t.Parallel()

	w := htmfunc.NewWriter(256)

	element := Canvas(attr.Join(attr.Class(attr.JoinValues("test", "other"))))

	err := element.RenderHtml(w)

	want := `<canvas class="test other"></canvas>`
	got := w.String()

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
