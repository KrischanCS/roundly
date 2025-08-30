package roundly

import "bytes"

type RenderOptions struct {
	Pretty          bool
	LineBreakMin    int
	ExplainYourself bool

	level        int
	funcBuffer   Writer
	elementStack []string
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
	if opts.ExplainYourself {
		err := writeExplanation(opts)
		if err != nil {
			return err
		}
	}

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
		err := writeAttributesWithOptions(w, attributes, opts)
		if err != nil {
			return err
		}
	}

	_, err = w.WriteString(">\n")
	if err != nil {
		return err
	}

	opts.IncreaseIndent()

	return nil
}

func writeAttributesWithOptions(w Writer, attributes Attribute, opts *RenderOptions) error {
	buf, err := bufferAttrsIndented(attributes, opts)
	if err != nil {
		return err
	}

	newLineCount := bytes.Count(buf, []byte{'\n'})

	if newLineCount < 3 { //nolint:mnd
		return writeFlattened(w, buf)
	}

	_, err = w.Write(buf)
	if err != nil {
		return err
	}

	return opts.WriteNewLineAndIndent(w)
}

func writeFlattened(w Writer, buf []byte) error {
	r := bytes.NewReader(buf)

	for b, err := r.ReadByte(); err == nil; b, err = r.ReadByte() {
		if b == '\n' {
			err = replaceLineBreakAndIndent(w, r)
			if err != nil {
				return err
			}

			continue
		}

		err := w.WriteByte(b)
		if err != nil {
			return err
		}
	}

	return nil
}

func replaceLineBreakAndIndent(w Writer, r *bytes.Reader) error {
	err := w.WriteByte(' ')
	if err != nil {
		return err
	}

	return consumeIndent(r)
}

func consumeIndent(r *bytes.Reader) error {
	b, err := r.ReadByte()
	for b == '\t' && err == nil {
		b, err = r.ReadByte()
	}

	err = r.UnreadByte()
	if err != nil {
		return err
	}

	return nil
}

func bufferAttrsIndented(attributes Attribute, opts *RenderOptions) ([]byte, error) {
	opts.IncreaseIndent()
	defer opts.DecreaseIndent()

	buf := make([]byte, 0, 128)
	tempW := bytes.NewBuffer(buf)

	err := attributes.RenderAttributeWithOptions(tempW, opts)
	if err != nil {
		return nil, err
	}

	return tempW.Bytes(), nil
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
	if value == "" {
		return nil
	}

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
