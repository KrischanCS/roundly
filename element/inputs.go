package element

import (
	"github.com/ch-schulz/htmfunc"
)

// Form creates a [form element].
//
// A form is a component of a web page that has form controls, such as text, buttons, checkboxes, range, or color picker
// controls. A user can interact with such a form, providing data that can then be sent to the server for further
// processing (e.g. returning the results of a search or calculation). No client-side scripting is needed in many cases,
// though an API is available so that scripts can augment the user experience or use forms for purposes other than
// submitting data to a server.
//
// [form element]: https://html.spec.whatwg.org/#the-form-element
func Form(attributes htmfunc.AttributeWriteFunc, children ...htmfunc.ElementWriteFunc) htmfunc.ElementWriteFunc {
	return htmfunc.WriteElement("form", attributes, children...)
}

// Label creates a [label element].
//
// The label element represents a caption in a user interface. The caption can be associated with a specific form
// control, known as the label element's labeled control, either using the for attribute, or by putting the form control
// inside the label element itself.
//
// [label element]: https://html.spec.whatwg.org/#the-label-element
func Label(attributes htmfunc.AttributeWriteFunc, children ...htmfunc.ElementWriteFunc) htmfunc.ElementWriteFunc {
	return htmfunc.WriteElement("label", attributes, children...)
}

// Input creates a [input element].
//
// The input element represents a typed data field, usually with a form control to allow the user to edit the data.
//
// The type attribute controls the data type (and associated control) of the element. It is an enumerated attribute.
// The following table lists the keywords and states for the attribute — the keywords in the left column map to the
// states in the cell in the second column on the same row as the keyword.
//
// [input element]: https://html.spec.whatwg.org/#the-input-element
func Input(attributes htmfunc.AttributeWriteFunc) htmfunc.ElementWriteFunc {
	return htmfunc.WriteVoidElement("input", attributes)
}

// Button creates a [button element].
//
// The button element represents a button labeled by its contents.
//
// [button element]: https://html.spec.whatwg.org/#the-button-element
func Button(attributes htmfunc.AttributeWriteFunc, children ...htmfunc.ElementWriteFunc) htmfunc.ElementWriteFunc {
	return htmfunc.WriteElement("button", attributes, children...)
}

// Select creates a [select element].
//
// The select element represents a control for selecting amongst a set of options.
//
// [select element]: https://html.spec.whatwg.org/#the-select-element
func Select(attributes htmfunc.AttributeWriteFunc, children ...htmfunc.ElementWriteFunc) htmfunc.ElementWriteFunc {
	return htmfunc.WriteElement("select", attributes, children...)
}

// Datalist creates a [datalist element].
//
// The datalist element represents a set of option elements that represent predefined options for other controls. In the
// rendering, the datalist element represents nothing and it, along with its children, should be hidden.
//
// [datalist element]: https://html.spec.whatwg.org/#the-datalist-element
func Datalist(attributes htmfunc.AttributeWriteFunc, children ...htmfunc.ElementWriteFunc) htmfunc.ElementWriteFunc {
	return htmfunc.WriteElement("datalist", attributes, children...)
}

// Optgroup creates a [optgroup element].
//
// The optgroup element represents a group of option elements with a common label.
//
// The element's group of option elements consists of the option elements that are children of the optgroup element.
//
// [optgroup element]: https://html.spec.whatwg.org/#the-optgroup-element
func Optgroup(attributes htmfunc.AttributeWriteFunc, children ...htmfunc.ElementWriteFunc) htmfunc.ElementWriteFunc {
	return htmfunc.WriteElement("optgroup", attributes, children...)
}

// Option creates a [option element].
//
// The option element represents an option in a select element or as part of a list of suggestions in a datalist
// element.
//
// [option element]: https://html.spec.whatwg.org/#the-option-element
func Option(attributes htmfunc.AttributeWriteFunc, children ...htmfunc.ElementWriteFunc) htmfunc.ElementWriteFunc {
	return htmfunc.WriteElement("option", attributes, children...)
}

// Textarea creates a [textarea element].
//
// The textarea element represents a multiline plain text edit control for the element's raw value. The contents of the
// control represent the control's default value.
//
// The raw value of a textarea control must be initially the empty string.
//
// [textarea element]: https://html.spec.whatwg.org/#the-textarea-element
func Textarea(attributes htmfunc.AttributeWriteFunc, children ...htmfunc.ElementWriteFunc) htmfunc.ElementWriteFunc {
	return htmfunc.WriteElement("textarea", attributes, children...)
}

// Output creates a [output element].
//
// The output element represents the result of a calculation performed by the application, or the result of a user
// action.
//
// [output element]: https://html.spec.whatwg.org/#the-output-element
func Output(attributes htmfunc.AttributeWriteFunc, children ...htmfunc.ElementWriteFunc) htmfunc.ElementWriteFunc {
	return htmfunc.WriteElement("output", attributes, children...)
}

// Progress creates a [progress element]
//
// The progress element represents the completion progress of a task. The progress is either indeterminate, indicating
// that progress is being made but that it is not clear how much more work remains to be done before the task is
// complete (e.g. because the task is waiting for a remote host to respond), or the progress is a number in the range
// zero to a maximum, giving the fraction of work that has so far been completed.
//
// There are two attributes that determine the current task completion represented by the element. The value attribute
// specifies how much of the task has been completed, and the max attribute specifies how much work the task requires in
// total. The units are arbitrary and not specified.
//
// [progress element]: https://html.spec.whatwg.org/#the-progress-element
func Progress(attributes htmfunc.AttributeWriteFunc, children ...htmfunc.ElementWriteFunc) htmfunc.ElementWriteFunc {
	return htmfunc.WriteElement("progress", attributes, children...)
}

// Meter creates a [meter element].
//
// The meter element represents a scalar measurement within a known range, or a fractional value; for example disk
// usage, the relevance of a query result, or the fraction of a voting population to have selected a particular
// candidate.
//
// This is also known as a gauge.
//
// The meter element should not be used to indicate progress (as in a progress bar). For that role, HTML provides a
// separate progress element.
//
// [meter element]: https://html.spec.whatwg.org/#the-meter-element
func Meter(attributes htmfunc.AttributeWriteFunc, children ...htmfunc.ElementWriteFunc) htmfunc.ElementWriteFunc {
	return htmfunc.WriteElement("meter", attributes, children...)
}

// Fieldset creates a [fieldset element].
//
// The fieldset element represents a set of form controls (or other content) grouped together, optionally with a
// caption. The caption is given by the first legend element that is a child of the fieldset element, if any. The
// remainder of the descendants form the group.
//
// [fieldset element]: https://html.spec.whatwg.org/#the-fieldset-element
func Fieldset(attributes htmfunc.AttributeWriteFunc, children ...htmfunc.ElementWriteFunc) htmfunc.ElementWriteFunc {
	return htmfunc.WriteElement("fieldset", attributes, children...)
}

// Legend creates a [legend element].
//
// The legend element represents a caption for the rest of the contents of the legend element's parent fieldset element,
// if any.
//
// [legend element]: https://html.spec.whatwg.org/#the-legend-element
func Legend(attributes htmfunc.AttributeWriteFunc, children ...htmfunc.ElementWriteFunc) htmfunc.ElementWriteFunc {
	return htmfunc.WriteElement("legend", attributes, children...)
}
