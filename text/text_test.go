package text

import (
	"html/template"
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

	type impl struct {
		name string
		impl func(text string) htmfunc.Element
	}

	impls := []impl{
		{"Text", Text},
		{"TextBaseline", TextBaseline},
		{"TextVariantSwitch", TextVariantSwitch},
		{"TextVariantChunk", TextVariantChunk},
		{"TextVariantChunkUnsafe", TextVariantChunkUnsafe},
		{"TextVariantIndexOfUnsafe", TextVariantIndexOfUnsafe},
		{"TextVariantDualModeUnsafe", TextVariantDualModeUnsafe},
	}

	for _, impl := range impls {
		for _, tt := range tests {
			t.Run(tt.name+"_"+impl.name, func(t *testing.T) {
				w := htmfunc.NewWriter()

				err := Text(tt.text).RenderElement(w)

				assert.NoError(t, err)
				assert.Equal(t, tt.want, w.String())
			})
		}
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
	type testCase struct {
		name string
		text string
	}

	tests := []testCase{
		{
			name: "No Characters escaped, 250 chars",
			text: `Lorem ipsum dolor sit amet, consectetur adipiscing
elit. Sed do eiusmod tempor incididunt ut labore 
et dolore magna aliqua. Ut enim ad minim veniam, 
quis nostrud exercitation ullamco laboris nisi 
ut aliquip ex ea commodo consequat. Dis aute irure.`,
		},
		{
			name: "Some Characters escaped, 250 chars",
			text: `Lorem ipsum dolor sit: "consectetur adipiscing
<elit>. Sed do eiusmod tempor incididunt & labore 
et dolore magna aliqua." Ut <enim> & <minim> veniam, 
quis nostrud exercitation ullamco laboris nisi 
ut aliquip ex ea \commodo \consequat. Dis aute ir.`,
		},
		{
			name: "Many Characters escaped, 250 chars",
			text: `<table id="mascot-tbl">
  <thead>
    <tr>
      <th>Language</th>
      <th>Name</th>
    </tr>
  </thead>
  <tbody>
    <tr><th scope="row">Go</th><td>The Go Gopher</td></tr>
    <tr><th scope="row">Rust</th><td>Ferris</td></tr>
  </tbody>
</table>`,
		},
		{
			name: "All Characters escaped, 250 chars",
			text: `&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&
&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\
&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\
&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\
&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\`,
		},
		{
			name: "Max dual mode switches, 250 chars",
			text: "&<>abcdefghijklmnop&<>abcdefghijklmnop&<>abcdefghijklmnop" +
				"&<>abcdefghijklmnop&<>abcdefghijklmnop&<>abcdefghijklmnop" +
				"&<>abcdefghijklmnop&<>abcdefghijklmnop&<>abcdefghijklmnop" +
				"&<>abcdefghijklmnop&<>abcdefghijklmnop&<>abcdefghijklmnop" +
				"&<>abcdefghijklmnop&<>",
		},
	}

	type impl struct {
		name string
		impl func(text string) htmfunc.Element
	}

	impls := []impl{
		//{"Text", Text},
		{"TextBaseline", TextBaseline},
		{"TextBaselineUnsafeErrCheck", TextBaselineUnsafeErrCheck},
		//{"TextVariantReplaceAll", TextVariantReplaceAll},
		//{"TextVariantReplaceAllUnsafe", TextVariantReplaceAllUnsafe},
		//{"TextVariantSwitch", TextVariantSwitch},
		//{"TextVariantChunk", TextVariantChunk},
		//{"TextVariantChunkUnsafe", TextVariantChunkUnsafe},
		//{"TextVariantIndexOfUnsafe", TextVariantIndexOfUnsafe},
		{"TextVariantDualModeUnsafe", TextVariantDualModeUnsafe},
	}

	for _, tt := range tests {

		for _, impl := range impls {
			b.Run(tt.name+"_"+impl.name, func(b *testing.B) {
				w := htmfunc.NewWriter()

				b.ReportAllocs()

				for b.Loop() {
					_ = impl.impl(tt.text).RenderElement(w)

					w.Reset()
				}
			})
		}
	}
}

func TestTextImplementations(t *testing.T) {
	type testCase struct {
		name string
		text string
	}

	tests := []testCase{
		{
			name: "No Characters escaped, 250 chars",
			text: `Lorem ipsum dolor sit amet, consectetur adipiscing
elit. Sed do eiusmod tempor incididunt ut labore 
et dolore magna aliqua. Ut enim ad minim veniam, 
quis nostrud exercitation ullamco laboris nisi 
ut aliquip ex ea commodo consequat. Dis aute irure.`,
		},
		{
			name: "Some Characters escaped, 250 chars",
			text: `Lorem ipsum dolor sit: "consectetur adipiscing
<elit>. Sed do eiusmod tempor incididunt & labore 
et dolore magna aliqua." Ut <enim> & <minim> veniam, 
quis nostrud exercitation ullamco laboris nisi 
ut aliquip ex ea \commodo \consequat. Dis aute ir.`,
		},
		{
			name: "Many Characters escaped, 250 chars",
			text: `<table id="mascot-tbl">
  <thead>
    <tr>
      <th>Language</th>
      <th>Name</th>
    </tr>
  </thead>
  <tbody>
    <tr><th scope="row">Go</th><td>The Go Gopher</td></tr>
    <tr><th scope="row">Rust</th><td>Ferris</td></tr>
  </tbody>
</table>`,
		},
		{
			name: "All Characters escaped, 250 chars",
			text: `&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&
&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\
&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\
&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\
&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\&<>"\`,
		},
		{
			name: "Max dual mode switches, 250 chars",
			text: "&<>abcdefghijklmnop&<>abcdefghijklmnop&<>abcdefghijklmnop" +
				"&<>abcdefghijklmnop&<>abcdefghijklmnop&<>abcdefghijklmnop" +
				"&<>abcdefghijklmnop&<>abcdefghijklmnop&<>abcdefghijklmnop" +
				"&<>abcdefghijklmnop&<>abcdefghijklmnop&<>abcdefghijklmnop" +
				"&<>abcdefghijklmnop&<>",
		},
	}

	type impl struct {
		name string
		impl func(text string) htmfunc.Element
	}

	impls := []impl{
		{"Text", Text},
		{"TextBaselineUnsafeErrCheck", TextBaselineUnsafeErrCheck},
		{"TextVariantSwitch", TextVariantSwitch},
		{"TextVariantChunk", TextVariantChunk},
		{"TextVariantChunkUnsafe", TextVariantChunkUnsafe},
		{"TextVariantIndexOfUnsafe", TextVariantIndexOfUnsafe},
		{"TextVariantDualModeUnsafe", TextVariantDualModeUnsafe},
	}

	for _, tt := range tests {

		for _, impl := range impls {
			t.Run(tt.name+"_"+impl.name, func(t *testing.T) {
				// Act
				got := impl.impl(tt.text).String()

				// Assert

				// Compare against std library escapting
				want := template.HTMLEscapeString(tt.text)
				assert.Equal(t, want, got)
			})
		}
	}
}

func BenchmarkRawTrusted_MostCharsNotEscaped(b *testing.B) {
	t := `"! \" # $ % & ' ( ) * + , - . / 0 1 2 3 4 5 6 7 8 9 : ; < = > ?
@ A B C D E F G H I J K L M N O P Q R S T U V W X Y Z [ \\ ] ^ _
 a b c d e f g h i j k l m n o p q r s t u v w x y z { | } ~"`

	w := htmfunc.NewWriter()

	b.ReportAllocs()

	for b.Loop() {
		_ = RawTrusted(t)(w) //nolint:errcheck
		w.Reset()
	}
}

func BenchmarkRawTrusted_MostCharsEscaped(b *testing.B) {
	t := `///|||\\\<<<>>>""""&&&&&abcde///|||\\\<<<>>>""""&&&&&abcde///|||\\\<<<>>>""""&&&&&abcde
///|||\\\<<<>>>""""&&&&&abcde///|||\\\<<<>>>""""&&&&&abcde
///|||\\\<<<>>>""""&&&&&abcde///|||\\\<<<>>`

	w := htmfunc.NewWriter()

	b.ReportAllocs()

	for b.Loop() {
		_ = RawTrusted(t)(w) //nolint:errcheck
		w.Reset()
	}
}
