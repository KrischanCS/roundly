// Package text Provides functions to add text to elements or attributes, either escaped of not.
//
// The not escaped version should only be used with trusted text and never with user input.

//go:build !htmfunc_text_std

package text

import (
	"bytes"
	"unsafe"

	"github.com/KrischanCS/htmfunc"
)

// Text represents exactly the given text without any extra tags. Html-specific characters will be
// escaped. If this is not wanted, [RawTrusted] may be used.
func Text(text string) htmfunc.Element {
	b := unsafe.Slice(unsafe.StringData(text), len(text))

	return func(w htmfunc.Writer) (err error) {
		return writeTextIndexOfMode(w, b)
	}
}

func writeTextIndexOfMode(w htmfunc.Writer, b []byte) (err error) {
	// This performs best, when characters which must be escaped are distributed sparsely over the
	// text. If they are close together, the performance would degrade significantly. In that case,
	// we switch to a SingleCheckMode.

	const shortDistance = 16
	const modeSwitchLimit = 3

	repeatedShortDistances := 0

	for len(b) > 0 {
		i := bytes.IndexAny(b, `&<>"'`)
		if i == -1 {
			_, err = w.Write(b)
			return err
		}

		if i < shortDistance {
			repeatedShortDistances++
			if repeatedShortDistances >= modeSwitchLimit {
				_, err = w.Write(b[:i])
				if err != nil {
					return err
				}

				return writeTextSingleCheckMode(w, b[i:])
			}
		} else {
			repeatedShortDistances = 0
		}

		_, err = w.Write(b[:i])
		if err != nil {
			return err
		}

		err = writeEscapedByte(w, b[i])
		if err != nil {
			return err
		}

		b = b[i+1:]
	}

	return nil
}

// Replacements for escaped chars, taken from stdlib html/template
var (
	htmlQuot = []byte("&#34;") // shorter than "&quot;"
	htmlApos = []byte("&#39;") // shorter than "&apos;" and apos was not in HTML until HTML5
	htmlAmp  = []byte("&amp;")
	htmlLt   = []byte("&lt;")
	htmlGt   = []byte("&gt;")
	htmlNull = []byte("\uFFFD")
)

func writeEscapedByte(w htmfunc.Writer, b byte) (err error) {
	switch b {
	case '&':
		_, err = w.Write(htmlAmp)
	case '<':
		_, err = w.Write(htmlLt)
	case '>':
		_, err = w.Write(htmlGt)
	case '\'':
		_, err = w.Write(htmlApos)
	case '"':
		_, err = w.Write(htmlQuot)
	case '\000':
		_, err = w.Write(htmlNull)
	default:
		err = w.WriteByte(b)
	}

	return err
}

func writeTextSingleCheckMode(w htmfunc.Writer, b []byte) (err error) {
	// This mode checks every character individually, which is usually slower than the IndexOfMode,
	// but it performs better when the text contains many characters that must be escaped,
	// since it avoids the double-checking.
	//
	// Switches back to IndexOfMode when it doesn't find any
	// characters that must be escaped for a while.

	const switchLimit = 16
	repeatedNonEscaped := 0

	for i, char := range b {
		switch char {
		case '&':
			repeatedNonEscaped = 0
			_, err = w.Write(htmlAmp)
		case '<':
			repeatedNonEscaped = 0
			_, err = w.Write(htmlLt)
		case '>':
			repeatedNonEscaped = 0
			_, err = w.Write(htmlGt)
		case '\'':
			repeatedNonEscaped = 0
			_, err = w.Write(htmlApos)
		case '"':
			repeatedNonEscaped = 0
			_, err = w.Write(htmlQuot)
		case '\000':
			repeatedNonEscaped = 0
			_, err = w.Write(htmlNull)
		default:
			repeatedNonEscaped++
			err = w.WriteByte(char)
		}

		if err != nil {
			return err
		}

		if repeatedNonEscaped >= switchLimit {
			return writeTextIndexOfMode(w, b[i+1:])
		}
	}

	return nil
}
