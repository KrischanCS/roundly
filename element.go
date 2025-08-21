// Package roundly provides a way to create type-safe, template like HTML
// components in pure go.
package roundly

type Element func(w Writer, options ...*RenderOptions) error

// RenderElement renders the element to the given Writer.
func (fn Element) RenderElement(w Writer) error {
	return fn(w)
}

func (fn Element) RenderElementWithOptions(w Writer, options *RenderOptions) error {
	return fn(w, options)
}

// String renders the element to a string. For most use cases RenderElement will be more efficient,
// this exists mainly for brevity in tests and examples.
func (fn Element) String() string {
	w := NewWriter()

	err := fn(w, &RenderOptions{
		Pretty: true,
	})
	if err != nil {
		panic("Writing to bufio.Writer failed unexpectedly: " + err.Error())
	}

	return w.String()
}

// WriteElement creates a normal html element, with open and closing tag, the given attributes in
// the opening tag and the given childNodes wrapped between the tags inside.
//
// The element is written without any extra linebreaks and indentation, thus is somewhat minified
func WriteElement(tag string, attributes Attribute, childNodes ...Element) Element {

	return func(w Writer, options ...*RenderOptions) error {
		if len(options) != 0 {
			return writeElementWithOptions(w, tag, attributes, childNodes, options[0])
		}

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
// This must only be used for elements which don't require a closing tag according to the
// [html standard].
//
// Usually it will be more convenient and readable to use the respective component functions
// provided by [el]. These are the specified void elements and component functions:
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
	return func(w Writer, opts ...*RenderOptions) error {
		if len(opts) == 0 {
			return writeOpenTag(w, tag, attributes)
		}

		return writeOpenTagWithOptions(w, tag, attributes, opts[0])
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
