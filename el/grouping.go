package el

import (
	"github.com/ch-schulz/htmfunc"
	"github.com/ch-schulz/htmfunc/attr"
)

// P creates a [p element].
//
// The p element represents a paragraph.
//
// [p element]: https://html.spec.whatwg.org/#the-p-element
func P(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("p", attributes, children...)
}

// Hr creates a [hr element].
//
// The hr element represents a paragraph-level thematic break, e.g., a scene change in a story, or a transition to
// another topic within a section of a reference book; alternatively, it represents a separator between a set of options
// of a select element.
//
// [hr element]: https://html.spec.whatwg.org/#the-hr-element
func Hr(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("hr", attributes, children...)
}

// Pre creates a [pre element].
//
// The pre element represents a block of preformatted text, in which structure is represented by typographic conventions
// rather than by elements.
//
// [pre element]: https://html.spec.whatwg.org/#the-pre-element
func Pre(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("pre", attributes, children...)
}

// Blockquote creates a [blockquote element].
//
// The blockquote element represents a section that is quoted from another source.
//
// Content inside a blockquote must be quoted from another source, whose address, if it has one, may be cited in the
// cite attribute.
//
// If the cite attribute is present, it must be a valid URL potentially surrounded by spaces.To obtain the corresponding
// citation link, the value of the attribute must be parsed relative to the element's node document. User agents may
// allow users to follow such citation links, but they are primarily intended for private use (e.g., by server-side
// scripts collecting statistics about a site's use of quotations), not for readers.
//
// The content of a blockquote may be abbreviated or may have context added in the conventional manner for the text's
// language.
//
// Attribution for the quotation, if any, must be placed outside the blockquote element.
//
// [blockquote element]: https://html.spec.whatwg.org/#the-blockquote-element
func Blockquote(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("blockquote", attributes, children...)
}

// Ol creates a [ol element].
//
// The ol element represents a list of items, where the items have been intentionally ordered, such that changing the
// order would change the meaning of the document.
//
// The items of the list are the li element child nodes of the ol element, in tree order.
//
// [ol element]: https://html.spec.whatwg.org/#the-ol-element
func Ol(attributes attr.Ls, listItems ...htmfunc.ListItem) htmfunc.Component {
	return htmfunc.Element("ol", attributes, listItems...)
}

// Ul creates a [ul element].
//
// The ul element represents a list of items, where the order of the items is not important — that is, where changing
// the order would not materially change the meaning of the document.
//
// The items of the list are the li element child nodes of the ul element.
//
// [ul element]: https://html.spec.whatwg.org/#the-ul-element
func Ul(attributes attr.Ls, listItems ...htmfunc.ListItem) htmfunc.Component {
	return htmfunc.Element("ul", attributes, listItems...)
}

// Menu creates a [menu element].
//
// The menu element represents a toolbar consisting of its contents, in the form of an unordered list of items
// (represented by li elements), each of which represents a command that the user can perform or activate.
//
// [menu element]: https://html.spec.whatwg.org/#the-menu-element
func Menu(attributes attr.Ls, listItems ...htmfunc.ListItem) htmfunc.Component {
	return htmfunc.Element("menu", attributes, listItems...)
}

// Li creates a [li element].
//
// The li element represents a list item. If its parent element is an ol, ul, or menu element, then the element is an
// item of the parent element's list, as defined for those elements. Otherwise, the list item has no defined
// list-related relationship to any other li element.
//
// The value attribute, if present, must be a valid integer. It is used to determine the ordinal value of the list item,
// when the li's list owner is an ol element.
//
// [li element]: https://html.spec.whatwg.org/#the-li-element
func Li(attributes attr.Ls, children ...htmfunc.Component) htmfunc.ListItem {
	return htmfunc.Element("li", attributes, children...)
}

// Dl creates a [dl element].
//
// The dl element represents an association list consisting of zero or more name-value groups (a description list).
// A name-value group consists of one or more names (dt elements, possibly as children of a div element child) followed
// by one or more values (dd elements, possibly as children of a div element child), ignoring any nodes other than dt
// and dd element children, and dt and dd elements that are children of div element children. Within a single dl
// element, there should not be more than one dt element for each name.
//
// Name-value groups may be terms and definitions, metadata topics and values, questions and answers, or any other
// groups of name-value data.
//
// [dl element]: https://html.spec.whatwg.org/#the-dl-element
func Dl(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("dl", attributes, children...)
}

// Dt creates a [dt element].
//
// The dt element represents the term, or name, part of a term-description group in a description list (dl element).
//
// [dt element]: https://html.spec.whatwg.org/#the-dt-element
func Dt(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("dt", attributes, children...)
}

// Dd creates a [dd element].
//
// The dd element represents the description, definition, or value, part of a term-description group in a description
// list (dl element).
//
// [dd element]: https://html.spec.whatwg.org/#the-dd-element
func Dd(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("dd", attributes, children...)
}

// Figure creates a [figure element]
//
// The figure element represents some flow content, optionally with a caption, that is self-contained (like a complete
// sentence) and is typically referenced as a single unit from the main flow of the document.
//
// [figure element]: https://html.spec.whatwg.org/#the-figu-element
func Figure(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("figure", attributes, children...)
}

// Figcaption creates a [figcaption element].
//
// The figcaption element represents a caption or legend for the rest of the contents of the figcaption element's parent
// figure element, if any.
//
// [figcaption element]: https://html.spec.whatwg.org/#the-figcaption-element
func Figcaption(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("figcaption", attributes, children...)
}

// Main creates a [main element]
//
// The main element represents the dominant contents of the document.
//
// A document must not have more than one main element that does not have the hidden attribute specified.
//
// A hierarchically correct main element is one whose ancestor elements are limited to html, body, div, form without an
// accessible name, and autonomous custom elements. Each main element must be a hierarchically correct main element.
//
// [main element]: https://html.spec.whatwg.org/#the-main-element
func Main(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("main", attributes, children...)
}

// Search creates a [search element].
//
// The search element represents a part of a document or application that contains a set of form controls or other
// content related to performing a search or filtering operation. This could be a search of the web site or application;
// a way of searching or filtering search results on the current web page; or a global or Internet-wide search function.
//
// [search element]: https://html.spec.whatwg.org/#the-search-element
func Search(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("search", attributes, children...)
}

// Div creates a [div element].
//
// The div element has no special meaning at all. It represents its children. It can be used with the class, lang, and
// title attributes to mark up semantics common to a group of consecutive elements. It can also be used in a dl element,
// wrapping groups of dt and dd elements.
//
// [div element]: https://html.spec.whatwg.org/#the-div-element
func Div(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("div", attributes, children...)
}
