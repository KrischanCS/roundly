package logic

import (
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/element"
	. "github.com/KrischanCS/roundly/text"
)

func TestIf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		condition bool
		element   roundly.Element
		want      string
	}{
		{
			false,
			H1(nil, Text("Test")),
			"",
		},
		{
			true,
			H1(nil, Text("Test")),
			"<h1>Test</h1>",
		},
	}

	for _, tc := range tests {
		t.Run(strconv.FormatBool(tc.condition), func(t *testing.T) {
			w := roundly.NewWriter()

			err := If(tc.condition, tc.element).RenderElement(w)

			assert.NoError(t, err)
			assert.Equal(t, tc.want, w.String())
		})
	}
}

func TestIfElse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		condition   bool
		elementIf   roundly.Element
		elementElse roundly.Element
		want        string
	}{
		{
			false,
			H1(nil, Text("TestIf")),
			H2(nil, Text("TestElse")),
			"<h2>TestElse</h2>",
		},
		{
			true,
			H1(nil, Text("Test")),
			H1(nil, Text("TestElse")),
			"<h1>Test</h1>",
		},
	}

	for _, tc := range tests {
		t.Run(strconv.FormatBool(tc.condition), func(t *testing.T) {
			w := roundly.NewWriter()

			err := IfElse(tc.condition, tc.elementIf, tc.elementElse).RenderElement(w)

			assert.NoError(t, err)
			assert.Equal(t, tc.want, w.String())
		})
	}
}

func TestIfValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		condition bool
		value     string
		want      string
	}{
		{
			false,
			"Test",
			"",
		},
		{
			true,
			"Test",
			"Test",
		},
	}

	for _, tc := range tests {
		t.Run(strconv.FormatBool(tc.condition), func(t *testing.T) {
			got := IfValue(tc.condition, tc.value)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestIfElseValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		condition bool
		valueIf   string
		valueElse string
		want      string
	}{
		{
			false,
			"TestIf",
			"TestElse",
			"TestElse",
		},
		{
			true,
			"TestIf",
			"TestElse",
			"TestIf",
		},
	}

	for _, tc := range tests {
		t.Run(strconv.FormatBool(tc.condition), func(t *testing.T) {
			got := IfElseValue(tc.condition, tc.valueIf, tc.valueElse)
			assert.Equal(t, tc.want, got)
		})
	}
}

func BenchmarkIf(b *testing.B) {
	w := roundly.NewWriter()

	r := rand.New(rand.NewPCG(1234, 5678)) //nolint:gosec
	conditions := make([]bool, b.N)

	for i := range b.N {
		conditions[i] = r.Float32() < 0.5
	}

	b.ResetTimer()
	b.ReportAllocs()

	var res []byte
	for i := range b.N { //nolint:wsl
		_ = If(conditions[i], H1(nil, Text("Test"))).RenderElement(w) //nolint:errcheck
		res = w.Bytes()
		w.Reset()
	}

	_ = res
}

func BenchmarkIfElse(b *testing.B) {
	w := roundly.NewWriter()

	r := rand.New(rand.NewPCG(1234, 5678)) //nolint:gosec
	conditions := make([]bool, b.N)

	for i := range b.N {
		conditions[i] = r.Float32() < 0.5
	}

	b.ResetTimer()
	b.ReportAllocs()

	var res []byte
	for i := range b.N { //nolint:wsl
		_ = IfElse(conditions[i],
			H1(nil, Text("TestIf")),
			H2(nil, Text("TestElse"))).
			RenderElement(w) //nolint:errcheck

		res = w.Bytes()
		w.Reset()
	}

	_ = res
}
