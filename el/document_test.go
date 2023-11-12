package el

import (
	"bytes"
	"github.com/ch-schulz/htmfunc"
	"github.com/ch-schulz/htmfunc/attr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHTML(t *testing.T) {
	t.Parallel()

	type args struct {
		lang string
		head htmfunc.Component
		body htmfunc.Component
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				lang: "en",
				head: Head(),
				body: Body(nil),
			},
			want: `<html lang="en"><head></head><body></body></html>`,
		},
		{
			name: "With Text",
			args: args{
				lang: "de",
				body: Body(attr.Ls{attr.Lang("de")}, Text("Dies ist ein Text zum Testen.")),
				head: Head(Title("Der Titel")),
			},
			want: `<html lang="de"><head><title>Der Titel</title></head><body lang="de">Dies ist ein Text zum Testen.</body></html>`, //nolint:lll
		},
		{
			name: "Multiple Components",
			args: args{
				lang: "en",
				head: Head(),
				body: Body(nil,
					Text("The quick brown fox jumped over the lazy dog."),
					Text("<p>Text Escaping Works!</p>"),
					TextNoEscape("<p>This will not be escaped!</p>")),
			},
			want: `<html lang="en"><head></head><body>The quick brown fox jumped over the lazy dog.&lt;p&gt;Text Escaping Works!&lt;/p&gt;<p>This will not be escaped!</p></body></html>`, //nolint:lll
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var b bytes.Buffer

			component := HTML(attr.Lang(tt.args.lang), tt.args.head, tt.args.body)

			err := component(&b)
			require.NoError(t, err)
			assert.Equal(t, tt.want, b.String())
		})
	}
}

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
