package cl

import "github.com/ch-schulz/htmfunc"

type Class func(w htmfunc.Writer) error

// C adds all given classes space separated to a class attribute.
func C(class string, classes ...string) Class {
	return func(w htmfunc.Writer) error {
		_, err := w.WriteString(class)
		if err != nil {
			return err
		}

		for _, c := range classes {
			err = w.WriteByte(' ')
			if err != nil {
				return err
			}

			_, err = w.WriteString(c)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

// Opt adds the given class if ok is true.
func Opt(ok bool, class string) Class {
	return func(w htmfunc.Writer) error {
		if !ok {
			return nil
		}

		_, err := w.WriteString(class)

		return err
	}
}
