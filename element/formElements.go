// Generated file. DO NOT EDIT.

package element

import (
	"github.com/KrischanCS/roundly"
)

// Button creates the button element - Button control [(More)]
//
// It belongs to the following categories:
//   - [flow] [phrasing] [interactive] [listed] [labelable] [submittable] [form-associated] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [phrasing]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals] [command] [commandfor] [disabled] [form] [formaction] [formenctype] [formmethod] [formnovalidate] [formtarget] [name] [popovertarget] [popovertargetaction] [type] [value]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/form-elements.html#the-button-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [interactive]: https://html.spec.whatwg.org/dev/dom.html#interactive-content-2
// [listed]: https://html.spec.whatwg.org/dev/forms.html#category-listed
// [labelable]: https://html.spec.whatwg.org/dev/forms.html#category-label
// [submittable]: https://html.spec.whatwg.org/dev/forms.html#category-submit
// [form-associated]: https://html.spec.whatwg.org/dev/forms.html#form-associated-element
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
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
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Button(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("button", attributes, children...)
}

// Datalist creates the datalist element - Container for options for [combo box control [(More)]
//
// It belongs to the following categories:
//   - [flow] [phrasing]
//
// It can be parent to the following elements/categories of elements:
//   - [phrasing] [option] [script-supporting elements]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/form-elements.html#the-datalist-element
// [combo box control]: https://html.spec.whatwg.org/dev/input.html#attr-input-list
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [option]: https://html.spec.whatwg.org/dev/form-elements.html#the-option-element
// [script-supporting elements]: https://html.spec.whatwg.org/dev/dom.html#script-supporting-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Datalist(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("datalist", attributes, children...)
}

// Fieldset creates the fieldset element - Group of form controls [(More)]
//
// It belongs to the following categories:
//   - [flow] [listed] [form-associated] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [legend] [flow]
//
// If can itself be a child of the following elements/categories of elements:
//   - [flow]
//
// The following attributes can be added to this element:
//   - [globals] [disabled] [form] [name]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/form-elements.html#the-fieldset-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [listed]: https://html.spec.whatwg.org/dev/forms.html#category-listed
// [form-associated]: https://html.spec.whatwg.org/dev/forms.html#form-associated-element
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [legend]: https://html.spec.whatwg.org/dev/form-elements.html#the-legend-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [disabled]: https://html.spec.whatwg.org/dev/form-elements.html#attr-fieldset-disabled
// [form]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [name]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-name
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Fieldset(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("fieldset", attributes, children...)
}

// Legend creates the legend element - Caption for [fieldset [(More)]
//
// It belongs to the following categories:
//   - none
//
// It can be parent to the following elements/categories of elements:
//   - [phrasing] [heading content]
//
// If can itself be a child of the following elements/categories of elements:
//   - [fieldset] [optgroup]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/form-elements.html#the-legend-element
// [fieldset]: https://html.spec.whatwg.org/dev/form-elements.html#the-fieldset-element
// [optgroup]: https://html.spec.whatwg.org/dev/form-elements.html#the-optgroup-element
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [heading content]: https://html.spec.whatwg.org/dev/dom.html#heading-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Legend(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("legend", attributes, children...)
}

// Meter creates the meter element - Gauge [(More)]
//
// It belongs to the following categories:
//   - [flow] [phrasing] [labelable] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [phrasing]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals] [value] [min] [max] [low] [high] [optimum]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/form-elements.html#the-meter-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [labelable]: https://html.spec.whatwg.org/dev/forms.html#category-label
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [value]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-value
// [min]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-min
// [max]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-max
// [low]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-low
// [high]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-high
// [optimum]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-optimum
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Meter(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("meter", attributes, children...)
}

// Optgroup creates the optgroup element - Group of options in a list box [(More)]
//
// It belongs to the following categories:
//   - [select element inner content elements]
//
// It can be parent to the following elements/categories of elements:
//   - [optgroup element inner content elements] [legend]
//
// If can itself be a child of the following elements/categories of elements:
//   - [select] [div]
//
// The following attributes can be added to this element:
//   - [globals] [disabled] [label]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/form-elements.html#the-optgroup-element
// [select element inner content elements]: https://html.spec.whatwg.org/dev/dom.html#select-element-inner-content-elements-2
// [select]: https://html.spec.whatwg.org/dev/form-elements.html#the-select-element
// [div]: https://html.spec.whatwg.org/dev/grouping-content.html#the-div-element
// [optgroup element inner content elements]: https://html.spec.whatwg.org/dev/dom.html#optgroup-element-inner-content-elements-2
// [legend]: https://html.spec.whatwg.org/dev/form-elements.html#the-legend-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [disabled]: https://html.spec.whatwg.org/dev/form-elements.html#attr-optgroup-disabled
// [label]: https://html.spec.whatwg.org/dev/form-elements.html#attr-optgroup-label
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Optgroup(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("optgroup", attributes, children...)
}

