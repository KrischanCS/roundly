package text

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/htmfunc"
)

func TestRawTrusted(t *testing.T) {
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

//nolint:funlen
func TestRawTrustedSamples(t *testing.T) {
	t.Parallel()

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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got := RawTrusted(tt.text).String()

			// Assert
			assert.Equal(t, tt.text, got,
				"RawTrusted should return the exact text without escaping.")
		})
	}
}

//nolint:funlen
func BenchmarkRawTrusted(b *testing.B) {
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

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			w := htmfunc.NewWriter()

			b.ReportAllocs()

			for b.Loop() {
				_ = RawTrusted(tt.text).RenderElement(w) //nolint:errcheck

				w.Reset()
			}
		})
	}
}
