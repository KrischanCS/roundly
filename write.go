package htmfunc

type ElementWriteFunc func(w Writer) error

func (fn ElementWriteFunc) RenderElement(w Writer) error {
	return fn(w)
}

type AttributeWriteFunc func(w Writer) error

func (fn AttributeWriteFunc) RenderAttribute(w Writer) error {
	return fn(w)
}

type ValueWriteFunc func(w Writer) error

func (fn ValueWriteFunc) RenderValue(w Writer) error {
	return fn(w)
}

// WriteElement creates a normal html element, with open and closing tag, the given attributes in the opening tag and
// the given childNodes wrapped between the tags inside.
func WriteElement(tag string, attributes AttributeRenderer, childNodes ...ElementRenderer) ElementRenderer {
	return ElementWriteFunc(func(w Writer) error {
		err := writeOpenTag(w, tag, attributes)
		if err != nil {
			return err
		}

		for _, node := range childNodes {
			err = node.RenderElement(w)
			if err != nil {
				return err
			}
		}

		err = writeCloseTag(w, tag)
		if err != nil {
			return err
		}

		return nil
	})
}

// WriteVoidElement creates a void element (An element without child nodes and closing tag).
//
// This must only be used for elements which don't require a closing tag according to the [html standard].
//
// Usually it will be more convenient and readable to use the respective component functions provided by [el]. These are
// the specified void elements and component functions:
//
//   - area		->	[el.Area]
//   - base		->	[el.Base]
//   - br		->	[el.Br]
//   - col		->	[el.Col]
//   - embed	->	[el.Embed]
//   - hr		->	[el.Hr]
//   - img		->	[el.Img]
//   - input	->	[el.Input]
//   - link		->	[el.Link]
//   - meta		->	[el.Meta]
//   - source	->	[el.Source]
//   - track	->	[el.Track]
//   - wbr		->	[el.Wbr]
//
// [html standard]: https://html.spec.whatwg.org/#void-elements
func WriteVoidElement(tag string, attributes AttributeRenderer) ElementRenderer {
	return ElementWriteFunc(func(w Writer) error {
		return writeOpenTag(w, tag, attributes)
	})
}

func writeOpenTag(w Writer, tag string, attributes AttributeRenderer) error {
	err := w.WriteByte('<')
	if err != nil {
		return err
	}

	_, err = w.WriteString(tag)
	if err != nil {
		return err
	}

	if attributes != nil {
		err = w.WriteByte(' ')
		if err != nil {
			return err
		}

		err = attributes.RenderAttribute(w)
		if err != nil {
			return err
		}
	}

	return w.WriteByte('>')
}

func writeCloseTag(w Writer, tag string) error {
	err := w.WriteByte('<')
	if err != nil {
		return err
	}

	err = w.WriteByte('/')
	if err != nil {
		return err
	}

	_, err = w.WriteString(tag)
	if err != nil {
		return err
	}

	return w.WriteByte('>')
}
