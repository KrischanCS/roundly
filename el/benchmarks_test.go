package el

import (
	"bytes"
	"github.com/ch-schulz/htmfunc/attr"
	"testing"
)

//nolint:errcheck
func BenchmarkHTML(b *testing.B) {
	var w bytes.Buffer

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		w.Reset()
		_ = HTML(attr.Lang("en"),
			Head(),
			Body(nil,
				Text("The quick brown fox jumped over the lazy dog."),
				Text("<p>Text with tags</p>")),
		)(&w)
	}

	b.StopTimer()

	got := w.String()
	want := `<html lang="en"><head></head><body>The quick brown fox jumped over the lazy dog.&lt;p&gt;Text with tags&lt;/p&gt;</body></html>` //nolint:lll

	if got != want {
		b.Errorf("Wanted: %s, got: %s", want, got)
	}
}

//nolint:errcheck
func BenchmarkHTMLNoEscaping(b *testing.B) {
	var w bytes.Buffer

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		w.Reset()
		_ = HTML(attr.Lang("en"),
			Head(),
			Body(nil,
				TextNoEscape("The quick brown fox jumped over the lazy dog."),
				TextNoEscape("<p>Text with tags</p>")),
		)(&w)
	}

	b.StopTimer()

	got := w.String()
	want := `<html lang="en"><head></head><body>The quick brown fox jumped over the lazy dog.<p>Text with tags</p></body></html>` //nolint:lll

	if got != want {
		b.Errorf("Wanted: %s, got: %s", want, got)
	}
}

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
					H(1, nil,
						Div(nil, TextNoEscape("Here could be your content")),
					),
				),
			),
		)

		_ = page(&w)
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageAndString(b *testing.B) {
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
					H(1, nil,
						Div(nil, TextNoEscape("Here could be your content")),
					),
				),
			),
		)

		_ = page(&w)
		_ = w.String()
	}

	_ = w.String()
}

//nolint:errcheck
func BenchmarkExamplePageRenderOnly(b *testing.B) {
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
				H(1, nil,
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

	_ = w.String()
}
