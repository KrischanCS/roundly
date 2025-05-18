// Generated file. DO NOT EDIT.

package element

import (
"github.com/KrischanCS/htmfunc"
)

// Button creates the button element - Button control [(More)]
//
// It belongs to the following categories:
//   - [flow]
//   - [phrasing]
//   - [interactive]
//   - [listed]
//   - [labelable]
//   - [submittable]
//   - [form-associated]
//   - [palpable]
//// It can be parent to the following elements/categories of elements:
//   - [phrasing]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [command]
//   - [commandfor]
//   - [disabled]
//   - [form]
//   - [formaction]
//   - [formenctype]
//   - [formmethod]
//   - [formnovalidate]
//   - [formtarget]
//   - [name]
//   - [popovertarget]
//   - [popovertargetaction]
//   - [type]
//   - [value]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [interactive]: https://html.spec.whatwg.org/dev/dom.html#interactive-content-2
// [listed]: https://html.spec.whatwg.org/dev/forms.html#category-listed
// [labelable]: https://html.spec.whatwg.org/dev/forms.html#category-label
// [submittable]: https://html.spec.whatwg.org/dev/forms.html#category-submit
// [form-associated]: https://html.spec.whatwg.org/dev/forms.html#form-associated-element
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [command]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-command
// [commandfor]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-commandfor
// [disabled]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-disabled
// [form]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [formaction]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formaction
// [formenctype]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formenctype
// [formmethod]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formmethod
// [formnovalidate]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formnovalidate
// [formtarget]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formtarget
// [name]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-name
// [popovertarget]: https://html.spec.whatwg.org/dev/popover.html#attr-popovertarget
// [popovertargetaction]: https://html.spec.whatwg.org/dev/popover.html#attr-popovertargetaction
// [type]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-type
// [value]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-value
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Button(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("button", attributes, children...)
}

// Datalist creates the datalist element - Container for options for [combo box control [(More)]
//
// It belongs to the following categories:
//   - [flow]
//   - [phrasing]
//// It can be parent to the following elements/categories of elements:
//   - [phrasing]
//   - [option]
//   - [script-supporting elements]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [combo box control]: https://html.spec.whatwg.org/dev/input.html#attr-input-list
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [option]: https://html.spec.whatwg.org/dev/form-elements.html#the-option-element
// [script-supporting elements]: https://html.spec.whatwg.org/dev/dom.html#script-supporting-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Datalist(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("datalist", attributes, children...)
}

// Fieldset creates the fieldset element - Group of form controls [(More)]
//
// It belongs to the following categories:
//   - [flow]
//   - [listed]
//   - [form-associated]
//   - [palpable]
//// It can be parent to the following elements/categories of elements:
//   - [legend]
//   - [flow]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [disabled]
//   - [form]
//   - [name]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [listed]: https://html.spec.whatwg.org/dev/forms.html#category-listed
// [form-associated]: https://html.spec.whatwg.org/dev/forms.html#form-associated-element
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [legend]: https://html.spec.whatwg.org/dev/form-elements.html#the-legend-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [disabled]: https://html.spec.whatwg.org/dev/form-elements.html#attr-fieldset-disabled
// [form]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [name]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-name
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Fieldset(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("fieldset", attributes, children...)
}

// Legend creates the legend element - Caption for [fieldset [(More)]
//
// It belongs to the following categories:
//   - none
//// It can be parent to the following elements/categories of elements:
//   - [phrasing]
//   - [heading content]
//
// If can itself be a child of the following elements/categories of elements:
//   - [fieldset]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [fieldset]: https://html.spec.whatwg.org/dev/form-elements.html#the-fieldset-element
// [fieldset]: https://html.spec.whatwg.org/dev/form-elements.html#the-fieldset-element
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [heading content]: https://html.spec.whatwg.org/dev/dom.html#heading-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Legend(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("legend", attributes, children...)
}

// Meter creates the meter element - Gauge [(More)]
//
// It belongs to the following categories:
//   - [flow]
//   - [phrasing]
//   - [labelable]
//   - [palpable]
//// It can be parent to the following elements/categories of elements:
//   - [phrasing]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [value]
//   - [min]
//   - [max]
//   - [low]
//   - [high]
//   - [optimum]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [labelable]: https://html.spec.whatwg.org/dev/forms.html#category-label
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [value]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-value
// [min]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-min
// [max]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-max
// [low]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-low
// [high]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-high
// [optimum]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-optimum
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Meter(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("meter", attributes, children...)
}

// Optgroup creates the optgroup element - Group of options in a list box [(More)]
//
// It belongs to the following categories:
//   - none
//// It can be parent to the following elements/categories of elements:
//   - [option]
//   - [script-supporting elements]
//
// If can itself be a child of the following elements/categories of elements:
//   - [select]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [disabled]
//   - [label]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [select]: https://html.spec.whatwg.org/dev/form-elements.html#the-select-element
// [option]: https://html.spec.whatwg.org/dev/form-elements.html#the-option-element
// [script-supporting elements]: https://html.spec.whatwg.org/dev/dom.html#script-supporting-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [disabled]: https://html.spec.whatwg.org/dev/form-elements.html#attr-optgroup-disabled
// [label]: https://html.spec.whatwg.org/dev/form-elements.html#attr-optgroup-label
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Optgroup(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("optgroup", attributes, children...)
}

