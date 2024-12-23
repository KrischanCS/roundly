package flow

import (
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ch-schulz/htmfunc"
	. "github.com/ch-schulz/htmfunc/element"
)

func TestIf(t *testing.T) {
	t.Parallel()

	tt := []struct {
		condition bool
		element   htmfunc.Element
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

	for _, tc := range tt {
		t.Run(strconv.FormatBool(tc.condition), func(t *testing.T) {
			w := htmfunc.NewWriter(4096)

			err := If(tc.condition, tc.element).RenderHTML(w)

			assert.NoError(t, err)
			assert.Equal(t, tc.want, w.String())
		})
	}
}

func BenchmarkIf(b *testing.B) {
	w := htmfunc.NewWriter(4096)

	r := rand.New(rand.NewPCG(1234, 5678)) //nolint:gosec
	conditions := make([]bool, b.N)

	for i := range b.N {
		conditions[i] = r.Float32() < 0.5
	}

	b.ResetTimer()
	b.ReportAllocs()

	var res []byte
	for i := range b.N { //nolint:wsl
		_ = If(conditions[i], H1(nil, Text("Test"))).RenderHTML(w) //nolint:errcheck
		res = w.Bytes()
		w.Reset()
	}

	_ = res
}
