package test

import (
	"strconv"
	"testing"

	"github.com/KrischanCS/go-toolbox/iterator"

	"github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/attribute"
	. "github.com/KrischanCS/roundly/element"
	. "github.com/KrischanCS/roundly/logic"
	. "github.com/KrischanCS/roundly/text"
)

//nolint:errcheck
func BenchmarkExamplePage(b *testing.B) {
	w := roundly.NewWriter()

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		w.Reset()

		page := Html(
			Lang("en"),
			Head(nil,
				Title(nil, RawTrusted("The Title of the Page")),
			),
			Body(nil,
				Nav(nil,
					A(HRef("/main"), RawTrusted("Main"),
						A(HRef("/details"), RawTrusted("Details")),
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
	w := roundly.NewWriter()

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		w.Reset()

		page := Html(
			Lang("en"),
			Head(
				nil,
				Title(
					nil, RawTrusted("The Title of the Page"),
				),
			),
			Body(nil,
				Nav(nil,
					A(HRef("/main"), RawTrusted("Main")),
					A(HRef("/details"), RawTrusted("Details")),
				),
				Main(nil,
					H1(nil,
						Div(nil, Text("Here could be your content")),
					),
					RangeSeq(iterator.FromToInclusive(1, 10), func(i int) roundly.Element {
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
	w := roundly.NewWriter()

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		w.Reset()

		page := Html(
			Lang("en"),
			Head(
				nil, RawTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(HRef("/main"), RawTrusted("Main")),
					A(HRef("/details"), RawTrusted("Details")),
				),
				Main(nil,
					H1(nil,
						Div(nil, RawTrusted("Here could be your content")),
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
	w := roundly.NewWriter()

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		w.Reset()

		page := Html(
			Lang("en"),
			Head(
				nil, RawTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(HRef("/main"), RawTrusted("Main")),
					A(HRef("/details"), RawTrusted("Details")),
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
	w := roundly.NewWriter()

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		w.Reset()

		page := Html(
			Lang("en"),
			Head(
				nil, RawTrusted("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(HRef("/main"), RawTrusted("Main")),
					A(HRef("/details"), RawTrusted("Details")),
				),
				Main(nil,
					H1(nil,
						Div(nil, RawTrusted("Here could be your content")),
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
	w := roundly.NewWriter()

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		w.Reset()

		page := Html(
			Lang("en"),
			Head(
				nil, Text("The Title of the Page"),
			),
			Body(nil,
				RangeSeq(iterator.FromToInclusive(1, 1), func(i int) roundly.Element {
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

	w := roundly.NewWriter()

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		page := Html(
			Lang("en"),
			Head(
				nil, RawTrusted("The Title of the Page"),
			),
			Body(nil,
				Div(Class("header"),
					Nav(nil,
						A(HRef("/main"), RawTrusted("Main")),
						A(HRef("/details"), RawTrusted("Details")),
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

func month(_ int, month monthDays) roundly.Element {
	return Div(Class("month"),
		H3(nil, Text(month.name)),
		Div(Class("days"),
			RangeInt(month.days, func(i int) roundly.Element {
				return Div(Class("day"), Text(strconv.Itoa(i+1)))
			}),
		),
	)
}
