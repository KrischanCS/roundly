// Copied by generator. DO NOT EDIT.
// Original file: src/text/text_stdlib.go

//go:build roundly_text_std

package html

import (
	"text/template"

	"github.com/KrischanCS/roundly"
)

// Text represents exactly the given text without any extra tags. Html-specific characters will be
// escaped. If this is not wanted, [RawTrusted] may be used.
func Text(text string) roundly.Element {
	return func(w roundly.Writer, opts *roundly.RenderOptions) (err error) {
		if opts != nil {
			textBytes = addIndentsAndLineBreaks(textBytes, opts)
		}

		template.HTMLEscape(w, []byte(text))
		return
	}
}
