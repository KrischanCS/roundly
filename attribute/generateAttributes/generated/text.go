package generated

import (
    "github.com/ch-schulz/htmfunc"
    "github.com/ch-schulz/htmfunc/attribute"
)

// Abbr creates the abbr attribute - Alternative label to use for the header cell when referencing the cell in other contexts
//
// It can be applied to the following elements:
//   - [[th]]
//
// Value constraints: [Text]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [th]: https://html.spec.whatwg.org/dev/tables.html#attr-th-abbr
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Abbr(abbr string) htmfunc.AttributeRenderer {
    return attribute.Attribute("abbr", abbr)
}

// Alt creates the alt attribute - Replacement text for use when images are not available
//
// It can be applied to the following elements:
//   - [[area]]
//   - [[img]]
//   - [[input]]
//
// Value constraints: [Text]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [area]: https://html.spec.whatwg.org/dev/image-maps.html#attr-area-alt
// [img]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-alt
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-alt
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Alt(alt string) htmfunc.AttributeRenderer {
    return attribute.Attribute("alt", alt)
}

// Content creates the content attribute - Value of the element
//
// It can be applied to the following elements:
//   - [[meta]]
//
// Value constraints: [Text]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meta]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-content
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Content(content string) htmfunc.AttributeRenderer {
    return attribute.Attribute("content", content)
}

// DirName creates the dirname attribute - Name of form control to use for sending the element's [directionality] in [form submission]
//
// It can be applied to the following elements:
//   - [[input]]
//   - [[textarea]]
//
// Value constraints: [Text]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-dirname
// [textarea]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-dirname
// [directionality]: https://html.spec.whatwg.org/dev/dom.html#the-directionality
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func DirName(dirName string) htmfunc.AttributeRenderer {
    return attribute.Attribute("dirname", dirName)
}

// Id creates the id attribute - The element's [ID]
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: [Text]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#the-id-attribute
// [ID]: https://html.spec.whatwg.org/dev/https://dom.spec.whatwg.org/#concept-id
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Id(id string) htmfunc.AttributeRenderer {
    return attribute.Attribute("id", id)
}

// Integrity creates the integrity attribute - Integrity metadata used in Subresource Integrity checks [[SRI]]
//
// It can be applied to the following elements:
//   - [[link]]
//   - [[script]]
//
// Value constraints: [Text]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-integrity
// [script]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-integrity
// [[SRI]]: https://html.spec.whatwg.org/dev/references.html#refsSRI
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Integrity(integrity string) htmfunc.AttributeRenderer {
    return attribute.Attribute("integrity", integrity)
}

// Label creates the label attribute - User-visible label
//
// It can be applied to the following elements:
//   - [[optgroup]]
//   - [[option]]
//   - [[track]]
//
// Value constraints: [Text]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [optgroup]: https://html.spec.whatwg.org/dev/form-elements.html#attr-optgroup-label
// [option]: https://html.spec.whatwg.org/dev/form-elements.html#attr-option-label
// [track]: https://html.spec.whatwg.org/dev/media.html#attr-track-label
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Label(label string) htmfunc.AttributeRenderer {
    return attribute.Attribute("label", label)
}

// Name creates the name attribute - Name of the element to use for [form submission] and in the [form.elements] API
//
// It can be applied to the following elements:
//   - [[button]]
//   - [[fieldset]]
//   - [[input]]
//   - [[output]]
//   - [[select]]
//   - [[textarea]]
//   - [[form-associated custom elements]]
//
// Value constraints: [Text]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-name
// [fieldset]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-name
// [input]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-name
// [output]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-name
// [select]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-name
// [textarea]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-name
// [form-associated custom elements]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-name
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [form.elements]: https://html.spec.whatwg.org/dev/forms.html#dom-form-elements
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Name(name string) htmfunc.AttributeRenderer {
    return attribute.Attribute("name", name)
}

// Name creates the name attribute - Name of group of mutually-exclusive [details] elements
//
// It can be applied to the following elements:
//   - [[details]]
//
// Value constraints: [Text]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [details]: https://html.spec.whatwg.org/dev/interactive-elements.html#attr-details-name
// [details]: https://html.spec.whatwg.org/dev/interactive-elements.html#the-details-element
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Name(name string) htmfunc.AttributeRenderer {
    return attribute.Attribute("name", name)
}

// Name creates the name attribute - Name of form to use in the [document.forms] API
//
// It can be applied to the following elements:
//   - [[form]]
//
// Value constraints: [Text]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [form]: https://html.spec.whatwg.org/dev/forms.html#attr-form-name
// [document.forms]: https://html.spec.whatwg.org/dev/dom.html#dom-document-forms
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Name(name string) htmfunc.AttributeRenderer {
    return attribute.Attribute("name", name)
}