// Option creates the option element - Option in a list box or combo box control [(More)]
//
// It belongs to the following categories:
//   - none
//// It can be parent to the following elements/categories of elements:
//   - [text]
//
// If can itself be a child of the following elements/categories of elements:
//   - [select]
//   - [datalist]
//   - [optgroup]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [disabled]
//   - [label]
//   - [selected]
//   - [value]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [select]: https://html.spec.whatwg.org/dev/form-elements.html#the-select-element
// [datalist]: https://html.spec.whatwg.org/dev/form-elements.html#the-datalist-element
// [optgroup]: https://html.spec.whatwg.org/dev/form-elements.html#the-optgroup-element
// [text]: https://html.spec.whatwg.org/dev/dom.html#text-content
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [disabled]: https://html.spec.whatwg.org/dev/form-elements.html#attr-option-disabled
// [label]: https://html.spec.whatwg.org/dev/form-elements.html#attr-option-label
// [selected]: https://html.spec.whatwg.org/dev/form-elements.html#attr-option-selected
// [value]: https://html.spec.whatwg.org/dev/form-elements.html#attr-option-value
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Option(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("option", attributes, children...)
}

// Output creates the output element - Calculated output value [(More)]
//
// It belongs to the following categories:
//   - [flow]
//   - [phrasing]
//   - [listed]
//   - [labelable]
//   - [resettable]
//   - [form-associated]
//   - [palpable]
//// It can be parent to the following elements/categories of elements:
//   - [phrasing]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [for]
//   - [form]
//   - [name]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [listed]: https://html.spec.whatwg.org/dev/forms.html#category-listed
// [labelable]: https://html.spec.whatwg.org/dev/forms.html#category-label
// [resettable]: https://html.spec.whatwg.org/dev/forms.html#category-reset
// [form-associated]: https://html.spec.whatwg.org/dev/forms.html#form-associated-element
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [for]: https://html.spec.whatwg.org/dev/form-elements.html#attr-output-for
// [form]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [name]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-name
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Output(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("output", attributes, children...)
}

// Progress creates the progress element - Progress bar [(More)]
//
// It belongs to the following categories:
//   - [flow]
//   - [phrasing]
//   - [labelable]
//   - [palpable]
//// It can be parent to the following elements/categories of elements:
//   - [phrasing]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [value]
//   - [max]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [labelable]: https://html.spec.whatwg.org/dev/forms.html#category-label
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [value]: https://html.spec.whatwg.org/dev/form-elements.html#attr-progress-value
// [max]: https://html.spec.whatwg.org/dev/form-elements.html#attr-progress-max
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Progress(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("progress", attributes, children...)
}

// Select creates the select element - List box control [(More)]
//
// It belongs to the following categories:
//   - [flow]
//   - [phrasing]
//   - [interactive]
//   - [listed]
//   - [labelable]
//   - [submittable]
//   - [resettable]
//   - [form-associated]
//   - [palpable]
//// It can be parent to the following elements/categories of elements:
//   - [option]
//   - [optgroup]
//   - [script-supporting elements]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [autocomplete]
//   - [disabled]
//   - [form]
//   - [multiple]
//   - [name]
//   - [required]
//   - [size]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [interactive]: https://html.spec.whatwg.org/dev/dom.html#interactive-content-2
// [listed]: https://html.spec.whatwg.org/dev/forms.html#category-listed
// [labelable]: https://html.spec.whatwg.org/dev/forms.html#category-label
// [submittable]: https://html.spec.whatwg.org/dev/forms.html#category-submit
// [resettable]: https://html.spec.whatwg.org/dev/forms.html#category-reset
// [form-associated]: https://html.spec.whatwg.org/dev/forms.html#form-associated-element
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [option]: https://html.spec.whatwg.org/dev/form-elements.html#the-option-element
// [optgroup]: https://html.spec.whatwg.org/dev/form-elements.html#the-optgroup-element
// [script-supporting elements]: https://html.spec.whatwg.org/dev/dom.html#script-supporting-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [autocomplete]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-autocomplete
// [disabled]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-disabled
// [form]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [multiple]: https://html.spec.whatwg.org/dev/form-elements.html#attr-select-multiple
// [name]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-name
// [required]: https://html.spec.whatwg.org/dev/form-elements.html#attr-select-required
// [size]: https://html.spec.whatwg.org/dev/form-elements.html#attr-select-size
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Select(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("select", attributes, children...)
}

// Textarea creates the textarea element - Multiline text controls [(More)]
//
// It belongs to the following categories:
//   - [flow]
//   - [phrasing]
//   - [interactive]
//   - [listed]
//   - [labelable]
//   - [submittable]
//   - [resettable]
//   - [form-associated]
//   - [palpable]
//// It can be parent to the following elements/categories of elements:
//   - [text]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals]
//   - [autocomplete]
//   - [cols]
//   - [dirname]
//   - [disabled]
//   - [form]
//   - [maxlength]
//   - [minlength]
//   - [name]
//   - [placeholder]
//   - [readonly]
//   - [required]
//   - [rows]
//   - [wrap]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [interactive]: https://html.spec.whatwg.org/dev/dom.html#interactive-content-2
// [listed]: https://html.spec.whatwg.org/dev/forms.html#category-listed
// [labelable]: https://html.spec.whatwg.org/dev/forms.html#category-label
// [submittable]: https://html.spec.whatwg.org/dev/forms.html#category-submit
// [resettable]: https://html.spec.whatwg.org/dev/forms.html#category-reset
// [form-associated]: https://html.spec.whatwg.org/dev/forms.html#form-associated-element
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [text]: https://html.spec.whatwg.org/dev/dom.html#text-content
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [autocomplete]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-autocomplete
// [cols]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-cols
// [dirname]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-dirname
// [disabled]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-disabled
// [form]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [maxlength]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-maxlength
// [minlength]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-minlength
// [name]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-name
// [placeholder]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-placeholder
// [readonly]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-readonly
// [required]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-required
// [rows]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-rows
// [wrap]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-wrap
//
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Textarea(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("textarea", attributes, children...)
}
