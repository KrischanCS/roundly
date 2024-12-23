package flow

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/bytebufferpool"

	"github.com/ch-schulz/htmfunc"
	. "github.com/ch-schulz/htmfunc/element"
)

func TestIf(t *testing.T) {

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
		t.Run(fmt.Sprint(tc.condition), func(t *testing.T) {
			w := bytebufferpool.Get()
			err := If(tc.condition, tc.element)(w)

			assert.NoError(t, err)
			assert.Equal(t, tc.want, w.String())
		})
	}
}

func BenchmarkIf(b *testing.B) {

	w := bytebufferpool.Get()

	r := rand.New(rand.NewPCG(1234, 5678))
	conditions := make([]bool, b.N)
	for i := range b.N {
		conditions[i] = r.Float32() < 0.5
	}

	b.ResetTimer()
	b.ReportAllocs()

	var res []byte
	for i := range b.N {
		_ = If(conditions[i], H1(nil, Text("Test")))(w)
		res = w.Bytes()
		w.Reset()
	}

	_ = res
}
