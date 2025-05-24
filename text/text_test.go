package text

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/htmfunc"
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
			w := htmfunc.NewWriter()

			err := Text(tt.text).RenderElement(w)

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
			w := htmfunc.NewWriter()

			err := RawTrusted(tt.text).RenderElement(w)

			assert.NoError(t, err)
			assert.Equal(t, tt.text, w.String())
		})
	}
}

func BenchmarkText(b *testing.B) {
	t := "! \" # $ % & ' ( ) * + , - . / 0 1 2 3 4 5 6 7 8 9 : ; < = > ?" +
		"@ A B C D E F G H I J K L M N O P Q R S T U V W X Y Z [ \\ ] ^ _" +
		"` a b c d e f g h i j k l m n o p q r s t u v w x y z { | } ~"

	w := htmfunc.NewWriter()

	b.ResetTimer()
	b.ReportAllocs()

	for range b.N {
		_ = Text(t)(w) //nolint:errcheck
		w.Reset()
	}
}

func BenchmarkTextTrusted(b *testing.B) {
	t := "! \" # $ % & ' ( ) * + , - . / 0 1 2 3 4 5 6 7 8 9 : ; < = > ?" +
		"@ A B C D E F G H I J K L M N O P Q R S T U V W X Y Z [ \\ ] ^ _" +
		"` a b c d e f g h i j k l m n o p q r s t u v w x y z { | } ~"

	w := htmfunc.NewWriter()

	b.ResetTimer()
	b.ReportAllocs()

	for range b.N {
		_ = RawTrusted(t)(w) //nolint:errcheck
		w.Reset()
	}
}
