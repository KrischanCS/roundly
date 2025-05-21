// Package htmfunc provides a way to create type-safe, template like HTML
// components in pure go.
package htmfunc

type Element func(w Writer) error

func (fn Element) RenderElement(w Writer) error {
	return fn(w)
}

// WriteElement creates a normal html element, with open and closing tag, the given attributes in the opening tag and
// the given childNodes wrapped between the tags inside.
func WriteElement(tag string, attributes Attribute, childNodes ...Element) Element {
	return func(w Writer) error {
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
	}
}

// WriteVoidElement creates a void element (An element without child nodes and closing tag).
//
// This must only be used for elements which don't require a closing tag according to the [html standard].
//
// Usually it will be more convenient and readable to use the respective component functions provided by [el]. These are
// the specified void elements and component functions:
//
//   - area		->	[element.Area]
//   - base		->	[element.Base]
//   - br		->	[element.Br]
//   - col		->	[element.Col]
//   - embed	->	[element.Embed]
//   - hr		->	[element.Hr]
//   - img		->	[element.Img]
//   - input	->	[element.Input]
//   - link		->	[element.Link]
//   - meta		->	[element.Meta]
//   - source	->	[element.Source]
//   - track	->	[element.Track]
//   - wbr		->	[element.Wbr]
//
// [html standard]: https://html.spec.whatwg.org/#void-elements
func WriteVoidElement(tag string, attributes Attribute) Element {
	return func(w Writer) error {
		return writeOpenTag(w, tag, attributes)
	}
}

func writeOpenTag(w Writer, tag string, attributes Attribute) error {
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
