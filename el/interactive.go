package el

import (
	"github.com/ch-schulz/htmfunc"
)

// Details creates a [details element].
//
// The details element represents a disclosure widget from which the user can obtain additional information or controls.
//
// [details element]: https://html.spec.whatwg.org/#the-details-element
func Details(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("details", attributes, children...)
}

// Summary creates a [summary element].
// The summary element represents a summary, caption, or legend for the rest of the contents of the summary element's
// parent details element, if any.
//
// A summary element is a summary for its parent details if the following algorithm returns true:
//   - If this summary element has no parent, then return false.
//   - Let parent be this summary element's parent.
//   - If parent is not a details element, then return false.
//   - If parent's first summary element child is not this summary element, then return false.
//   - Return true.
//
// The activation behavior of summary elements is to run the following steps:
//   - If this summary element is not the summary for its parent details, then return.
//   - Let parent be this summary element's parent.
//   - If the open attribute is present on parent, then remove it. Otherwise, set parent's open attribute to the empty
//     string.
//
// [summary element]: https://html.spec.whatwg.org/#the-summary-element
func Summary(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("summary", attributes, children...)
}

// Dialog creates a [dialog element].
//
// The dialog element represents a transitory part of an application, in the form of a small window ("dialog box"),
// which the user interacts with to perform a task or gather information. Once the user is done, the dialog can be
// automatically closed by the application, or manually closed by the user.
//
// [dialog element]: https://html.spec.whatwg.org/#the-dialog-element
func Dialog(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("dialog", attributes, children...)
}
