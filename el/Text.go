package el

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
	CharEntity = [...]string{
		"&amp;",
		"&#34;",
		"&#39;",
		"&lt;",
		"&gt;",
	}
)

// Text represents exactly the given text without any extra tags. Html characters will be escaped. If this is not
// wanted, [TextNoEscape] may be used.
func Text(text string) htmfunc.Component {
	return func(w htmfunc.Writer) error {
		var err error
		for _, c := range []byte(text) {
			if n := indexOf(escapeChar, c); n >= 0 {
				_, err = w.WriteString(CharEntity[n])
				if err != nil {
					return err
				}
				continue
			}

			err = w.WriteByte(c)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func indexOf(chars [5]byte, char byte) int {
	for i := 0; i < 5; i++ {
		if chars[i] == char {
			return i
		}
	}
	return -1
}

// TextNoEscape is equivalent to [Text], but won't escape the input. Only use with safe text and never with user input.
func TextNoEscape(text string) htmfunc.Component {
	return func(w htmfunc.Writer) error {
		_, err := w.WriteString(text)
		return err
	}
}
