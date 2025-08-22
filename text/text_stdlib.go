//go:build roundly_text_std

// Package text Provides functions to add text to elements or attributes, either escaped of not.
//
// The not escaped version should only be used with trusted text and never with user input.
package text

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
