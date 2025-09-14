package roundly

type Attribute func(w Writer, options *RenderOptions) error

func (fn Attribute) RenderAttribute(w Writer) error {
	return fn(w, nil)
}

func (fn Attribute) RenderAttributeWithOptions(w Writer, opts *RenderOptions) error {
	return fn(w, opts)
}

func WriteAttribute(name string, value string) Attribute {
	return func(w Writer, opts *RenderOptions) error {
		if opts == nil {
			return writeAttribute(w, name, value)
		}

		return writeAttributeWithOptions(w, name, value, opts)
	}
}

func writeAttribute(w Writer, name string, value string) error {
	if value == "" {
		return nil
	}

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

func WriteMultiValueAttribute(name string, delimiter byte, values ...string) Attribute {
	return func(w Writer, opts *RenderOptions) error {
		if opts == nil {
			return writeMultiValueAttribute(w, name, delimiter, values)
		}

		return WriteMultiValueAttributeWithOptions(w, name, delimiter, values, opts)
	}
}

func writeMultiValueAttribute(w Writer, name string, delimiter byte, values []string) error {
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

func WriteBoolAttribute(name string) Attribute {
	return func(w Writer, opts *RenderOptions) error {
		if opts == nil || !opts.Pretty {
			err := w.WriteByte(' ')
			if err != nil {
				return err
			}

			_, err = w.WriteString(name)
			return err
		}

		err := w.WriteByte('\n')
		if err != nil {
			return err
		}

		err = opts.WriteIndent(w)
		if err != nil {
			return err
		}

		_, err = w.WriteString(name)

		return err
	}
}

func writeStringsSeparated(w Writer, delimiter byte, values []string) error {
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
		err = w.WriteByte(delimiter)
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
