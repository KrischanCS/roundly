package roundly

// Document creates an html document with doctype and html root.
func Document(doctype string, html Element) Element {
	return func(w Writer, opts ...*RenderOptions) error {
		if len(opts) == 0 {
			err := Doctype(doctype).RenderElement(w)
			if err != nil {
				return err
			}

			return html.RenderElement(w)
		}

		err := Doctype(doctype).RenderElementWithOptions(w, opts[0])
		if err != nil {
			return err
		}

		return html.RenderElementWithOptions(w, opts[0])
	}
}

// Doctype creates the mandatory [doctype tag].
//
// [doctype tag]: https://html.spec.whatwg.org/#the-doctype
func Doctype(doctype string) Element {
	return func(w Writer, opts ...*RenderOptions) error {
		_, err := w.WriteString("<!doctype ")
		if err != nil {
			return err
		}

		_, err = w.WriteString(doctype)
		if err != nil {
			return err
		}

		err = w.WriteByte('>')

		return err
	}
}
