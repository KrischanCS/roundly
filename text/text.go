// Package text Provides functions to add text to elements or attributes, either escaped of not.
//
// The not escaped version should only be used with trusted text and never with user input.
package text

import (
	"bytes"
	"slices"
	"text/template"
	"unsafe"

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

func TextBaseline(text string) htmfunc.Element {
	return func(w htmfunc.Writer) (err error) {
		template.HTMLEscape(w, []byte(text))
		return nil
	}
}

var (
	htmlQuot = []byte("&#34;") // shorter than "&quot;"
	htmlApos = []byte("&#39;") // shorter than "&apos;" and apos was not in HTML until HTML5
	htmlAmp  = []byte("&amp;")
	htmlLt   = []byte("&lt;")
	htmlGt   = []byte("&gt;")
	htmlNull = []byte("\uFFFD")
)

func TextBaselineUnsafeErrCheck(text string) htmfunc.Element {
	b := unsafe.Slice(unsafe.StringData(text), len(text))

	return func(w htmfunc.Writer) (err error) {
		last := 0
		for i, c := range b {
			var html []byte
			switch c {
			case '\000':
				html = htmlNull
			case '"':
				html = htmlQuot
			case '\'':
				html = htmlApos
			case '&':
				html = htmlAmp
			case '<':
				html = htmlLt
			case '>':
				html = htmlGt
			default:
				continue
			}

			_, err = w.Write(b[last:i])
			if err != nil {
				return err
			}

			_, err := w.Write(html)
			if err != nil {
				return err
			}

			last = i + 1
		}

		_, err = w.Write(b[last:])
		return err
	}
}

// Text represents exactly the given text without any extra tags. Html-specific characters will be
// escaped. If this is not wanted, [RawTrusted] may be used.
func Text(text string) htmfunc.Element {
	return func(w htmfunc.Writer) (err error) {
		for _, char := range []byte(text) {
			i := bytes.IndexByte(escapeChar, char)

			switch i {
			case -1:
				err = w.WriteByte(char)
				if err != nil {
					return err
				}
			default:
				_, err = w.WriteString(charEntity[i])
				if err != nil {
					return err
				}
			}
		}

		return nil

	}
}

func TextVariantSwitch(text string) htmfunc.Element {
	return func(w htmfunc.Writer) (err error) {
		for _, b := range []byte(text) {
			switch b {
			case '&':
				_, err = w.WriteString("&amp;")
			case '<':
				_, err = w.WriteString("&lt;")
			case '>':
				_, err = w.WriteString("&gt;")
			case '\'':
				_, err = w.WriteString("&#39;")
			case '"':
				_, err = w.WriteString("&#34;")
			default:
				err = w.WriteByte(b)
			}

			if err != nil {
				return err
			}
		}

		return nil
	}
}

func TextVariantChunk(text string) htmfunc.Element {
	return func(w htmfunc.Writer) (err error) {
		chunks := slices.Chunk([]byte(text), 64)

		for chunk := range chunks {
			if bytes.ContainsAny(chunk, `&<>"\`) {
				err = writeEscaped(w, chunk)
			} else {
				_, err = w.Write(chunk)
			}

			if err != nil {
				return err
			}
		}

		return nil
	}
}

func TextVariantChunkUnsafe(text string) htmfunc.Element {
	return func(w htmfunc.Writer) (err error) {
		b := unsafe.Slice(unsafe.StringData(text), len(text))
		chunks := slices.Chunk(b, 64)

		for chunk := range chunks {
			if bytes.ContainsAny(chunk, `&<>"'`) {
				err = writeEscaped(w, chunk)
			} else {
				_, err = w.Write(chunk)
			}

			if err != nil {
				return err
			}
		}

		return nil
	}
}

func writeEscaped(w htmfunc.Writer, chunk []byte) (err error) {
	for _, b := range chunk {
		switch b {
		case '&':
			_, err = w.WriteString("&amp;")
		case '<':
			_, err = w.WriteString("&lt;")
		case '>':
			_, err = w.WriteString("&gt;")
		case '\'':
			_, err = w.WriteString("&#39;")
		case '"':
			_, err = w.WriteString("&#34;")
		default:
			err = w.WriteByte(b)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func TextVariantIndexOfUnsafe(text string) htmfunc.Element {
	return func(w htmfunc.Writer) (err error) {
		b := unsafe.Slice(unsafe.StringData(text), len(text))

		for len(b) > 0 {
			i := bytes.IndexAny(b, `&<>"'`)
			if i == -1 {
				_, err = w.Write(b)
				return err
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
}

func TextVariantDualModeUnsafe(text string) htmfunc.Element {
	b := unsafe.Slice(unsafe.StringData(text), len(text))

	return func(w htmfunc.Writer) (err error) {
		return writeTextIndexOfMode(w, b)
	}
}

func writeTextIndexOfMode(w htmfunc.Writer, b []byte) (err error) {
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

func writeTextSingleCheckMode(w htmfunc.Writer, b []byte) error {
	// First byte is always to-be-escaped when writeTextSingleCheckMode is called.
	err := writeEscapedByte(w, b[0])
	if err != nil {
		return err
	}

	b = b[1:]

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
