package el

import (
	"github.com/ch-schulz/htmfunc"
)

// Document creates an html document with doctype and html root.
func Document(doctype string, html htmfunc.Element) htmfunc.Element {
	return func(w htmfunc.Writer) error {
		err := Doctype(doctype)(w)
		if err != nil {
			return err
		}

		return html(w)
	}
}

// Doctype creates the mandatory [doctype tag].
//
// [doctype tag]: https://html.spec.whatwg.org/#the-doctype
func Doctype(doctype string) htmfunc.Element {
	return func(w htmfunc.Writer) error {
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

// HTML creates the [html element], which represents the root of an HTML document
//
// [html element]: https://html.spec.whatwg.org/#the-html-element
func HTML(lang htmfunc.Attribute, head, body htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("html", lang, head, body)
}

// Head creates the [head element], which represents a collection of metadata for the Document.
//
// # Mandatory child elements
//
//   - [Title]
//   - [Style]
//   - [Base]
//   - [Link]
//   - [Meta]
//   - [Script]
//   - [Noscript]
//
// [head element]: https://html.spec.whatwg.org/#the-head-element
func Head(childNodes ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("head", nil, childNodes...)
}

// Title creates the [title element], which The title element represents the document's title or name.
//
// Authors should use titles that identify their documents even when they are used out of context, for example in a
// user's history or bookmarks, or in search results. The document's title is often different from its first heading,
// since the first heading does not have to stand alone when taken out of context.
//
// There must be no more than one title element per document.
//
// [title element]: https://html.spec.whatwg.org/#the-head-element
func Title(title string) htmfunc.Element {
	return htmfunc.WriteElement("title", nil, Text(title))
}

// TitleTrusted is equivalent to [Title], but does not escape the given string.
//
// This is more efficient, but only use it if the given string is safe and not possibly influenced by user input.
func TitleTrusted(title string) htmfunc.Element {
	return htmfunc.WriteElement("title", nil, TextTrusted(title))
}

// Base creates the [base element]. The base element allows authors to specify the document base URL for the purposes of
// parsing URLs, and the name of the default navigable for the purposes of following hyperlinks. The element does not
// represent any content beyond this information.
//
// There must be no more than one base element per document.
//
// A base element must have either an href attribute, a target attribute, or both.
//
// [base element]: https://html.spec.whatwg.org/#the-base-element
func Base(attributes htmfunc.Attribute) htmfunc.Element {
	return htmfunc.WriteVoidElement("base", attributes)
}

// Link creates the [link element]. The link element allows authors to link their document to other resources.
//
// The address of the link(s) is given by the href attribute. If the href attribute is present, then its value must be a
// valid non-empty URL potentially surrounded by spaces. One or both of the href or imagesrcset attributes must be
// present.
//
// [link element]: https://html.spec.whatwg.org/#the-link-element
func Link(attributes htmfunc.Attribute) htmfunc.Element {
	return htmfunc.WriteVoidElement("link", attributes)
}

// Meta creates the [meta element]. Defines metadata about an HTML document. he meta element represents various kinds of
// metadata that cannot be expressed using the title, base, link, style, and script elements.
//
// The meta element can represent document-level metadata with the name attribute, pragma directives with the http-equiv
// attribute, and the file's character encoding declaration when an HTML document is serialized to string form (e.g. for
// transmission over the network or for disk storage) with the charset attribute.
//
// Exactly one of the name, http-equiv, charset, and itemprop attributes must be specified.
//
// If either name, http-equiv, or itemprop is specified, then the content attribute must also be specified. Otherwise,
// it must be omitted.
//
// [meta element]: https://html.spec.whatwg.org/#the-meta-element
func Meta(attributes htmfunc.Attribute) htmfunc.Element {
	return htmfunc.WriteVoidElement("meta", attributes)
}

// Style creates the [style element].
//
// The style element allows authors to embed CSS style sheets in their documents. The style element is one of several
// inputs to the styling processing model. The element does not represent content for the user.
//
// [style element]: https://html.spec.whatwg.org/#the-style-element
func Style(attributes htmfunc.Attribute, css string) htmfunc.Element {
	return htmfunc.WriteElement("style", attributes, Text(css))
}

// StyleTrusted is equivalent to [Style], but does not escape the given string.
//
// This is more efficient, but only use it if the given string is safe and not possibly influenced by user input.
func StyleTrusted(attributes htmfunc.Attribute, css string) htmfunc.Element {
	return htmfunc.WriteElement("style", attributes, TextTrusted(css))
}
