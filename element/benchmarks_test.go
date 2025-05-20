package element

import (
	"strconv"
	"testing"

	"github.com/KrischanCS/go-toolbox/iterator"

	"github.com/KrischanCS/htmfunc"
	. "github.com/KrischanCS/htmfunc/attribute"
	. "github.com/KrischanCS/htmfunc/flow"
	. "github.com/KrischanCS/htmfunc/text"
)

//nolint:errcheck
func BenchmarkExamplePage(b *testing.B) {
	w := htmfunc.NewWriter(4096)

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		w.Reset()

		page := Html(
			Lang("en"),
			Head(nil,
				Title(nil, TextTrusted("The Title of the Page")),
			),
			Body(nil,
				Nav(nil,
					A(HRef("/main"), TextTrusted(("Main")),
						A(HRef("/details"), TextTrusted(("Details"))),
						Main(nil,
							H1(nil,
								Div(nil, Text("Here could be your content")),
							),
						),
					),
				),
			),
		)

		_ = page.RenderElement(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageRange10(b *testing.B) {
	w := htmfunc.NewWriter(4096)

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		w.Reset()

		page := Html(
			Lang("en"),
			Head(
				nil,
				Title(
					nil, TextTrusted("The Title of the Page"),
				),
			),
			Body(nil,
				Nav(nil,
					A(HRef("/main"), TextTrusted("Main")),
					A(HRef("/details"), TextTrusted("Details")),
				),
				Main(nil,
					H1(nil,
						Div(nil, Text("Here could be your content")),
					),
					RangeIter(iterator.FromToInclusive(1, 10), func(i int) htmfunc.Element {
						return Div(nil, Text(strconv.Itoa(i)))
					}),
				),
			),
		)

		_ = page.RenderElement(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageNoEscape(b *testing.B) {
	w := htmfunc.NewWriter(4096)

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		w.Reset()

		page := Html(
			Lang("en"),
			Head(
				nil, TextTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(HRef("/main"), TextTrusted("Main")),
					A(HRef("/details"), TextTrusted("Details")),
				),
				Main(nil,
					H1(nil,
						Div(nil, TextTrusted("Here could be your content")),
					),
				),
			),
		)

		_ = page.RenderElement(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageWriteOnly(b *testing.B) {
	w := htmfunc.NewWriter(4096)

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		w.Reset()

		page := Html(
			Lang("en"),
			Head(
				nil, TextTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(HRef("/main"), TextTrusted("Main")),
					A(HRef("/details"), TextTrusted("Details")),
				),
				Main(nil,
					H1(nil,
						Div(nil, Text("Here could be your content")),
					),
				),
			),
		)

		_ = page.RenderElement(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageWriteOnlyNoEscape(b *testing.B) {
	w := htmfunc.NewWriter(4096)

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		w.Reset()

		page := Html(
			Lang("en"),
			Head(
				nil, TextTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(HRef("/main"), TextTrusted("Main")),
					A(HRef("/details"), TextTrusted("Details")),
				),
				Main(nil,
					H1(nil,
						Div(nil, TextTrusted("Here could be your content")),
					),
				),
			),
		)

		_ = page.RenderElement(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkRange(b *testing.B) {
	w := htmfunc.NewWriter(4096)

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		w.Reset()

		page := Html(
			Lang("en"),
			Head(
				nil, Text("The Title of the Page"),
			),
			Body(nil,
				RangeIter(iterator.FromToInclusive(1, 1), func(i int) htmfunc.Element {
					return Div(nil, Text(strconv.Itoa(i)))
				}),
			),
		)

		_ = page.RenderElement(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkYearCalendar(b *testing.B) {
	months := []monthDays{
		{"January", 31},
		{"February", 29},
		{"March", 31},
		{"April", 30},
		{"May", 31},
		{"June", 30},
		{"July", 31},
		{"August", 31},
		{"September", 30},
		{"October", 31},
		{"November", 30},
		{"December", 31},
	}

	w := htmfunc.NewWriter(4096)

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		page := Html(
			Lang("en"),
			Head(
				nil, TextTrusted("The Title of the Page"),
			),
			Body(nil,
				Div(Class("header"),
					Nav(nil,
						A(HRef("/main"), TextTrusted("Main")),
						A(HRef("/details"), TextTrusted("Details")),
					),
					H1(nil, Text("Calendar")),
					H2(nil, Text("2024")),
				),
				Main(nil,
					Div(Class("year"),
						Range(months, month),
					),
				),
			),
		)

		w.Reset()

		_ = page.RenderElement(w)
	}

	_ = w.String()
}

type monthDays struct {
	name string
	days int
}

func month(_ int, month monthDays) htmfunc.Element {
	return Div(Class("month"),
		H3(nil, Text(month.name)),
		Div(Class("days"),
			RangeInt(month.days, func(i int) htmfunc.Element {
				return Div(Class("day"), Text(strconv.Itoa(i+1)))
			}),
		),
	)
}
