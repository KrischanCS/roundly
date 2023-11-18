package el

import (
	"bytes"
	"github.com/ch-schulz/htmfunc/attr"
	"testing"
)

//nolint:errcheck
func BenchmarkExamplePage(b *testing.B) {
	var w bytes.Buffer

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()

		page := HTML(
			attr.Lang("en"),
			Head(
				TitleNoEscape("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(attr.Ls{attr.HRef("/main")}, "Main"),
					A(attr.Ls{attr.HRef("/details")}, "Details"),
				),
				Main(nil,
					H1(nil,
						Div(nil, Text("Here could be your content")),
					),
				),
			),
		)

		_ = page(&w)
	}

	b.StopTimer()
	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageNoEscape(b *testing.B) {
	var w bytes.Buffer

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()

		page := HTML(
			attr.Lang("en"),
			Head(
				TitleNoEscape("The Title of the Page"),
			),
			Body(nil,
				Nav(nil,
					A(attr.Ls{attr.HRef("/main")}, "Main"),
					A(attr.Ls{attr.HRef("/details")}, "Details"),
				),
				Main(nil,
					H1(nil,
						Div(nil, TextNoEscape("Here could be your content")),
					),
				),
			),
		)

		_ = page(&w)
	}

	b.StopTimer()
	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageWriteOnly(b *testing.B) {
	page := HTML(
		attr.Lang("en"),
		Head(
			TitleNoEscape("The Title of the Page"),
		),
		Body(nil,
			Nav(nil,
				A(attr.Ls{attr.HRef("/main")}, "Main"),
				A(attr.Ls{attr.HRef("/details")}, "Details"),
			),
			Main(nil,
				H1(nil,
					Div(nil, Text("Here could be your content")),
				),
			),
		),
	)

	var w bytes.Buffer

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()
		_ = page(&w)
	}

	b.StopTimer()
	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageWriteOnlyNoEscape(b *testing.B) {
	page := HTML(
		attr.Lang("en"),
		Head(
			TitleNoEscape("The Title of the Page"),
		),
		Body(nil,
			Nav(nil,
				A(attr.Ls{attr.HRef("/main")}, "Main"),
				A(attr.Ls{attr.HRef("/details")}, "Details"),
			),
			Main(nil,
				H1(nil,
					Div(nil, TextNoEscape("Here could be your content")),
				),
			),
		),
	)

	var w bytes.Buffer

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Reset()
		_ = page(&w)
	}

	b.StopTimer()
	_ = w.String()
}
