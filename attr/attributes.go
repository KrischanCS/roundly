package attr

import (
	"github.com/ch-schulz/htmfunc"
)

// Ls is a type for convenience, so one does not always have to write `[]htmfunc.Attribute` in all of [el]s method
// calls.
type Ls []htmfunc.Attribute

func Lang(language string) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		return Attribute(w, "lang", language)
	}
}

func HRef(href string) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		return Attribute(w, "href", href)
	}
}

func Value(value string) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		return Attribute(w, "value", value)
	}
}

func DateTime(dateTime string) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		return Attribute(w, "datetime", dateTime)
	}
}

func Dir(direction htmfunc.TextDirection) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		return Attribute(w, "dir", string(direction))
	}
}

func Attribute(w htmfunc.Writer, key string, values ...string) error {
	_, err := w.WriteString(key)
	if err != nil {
		return err
	}

	_, err = w.WriteString(`="`)
	if err != nil {
		return err
	}

	if len(values) != 0 {
		_, err = w.WriteString(values[0])
		if err != nil {
			return err
		}
	}

	for _, v := range values[1:] {
		err = w.WriteByte(' ')
		if err != nil {
			return err
		}

		_, err = w.WriteString(v)
		if err != nil {
			return err
		}
	}

	return w.WriteByte('"')
}
