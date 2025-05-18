package element

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/htmfunc"
	"github.com/KrischanCS/htmfunc/attribute"
	"github.com/KrischanCS/htmfunc/text"
)

func TestTextSemantics(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		A,
		Em,
		Strong,
		Small,
		S,
		Cite,
		Q,
		Dfn,
		Abbr,
		Ruby,
		Rt,
		Rp,
		Code,
		Var,
		Samp,
		Kbd,
		Sub,
		Sup,
		I,
		B,
		U,
		Mark,
		Bdi,
		Span,
		Wbr,
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}

func TestData(t *testing.T) {
	t.Parallel()

	w := htmfunc.NewWriter(64)

	err := Data(attribute.Value_Data("42"), text.TextTrusted("42")).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<data value="42">42</data>`, w.String())
}

func TestTimeAttribute(t *testing.T) {
	t.Parallel()

	w := htmfunc.NewWriter(64)

	err := Time(
		attribute.Attributes(
			attribute.DateTime_Time("2024-12-24T12:34:56Z"),
			attribute.Class("time"),
		),
		text.TextTrusted("24.12.2024 12:34:56"),
	).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t,
		`<time datetime="2024-12-24T12:34:56Z" class="time">24.12.2024 12:34:56</time>`,
		w.String())
}

func TestBdo_RTL(t *testing.T) {
	t.Parallel()

	w := htmfunc.NewWriter(64)

	err := Bdo(attribute.Dir_Bdo(string(htmfunc.RightToLeft)), text.TextTrusted("مرحباً بالعالم")).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<bdo dir="rtl">مرحباً بالعالم</bdo>`, w.String())
}

func TestBdo_LTR(t *testing.T) {
	t.Parallel()

	w := htmfunc.NewWriter(64)

	err := Bdo(attribute.Dir_Bdo(string(htmfunc.LeftToRight)), text.TextTrusted("مرحباً بالعالم")).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<bdo dir="ltr">مرحباً بالعالم</bdo>`, w.String())
}

func TestBr(t *testing.T) {
	t.Parallel()

	w := htmfunc.NewWriter(64)

	err := Br(nil).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<br>`, w.String())
}
