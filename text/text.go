// Package text Provides functions to add text to elements or attributes, either escaped of not.
//
// The not escaped version should only be used with trusted text and never with user input.
package text

import (
	"bytes"

	"github.com/KrischanCS/htmfunc"
)

var (
	escapeChar = []byte{ //nolint:gochecknoglobals
		'&',
		'"',
		'\'',
		'<',
		'>',
	}
	charEntity = []string{ //nolint:gochecknoglobals
		"&amp;",
		"&#34;",
		"&#39;",
		"&lt;",
		"&gt;",
	}
)

// Text represents exactly the given text without any extra tags. Html-specific characters will be
// escaped. If this is not wanted, [RawTrusted] may be used.
func Text(text string) htmfunc.Element {
	return func(w htmfunc.Writer) error {
		for _, char := range []byte(text) {
			err := checkEscapeAndWrite(w, char)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

// RawTrusted writes the given text without escaping. This must never be used with unescaped user
// input.
//
// It has two major use cases:
//  1. Adding html snippets from trusted sources.
//  2. Adding text that is already escaped, e.g. from a database or a template engine.
//
// It is slightly more performant than [Text], but to a margin that it should seldom be a reason to
// use it. The main reason to use this is to avoid double escaping of text.
func RawTrusted(text string) htmfunc.Element {
	return func(w htmfunc.Writer) error {
		_, err := w.WriteString(text)
		return err
	}
}

func checkEscapeAndWrite(w htmfunc.Writer, char byte) error {
	i := bytes.IndexByte(escapeChar, char)

	if i == -1 {
		return w.WriteByte(char)
	}

	_, err := w.WriteString(charEntity[i])
	if err != nil {
		return err
	}

	return nil
}
