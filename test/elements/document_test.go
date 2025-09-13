package elements

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/html"
)

func TestDocument(t *testing.T) {
	t.Parallel()

	w := roundly.NewWriter()

	doc := roundly.Document("html", Html(Lang("en"), Head(nil), Body(nil)))

	err := doc.RenderElement(w)

	assert.NoError(t, err)
	assert.Equal(t, `<!doctype html><html lang="en"><head></head><body></body></html>`, w.String())
}

func TestHtml(t *testing.T) {
	t.Parallel()

	type args struct {
		lang string
		head roundly.Element
		body roundly.Element
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
				head: Head(nil),
				body: Body(nil),
			},
			want: `<html lang="en"><head></head><body></body></html>`,
		},
		{
			name: "With Text",
			args: args{
				lang: "de",
				body: Body(Lang("de"), Text("Dies ist ein Text zum Testen.")),
				head: Head(nil, Title(nil, Text("Der Titel"))),
			},
			want: `<html lang="de"><head><title>Der Titel</title></head><body lang="de">Dies ist ein Text zum Testen.</body></html>`, //nolint:nolintlint,lll
		},
		{
			name: "Multiple Components",
			args: args{
				lang: "en",
				head: Head(nil),
				body: Body(nil,
					Text("The quick brown fox jumped over the lazy dog."),
					Text("<p>Text Escaping Works!</p>"),
					RawTrusted("<p>This will not be escaped!</p>")),
			},
			want: `<html lang="en"><head></head><body>The quick brown fox jumped over the lazy dog.&lt;p&gt;Text Escaping Works!&lt;/p&gt;<p>This will not be escaped!</p></body></html>`, //nolint:nolintlint,lll
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := roundly.NewWriter()
			component := Html(Lang(tt.args.lang), tt.args.head, tt.args.body)

			err := component.RenderElement(w)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

func TestBase(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		attributes roundly.Attribute
		want       string
	}{
		{
			name:       "bare tag",
			attributes: nil,
			want:       "<base>",
		},
		{
			name:       "with href",
			attributes: HRef("https://example.com/index.html"),
			want:       `<base href="https://example.com/index.html">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := roundly.NewWriter()

			component := Base(tt.attributes)

			err := component.RenderElement(w)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

//nolint:lll
func TestDoctype(t *testing.T) {
	t.Parallel()

	type args struct {
		doctype string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Html 5",
			args: args{
				doctype: "html",
			},
			want: "<!doctype html>",
		},
		{
			name: "Html 4",
			args: args{
				doctype: `Html PUBLIC "-//W3C//DTD Html 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd"`,
			},
			want: `<!doctype Html PUBLIC "-//W3C//DTD Html 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := roundly.NewWriter()

			component := roundly.Doctype(tt.args.doctype)

			err := component.RenderElement(w)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

func TestHead(t *testing.T) {
	t.Parallel()

	type args struct {
		childNodes []roundly.Element
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				childNodes: nil,
			},
			want: "<head></head>",
		},
		{
			name: "title only",
			args: args{
				childNodes: []roundly.Element{
					Title(nil, Text("The Title")),
				},
			},
			want: "<head><title>The Title</title></head>",
		},
		{
			name: "empty",
			args: args{
				childNodes: []roundly.Element{
					Title(nil, Text("The Title")),
					Link(Attributes(HRef("/style.css"), Rel("stylesheet"))),
				},
			},
			want: `<head><title>The Title</title><link href="/style.css" rel="stylesheet"></head>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := roundly.NewWriter()

			component := Head(nil, tt.args.childNodes...)

			err := component.RenderElement(w)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

func TestMeta(t *testing.T) {
	t.Parallel()

	w := roundly.NewWriter()

	component := Meta(Attributes(Name("keywords"), Content(
		"british,type face,font,fonts,highway,highways")))

	err := component.RenderElement(w)
	assert.NoError(t, err)

	want := `<meta name="keywords" content="british,type face,font,fonts,highway,highways">`

	assert.Equal(t, want, w.String())
}

//nolint:dupl
func TestStyle(t *testing.T) {
	t.Parallel()

	type args struct {
		attributes roundly.Attribute
		css        string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				attributes: nil,
				css:        "",
			},
			want: "<style></style>",
		},
		{
			name: "css",
			args: args{
				attributes: Type("text/css"),
				css:        `body{background-color: firebrick}`,
			},
			want: `<style type="text/css">body{background-color: firebrick}</style>`,
		},
		{
			name: "css escaping",
			args: args{
				attributes: Type("text/css"),
				css:        `body>div{background-color: firebrick}`,
			},
			want: `<style type="text/css">body&gt;div{background-color: firebrick}</style>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := roundly.NewWriter()

			component := Style(tt.args.attributes, Text(tt.args.css))

			err := component.RenderElement(w)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

//nolint:dupl
func TestStyleTrusted(t *testing.T) {
	t.Parallel()

	type args struct {
		attributes roundly.Attribute
		css        string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				attributes: nil,
				css:        "",
			},
			want: "<style></style>",
		},
		{
			name: "css",
			args: args{
				attributes: Type("text/css"),
				css:        `body{background-color: firebrick}`,
			},
			want: `<style type="text/css">body{background-color: firebrick}</style>`,
		},
		{
			name: "css escaping",
			args: args{
				attributes: Type("text/css"),
				css:        `body>div{background-color: firebrick}`,
			},
			want: `<style type="text/css">body>div{background-color: firebrick}</style>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := roundly.NewWriter()

			component := Style(tt.args.attributes, RawTrusted(tt.args.css))

			err := component.RenderElement(w)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

func TestTitle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		title string
		want  string
	}{
		{
			name:  "empty",
			title: "",
			want:  "<title></title>",
		},
		{
			name:  "simple title",
			title: "The Title",
			want:  `<title>The Title</title>`,
		},
		{
			name:  "simple title",
			title: "The Title & some more",
			want:  `<title>The Title &amp; some more</title>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := roundly.NewWriter()

			component := Title(nil, Text(tt.title))

			err := component.RenderElement(w)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

func TestTitleTrusted(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		title string
		want  string
	}{
		{
			name:  "empty",
			title: "",
			want:  "<title></title>",
		},
		{
			name:  "simple title",
			title: "The Title",
			want:  `<title>The Title</title>`,
		},
		{
			name:  "simple title",
			title: "The Title & some more",
			want:  `<title>The Title & some more</title>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := roundly.NewWriter()

			component := Title(nil, RawTrusted(tt.title))

			err := component.RenderElement(w)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}
