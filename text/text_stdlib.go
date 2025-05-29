// Package text Provides functions to add text to elements or attributes, either escaped of not.
//
// The not escaped version should only be used with trusted text and never with user input.

//go:build htmfunc_text_std

package text

import (
	"text/template"

	"github.com/KrischanCS/htmfunc"
)

// Text represents exactly the given text without any extra tags. Html-specific characters will be
// escaped. If this is not wanted, [RawTrusted] may be used.
func Text(text string) htmfunc.Element {
	return func(w htmfunc.Writer) (err error) {
		template.HTMLEscape(w, []byte(text))
		return
	}
}
