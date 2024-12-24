package element

import (
	"time"

	"github.com/ch-schulz/htmfunc"
	"github.com/ch-schulz/htmfunc/attribute"
)

// A creates a [a element].
//
// If the a element has an href attribute, then it represents a hyperlink (a hypertext anchor) labeled by its contents.
//
// If the a element has no href attribute, then the element represents a placeholder for where a link might otherwise
// have been placed, if it had been relevant, consisting of just the element's contents.
//
// The target, download, ping, rel, hreflang, type, and referrerpolicy attributes must be omitted if the href attribute
// is not present.
//
// If the itemprop attribute is specified on an a element, then the href attribute must also be specified.
//
// [a element]: https://html.spec.whatwg.org/#the-a-element
func A(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("a", attributes, children...)
}

// Em creates a [em element].
//
// The em element represents stress emphasis of its contents.
//
// The level of stress that a particular piece of content has is given by its number of ancestor em elements.
//
// [em element]: https://html.spec.whatwg.org/#the-em-element
func Em(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("em", attributes, children...)
}

// Strong creates a [strong element].
//
// The strong element represents strong importance, seriousness, or urgency for its contents.
//
// Importance: the strong element can be used in a heading, caption, or paragraph to distinguish the part that really
// matters from other parts that might be more detailed, more jovial, or merely boilerplate. (This is distinct from
// marking up subheadings, for which the hgroup element is appropriate.)
//
// [strong element]: https://html.spec.whatwg.org/#the-strong-element
func Strong(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("strong", attributes, children...)
}

// Small creates a [small element].
//
// The small element represents side comments such as small print.
//
// [small element]: https://html.spec.whatwg.org/#the-small-element
func Small(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("small", attributes, children...)
}

// S creates a [s element].
//
// The s element represents contents that are no longer accurate or no longer relevant.
//
// The small element represents side comments such as small print.
//
// [s element]: https://html.spec.whatwg.org/#the-s-element
func S(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("s", attributes, children...)
}

// Cite creates a [cite element].
//
// The cite element represents the title of a work (e.g. a book, a paper, an essay, a poem, a score, a song, a script,
// a film, a TV show, a game, a sculpture, a painting, a theatre production, a play, an opera, a musical, an exhibition,
// a legal case report, a computer program, etc.). This can be a work that is being quoted or referenced in detail
// (i.e., a citation), or it can just be a work that is mentioned in passing.
//
// A person's name is not the title of a work — even if people call that person a piece of work — and the element must
// therefore not be used to mark up people's names. (In some cases, the b element might be appropriate for names; e.g.
// in a gossip article where the names of famous people are keywords rendered with a different style to draw attention
// to them. In other cases, if an element is really needed, the span element can be used.)
//
// [cite element]: https://html.spec.whatwg.org/#the-cite-element
func Cite(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("cite", attributes, children...)
}

// Q creates a [q element].
//
// The q element represents some phrasing content quoted from another source.
//
// Quotation punctuation (such as quotation marks) that is quoting the contents of the element must not appear
// immediately before, after, or inside q elements; they will be inserted into the rendering by the user agent.
//
// Content inside a q element must be quoted from another source, whose address, if it has one, may be cited in the cite
// attr. The source may be fictional, as when quoting characters in a novel or screenplay.
//
// If the cite attribute is present, it must be a valid URL potentially surrounded by spaces. To obtain the
// corresponding citation link, the value of the attribute must be parsed relative to the element's node document.
// User agents may allow users to follow such citation links, but they are primarily intended for private use (e.g., by
// server-side scripts collecting statistics about a site's use of quotations), not for readers.
//
// The q element must not be used in place of quotation marks that do not represent quotes; for example, it is
// inappropriate to use the q element for marking up sarcastic statements.
//
// The use of q elements to mark up quotations is entirely optional; using explicit quotation punctuation without q
// elements is just as correct.
//
// [q element]: https://html.spec.whatwg.org/#the-q-element
func Q(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("q", attributes, children...)
}

