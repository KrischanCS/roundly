package htmfunc

// Element creates a normal html element, with open and closing tag, the given attributes in the opening tag and
// the given childNodes wrapped between the tags inside.
func Element(tag string, attributes []Attribute, childNodes ...Component) Component {
	return func(w Writer) error {
		err := writeOpenTag(w, tag, attributes)
		if err != nil {
			return err
		}

		for _, node := range childNodes {
			err = node(w)
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

// VoidElement creates a void element (An element without child nodes and closing tag).
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
func VoidElement(tag string, attributes []Attribute) Component {
	return func(w Writer) error {
		return writeOpenTag(w, tag, attributes)
	}
}

func writeOpenTag(w Writer, tag string, attributes []Attribute) error {
	err := w.WriteByte('<')
	if err != nil {
		return err
	}

	_, err = w.WriteString(tag)
	if err != nil {
		return err
	}

	err = writeAttributes(w, attributes)
	if err != nil {
		return err
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

func writeAttributes(w Writer, attributes []Attribute) error {
	for _, attribute := range attributes {
		err := w.WriteByte(' ')
		if err != nil {
			return err
		}

		err = attribute(w)
		if err != nil {
			return err
		}
	}

	return nil
}
