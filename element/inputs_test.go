package element

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/htmfunc"
	. "github.com/KrischanCS/htmfunc/attribute"
)

func TestInputs(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		Form,
		Label,
		Button,
		Select,
		Datalist,
		Optgroup,
		Option,
		Textarea,
		Output,
		Progress,
		Meter,
		Fieldset,
		Legend,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}

func TestInput(t *testing.T) {
	t.Parallel()

	want := `<input value="test" class="rounded" type="button">`

	w := htmfunc.NewWriter(256)

	e := Input(Attributes(
		Value("test"),
		Class("rounded"),
		Type("button"),
	))

	err := e.RenderElement(w)

	got := w.String()

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