// Dfn creates a [dfn element]
//
// The dfn element represents the defining instance of a term. The paragraph, description list group, or section that is
// the nearest ancestor of the dfn element must also contain the definition(s) for the term given by the dfn element.
//
// Defining term: if the dfn element has a title attribute, then the exact value of that attribute is the term being
// defined. Otherwise, if it contains exactly one element child node and no child Text nodes, and that child element is
// an abbr element with a title attribute, then the exact value of that attribute is the term being defined.Otherwise,
// it is the descendant text content of the dfn element that gives the term being defined.
//
// If the title attribute of the dfn element is present, then it must contain only the term being defined.
//
// An a element that links to a dfn element represents an instance of the term defined by the dfn element.
//
// [dfn element]: https://html.spec.whatwg.org/#the-dfn-element
func Dfn(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("dfn", attributes, children...)
}

// Abbr creates a [abbr element].
//
// The abbr element represents an abbreviation or acronym, optionally with its expansion. The title attribute may be
// used to provide an expansion of the abbreviation. The attribute, if specified, must contain an expansion of the
// abbreviation, and nothing else.
//
// [abbr element]: https://html.spec.whatwg.org/#the-abbr-element
func Abbr(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("abbr", attributes, children...)
}

// Ruby creates a [ruby element].
//
// The ruby element allows one or more spans of phrasing content to be marked with ruby annotations.Ruby annotations are
// short runs of text presented alongside base text, primarily used in East Asian typography as a guide for
// pronunciation or to include other annotations.In Japanese, this form of typography is also known as furigana.
//
// The content model of ruby elements consists of one or more of the following sequences:
//
// One or the other of the following:
//
//   - Phrasing content, but with no ruby elements and with no ruby element descendants
//   - A single ruby element that itself has no ruby element descendants
//
// One or the other of the following:
//
//   - One or more rt elements
//   - An rp element followed by one or more rt elements, each of which is itself followed by an rp element
//
// The ruby and rt elements can be used for a variety of kinds of annotations, including in particular (though by no
// means limited to) those described below.For more details on Japanese Ruby in particular, and how to render Ruby for
// Japanese, see Requirements for Japanese Text Layout. [JLREQ]
//
// [ruby element]: https://html.spec.whatwg.org/#the-ruby-element
// [JLREQ]: https://html.spec.whatwg.org/#refsJLREQ
func Ruby(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("ruby", attributes, children...)
}

// Rt creates a [rt element].
//
// The rt element marks the ruby text component of a ruby annotation. When it is the child of a ruby element, it doesn't
// represent anything itself, but the ruby element uses it as part of determining what it represents.
//
// An rt element that is not a child of a ruby element represents the same thing as its children.
//
// [rt element]: https://html.spec.whatwg.org/#the-rt-element
func Rt(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("rt", attributes, children...)
}

// Rp creates a [rp element].
//
// The rp element can be used to provide parentheses or other content around a ruby text component of a ruby annotation,
// to be shown by user agents that don't support ruby annotations.
//
// An rp element that is a child of a ruby element represents nothing. An rp element whose parent element is not a ruby
// element represents its children.
//
// [rp element]: https://html.spec.whatwg.org/#the-rp-element
func Rp(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("rp", attributes, children...)
}

// Data creates a [data element].
//
// The data element represents its contents, along with a machine-readable form of those contents in the value
// attr.
//
// The value attribute must be present. Its value must be a representation of the element's contents in a
// machine-readable format.
//
// [data element]: https://html.spec.whatwg.org/#the-data-element
func Data(value string, attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	attributes = attribute.Join(attribute.Value(value), attributes)
	return htmfunc.WriteElement("data", attributes, children...)
}

// TimeMachineReadableAsContent creates a [time element].
//
// The content of the element will be formatted according to [time.RFC3339]
//
// The time element represents its contents, along with a machine-readable form of those contents in the datetime
// attr. The kind of content is limited to various kinds of dates, times, time-zone offsets, and durations, as
// described below.
//
// The datetime attribute may be present. If present, its value must be a representation of the element's contents in a
// machine-readable format.
//
// A time element that does not have a datetime content attribute must not have any element descendants.
//
// The datetime value of a time element is the value of the element's datetime content attribute, if it has one,
// otherwise the child text content of the time element.
//
// [time element]: https://html.spec.whatwg.org/#the-time-element
func TimeMachineReadableAsContent(attributes htmfunc.Attribute, t time.Time) htmfunc.Element {
	return htmfunc.WriteElement("time", attributes, TextTrusted(t.Format(time.RFC3339)))
}

