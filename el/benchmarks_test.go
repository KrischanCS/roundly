package el

import (
	"github.com/ch-schulz/htmfunc"
	"github.com/ch-schulz/htmfunc/attr"
	"github.com/valyala/bytebufferpool"
	"strconv"
	"testing"
)

//nolint:errcheck
func BenchmarkExamplePage(b *testing.B) {
	w := bytebufferpool.Get()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()

		page := HTML(
			attr.Lang("en"),
			Head(
				TitleTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(attr.HRef("/main"), "Main"),
					A(attr.HRef("/details"), "Details"),
				),
				Main(nil,
					H1(nil,
						Div(nil, Text("Here could be your content")),
					),
				),
			),
		)

		_ = page(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageRange10(b *testing.B) {
	w := bytebufferpool.Get()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()

		page := HTML(
			attr.Lang("en"),
			Head(
				TitleTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(attr.HRef("/main"), "Main"),
					A(attr.HRef("/details"), "Details"),
				),
				Main(nil,
					H1(nil,
						Div(nil, Text("Here could be your content")),
					),
					Range(IteratorFromTo(1, 10), func(i int) htmfunc.Element {
						return Div(nil, Text(strconv.Itoa(i)))
					}),
				),
			),
		)

		_ = page(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageNoEscape(b *testing.B) {
	w := bytebufferpool.Get()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()

		page := HTML(
			attr.Lang("en"),
			Head(
				TitleTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(attr.HRef("/main"), "Main"),
					A(attr.HRef("/details"), "Details"),
				),
				Main(nil,
					H1(nil,
						Div(nil, TextTrusted("Here could be your content")),
					),
				),
			),
		)

		_ = page(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageWriteOnly(b *testing.B) {
	page := HTML(
		attr.Lang("en"),
		Head(
			TitleTrusted("The Title of the Page"),
		),
		Body(nil,
			Nav(nil,
				A(attr.HRef("/main"), "Main"),
				A(attr.HRef("/details"), "Details"),
			),
			Main(nil,
				H1(nil,
					Div(nil, Text("Here could be your content")),
				),
			),
		),
	)

	w := bytebufferpool.Get()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()
		_ = page(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageWriteOnlyNoEscape(b *testing.B) {
	page := HTML(
		attr.Lang("en"),
		Head(
			TitleTrusted("The Title of the Page"),
		),
		Body(nil,
			Nav(nil,
				A(attr.HRef("/main"), "Main"),
				A(attr.HRef("/details"), "Details"),
			),
			Main(nil,
				H1(nil,
					Div(nil, TextTrusted("Here could be your content")),
				),
			),
		),
	)

	w := bytebufferpool.Get()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()
		_ = page(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkRange(b *testing.B) {
	w := bytebufferpool.Get()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()

		page := HTML(
			attr.Lang("en"),
			Head(
				Title("The Title of the Page"),
			),
			Body(nil,
				Range(IteratorFromTo(1, 1), func(i int) htmfunc.Element {
					return Div(nil, Text(strconv.Itoa(i)))
				}),
			),
		)

		_ = page(w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkYearCalendar(b *testing.B) {
	months := []struct {
		name string
		days int
	}{
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

	page := HTML(
		attr.Lang("en"),
		Head(
			TitleTrusted("The Title of the Page"),
		),
		Body(nil,
			Div(attr.Class(attr.JoinValues("header")),
				Nav(nil,
					A(attr.HRef("/main"), "Main"),
					A(attr.HRef("/details"), "Details"),
				),
				H1(nil, Text("Calendar")),
				H2(nil, Text("2024")),
			),
			Main(nil,
				Div(attr.Class(attr.JoinValues("year")),
					Range(IteratorOf(months...), func(month struct {
						name string
						days int
					}) htmfunc.Element {
						return Div(attr.Class(attr.JoinValues("month")),
							H3(nil, Text(month.name)),
							Div(attr.Class(attr.JoinValues("days")),
								Range(
									IteratorFromTo(1, month.days), func(i int) htmfunc.Element {
										return Div(attr.Class(attr.JoinValues("days")),
											Text(strconv.Itoa(i)),
										)
									},
								),
							),
						)
					})),
			),
		),
	)

	w := bytebufferpool.Get()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()
		_ = page(w)
	}

	_ = w.String()
}
