package elementArch

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KrischanCS/htmfunc"
	"github.com/KrischanCS/htmfunc/attribute"
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

	err := Data("42", nil, TextTrusted("42")).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<data value="42">42</data>`, w.String())
}

func TestTimeMachineReadableAsContent(t *testing.T) {
	t.Parallel()

	w := htmfunc.NewWriter(64)

	ts, err := time.Parse(time.RFC3339, "2024-12-24T12:34:56Z")
	require.NoError(t, err)

	err = TimeMachineReadableAsContent(nil, ts).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<time>2024-12-24T12:34:56Z</time>`, w.String())
}

func TestTimeAttribute(t *testing.T) {
	t.Parallel()

	w := htmfunc.NewWriter(64)

	ts, err := time.Parse(time.RFC3339, "2024-12-24T12:34:56Z")
	require.NoError(t, err)

	err = TimeAttribute(
		attribute.Class("time"),
		ts,
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

	err := Bdo(htmfunc.RightToLeft, nil, TextTrusted("مرحباً بالعالم")).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<bdo dir="rtl">مرحباً بالعالم</bdo>`, w.String())
}

func TestBdo_LTR(t *testing.T) {
	t.Parallel()

	w := htmfunc.NewWriter(64)

	err := Bdo(htmfunc.LeftToRight, nil, TextTrusted("مرحباً بالعالم")).RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<bdo dir="ltr">مرحباً بالعالم</bdo>`, w.String())
}

func TestBr(t *testing.T) {
	t.Parallel()

	w := htmfunc.NewWriter(64)

	err := Br().RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<br>`, w.String())
}
