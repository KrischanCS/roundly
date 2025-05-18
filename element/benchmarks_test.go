package element

import (
	"strconv"
	"testing"

	"github.com/KrischanCS/htmfunc"
	"github.com/KrischanCS/htmfunc/attribute"
	"github.com/KrischanCS/htmfunc/flow"
	"github.com/KrischanCS/htmfunc/iters"
	"github.com/KrischanCS/htmfunc/text"
)

//nolint:errcheck
func BenchmarkExamplePage(b *testing.B) {
	w := htmfunc.NewWriter(4096)

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		w.Reset()

		page := Html(
			attribute.Lang("en"),
			Head(nil,
				Title(nil, text.TextTrusted("The Title of the Page")),
			),
			Body(nil,
				Nav(nil,
					A(attribute.HRef_AArea("/main"), text.TextTrusted(("Main")),
						A(attribute.HRef_AArea("/details"), text.TextTrusted(("Details"))),
						Main(nil,
							H1(nil,
								Div(nil, text.Text("Here could be your content")),
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
			attribute.Lang("en"),
			Head(
				nil,
				Title(
					nil,
					text.TextTrusted("The Title of the Page"),
				),
			),
			Body(nil,
				Nav(nil,
					A(attribute.HRef_AArea("/main"), text.TextTrusted("Main")),
					A(attribute.HRef_AArea("/details"), text.TextTrusted("Details")),
				),
				Main(nil,
					H1(nil,
						Div(nil, text.Text("Here could be your content")),
					),
					flow.RangeIter(iters.FromToInclusive(1, 10), func(_, i int) htmfunc.Element {
						return Div(nil, text.Text(strconv.Itoa(i)))
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
			attribute.Lang("en"),
			Head(
				nil,
				text.TextTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(attribute.HRef_AArea("/main"), text.TextTrusted("Main")),
					A(attribute.HRef_AArea("/details"), text.TextTrusted("Details")),
				),
				Main(nil,
					H1(nil,
						Div(nil, text.TextTrusted("Here could be your content")),
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
			attribute.Lang("en"),
			Head(
				nil,
				text.TextTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(attribute.HRef_AArea("/main"), text.TextTrusted("Main")),
					A(attribute.HRef_AArea("/details"), text.TextTrusted("Details")),
				),
				Main(nil,
					H1(nil,
						Div(nil, text.Text("Here could be your content")),
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
			attribute.Lang("en"),
			Head(
				nil,
				text.TextTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(attribute.HRef_AArea("/main"), text.TextTrusted("Main")),
					A(attribute.HRef_AArea("/details"), text.TextTrusted("Details")),
				),
				Main(nil,
					H1(nil,
						Div(nil, text.TextTrusted("Here could be your content")),
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
			attribute.Lang("en"),
			Head(
				nil,
				text.Text("The Title of the Page"),
			),
			Body(nil,
				flow.RangeIter(iters.FromToInclusive(1, 1), func(_, i int) htmfunc.Element {
					return Div(nil, text.Text(strconv.Itoa(i)))
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
			attribute.Lang("en"),
			Head(
				nil,
				text.TextTrusted("The Title of the Page"),
			),
			Body(nil,
				Div(attribute.Class("header"),
					Nav(nil,
						A(attribute.HRef_AArea("/main"), text.TextTrusted("Main")),
						A(attribute.HRef_AArea("/details"), text.TextTrusted("Details")),
					),
					H1(nil, text.Text("Calendar")),
					H2(nil, text.Text("2024")),
				),
				Main(nil,
					Div(attribute.Class("year"),
						flow.Range(months, month),
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
	return Div(attribute.Class("month"),
		H3(nil, text.Text(month.name)),
		Div(attribute.Class("days"),
			flow.RangeInt(month.days, func(i int) htmfunc.Element {
				return Div(attribute.Class("day"),
					text.Text(strconv.Itoa(i+1)),
				)
			}),
		),
	)
}
