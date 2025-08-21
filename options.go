package roundly

import "strings"

type RenderOptions struct {
	Pretty bool
	Level  int
	Indent string
}

func writeElementWithOptions(w Writer, tag string, attributes Attribute, childNodes []Element, opts *RenderOptions) error {
	err := writeOpenTagWithOptions(w, tag, attributes, opts)
	if err != nil {
		return err
	}

	for _, node := range childNodes {
		err = node.RenderElementWithOptions(w, opts)
		if err != nil {
			return err
		}
	}

	err = writeCloseTag(w, tag)
	if err != nil {
		return err
	}

	return nil
}

func writeOpenTagWithOptions(w Writer, tag string, attributes Attribute, opts *RenderOptions) error {
	indentation := "\t"
	if opts.Indent != "" {
		indentation = opts.Indent
	}

	if opts.Pretty {
		_, err := w.WriteString(strings.Repeat(indentation, opts.Level))
		if err != nil {
			return err
		}
	}

	err := w.WriteByte('<')
	if err != nil {
		return err
	}

	_, err = w.WriteString(tag)
	if err != nil {
		return err
	}

	if attributes != nil {
		err = attributes.RenderAttribute(w)
		if err != nil {
			return err
		}
	}

	return w.WriteByte('>')
}

func WriteAttributeWithOptions(w Writer, name string, value string, options *RenderOptions) error {
	// TODO this is currently identical with WriteAttribute
	err := w.WriteByte(' ')
	if err != nil {
		return err
	}

	_, err = w.WriteString(name)
	if err != nil {
		return err
	}

	_, err = w.WriteString(`="`)
	if err != nil {
		return err
	}

	_, err = w.WriteString(value)
	if err != nil {
		return err
	}

	return w.WriteByte('"')
}

func WriteMultiValueAttributeWithOptions(w Writer, name string, delimiter byte, values []string, options *RenderOptions) error {
	// TODO this is currently identical with WriteMultiValueAttribute
	err := w.WriteByte(' ')
	if err != nil {
		return err
	}

	_, err = w.WriteString(name)
	if err != nil {
		return err
	}

	_, err = w.WriteString(`="`)
	if err != nil {
		return err
	}

	err = writeStringsSeparated(w, delimiter, values)
	if err != nil {
		return err
	}

	return w.WriteByte('"')
}
