package logic

import (
	"fmt"
	"iter"
	"strconv"
	"testing"

	"github.com/KrischanCS/go-toolbox/iterator"
	"github.com/stretchr/testify/assert"
	"github.com/yosssi/gohtml"

	"github.com/KrischanCS/htmfunc"
	. "github.com/KrischanCS/htmfunc/attribute"
	. "github.com/KrischanCS/htmfunc/element"
	. "github.com/KrischanCS/htmfunc/text"
)

//nolint:errcheck
func ExampleRange() {
	months := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}

	list := Ol(nil,
		Range(months, func(_ int, month string) htmfunc.Element {
			return Li(nil, Text(month))
		}),
	)

	w := htmfunc.NewWriter()
	_ = list.RenderElement(w)

	fmt.Println(gohtml.Format(w.String()))

	// Output:
	// <ol>
	//   <li>
	//     January
	//   </li>
	//   <li>
	//     February
	//   </li>
	//   <li>
	//     March
	//   </li>
	//   <li>
	//     April
	//   </li>
	//   <li>
	//     May
	//   </li>
	//   <li>
	//     June
	//   </li>
	//   <li>
	//     July
	//   </li>
	//   <li>
	//     August
	//   </li>
	//   <li>
	//     September
	//   </li>
	//   <li>
	//     October
	//   </li>
	//   <li>
	//     November
	//   </li>
	//   <li>
	//     December
	//   </li>
	// </ol>
}

//nolint:errcheck
func ExampleRangeInt() {
	list := Ol(nil,
		RangeInt(5, func(i int) htmfunc.Element {
			return Li(nil, Text(strconv.Itoa(i)))
		}),
	)

	w := htmfunc.NewWriter()
	_ = list.RenderElement(w)

	fmt.Println(gohtml.Format(w.String()))

	// Output:
	// <ol>
	//   <li>
	//     0
	//   </li>
	//   <li>
	//     1
	//   </li>
	//   <li>
	//     2
	//   </li>
	//   <li>
	//     3
	//   </li>
	//   <li>
	//     4
	//   </li>
	// </ol>
}

//nolint:errcheck
func ExampleRangeSeq() {
	list := Ol(nil,
		RangeSeq(iterator.FromStepTo(0, 0.1, 0.5), func(f float64) htmfunc.Element {
			return Li(nil, Text(strconv.FormatFloat(f, 'f', 1, 64)))
		}),
	)

	w := htmfunc.NewWriter()
	_ = list.RenderElement(w)

	fmt.Println(gohtml.Format(w.String()))

	// Output:
	// <ol>
	//   <li>
	//     0.0
	//   </li>
	//   <li>
	//     0.1
	//   </li>
	//   <li>
	//     0.2
	//   </li>
	//   <li>
	//     0.3
	//   </li>
	//   <li>
	//     0.4
	//   </li>
	// </ol>
}

//nolint:funlen
func TestRange(t *testing.T) {
	t.Parallel()

	type args[T any] struct {
		items     []T
		component func(int, T) htmfunc.Element
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
				component: func(i int, s string) htmfunc.Element {
					return Li(nil,
						Div(Class("number"), Text(strconv.Itoa(i+1))),
						Div(Class("value"), Text(s)),
					)
				},
			},
			"",
		},
		{
			"oneItem",
			args[string]{
				items: []string{"apples"},
				component: func(i int, s string) htmfunc.Element {
					return Li(nil,
						Div(Class("number"), Text(strconv.Itoa(i+1))),
						Div(Class("value"), Text(s)),
					)
				},
			},
			`<li><div class="number">1</div><div class="value">apples</div></li>`,
		},
		{
			"ThreeItems",
			args[string]{
				items: []string{"apples", "bananas", "oranges"},
				component: func(i int, s string) htmfunc.Element {
					return Li(nil,
						Div(Class("number"), Text(strconv.Itoa(i+1))),
						Div(Class("value"), Text(s)),
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
			w := htmfunc.NewWriter()

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
		component func(int) htmfunc.Element
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
				component: func(i int) htmfunc.Element {
					return Li(nil, Text(strconv.Itoa(i)))
				},
			},
			want: "",
		},
		{
			name: "1",
			args: args{
				limit: 1,
				component: func(i int) htmfunc.Element {
					return Li(nil, Text(strconv.Itoa(i)))
				},
			},
			want: "<li>0</li>",
		},
		{
			name: "3",
			args: args{
				limit: 3,
				component: func(i int) htmfunc.Element {
					return Li(nil, Text(strconv.Itoa(i)))
				},
			},
			want: "<li>0</li><li>1</li><li>2</li>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := htmfunc.NewWriter()

			err := RangeInt(tt.args.limit, tt.args.component).RenderElement(w)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

//nolint:funlen
func TestRangeIter(t *testing.T) {
	t.Parallel()

	type args struct {
		seq       iter.Seq[int]
		component func(int) htmfunc.Element
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
				component: func(s int) htmfunc.Element {
					return Li(nil,
						Text(fmt.Sprintf("%c", 'a'+s)))
				},
			},
			want: "",
		},
		{
			name: "1",
			args: args{
				seq: iterator.FromToInclusive(7, 7),
				component: func(s int) htmfunc.Element {
					return Li(nil,
						Text(fmt.Sprintf("%c", 'a'+s-1)))
				},
			},
			want: `<li>g</li>`,
		},
		{
			name: "5",
			args: args{
				seq: iterator.FromTo(3, 8),
				component: func(s int) htmfunc.Element {
					return Li(nil,
						Text(fmt.Sprintf("%c", 'a'+s-1)))
				},
			},
			want: `<li>c</li>` +
				`<li>d</li>` +
				`<li>e</li>` +
				`<li>f</li>` +
				`<li>g</li>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := htmfunc.NewWriter()

			err := RangeSeq(tt.args.seq, tt.args.component).RenderElement(w)

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

	w := htmfunc.NewWriter()

	b.ResetTimer()
	b.ReportAllocs()

	var res []byte

	for b.Loop() {
		_ = Range(grid, func(_ int, row []int) htmfunc.Element { //nolint:errcheck
			return Div(Class("row"),
				Range(row, func(_ int, i int) htmfunc.Element {
					return Div(Class("col"),
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
	w := htmfunc.NewWriter()

	b.ResetTimer()
	b.ReportAllocs()

	var res []byte

	for b.Loop() {
		_ = RangeInt(10, func(row int) htmfunc.Element { //nolint:errcheck
			return Div(Class("row"),
				RangeInt(20, func(col int) htmfunc.Element {
					return Div(Class("col"),
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
	w := htmfunc.NewWriter()

	b.ResetTimer()
	b.ReportAllocs()

	var res []byte

	for b.Loop() {
		_ = RangeSeq(iterator.FromTo(0, 10), func(row int) htmfunc.Element { //nolint:errcheck
			return Div(Class("row"),
				RangeSeq(iterator.FromTo(0, 20), func(col int) htmfunc.Element {
					return Div(Class("col"),
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
