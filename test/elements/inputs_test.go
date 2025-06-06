package elements

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/attribute"
	"github.com/KrischanCS/roundly/element"
)

func TestInputs(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		element.Form,
		element.Label,
		element.Button,
		element.Select,
		element.Datalist,
		element.Optgroup,
		element.Option,
		element.Textarea,
		element.Output,
		element.Progress,
		element.Meter,
		element.Fieldset,
		element.Legend,
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

	w := roundly.NewWriter()

	e := element.Input(Attributes(
		Value("test"),
		Class("rounded"),
		Type("button"),
	))

	err := e.RenderElement(w)

	got := w.String()

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
