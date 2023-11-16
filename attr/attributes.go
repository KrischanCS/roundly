package attr

import (
	"github.com/ch-schulz/htmfunc"
	"github.com/ch-schulz/htmfunc/attr/cl"
)

// Ls is a type for convenience, so one does not always have to write `[]htmfunc.Attribute` in all of [el]s method
// calls.
type Ls []htmfunc.Attribute

func Lang(language string) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		return Attribute("lang", language)(w)
	}
}

func HRef(href string) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		return Attribute("href", href)(w)
	}
}

func Value(value string) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		return Attribute("value", value)(w)
	}
}

func DateTime(dateTime string) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		return Attribute("datetime", dateTime)(w)
	}
}

func Dir(direction htmfunc.TextDirection) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		return Attribute("dir", string(direction))(w)
	}
}

func Rel(relation string) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		return Attribute("rel", relation)(w)
	}
}

func Name(name string) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		return Attribute("name", name)(w)
	}
}

func Content(content string) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		return Attribute("content", content)(w)
	}
}

func Type(t string) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		return Attribute("type", t)(w)
	}
}

func Class(classes ...cl.Class) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		_, err := w.WriteString(`class="`)
		if err != nil {
			return err
		}

		err = writeClassesPaceSeparated(w, classes)
		if err != nil {
			return err
		}

		return w.WriteByte('"')
	}
}

func Attribute(key string, values ...string) func(w htmfunc.Writer) error {
	return func(w htmfunc.Writer) error {
		_, err := w.WriteString(key)
		if err != nil {
			return err
		}

		_, err = w.WriteString(`="`)
		if err != nil {
			return err
		}

		err = writeStringsSpaceSeparated(w, values)
		if err != nil {
			return err
		}

		return w.WriteByte('"')
	}
}

func writeStringsSpaceSeparated(w htmfunc.Writer, values []string) error {
	if len(values) != 0 {
		_, err := w.WriteString(values[0])
		if err != nil {
			return err
		}
	}

	for _, v := range values[1:] {
		err := w.WriteByte(' ')
		if err != nil {
			return err
		}

		_, err = w.WriteString(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeClassesPaceSeparated(w htmfunc.Writer, classes []cl.Class) error {
	if len(classes) != 0 {
		err := classes[0](w)
		if err != nil {
			return err
		}
	}

	for _, c := range classes[1:] {
		err := w.WriteByte(' ')
		if err != nil {
			return err
		}

		err = c(w)
		if err != nil {
			return err
		}
	}

	return nil
}
