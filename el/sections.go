package el

import (
	"fmt"
	"github.com/ch-schulz/htmfunc"
	"github.com/ch-schulz/htmfunc/attr"
)

// Body creates the [body element] The body element represents the contents of the document.
//
// In conforming documents, there is only one body element. The document.body IDL attribute provides scripts with easy
// access to a document's body element.
//
// [body element]: https://html.spec.whatwg.org/#the-body-element
func Body(attributes []htmfunc.Attribute, childNodes ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("body", attributes, childNodes...)
}

// Article creates a [article element]
//
// The article element represents a complete, or self-contained, composition in a document, page, application, or site
// and that is, in principle, independently distributable or reusable, e.g. in syndication. This could be a forum post,
// a magazine or newspaper article, a blog entry, a user-submitted comment, an interactive widget or gadget, or any
// other independent item of content.
//
// When article elements are nested, the inner article elements represent articles that are in principle related to the
// contents of the outer article.For instance, a blog entry on a site that accepts user-submitted comments could
// represent the comments as article elements nested within the article element for the blog entry.
//
// Author information associated with an article element (q.v.the address element) does not apply to nested article
// elements.
//
// [article element]: https://html.spec.whatwg.org/#the-article-element
func Article(attributes attr.Ls, childNodes ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("article", attributes, childNodes...)
}

// Section creates a [section element].
//
// The section element represents a generic section of a document or application. A section, in this context, is a
// thematic grouping of content, typically with a heading.
//
// [section element]: https://html.spec.whatwg.org/#the-section-element
func Section(attributes attr.Ls, childNodes ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("section", attributes, childNodes...)
}

// Nav creates a [nav element].
//
// The nav element represents a section of a page that links to other pages or to parts within the page: a section with
// navigation links.
//
// [nav element]: https://html.spec.whatwg.org/#the-nav-element
func Nav(attributes attr.Ls, childNodes ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("nav", attributes, childNodes...)
}

// Aside creates a [aside element].
//
// The aside element represents a section of a page that consists of content that is tangentially related to the content
// around the aside element, and which could be considered separate from that content. Such sections are often
// represented as sidebars in printed typography.
//
// The element can be used for typographical effects like pull quotes or sidebars, for advertising, for groups of nav
// elements, and for other content that is considered separate from the main content of the page.
//
// [aside element]: https://html.spec.whatwg.org/#the-aside-element
func Aside(attributes attr.Ls, childNodes ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("aside", attributes, childNodes...)
}

// H creates a [heading element] (h1, h2, ..., h6).
//
// H will create invalid headings when given numbers > 6
//
// These elements represent headings for their sections.
//
// [heading element]: https://html.spec.whatwg.org/#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements
func H(level uint, attributes attr.Ls, childNodes ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement(fmt.Sprintf("h%d", level), attributes, childNodes...)
}

// Hgroup creates a [hgroup element].
//
// The hgroup element represents a heading and related content. The element may be used to group an h1–h6 element with
// one or more p elements containing content representing a subheading, alternative title, or tagline.
//
// [hgroup element]: https://html.spec.whatwg.org/#the-hgroup-element
func Hgroup(attributes attr.Ls, childNodes ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("hgroup", attributes, childNodes...)
}

// Header creates a [header element].
//
// The header element represents a group of introductory or navigational aids.
//
// [header element]: https://html.spec.whatwg.org/#the-header-element
func Header(attributes attr.Ls, childNodes ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("header", attributes, childNodes...)
}

// Footer creates a [footer element].
//
// The footer element represents a footer for its nearest ancestor sectioning content element, or for the body element
// if there is no such ancestor. A footer typically contains information about its section such as who wrote it, links
// to related documents, copyright data, and the like.
//
// When the footer element contains entire sections, they represent appendices, indices, long colophons, verbose license
// agreements, and other such content.
//
// [footer element]: https://html.spec.whatwg.org/#the-footer-element
func Footer(attributes attr.Ls, childNodes ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("footer", attributes, childNodes...)
}

// Address creates a [address element].
//
// The address element represents the contact information for its nearest article or body element ancestor. If that is
// the body element, then the contact information applies to the document as a whole.
//
// [address element]: https://html.spec.whatwg.org/#the-address-element
func Address(attributes attr.Ls, childNodes ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("address", attributes, childNodes...)
}
