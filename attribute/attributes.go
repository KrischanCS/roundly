package attribute

import (
	"github.com/ch-schulz/htmfunc"
)

func Attribute(name string, values ...string) htmfunc.AttributeRenderer {
	return htmfunc.WriteAttributeFunc(func(w htmfunc.Writer) error {
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
	})
}

func MultiValueAttribute(name string, values []htmfunc.ValueRenderer) htmfunc.AttributeRenderer {
	return htmfunc.WriteAttributeFunc(func(w htmfunc.Writer) error {
		_, err := w.WriteString(name + `="`)
		if err != nil {
			return err
		}

		if len(values) == 0 {
			_, err = w.WriteString(`"`)
			return err
		}

		err = values[0].RenderValue(w)
		if err != nil {
			return err
		}

		for _, value := range values[1:] {
			err = w.WriteByte(' ')
			if err != nil {
				return err
			}

			err = value.RenderValue(w)
			if err != nil {
				return err
			}
		}

		return w.WriteByte('"')
	})
}

// Join joins the given attributes with spaces.
func Join(attributes ...htmfunc.AttributeRenderer) htmfunc.AttributeRenderer {
	return htmfunc.WriteAttributeFunc(func(w htmfunc.Writer) error {
		return WriteSpaceSeperated(w, attributes...)
	})
}

func Lang(language string) htmfunc.AttributeRenderer {
	return Attribute("lang", language)
}

func Src(source string) htmfunc.AttributeRenderer {
	return Attribute("src", source)
}

func Charset(charset string) htmfunc.AttributeRenderer {
	return Attribute("charset", charset)
}

func HRef(href string) htmfunc.AttributeRenderer {
	return Attribute("href", href)
}

func Value(value string) htmfunc.AttributeRenderer {
	return Attribute("value", value)
}

func DateTime(dateTime string) htmfunc.AttributeRenderer {
	return Attribute("datetime", dateTime)
}

func Dir(direction htmfunc.TextDirection) htmfunc.AttributeRenderer {
	return Attribute("dir", string(direction))
}

func Rel(relation string) htmfunc.AttributeRenderer {
	return Attribute("rel", relation)
}

func Name(name string) htmfunc.AttributeRenderer {
	return Attribute("name", name)
}

func Content(content string) htmfunc.AttributeRenderer {
	return Attribute("content", content)
}

func Type(t string) htmfunc.AttributeRenderer {
	return Attribute("type", t)
}

func Class(classes ...htmfunc.ValueRenderer) htmfunc.AttributeRenderer {
	return MultiValueAttribute("class", classes)
}

func Id(id string) htmfunc.AttributeRenderer {
	return Attribute("id", id)
}

func BooleanAttribute(name string) htmfunc.AttributeRenderer {
	return htmfunc.WriteAttributeFunc(func(w htmfunc.Writer) error {
		_, err := w.WriteString(name)
		return err
	})
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

func WriteSpaceSeperated(w htmfunc.Writer, attributes ...htmfunc.AttributeRenderer) (err error) {
	if len(attributes) == 0 {
		return
	}

	err = attributes[0].RenderAttribute(w)
	if err != nil {
		return err
	}

	for _, attribute := range attributes[1:] {
		if attribute == nil {
			continue
		}

		err = w.WriteByte(' ')
		if err != nil {
			return err
		}

		err = attribute.RenderAttribute(w)
		if err != nil {
			return err
		}
	}

	return
}
