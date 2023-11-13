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
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer

			component := HTML(attr.Lang(tt.args.lang), tt.args.head, tt.args.body)

			err := component(&b)
			require.NoError(t, err)
			assert.Equal(t, tt.want, b.String())
		})
	}
}

func TestBase(t *testing.T) {
	t.Parallel()

	type args struct {
		attributes attr.Ls
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
				attributes: attr.Ls{attr.HRef("https://example.com/index.html")},
			},
			want: `<base href="https://example.com/index.html">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer

			component := Base(tt.args.attributes)

			err := component(&b)
			require.NoError(t, err)
			assert.Equal(t, tt.want, b.String())
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
			var b bytes.Buffer

			component := Doctype(tt.args.doctype)

			err := component(&b)
			require.NoError(t, err)
			assert.Equal(t, tt.want, b.String())
		})
	}
}

func TestHead(t *testing.T) {
	t.Parallel()

	type args struct {
		childNodes []htmfunc.Component
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
				childNodes: []htmfunc.Component{
					Title("The Title"),
				},
			},
			want: "<head><title>The Title</title></head>",
		},
		{
			name: "empty",
			args: args{
				childNodes: []htmfunc.Component{
					Title("The Title"),
					Link(attr.Ls{attr.HRef("/style.css"), attr.Rel("stylesheet")}),
				},
			},
			want: `<head><title>The Title</title><link href="/style.css" rel="stylesheet"></head>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer

			component := Head(tt.args.childNodes...)

			err := component(&b)
			require.NoError(t, err)
			assert.Equal(t, tt.want, b.String())
		})
	}
}

func TestMeta(t *testing.T) {
	t.Parallel()

	var b bytes.Buffer

	component := Meta(attr.Ls{attr.Name("keywords"), attr.Content("british,type face,font,fonts,highway,highways")})

	err := component(&b)
	require.NoError(t, err)

	want := `<meta name="keywords" content="british,type face,font,fonts,highway,highways">`

	assert.Equal(t, want, b.String())

}

func TestStyle(t *testing.T) {
	t.Parallel()

	type args struct {
		attributes attr.Ls
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
				attributes: attr.Ls{attr.Type("text/css")},
				css:        `body{background-color: firebrick}`,
			},
			want: `<style type="text/css">body{background-color: firebrick}</style>`,
		},
		{
			name: "css escaping",
			args: args{
				attributes: attr.Ls{attr.Type("text/css")},
				css:        `body>div{background-color: firebrick}`,
			},
			want: `<style type="text/css">body&gt;div{background-color: firebrick}</style>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer

			component := Style(tt.args.attributes, tt.args.css)

			err := component(&b)
			require.NoError(t, err)
			assert.Equal(t, tt.want, b.String())
		})
	}
}

func TestStyleNoEscape(t *testing.T) {
	t.Parallel()

	type args struct {
		attributes attr.Ls
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
				attributes: attr.Ls{attr.Type("text/css")},
				css:        `body{background-color: firebrick}`,
			},
			want: `<style type="text/css">body{background-color: firebrick}</style>`,
		},
		{
			name: "css escaping",
			args: args{
				attributes: attr.Ls{attr.Type("text/css")},
				css:        `body>div{background-color: firebrick}`,
			},
			want: `<style type="text/css">body>div{background-color: firebrick}</style>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer

			component := StyleNoEscape(tt.args.attributes, tt.args.css)

			err := component(&b)
			require.NoError(t, err)
			assert.Equal(t, tt.want, b.String())
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
			var b bytes.Buffer

			component := Title(tt.args.title)

			err := component(&b)
			require.NoError(t, err)
			assert.Equal(t, tt.want, b.String())
		})
	}
}

func TestTitleNoEscape(t *testing.T) {
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
			var b bytes.Buffer

			component := TitleNoEscape(tt.args.title)

			err := component(&b)
			require.NoError(t, err)
			assert.Equal(t, tt.want, b.String())
		})
	}
}