// TimeAttribute creates a [time element].
//
// The datetime attribute will be filled with the given time in [time.RFC3339] format
//
// The time element represents its contents, along with a machine-readable form of those contents in the datetime
// attr. The kind of content is limited to various kinds of dates, times, time-zone offsets, and durations, as
// described below.
//
// The datetime attribute may be present. If present, its value must be a representation of the element's contents in a
// machine-readable format.
//
// A time element that does not have a datetime content attribute must not have any element descendants.
//
// The datetime value of a time element is the value of the element's datetime content attribute, if it has one,
// otherwise the child text content of the time element.
//
// [time element]: https://html.spec.whatwg.org/#the-time-element
func TimeAttribute(attributes htmfunc.Attribute, t time.Time, childNodes ...htmfunc.Element) htmfunc.Element {
	attributes = attribute.Join(attribute.DateTime(t.Format(time.RFC3339)), attributes)
	return htmfunc.WriteElement("time", attributes, childNodes...)
}

// Code creates a [code element].
//
// The code element represents a fragment of computer code. This could be an XML element name, a filename, a computer
// program, or any other string that a computer would recognize.
//
// There is no formal way to indicate the language of computer code being marked up. Authors who wish to mark code
// elements with the language used, e.g. so that syntax highlighting scripts can use the right rules, can use the class
// attribute, e.g. by adding a class prefixed with "language-" to the element.
//
// [code element]: https://html.spec.whatwg.org/#the-code-element
func Code(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("code", attributes, children...)
}

// Var creates a [var element].
//
// The var element represents a variable. This could be an actual variable in a mathematical expression or programming
// context, an identifier representing a constant, a symbol identifying a physical quantity, a function parameter, or
// just be a term used as a placeholder in prose.
//
// [var element]: https://html.spec.whatwg.org/#the-var-element
func Var(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("var", attributes, children...)
}

// Samp creates a [samp element].
//
// The samp element represents sample or quoted output from another program or computing system.
//
// [samp element]: https://html.spec.whatwg.org/#the-samp-element
func Samp(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("samp", attributes, children...)
}

// Kbd creates a [kbd element].
//
// The kbd element represents user input (typically keyboard input, although it may also be used to represent other
// input, such as voice commands).
//
// When the kbd element is nested inside a samp element, it represents the input as it was echoed by the system.
//
// When the kbd element contains a samp element, it represents input based on system output, for example invoking a menu
// item.
//
// When the kbd element is nested inside another kbd element, it represents an actual key or other single unit of input
// as appropriate for the input mechanism.
//
// [kbd element]: https://html.spec.whatwg.org/#the-kbd-element
func Kbd(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("kbd", attributes, children...)
}

// Sub creates a [sub element].
//
// The sup element represents a superscript and the sub element represents a subscript.
//
// These elements must be used only to mark up typographical conventions with specific meanings, not for typographical
// presentation for presentation's sake. For example, it would be inappropriate for the sub and sup elements to be used
// in the name of the LaTeX document preparation system. In general, authors should use these elements only if the
// absence of those elements would change the meaning of the content.
//
// In certain languages, superscripts are part of the typographical conventions for some abbreviations.
//
// [sub element]: https://html.spec.whatwg.org/#the-sub-element
func Sub(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("sub", attributes, children...)
}

// Sup creates a [sup element].
//
// The sup element represents a superscript and the sub element represents a subscript.
//
// These elements must be used only to mark up typographical conventions with specific meanings, not for typographical
// presentation for presentation's sake. For example, it would be inappropriate for the sub and sup elements to be used
// in the name of the LaTeX document preparation system. In general, authors should use these elements only if the
// absence of those elements would change the meaning of the content.
//
// In certain languages, superscripts are part of the typographical conventions for some abbreviations.
//
// [sup element]: https://html.spec.whatwg.org/#the-sup-element
func Sup(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("sup", attributes, children...)
}

