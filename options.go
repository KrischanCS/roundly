package roundly

import "bytes"

type RenderOptions struct {
	Pretty       bool
	LineBreakMin int

	level int
}

func (r *RenderOptions) IncreaseIndent() {
	r.level++
}

func (r *RenderOptions) DecreaseIndent() {
	r.level--
}

func (r *RenderOptions) WriteIndent(w Writer) (err error) {
	for range r.level {
		err = w.WriteByte('\t')
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RenderOptions) WriteNewLineAndIndent(w Writer) error {
	err := w.WriteByte('\n')
	if err != nil {
		return err
	}

	return r.WriteIndent(w)
}

func writeElementWithOptions(
	w Writer,
	tag string,
	attributes Attribute,
	childNodes []Element,
	opts *RenderOptions,
) error {
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

	err = writeCloseTagWithOptions(w, tag, opts)
	if err != nil {
		return err
	}

	return nil
}

func writeOpenTagWithOptions(
	w Writer,
	tag string,
	attributes Attribute,
	opts *RenderOptions,
) error {
	err := opts.WriteIndent(w)
	if err != nil {
		return err
	}

	err = w.WriteByte('<')
	if err != nil {
		return err
	}

	_, err = w.WriteString(tag)
	if err != nil {
		return err
	}

	if attributes != nil {
		buf := make([]byte, 0, 128)
		tempW := bytes.NewBuffer(buf)

		opts.IncreaseIndent()
		err = attributes.RenderAttributeWithOptions(tempW, opts)
		opts.DecreaseIndent()

		if err != nil {
			return err
		}

		buf = tempW.Bytes()
		newLines := bytes.Count(buf, []byte{'\n'})
		if newLines >= 3 {
			_, err = w.Write(buf)
			if err != nil {
				return err
			}

			err = opts.WriteNewLineAndIndent(w)
		} else {
			lastWasLineBreak := false
			for _, b := range buf {
				if b == '\n' {
					err = w.WriteByte(' ')
					if err != nil {
						return err
					}
					lastWasLineBreak = true
					continue
				}

				if lastWasLineBreak && b == '\t' {
					continue
				}

				lastWasLineBreak = false

				err = w.WriteByte(b)
				if err != nil {
					return err
				}
			}
		}
	}

	_, err = w.WriteString(">\n")
	if err != nil {
		return err
	}

	opts.IncreaseIndent()

	return nil
}

func writeCloseTagWithOptions(w Writer, tag string, opts *RenderOptions) error {
	if !opts.Pretty {
		return writeCloseTag(w, tag)
	}

	opts.DecreaseIndent()

	err := opts.WriteIndent(w)
	if err != nil {
		return err
	}

	err = writeCloseTag(w, tag)
	if err != nil {
		return err
	}

	err = w.WriteByte('\n')
	if err != nil {
		return err
	}

	return nil
}

func writeAttributeWithOptions(w Writer, name string, value string, options *RenderOptions) error {
	if !options.Pretty {
		return writeAttribute(w, name, value)
	}

	err := options.WriteNewLineAndIndent(w)
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

func WriteMultiValueAttributeWithOptions(
	w Writer,
	name string,
	delimiter byte,
	values []string,
	options *RenderOptions,
) error {
	if !options.Pretty {
		return writeMultiValueAttribute(w, name, delimiter, values)
	}

	err := options.WriteNewLineAndIndent(w)
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
