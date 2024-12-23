package element

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/bytebufferpool"

	"github.com/ch-schulz/htmfunc"
	attr "github.com/ch-schulz/htmfunc/attribute"
)

func TestHTML(t *testing.T) {
	t.Parallel()

	type args struct {
		lang string
		head htmfunc.Element
		body htmfunc.Element
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
				body: Body(attr.Lang("de"), Text("Dies ist ein Text zum Testen.")),
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
					TextTrusted("<p>This will not be escaped!</p>")),
			},
			want: `<html lang="en"><head></head><body>The quick brown fox jumped over the lazy dog.&lt;p&gt;Text Escaping Works!&lt;/p&gt;<p>This will not be escaped!</p></body></html>`, //nolint:lll
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := bytebufferpool.Get()
			component := HTML(attr.Lang(tt.args.lang), tt.args.head, tt.args.body)

			err := component.RenderHTML(w)
			require.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

func TestBase(t *testing.T) {
	t.Parallel()

	type args struct {
		attributes htmfunc.Attribute
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "bare tag",
			args: args{},
			want: "<base>",
		},
		{
			name: "with href",
			args: args{
				attributes: attr.HRef("https://example.com/index.html"),
			},
			want: `<base href="https://example.com/index.html">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := bytebufferpool.Get()

			component := Base(tt.args.attributes)

			err := component.RenderHTML(w)
			require.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

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
				doctype: `HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd"`,
			},
			want: `<!doctype HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := bytebufferpool.Get()

			component := Doctype(tt.args.doctype)

			err := component.RenderHTML(w)
			require.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

func TestHead(t *testing.T) {
	t.Parallel()

	type args struct {
		childNodes []htmfunc.Element
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
				childNodes: []htmfunc.Element{
					Title("The Title"),
				},
			},
			want: "<head><title>The Title</title></head>",
		},
		{
			name: "empty",
			args: args{
				childNodes: []htmfunc.Element{
					Title("The Title"),
					Link(attr.Join(attr.HRef("/style.css"), attr.Rel("stylesheet"))),
				},
			},
			want: `<head><title>The Title</title><link href="/style.css" rel="stylesheet"></head>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := bytebufferpool.Get()

			component := Head(tt.args.childNodes...)

			err := component.RenderHTML(w)
			require.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

func TestMeta(t *testing.T) {
	t.Parallel()

	w := bytebufferpool.Get()

	component := Meta(attr.Join(attr.Name("keywords"), attr.Content("british,type face,font,fonts,highway,"+
		"highways")))

	err := component.RenderHTML(w)
	require.NoError(t, err)

	want := `<meta name="keywords" content="british,type face,font,fonts,highway,highways">`

	assert.Equal(t, want, w.String())
}

//nolint:dupl
func TestStyle(t *testing.T) {
	t.Parallel()

	type args struct {
		attributes htmfunc.Attribute
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
				attributes: attr.Type("text/css"),
				css:        `body{background-color: firebrick}`,
			},
			want: `<style type="text/css">body{background-color: firebrick}</style>`,
		},
		{
			name: "css escaping",
			args: args{
				attributes: attr.Type("text/css"),
				css:        `body>div{background-color: firebrick}`,
			},
			want: `<style type="text/css">body&gt;div{background-color: firebrick}</style>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := bytebufferpool.Get()

			component := Style(tt.args.attributes, tt.args.css)

			err := component.RenderHTML(w)
			require.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

//nolint:dupl
func TestStyleTrusted(t *testing.T) {
	t.Parallel()

	type args struct {
		attributes htmfunc.Attribute
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
				attributes: attr.Type("text/css"),
				css:        `body{background-color: firebrick}`,
			},
			want: `<style type="text/css">body{background-color: firebrick}</style>`,
		},
		{
			name: "css escaping",
			args: args{
				attributes: attr.Type("text/css"),
				css:        `body>div{background-color: firebrick}`,
			},
			want: `<style type="text/css">body>div{background-color: firebrick}</style>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := bytebufferpool.Get()

			component := StyleTrusted(tt.args.attributes, tt.args.css)

			err := component.RenderHTML(w)
			require.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

func TestTitle(t *testing.T) {
	t.Parallel()

	type args struct {
		title string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{},
			want: "<title></title>",
		},
		{
			name: "simple title",
			args: args{
				title: "The Title",
			},
			want: `<title>The Title</title>`,
		},
		{
			name: "simple title",
			args: args{
				title: "The Title & some more",
			},
			want: `<title>The Title &amp; some more</title>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := bytebufferpool.Get()

			component := Title(tt.args.title)

			err := component.RenderHTML(w)
			require.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

func TestTitleTrusted(t *testing.T) {
	t.Parallel()

	type args struct {
		title string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{},
			want: "<title></title>",
		},
		{
			name: "simple title",
			args: args{
				title: "The Title",
			},
			want: `<title>The Title</title>`,
		},
		{
			name: "simple title",
			args: args{
				title: "The Title & some more",
			},
			want: `<title>The Title & some more</title>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := bytebufferpool.Get()

			component := TitleTrusted(tt.args.title)

			err := component.RenderHTML(w)
			require.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}