// I creates an [i element].
//
// The i element represents a span of text in an alternate voice or mood, or otherwise offset from the normal prose in a
// manner indicating a different quality of text, such as a taxonomic designation, a technical term, an idiomatic phrase
// from another language, transliteration, a thought, or a ship name in Western texts.
//
// Terms in languages different from the main text should be annotated with lang attributes (or, in XML, lang attributes
// in the XML namespace).
//
// [i element]: https://html.spec.whatwg.org/#the-i-element
func I(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("i", attributes, children...)
}

// B creates a [b element].
//
// The b element represents a span of text to which attention is being drawn for utilitarian purposes without conveying
// any extra importance and with no implication of an alternate voice or mood, such as key words in a document abstract,
// product names in a review, actionable words in interactive text-driven software, or an article lede.
//
// [b element]: https://html.spec.whatwg.org/#the-b-element
func B(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("b", attributes, children...)
}

// U creates a [u element].
//
// The u element represents a span of text with an unarticulated, though explicitly rendered, non-textual annotation,
// such as labeling the text as being a proper name in Chinese text (a Chinese proper name mark), or labeling the text
// as being misspelt.
//
// In most cases, another element is likely to be more appropriate: for marking stress emphasis, the em element should
// be used; for marking key words or phrases either the b element or the mark element should be used, depending on the
// context; for marking book titles, the cite element should be used; for labeling text with explicit textual
// annotations, the ruby element should be used; for technical terms, taxonomic designation, transliteration, a thought,
// or for labeling ship names in Western texts, the i element should be used.
//
// [u element]: https://html.spec.whatwg.org/#the-u-element
func U(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("u", attributes, children...)
}

// Mark creates a [mark element].
//
// The mark element represents a run of text in one document marked or highlighted for reference purposes, due to its
// relevance in another context. When used in a quotation or other block of text referred to from the prose, it
// indicates a highlight that was not originally present but which has been added to bring the reader's attention to a
// part of the text that might not have been considered important by the original author when the block was originally
// written, but which is now under previously unexpected scrutiny. When used in the main prose of a document, it
// indicates a part of the document that has been highlighted due to its likely relevance to the user's current
// activity.
//
// [mark element]: https://html.spec.whatwg.org/#the-mark-element
func Mark(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("mark", attributes, children...)
}

// Bdi creates a [bdi element].
//
// The bdi element represents a span of text that is to be isolated from its surroundings for the purposes of
// [bidirectional text formatting].
//
// [bdi element]: https://html.spec.whatwg.org/#the-bdi-element
// [bidirectional text formatting]: https://html.spec.whatwg.org/#refsBIDI
func Bdi(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("bdi", attributes, children...)
}

// Bdo creates a [bdo element].
//
// The bdo element represents explicit text directionality formatting control for its children.It allows authors to
// override the Unicode bidirectional algorithm by explicitly specifying a direction override.[BIDI]
//
// Authors must specify the dir attribute on this element, with the value ltr to specify a left-to-right override and
// with the value rtl to specify a right-to-left override. The auto value must not be specified.
//
// [bdo element]: https://html.spec.whatwg.org/#the-bdo-element
func Bdo(direction htmfunc.TextDirection, attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	attributes = attribute.Join(attribute.Dir(direction), attributes)
	return htmfunc.WriteElement("bdo", attributes, children...)
}

// Span creates a [span element].
//
// The span element doesn't mean anything on its own, but can be useful when used together with the global attributes,
// e.g. class, lang, or dir. It represents its children.
//
// [span element]: https://html.spec.whatwg.org/#the-span-element
func Span(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("span", attributes, children...)
}

// Br creates a [br element].
//
// [br element]: https://html.spec.whatwg.org/#the-br-element
func Br() htmfunc.Element {
	return htmfunc.WriteVoidElement("br", nil)
}

// Wbr creates a [wbr element].
//
// The wbr element represents a line break opportunity.
//
// Any content inside wbr elements must not be considered part of the surrounding text.
//
// [wbr element]: https://html.spec.whatwg.org/#the-wbr-element
func Wbr(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("wbr", attributes, children...)
}
