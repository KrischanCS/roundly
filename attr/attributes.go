package attr

import (
	"github.com/ch-schulz/htmfunc"
)

// Ls is a type for convenience, so one does not always have to write `[]htmfunc.Attribute` in all of [el]s method
// calls.
type Ls []htmfunc.Attribute

func Lang(language string) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		_, err := w.WriteString(`lang="`)
		if err != nil {
			return err
		}

		_, err = w.WriteString(language)
		if err != nil {
			return err
		}

		err = w.WriteByte('"')
		if err != nil {
			return err
		}

		return nil
	}
}

func HRef(href string) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		_, err := w.WriteString(`href="`)
		if err != nil {
			return err
		}

		_, err = w.WriteString(href)
		if err != nil {
			return err
		}

		err = w.WriteByte('"')
		if err != nil {
			return err
		}

		return nil
	}
}
