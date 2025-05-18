package htmfunc

// Document creates an html document with doctype and html root.
func Document(doctype string, html Element) Element {
	return func(w Writer) error {
		err := Doctype(doctype).RenderElement(w)
		if err != nil {
			return err
		}

		return html.RenderElement(w)
	}
}

// Doctype creates the mandatory [doctype tag].
//
// [doctype tag]: https://html.spec.whatwg.org/#the-doctype
func Doctype(doctype string) Element {
	return func(w Writer) error {
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
