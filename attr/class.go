package attr

import "github.com/ch-schulz/htmfunc"

// Values writes all given strings space separated to an attribute.
func Values(value string, values ...string) htmfunc.Value {
	return func(w htmfunc.Writer) error {
		_, err := w.WriteString(value)
		if err != nil {
			return err
		}

		for _, c := range values {
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

// ValueIf writes the given value if condition is true.
func ValueIf(condition bool, value string) htmfunc.Value {
	return func(w htmfunc.Writer) error {
		if !condition {
			return nil
		}

		_, err := w.WriteString(value)

		return err
	}
}
