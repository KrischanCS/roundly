package element

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ch-schulz/htmfunc"
)

func TestText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "empty",
			text: "",
			want: "",
		},
		{
			name: "simple",
			text: "hello world",
			want: "hello world",
		},
		{
			name: "escape <",
			text: "hello<world",
			want: "hello&lt;world",
		},
		{
			name: "escape >",
			text: "hello>world",
			want: "hello&gt;world",
		},
		{
			name: "escape &",
			text: "hello&world",
			want: "hello&amp;world",
		},
		{
			name: "escape '",
			text: "hello'world",
			want: "hello&#39;world",
		},
		{
			name: "escape \"",
			text: "hello\"world",
			want: "hello&#34;world",
		},
		{
			name: "escape all",
			text: `<div>'Test'>"test" & 1<3</div>`,
			want: "&lt;div&gt;&#39;Test&#39;&gt;&#34;test&#34; &amp; 1&lt;3&lt;/div&gt;",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := htmfunc.NewWriter(16)

			err := Text(tt.text).RenderHTML(w)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

func TestTextTrusted(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		text string
	}{
		{
			name: "empty",
			text: "",
		},
		{
			name: "simple",
			text: "hello world",
		},
		{
			name: "escape <",
			text: "hello<world",
		},
		{
			name: "escape >",
			text: "hello>world",
		},
		{
			name: "escape &",
			text: "hello&world",
		},
		{
			name: "escape '",
			text: "hello'world",
		},
		{
			name: "escape \"",
			text: "hello\"world",
		},
		{
			name: "escape all",
			text: `<div>'Test'>"test" & 1<3</div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := htmfunc.NewWriter(16)

			err := TextTrusted(tt.text).RenderHTML(w)

			assert.NoError(t, err)
			assert.Equal(t, tt.text, w.String())
		})
	}
}
