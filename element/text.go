package element

import (
	"github.com/ch-schulz/htmfunc"
)

var (
	escapeChar = [...]byte{ //nolint:gochecknoglobals
		'&',
		'"',
		'\'',
		'<',
		'>',
	}
	charEntity = [...]string{ //nolint:gochecknoglobals
		"&amp;",
		"&#34;",
		"&#39;",
		"&lt;",
		"&gt;",
	}
)

// Text represents exactly the given text without any extra tags.
// Html characters will be escaped. If this is not wanted,
// [TextTrusted] may be used.
func Text(text string) htmfunc.Element {
	return func(w htmfunc.Writer) error {
		for _, r := range []byte(text) {
			if i := indexOf(escapeChar, r); i >= 0 {
				_, err := w.WriteString(charEntity[i])
				if err != nil {
					return err
				}

				continue
			}

			err := w.WriteByte(r)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func indexOf(chars [5]byte, char byte) int {
	for i := range 5 {
		if chars[i] == char {
			return i
		}
	}

	return -1
}

// TextTrusted is equivalent to [Text], but won't escape the input. Only use with safe text and never with user input.
func TextTrusted(text string) htmfunc.Element {
	return func(w htmfunc.Writer) error {
		_, err := w.WriteString(text)
		return err
	}
}
