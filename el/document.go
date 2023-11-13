package el

import (
	"github.com/ch-schulz/htmfunc"
	"github.com/ch-schulz/htmfunc/attr"
)

// Doctype creates the mandatory [doctype tag].
//
// [doctype tag]: https://html.spec.whatwg.org/#the-doctype
func Doctype(doctype string) htmfunc.Component {
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
func HTML(lang htmfunc.Attribute, head, body htmfunc.Component) htmfunc.Component {
	tag := "html"
	return htmfunc.Element(tag, []htmfunc.Attribute{lang}, head, body)
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
func Head(childNodes ...htmfunc.Component) htmfunc.Component {
	tag := "head"
	return htmfunc.Element(tag, nil, childNodes...)
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
func Title(title string) htmfunc.Component {
	tag := "title"
	return htmfunc.Element(tag, nil, Text(title))
}

// TitleNoEscape is equivalent to [Title], but does not escape the given string.
//
// This is more efficient, but only use it if the given string is safe and not possibly influenced by user input.
func TitleNoEscape(title string) htmfunc.Component {
	return htmfunc.Element("title", nil, TextNoEscape(title))
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
func Base(attributes attr.Ls) htmfunc.Component {
	return htmfunc.VoidElement("base", attributes)
}

// Link creates the [link element]. The link element allows authors to link their document to other resources.
//
// The address of the link(s) is given by the href attribute. If the href attribute is present, then its value must be a
// valid non-empty URL potentially surrounded by spaces. One or both of the href or imagesrcset attributes must be
// present.
//
// [link element]: https://html.spec.whatwg.org/#the-link-element
func Link(attributes attr.Ls) htmfunc.Component {
	return htmfunc.VoidElement("link", attributes)
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
func Meta(attributes attr.Ls) htmfunc.Component {
	return htmfunc.VoidElement("meta", attributes)
}

// Style creates the [style element].
//
// The style element allows authors to embed CSS style sheets in their documents. The style element is one of several
// inputs to the styling processing model. The element does not represent content for the user.
//
// [style element]: https://html.spec.whatwg.org/#the-style-element
func Style(attributes attr.Ls, css string) htmfunc.Component {
	return htmfunc.Element("style", attributes, Text(css))
}

// StyleNoEscape is equivalent to [Style], but does not escape the given string.
//
// This is more efficient, but only use it if the given string is safe and not possibly influenced by user input.
func StyleNoEscape(attributes attr.Ls, css string) htmfunc.Component {
	return htmfunc.Element("style", attributes, TextNoEscape(css))
}
