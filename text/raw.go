package text

import "github.com/KrischanCS/htmfunc"

// RawTrusted writes the given text without escaping. This must never be used with unescaped user
// input or text from any not trusted source.
//
// It has two major use cases:
//  1. Adding html snippets from trusted sources, must not be escaped on purpose.
//  2. Adding text that is already escaped, avoiding double escaping & better performance.
func RawTrusted(text string) htmfunc.Element {
	return func(w htmfunc.Writer) error {
		_, err := w.WriteString(text)
		return err
	}
}
