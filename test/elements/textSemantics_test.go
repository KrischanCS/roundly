package elements

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/html"
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
	}

	for _, element := range elements {
		t.Run(getFunctionName(element), func(t *testing.T) {
			elementTest(t, element)
		})
	}
}

func TestData(t *testing.T) {
	t.Parallel()

	w := roundly.NewWriter()

	err := Data(Value("42"), RawTrusted("42")).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<data value="42">42</data>`, w.String())
}

func TestTimeAttribute(t *testing.T) {
	t.Parallel()

	w := roundly.NewWriter()

	err := Time(
		Attributes(
			DateTime("2024-12-24T12:34:56Z"),
			Class("time"),
		),
		RawTrusted("24.12.2024 12:34:56"),
	).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t,
		`<time datetime="2024-12-24T12:34:56Z" class="time">24.12.2024 12:34:56</time>`,
		w.String())
}

func TestBdo_Rtl(t *testing.T) {
	t.Parallel()

	w := roundly.NewWriter()

	err := Bdo(DirRtl(), RawTrusted("مرحباً بالعالم")).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<bdo dir="rtl">مرحباً بالعالم</bdo>`, w.String())
}

func TestBdo_Ltr(t *testing.T) {
	t.Parallel()

	w := roundly.NewWriter()

	err := Bdo(DirLtr(), RawTrusted("مرحباً بالعالم")).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<bdo dir="ltr">مرحباً بالعالم</bdo>`, w.String())
}

func TestBr(t *testing.T) {
	t.Parallel()

	w := roundly.NewWriter()

	err := Br(nil).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<br>`, w.String())
}