// Option creates the option element - Option in a list box or combo box control [(More)]
//
// It belongs to the following categories:
//   - [select element inner content elements] [optgroup element inner content elements]
//
// It can be parent to the following elements/categories of elements:
//   - [text] [option element inner content elements]
//
// If can itself be a child of the following elements/categories of elements:
//   - [select] [datalist] [optgroup] [div]
//
// The following attributes can be added to this element:
//   - [globals] [disabled] [label] [selected] [value]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/form-elements.html#the-option-element
// [select element inner content elements]: https://html.spec.whatwg.org/dev/dom.html#select-element-inner-content-elements-2
// [optgroup element inner content elements]: https://html.spec.whatwg.org/dev/dom.html#optgroup-element-inner-content-elements-2
// [select]: https://html.spec.whatwg.org/dev/form-elements.html#the-select-element
// [datalist]: https://html.spec.whatwg.org/dev/form-elements.html#the-datalist-element
// [optgroup]: https://html.spec.whatwg.org/dev/form-elements.html#the-optgroup-element
// [div]: https://html.spec.whatwg.org/dev/grouping-content.html#the-div-element
// [text]: https://html.spec.whatwg.org/dev/dom.html#text-content
// [option element inner content elements]: https://html.spec.whatwg.org/dev/dom.html#option-element-inner-content-elements-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [disabled]: https://html.spec.whatwg.org/dev/form-elements.html#attr-option-disabled
// [label]: https://html.spec.whatwg.org/dev/form-elements.html#attr-option-label
// [selected]: https://html.spec.whatwg.org/dev/form-elements.html#attr-option-selected
// [value]: https://html.spec.whatwg.org/dev/form-elements.html#attr-option-value
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Option(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("option", attributes, children...)
}

// Output creates the output element - Calculated output value [(More)]
//
// It belongs to the following categories:
//   - [flow] [phrasing] [listed] [labelable] [resettable] [form-associated] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [phrasing]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals] [for] [form] [name]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/form-elements.html#the-output-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [listed]: https://html.spec.whatwg.org/dev/forms.html#category-listed
// [labelable]: https://html.spec.whatwg.org/dev/forms.html#category-label
// [resettable]: https://html.spec.whatwg.org/dev/forms.html#category-reset
// [form-associated]: https://html.spec.whatwg.org/dev/forms.html#form-associated-element
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [for]: https://html.spec.whatwg.org/dev/form-elements.html#attr-output-for
// [form]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [name]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-name
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Output(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("output", attributes, children...)
}

// Progress creates the progress element - Progress bar [(More)]
//
// It belongs to the following categories:
//   - [flow] [phrasing] [labelable] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [phrasing]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals] [value] [max]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/form-elements.html#the-progress-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [labelable]: https://html.spec.whatwg.org/dev/forms.html#category-label
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [value]: https://html.spec.whatwg.org/dev/form-elements.html#attr-progress-value
// [max]: https://html.spec.whatwg.org/dev/form-elements.html#attr-progress-max
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Progress(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("progress", attributes, children...)
}

// Select creates the select element - List box control [(More)]
//
// It belongs to the following categories:
//   - [flow] [phrasing] [interactive] [listed] [labelable] [submittable] [resettable] [form-associated] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [select element inner content elements] [button]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals] [autocomplete] [disabled] [form] [multiple] [name] [required] [size]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/form-elements.html#the-select-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [interactive]: https://html.spec.whatwg.org/dev/dom.html#interactive-content-2
// [listed]: https://html.spec.whatwg.org/dev/forms.html#category-listed
// [labelable]: https://html.spec.whatwg.org/dev/forms.html#category-label
// [submittable]: https://html.spec.whatwg.org/dev/forms.html#category-submit
// [resettable]: https://html.spec.whatwg.org/dev/forms.html#category-reset
// [form-associated]: https://html.spec.whatwg.org/dev/forms.html#form-associated-element
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [select element inner content elements]: https://html.spec.whatwg.org/dev/dom.html#select-element-inner-content-elements-2
// [button]: https://html.spec.whatwg.org/dev/form-elements.html#the-button-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [autocomplete]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-autocomplete
// [disabled]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-disabled
// [form]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [multiple]: https://html.spec.whatwg.org/dev/form-elements.html#attr-select-multiple
// [name]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-name
// [required]: https://html.spec.whatwg.org/dev/form-elements.html#attr-select-required
// [size]: https://html.spec.whatwg.org/dev/form-elements.html#attr-select-size
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Select(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("select", attributes, children...)
}

// Selectedcontent creates the selectedcontent element - Mirrors content from an [option [(More)]
//
// It belongs to the following categories:
//   - none
//
// It can be parent to the following elements/categories of elements:
//   - empty
//
// If can itself be a child of the following elements/categories of elements:
//   - [button]
//
// The following attributes can be added to this element:
//   - [globals]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/form-elements.html#the-selectedcontent-element
// [option]: https://html.spec.whatwg.org/dev/form-elements.html#the-option-element
// [button]: https://html.spec.whatwg.org/dev/form-elements.html#the-button-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Selectedcontent(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("selectedcontent", attributes, children...)
}

// Textarea creates the textarea element - Multiline text controls [(More)]
//
// It belongs to the following categories:
//   - [flow] [phrasing] [interactive] [listed] [labelable] [submittable] [resettable] [form-associated] [palpable]
//
// It can be parent to the following elements/categories of elements:
//   - [text]
//
// If can itself be a child of the following elements/categories of elements:
//   - [phrasing]
//
// The following attributes can be added to this element:
//   - [globals] [autocomplete] [cols] [dirname] [disabled] [form] [maxlength] [minlength] [name] [placeholder] [readonly] [required] [rows] [wrap]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/form-elements.html#the-textarea-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [interactive]: https://html.spec.whatwg.org/dev/dom.html#interactive-content-2
// [listed]: https://html.spec.whatwg.org/dev/forms.html#category-listed
// [labelable]: https://html.spec.whatwg.org/dev/forms.html#category-label
// [submittable]: https://html.spec.whatwg.org/dev/forms.html#category-submit
// [resettable]: https://html.spec.whatwg.org/dev/forms.html#category-reset
// [form-associated]: https://html.spec.whatwg.org/dev/forms.html#form-associated-element
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
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
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Textarea(attributes roundly.Attribute, children ...roundly.Element) roundly.Element {
    return roundly.WriteElement("textarea", attributes, children...)
}
