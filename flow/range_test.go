package flow

import (
	"fmt"
	"iter"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ch-schulz/htmfunc"
	attr "github.com/ch-schulz/htmfunc/attribute"
	. "github.com/ch-schulz/htmfunc/element"
	"github.com/ch-schulz/htmfunc/iters"
)

func TestRange(t *testing.T) {
	t.Parallel()

	type args[T any] struct {
		items     []T
		component func(int, T) htmfunc.ElementRenderer
	}

	type testCase[T any] struct {
		name string
		args args[T]
		want string
	}

	tests := []testCase[string]{
		{
			"zeroItems",
			args[string]{
				items: nil,
				component: func(i int, s string) htmfunc.ElementRenderer {
					return Li(nil,
						Div(attr.Class(attr.JoinValues("number")), Text(strconv.Itoa(i+1))),
						Div(attr.Class(attr.JoinValues("value")), Text(s)),
					)
				},
			},
			"",
		},
		{
			"oneItem",
			args[string]{
				items: []string{"apples"},
				component: func(i int, s string) htmfunc.ElementRenderer {
					return Li(nil,
						Div(attr.Class(attr.JoinValues("number")), Text(strconv.Itoa(i+1))),
						Div(attr.Class(attr.JoinValues("value")), Text(s)),
					)
				},
			},
			`<li><div class="number">1</div><div class="value">apples</div></li>`,
		},
		{
			"ThreeItems",
			args[string]{
				items: []string{"apples", "bananas", "oranges"},
				component: func(i int, s string) htmfunc.ElementRenderer {
					return Li(nil,
						Div(attr.Class(attr.JoinValues("number")), Text(strconv.Itoa(i+1))),
						Div(attr.Class(attr.JoinValues("value")), Text(s)),
					)
				},
			},
			`<li><div class="number">1</div><div class="value">apples</div></li>` +
				`<li><div class="number">2</div><div class="value">bananas</div></li>` +
				`<li><div class="number">3</div><div class="value">oranges</div></li>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := htmfunc.NewWriter(4096)

			err := Range(tt.args.items, tt.args.component).RenderElement(w)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

func TestRangeInt(t *testing.T) {
	t.Parallel()

	type args struct {
		limit     int
		component func(int) htmfunc.ElementRenderer
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{
				limit: 0,
				component: func(i int) htmfunc.ElementRenderer {
					return Li(nil, Text(strconv.Itoa(i)))
				},
			},
			want: "",
		},
		{
			name: "1",
			args: args{
				limit: 1,
				component: func(i int) htmfunc.ElementRenderer {
					return Li(nil, Text(strconv.Itoa(i)))
				},
			},
			want: "<li>0</li>",
		},
		{
			name: "3",
			args: args{
				limit: 3,
				component: func(i int) htmfunc.ElementRenderer {
					return Li(nil, Text(strconv.Itoa(i)))
				},
			},
			want: "<li>0</li><li>1</li><li>2</li>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := htmfunc.NewWriter(4096)

			err := RangeInt(tt.args.limit, tt.args.component).RenderElement(w)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

func TestRangeIter(t *testing.T) {
	t.Parallel()

	type args struct {
		seq       iter.Seq2[int, int]
		component func(int, int) htmfunc.ElementRenderer
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nil",
			args: args{
				seq: nil,
				component: func(i int, s int) htmfunc.ElementRenderer {
					return Li(nil,
						Text(fmt.Sprintf("%d - %c", i, 'a'+i)))
				},
			},
			want: "",
		},
		{
			name: "1",
			args: args{
				seq: iters.FromToInclusive(7, 7),
				component: func(i int, s int) htmfunc.ElementRenderer {
					return Li(nil,
						Text(fmt.Sprintf("%d - %c", i+1, 'a'+s-1)))
				},
			},
			want: `<li>1 - g</li>`,
		},
		{
			name: "5",
			args: args{
				seq: iters.FromTo(3, 8),
				component: func(i int, s int) htmfunc.ElementRenderer {
					return Li(nil,
						Text(fmt.Sprintf("%d - %c", i+1, 'a'+s-1)))
				},
			},
			want: `<li>1 - c</li>` +
				`<li>2 - d</li>` +
				`<li>3 - e</li>` +
				`<li>4 - f</li>` +
				`<li>5 - g</li>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := htmfunc.NewWriter(4096)

			err := RangeIter(tt.args.seq, tt.args.component).RenderElement(w)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

func BenchmarkRange(b *testing.B) {
	grid := make([][]int, 10)
	for i := range 10 {
		grid[i] = make([]int, 20)
		for j := range 20 {
			grid[i][j] = i*100 + j
		}
	}

	w := htmfunc.NewWriter(4096)

	b.ResetTimer()
	b.ReportAllocs()

	var res []byte

	for range b.N {
		_ = Range(grid, func(_ int, row []int) htmfunc.ElementRenderer { //nolint:errcheck
			return Div(attr.Class(attr.JoinValues("row")),
				Range(row, func(_ int, i int) htmfunc.ElementRenderer {
					return Div(attr.Class(attr.JoinValues("col")),
						Text(strconv.Itoa(i)),
					)
				}),
			)
		}).RenderElement(w)

		res = w.Bytes()
		w.Reset()
	}

	_ = res
}

func BenchmarkRangeInt(b *testing.B) {
	w := htmfunc.NewWriter(4096)

	b.ResetTimer()
	b.ReportAllocs()

	var res []byte

	for range b.N {
		_ = RangeInt(10, func(row int) htmfunc.ElementRenderer { //nolint:errcheck
			return Div(attr.Class(attr.JoinValues("row")),
				RangeInt(20, func(col int) htmfunc.ElementRenderer {
					return Div(attr.Class(attr.JoinValues("col")),
						Text(strconv.Itoa(row*100+col)),
					)
				}),
			)
		}).RenderElement(w)

		res = w.Bytes()
		w.Reset()
	}

	_ = res
}

func BenchmarkRangeIter(b *testing.B) {
	w := htmfunc.NewWriter(4096)

	b.ResetTimer()
	b.ReportAllocs()

	var res []byte

	for range b.N {
		_ = RangeIter(iters.FromTo(0, 10), func(_ int, row int) htmfunc.ElementRenderer { //nolint:errcheck
			return Div(attr.Class(attr.JoinValues("row")),
				RangeIter(iters.FromTo(0, 20), func(_ int, col int) htmfunc.ElementRenderer {
					return Div(attr.Class(attr.JoinValues("col")),
						Text(strconv.Itoa(row*100+col)),
					)
				}),
			)
		}).RenderElement(w)

		res = w.Bytes()
		w.Reset()
	}

	_ = res
}
