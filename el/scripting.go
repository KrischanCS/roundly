package el

import (
	"github.com/ch-schulz/htmfunc"
)

// Script creates a [script element].
//
// The script element allows authors to include dynamic script and data blocks in their documents. The element does not
// represent content for the user.
//
// [script element]: https://html.spec.whatwg.org/#the-script-element
func Script(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("script", attributes, children...)
}

// Noscript creates a [noscript element].
//
// The noscript element represents nothing if scripting is enabled, and represents its children if scripting is
// disabled. It is used to present different markup to user agents that support scripting and those that don't support
// scripting, by affecting how the document is parsed.
//
// [noscript element]: https://html.spec.whatwg.org/#the-noscript-element
func Noscript(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("noscript", attributes, children...)
}

// Template creates a [template element].
//
// The template element is used to declare fragments of HTML that can be cloned and inserted in the document by script.
//
// In a rendering, the template element represents nothing.
//
// [template element]: https://html.spec.whatwg.org/#the-template-element
func Template(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("template", attributes, children...)
}

// Slot creates a [slot element].
//
// The slot element defines a slot. It is typically used in a shadow tree. A slot element represents its assigned nodes,
// if any, and its contents otherwise.
//
// [template element]: https://html.spec.whatwg.org/#the-template-element
func Slot(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("slot", attributes, children...)
}
