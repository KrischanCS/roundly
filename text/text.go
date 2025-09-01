//go:build !roundly_text_std

// Package text Provides functions to add text to elements or attributes, either escaped of not.
//
// The not escaped version should only be used with trusted text and never with user input.
package text

import (
	"bytes"
	"unsafe"

	"github.com/KrischanCS/roundly"
)

// Text represents exactly the given text without any extra tags. Html-specific characters will be
// escaped. If this is not wanted, [RawTrusted] may be used.
func Text(text string) roundly.Element {
	// This function is optimized for performance, since it is used very often in templates.
	//
	// The implementation uses unsafe (described below) and is relatively complex. It beats the
	// stdlib's template.HTMLEscape in most scenarios in runtime, and clearly in bytes allocated.
	//
	// Users can opt-out of this optimization by using the roundly_text_std build tag.
	//
	// # Modes
	//
	// The implementation uses a two-mode approach, which switches between two different ways of
	// finding characters that must be escaped:
	//
	// 1. IndexOfMode: Uses bytes.IndexAny to find the next character that must be escaped and
	//    checks that then individually. This writes lots of bytes at once, but checks the byte to
	//    be escaped twice (Index of and individually in switch-case).
	//    It is very fast when characters to be escaped are distributed sparsely over the text.
	//
	// 2. CheckEachMode: Checks each character individually, which avoids the double-checking
	//    by checking each character individually, it is fast when there is a high density of
	//    characters to be escaped.
	//
	// # Usage of unsafe
	//
	// Accessing the underlying byte slice through unsafe to avoid the extra allocation.
	// Since printing text to html safely escaped is a very common operation, keeping this function
	// fast & low on allocations is important.
	//
	// The underlying array is never modified, and it must be kept like this when this function is
	// edited.
	textBytes := unsafe.Slice(unsafe.StringData(text), len(text)) //nolint:gosec // explained abov

	return func(w roundly.Writer, opts *roundly.RenderOptions) (err error) {
		if opts != nil {
			textBytes = addIndentsAndLineBreaks(textBytes, opts)
		}

		return writeTextFindNextMode(w, textBytes)
	}
}

func writeTextFindNextMode(w roundly.Writer, text []byte) (err error) {
	// This performs best, when characters which must be escaped are distributed sparsely over the
	// text. If they are close together, the performance would degrade significantly. In that case,
	// we switch to a writeTextCheckEachMode.
	const escapedBytes = "&<>\"'\000"

	const (
		shortDistance   = 16
		modeSwitchLimit = 3
	)

	repeatedShortDistances := 0

	for len(text) > 0 {
		i := bytes.IndexAny(text, escapedBytes)
		if i == -1 {
			_, err = w.Write(text)
			return err
		}

		if i < shortDistance {
			repeatedShortDistances++
		} else {
			repeatedShortDistances = 0
		}

		err = writeChunk(w, text[:i], text[i])
		if err != nil {
			return err
		}

		text = text[i+1:]

		if repeatedShortDistances >= modeSwitchLimit {
			return writeTextCheckEachMode(w, text)
		}
	}

	return nil
}

func writeChunk(w roundly.Writer, unescapedBytes []byte, byteToEscape byte) (err error) {
	_, err = w.Write(unescapedBytes)
	if err != nil {
		return err
	}

	return writeEscapedByte(w, byteToEscape)
}

// Replacements for escaped chars, taken from stdlib html/template.
//
//nolint:gochecknoglobals
var (
	htmlQuot = []byte("&#34;") // shorter than "&quot;"
	htmlApos = []byte("&#39;") // shorter than "&apos;" and apos was not in HTML until HTML5
	htmlAmp  = []byte("&amp;")
	htmlLt   = []byte("&lt;")
	htmlGt   = []byte("&gt;")
	htmlNull = []byte("\uFFFD")
)

func writeEscapedByte(w roundly.Writer, text byte) (err error) {
	switch text {
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
		err = w.WriteByte(text)
	}

	return err
}

func writeTextCheckEachMode(w roundly.Writer, text []byte) (err error) {
	// This mode checks every character individually, which is usually slower than the IndexOfMode,
	// but it performs better when the text contains many characters that must be escaped,
	// since it avoids the double-checking.
	//
	// We switch back to FindNextMode after a certain amount of consecutive characters that didn't
	// need escaping.
	const switchLimit = 16

	repeatedNonEscaped := 0

	for i, char := range text {
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
			return writeTextFindNextMode(w, text[i+1:])
		}
	}

	return nil
}

//nolint:errcheck // writing to bytes.Buffer never fails
func addIndentsAndLineBreaks(textBytes []byte, opts *roundly.RenderOptions) []byte {
	if !opts.Pretty {
		return textBytes
	}

	if opts.LineBreakMin == 0 {
		opts.LineBreakMin = 80
	}

	buf := make([]byte, 0, len(textBytes)+len(textBytes)/opts.LineBreakMin+16)
	w := bytes.NewBuffer(buf)

	_ = opts.WriteIndent(w)

	charsSinceBreak := 0
	for _, char := range textBytes {
		if char == '\n' {
			charsSinceBreak = 0
			_ = w.WriteByte(char)
			_ = opts.WriteIndent(w)

			continue
		}

		charsSinceBreak++
		if charsSinceBreak <= opts.LineBreakMin {
			_ = w.WriteByte(char)
			continue
		}

		switch char {
		case '\t':
		case ' ':
			w.WriteByte('\n')
			_ = opts.WriteIndent(w)
			charsSinceBreak = 0
		default:
			w.WriteByte(char)
		}
	}

	w.WriteByte('\n')

	return w.Bytes()
}