// Name creates the name attribute - Name of [image map] to [reference] from the [usemap] attribute
//
// It can be applied to the following elements:
//   - [[map]]
//
// Value constraints: [Text]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [map]: https://html.spec.whatwg.org/dev/image-maps.html#attr-map-name
// [image map]: https://html.spec.whatwg.org/dev/image-maps.html#image-map
// [reference]: https://html.spec.whatwg.org/dev/dom.html#referenced
// [usemap]: https://html.spec.whatwg.org/dev/image-maps.html#attr-hyperlink-usemap
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Name(name string) htmfunc.AttributeRenderer {
    return attribute.Attribute("name", name)
}

// Name creates the name attribute - Metadata name
//
// It can be applied to the following elements:
//   - [[meta]]
//
// Value constraints: [Text]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meta]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-name
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Name(name string) htmfunc.AttributeRenderer {
    return attribute.Attribute("name", name)
}

// Name creates the name attribute - Name of shadow tree slot
//
// It can be applied to the following elements:
//   - [[slot]]
//
// Value constraints: [Text]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [slot]: https://html.spec.whatwg.org/dev/scripting.html#attr-slot-name
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Name(name string) htmfunc.AttributeRenderer {
    return attribute.Attribute("name", name)
}

// Nonce creates the nonce attribute - Cryptographic nonce used in Content Security Policy checks [[CSP]]
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: [Text]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#attr-nonce
// [[CSP]]: https://html.spec.whatwg.org/dev/references.html#refsCSP
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Nonce(nonce string) htmfunc.AttributeRenderer {
    return attribute.Attribute("nonce", nonce)
}

// PlaceHolder creates the placeholder attribute - User-visible label to be placed within the form control
//
// It can be applied to the following elements:
//   - [[input]]
//   - [[textarea]]
//
// Value constraints: [Text]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-placeholder
// [textarea]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-placeholder
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func PlaceHolder(placeHolder string) htmfunc.AttributeRenderer {
    return attribute.Attribute("placeholder", placeHolder)
}

// Slot creates the slot attribute - The element's desired slot
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: [Text]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#the-id-attribute
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Slot(slot string) htmfunc.AttributeRenderer {
    return attribute.Attribute("slot", slot)
}

// Title creates the title attribute - Advisory information for the element
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: [Text]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#attr-title
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Title(title string) htmfunc.AttributeRenderer {
    return attribute.Attribute("title", title)
}

// Title creates the title attribute - Full term or expansion of abbreviation
//
// It can be applied to the following elements:
//   - [[abbr]]
//   - [[dfn]]
//
// Value constraints: [Text]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [abbr]: https://html.spec.whatwg.org/dev/text-level-semantics.html#attr-abbr-title
// [dfn]: https://html.spec.whatwg.org/dev/text-level-semantics.html#attr-dfn-title
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Title(title string) htmfunc.AttributeRenderer {
    return attribute.Attribute("title", title)
}

// Title creates the title attribute - Description of pattern (when used with [pattern] attribute)
//
// It can be applied to the following elements:
//   - [[input]]
//
// Value constraints: [Text]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-title
// [pattern]: https://html.spec.whatwg.org/dev/input.html#attr-input-pattern
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Title(title string) htmfunc.AttributeRenderer {
    return attribute.Attribute("title", title)
}

// Title creates the title attribute - Title of the link
//
// It can be applied to the following elements:
//   - [[link]]
//
// Value constraints: [Text]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-title
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Title(title string) htmfunc.AttributeRenderer {
    return attribute.Attribute("title", title)
}

// Title creates the title attribute - [CSS style sheet set name]
//
// It can be applied to the following elements:
//   - [[link]]
//   - [[style]]
//
// Value constraints: [Text]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-title
// [style]: https://html.spec.whatwg.org/dev/semantics.html#attr-style-title
// [CSS style sheet set name]: https://html.spec.whatwg.org/dev/https://drafts.csswg.org/cssom/#css-style-sheet-set-name
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Title(title string) htmfunc.AttributeRenderer {
    return attribute.Attribute("title", title)
}

// Value creates the value attribute - Value to be used for [form submission]
//
// It can be applied to the following elements:
//   - [[button]]
//   - [[option]]
//
// Value constraints: [Text]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-value
// [option]: https://html.spec.whatwg.org/dev/form-elements.html#attr-option-value
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Value(value string) htmfunc.AttributeRenderer {
    return attribute.Attribute("value", value)
}

// Value creates the value attribute - Machine-readable value
//
// It can be applied to the following elements:
//   - [[data]]
//
// Value constraints: [Text]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [data]: https://html.spec.whatwg.org/dev/text-level-semantics.html#attr-data-value
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Value(value string) htmfunc.AttributeRenderer {
    return attribute.Attribute("value", value)
}
