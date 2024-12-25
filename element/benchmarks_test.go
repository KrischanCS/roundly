package element

import (
	"strconv"
	"testing"

	"github.com/ch-schulz/htmfunc"
	"github.com/ch-schulz/htmfunc/attribute"
	"github.com/ch-schulz/htmfunc/flow"
	"github.com/ch-schulz/htmfunc/iters"
)

//nolint:errcheck
func BenchmarkExamplePage(b *testing.B) {
	w := htmfunc.NewWriter(4096)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()

		page := Html(
			attribute.Lang("en"),
			Head(
				TitleTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(attribute.HRef("/main"), TextTrusted("Main")),
					A(attribute.HRef("/details"), TextTrusted("Details")),
				),
				Main(nil,
					H1(nil,
						Div(nil, Text("Here could be your content")),
					),
				),
			),
		)

		_ = page.RenderHtml(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageRange10(b *testing.B) {
	w := htmfunc.NewWriter(4096)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()

		page := Html(
			attribute.Lang("en"),
			Head(
				TitleTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(attribute.HRef("/main"), TextTrusted("Main")),
					A(attribute.HRef("/details"), TextTrusted("Details")),
				),
				Main(nil,
					H1(nil,
						Div(nil, Text("Here could be your content")),
					),
					flow.RangeIter(iters.FromToInclusive(1, 10), func(_, i int) htmfunc.Element {
						return Div(nil, Text(strconv.Itoa(i)))
					}),
				),
			),
		)

		_ = page.RenderHtml(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageNoEscape(b *testing.B) {
	w := htmfunc.NewWriter(4096)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()

		page := Html(
			attribute.Lang("en"),
			Head(
				TitleTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(attribute.HRef("/main"), TextTrusted("Main")),
					A(attribute.HRef("/details"), TextTrusted("Details")),
				),
				Main(nil,
					H1(nil,
						Div(nil, TextTrusted("Here could be your content")),
					),
				),
			),
		)

		_ = page.RenderHtml(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageWriteOnly(b *testing.B) {
	w := htmfunc.NewWriter(4096)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()

		page := Html(
			attribute.Lang("en"),
			Head(
				TitleTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(attribute.HRef("/main"), TextTrusted("Main")),
					A(attribute.HRef("/details"), TextTrusted("Details")),
				),
				Main(nil,
					H1(nil,
						Div(nil, Text("Here could be your content")),
					),
				),
			),
		)

		_ = page.RenderHtml(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageWriteOnlyNoEscape(b *testing.B) {
	w := htmfunc.NewWriter(4096)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()

		page := Html(
			attribute.Lang("en"),
			Head(
				TitleTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(attribute.HRef("/main"), TextTrusted("Main")),
					A(attribute.HRef("/details"), TextTrusted("Details")),
				),
				Main(nil,
					H1(nil,
						Div(nil, TextTrusted("Here could be your content")),
					),
				),
			),
		)

		_ = page.RenderHtml(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkRange(b *testing.B) {
	w := htmfunc.NewWriter(4096)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()

		page := Html(
			attribute.Lang("en"),
			Head(
				Title("The Title of the Page"),
			),
			Body(nil,
				flow.RangeIter(iters.FromToInclusive(1, 1), func(_, i int) htmfunc.Element {
					return Div(nil, Text(strconv.Itoa(i)))
				}),
			),
		)

		_ = page.RenderHtml(w)
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
				TitleTrusted("The Title of the Page"),
			),
			Body(nil,
				Div(attribute.Class(attribute.JoinValues("header")),
					Nav(nil,
						A(attribute.HRef("/main"), TextTrusted("Main")),
						A(attribute.HRef("/details"), TextTrusted("Details")),
					),
					H1(nil, Text("Calendar")),
					H2(nil, Text("2024")),
				),
				Main(nil,
					Div(attribute.Class(attribute.JoinValues("year")),
						flow.Range(months, month),
					),
				),
			),
		)

		w.Reset()

		_ = page.RenderHtml(w)
	}

	_ = w.String()
}

type monthDays struct {
	name string
	days int
}

func month(_ int, month monthDays) htmfunc.Element {
	return Div(attribute.Class(attribute.JoinValues("month")),
		H3(nil, Text(month.name)),
		Div(attribute.Class(attribute.JoinValues("days")),
			flow.RangeInt(month.days, func(i int) htmfunc.Element {
				return Div(attribute.Class(attribute.JoinValues("day")),
					Text(strconv.Itoa(i+1)),
				)
			}),
		),
	)
}
