package el

import (
	"github.com/ch-schulz/htmfunc"
	"html"
)

// Text represents exactly the given text without any extra tags. Html characters will be escaped. If this is not
// wanted, [TextNoEscape] may be used.
func Text(text string) htmfunc.Component {
	return func(w htmfunc.Writer) error {
		escaped := html.EscapeString(text)

		_, err := w.WriteString(escaped)

		return err
	}
}

// TextNoEscape is equivalent to [Text], but won't escape the input. Only use with safe text and never with user input.
func TextNoEscape(text string) htmfunc.Component {
	return func(w htmfunc.Writer) error {
		_, err := w.WriteString(text)
		return err
	}
}
