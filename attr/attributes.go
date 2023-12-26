package attr

import (
	"github.com/ch-schulz/htmfunc"
)

// Join joins the given attributes with spaces
func Join(attributes ...htmfunc.Attribute) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		return WriteSpaceSeperated(w, attributes...)
	}
}

func Lang(language string) htmfunc.Attribute {
	return Attribute("lang", language)
}

func Src(source string) htmfunc.Attribute {
	return Attribute("src", source)
}

func Charset(charset string) htmfunc.Attribute {
	return Attribute("charset", charset)
}

func HRef(href string) htmfunc.Attribute {
	return Attribute("href", href)
}

func Value(value string) htmfunc.Attribute {
	return Attribute("value", value)
}

func DateTime(dateTime string) htmfunc.Attribute {
	return Attribute("datetime", dateTime)
}

func Dir(direction htmfunc.TextDirection) htmfunc.Attribute {
	return Attribute("dir", string(direction))
}

func Rel(relation string) htmfunc.Attribute {
	return Attribute("rel", relation)
}

func Name(name string) htmfunc.Attribute {
	return Attribute("name", name)
}

func Content(content string) htmfunc.Attribute {
	return Attribute("content", content)
}

func Type(t string) htmfunc.Attribute {
	return Attribute("type", t)
}

func Class(classes htmfunc.Value) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		_, err := w.WriteString(`class="`)
		if err != nil {
			return err
		}

		err = classes(w)
		if err != nil {
			return err
		}

		return w.WriteByte('"')
	}
}

func Id(id string) htmfunc.Attribute {
	return Attribute("id", id)
}

func Attribute(name string, values ...string) func(w htmfunc.Writer) error {
	return func(w htmfunc.Writer) error {
		_, err := w.WriteString(name)
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

func BooleanAttribute(name string) func(w htmfunc.Writer) error {
	return func(w htmfunc.Writer) error {
		_, err := w.WriteString(name)
		return err
	}
}

func writeStringsSpaceSeparated(w htmfunc.Writer, values []string) error {
	switch len(values) {
	case 0:
		return nil
	case 1:
		_, err := w.WriteString(values[0])
		return err
	}

	_, err := w.WriteString(values[0])
	if err != nil {
		return err
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

func WriteSpaceSeperated(w htmfunc.Writer, attrs ...htmfunc.Attribute) (err error) {
	if len(attrs) == 0 {
		return
	}

	err = attrs[0](w)
	if err != nil {
		return err
	}

	for _, c := range attrs[1:] {
		err = w.WriteByte(' ')
		if err != nil {
			return err
		}

		err = c(w)
		if err != nil {
			return err
		}
	}

	return
}
