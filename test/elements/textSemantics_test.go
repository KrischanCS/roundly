package elements

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/htmfunc"
	. "github.com/KrischanCS/htmfunc/attribute"
	"github.com/KrischanCS/htmfunc/element"
	. "github.com/KrischanCS/htmfunc/text"
)

func TestTextSemantics(t *testing.T) {
	t.Parallel()

	elements := []elementFunc{
		element.A,
		element.Em,
		element.Strong,
		element.Small,
		element.S,
		element.Cite,
		element.Q,
		element.Dfn,
		element.Abbr,
		element.Ruby,
		element.Rt,
		element.Rp,
		element.Code,
		element.Var,
		element.Samp,
		element.Kbd,
		element.Sub,
		element.Sup,
		element.I,
		element.B,
		element.U,
		element.Mark,
		element.Bdi,
		element.Span,
		element.Wbr,
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

	err := element.Data(Value("42"), TextTrusted("42")).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<data value="42">42</data>`, w.String())
}

func TestTimeAttribute(t *testing.T) {
	t.Parallel()

	w := htmfunc.NewWriter(64)

	err := element.Time(
		Attributes(
			DateTime("2024-12-24T12:34:56Z"),
			Class("time"),
		),
		TextTrusted("24.12.2024 12:34:56"),
	).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t,
		`<time datetime="2024-12-24T12:34:56Z" class="time">24.12.2024 12:34:56</time>`,
		w.String())
}

func TestBdo_RTL(t *testing.T) {
	t.Parallel()

	w := htmfunc.NewWriter(64)

	err := element.Bdo(DirRtl(), TextTrusted("مرحباً بالعالم")).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<bdo dir="rtl">مرحباً بالعالم</bdo>`, w.String())
}

func TestBdo_LTR(t *testing.T) {
	t.Parallel()

	w := htmfunc.NewWriter(64)

	err := element.Bdo(DirLtr(), TextTrusted("مرحباً بالعالم")).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<bdo dir="ltr">مرحباً بالعالم</bdo>`, w.String())
}

func TestBr(t *testing.T) {
	t.Parallel()

	w := htmfunc.NewWriter(64)

	err := element.Br(nil).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<br>`, w.String())
}
