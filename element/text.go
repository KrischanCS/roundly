package element

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

// Text represents exactly the given text without any extra tags. Html-specific
// characters will be escaped. If this is not wanted, [TextTrusted] may be used.
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

// TextTrusted is equivalent to [Text], but won't escape the input. Only use with safe text and never with user input.
func TextTrusted(text string) htmfunc.Element {
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
