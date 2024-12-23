package element

import (
	"github.com/ch-schulz/htmfunc"
)

var (
	escapeChar = [...]byte{
		'&',
		'"',
		'\'',
		'<',
		'>',
	}
	charEntity = [...]string{
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
type Text string

func (t Text) RenderHTML(w htmfunc.Writer) error {
	for _, r := range []byte(t) {
		if n := indexOf(escapeChar, r); n >= 0 {
			_, err := w.WriteString(charEntity[n])
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

func indexOf(chars [5]byte, char byte) int {
	for i := 0; i < 5; i++ {
		if chars[i] == char {
			return i
		}
	}

	return -1
}

// TextTrusted is equivalent to [Text], but won't escape the input. Only use with safe text and never with user input.
type TextTrusted string

func (t TextTrusted) RenderHTML(w htmfunc.Writer) error {
	_, err := w.WriteString(string(t))
	return err
}
