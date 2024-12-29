package generated

import (
    "github.com/ch-schulz/htmfunc"
    "github.com/ch-schulz/htmfunc/attribute"
)

// Accept creates the accept attribute - Hint for expected file type in [file upload controls]
//
// It can be applied to the following elements:
//   - [[input]]
//
// Value constraints: [Set of comma-separated tokens]* consisting of [valid MIME type strings with no parameters] or audio/*, video/*, or image/*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-accept
// [file upload controls]: https://html.spec.whatwg.org/dev/input.html#file-upload-state-(type=file)
// [Set of comma-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#set-of-comma-separated-tokens
// [valid MIME type strings with no parameters]: https://html.spec.whatwg.org/dev/https://mimesniff.spec.whatwg.org/#valid-mime-type-with-no-parameters
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Accept(accept string) htmfunc.AttributeRenderer {
    return attribute.Attribute("accept", accept)
}

// AcceptCharset creates the accept-charset attribute - Character encodings to use for [form submission]
//
// It can be applied to the following elements:
//   - [[form]]
//
// Value constraints: [ASCII case-insensitive] match for "UTF-8"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [form]: https://html.spec.whatwg.org/dev/forms.html#attr-form-accept-charset
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [ASCII case-insensitive]: https://html.spec.whatwg.org/dev/https://infra.spec.whatwg.org/#ascii-case-insensitive
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func AcceptCharset(acceptCharset string) htmfunc.AttributeRenderer {
    return attribute.Attribute("accept-charset", acceptCharset)
}

// AccessKey creates the accesskey attribute - Keyboard shortcut to activate or focus element
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: [Ordered set of unique space-separated tokens], none of which are [identical to] another, each consisting of one code point in length
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/interaction.html#the-accesskey-attribute
// [Ordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#ordered-set-of-unique-space-separated-tokens
// [identical to]: https://html.spec.whatwg.org/dev/https://infra.spec.whatwg.org/#string-is
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func AccessKey(accessKey string) htmfunc.AttributeRenderer {
    return attribute.Attribute("accesskey", accessKey)
}

// Action creates the action attribute - [URL] to use for [form submission]
//
// It can be applied to the following elements:
//   - [[form]]
//
// Value constraints: [Valid non-empty URL potentially surrounded by spaces]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [form]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-action
// [URL]: https://html.spec.whatwg.org/dev/https://url.spec.whatwg.org/#concept-url
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [Valid non-empty URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Action(action string) htmfunc.AttributeRenderer {
    return attribute.Attribute("action", action)
}

// Allow creates the allow attribute - [Permissions policy] to be applied to the [iframe]'s contents
//
// It can be applied to the following elements:
//   - [[iframe]]
//
// Value constraints: [Serialized permissions policy]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-allow
// [Permissions policy]: https://html.spec.whatwg.org/dev/https://w3c.github.io/webappsec-feature-policy/#permissions-policy
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#the-iframe-element
// [Serialized permissions policy]: https://html.spec.whatwg.org/dev/https://w3c.github.io/webappsec-feature-policy/#serialized-permissions-policy
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Allow(allow string) htmfunc.AttributeRenderer {
    return attribute.Attribute("allow", allow)
}

// As creates the as attribute - [Potential destination] for a preload request (for [rel]="[preload]" and [rel]="[modulepreload]")
//
// It can be applied to the following elements:
//   - [[link]]
//
// Value constraints: [Potential destination], for [rel]="[preload]"; [script-like destination], for [rel]="[modulepreload]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-as
// [Potential destination]: https://html.spec.whatwg.org/dev/https://fetch.spec.whatwg.org/#concept-potential-destination
// [rel]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [preload]: https://html.spec.whatwg.org/dev/links.html#link-type-preload
// [rel]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [modulepreload]: https://html.spec.whatwg.org/dev/links.html#link-type-modulepreload
// [Potential destination]: https://html.spec.whatwg.org/dev/https://fetch.spec.whatwg.org/#concept-potential-destination
// [rel]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [preload]: https://html.spec.whatwg.org/dev/links.html#link-type-preload
// [script-like destination]: https://html.spec.whatwg.org/dev/https://fetch.spec.whatwg.org/#request-destination-script-like
// [rel]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [modulepreload]: https://html.spec.whatwg.org/dev/links.html#link-type-modulepreload
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func As(as string) htmfunc.AttributeRenderer {
    return attribute.Attribute("as", as)
}

// AutoComplete creates the autocomplete attribute - Hint for form autofill feature
//
// It can be applied to the following elements:
//   - [[input]]
//   - [[select]]
//   - [[textarea]]
//
// Value constraints: [Autofill field] name and related tokens*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-autocomplete
// [select]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-autocomplete
// [textarea]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-autocomplete
// [Autofill field]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#autofill-field
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func AutoComplete(autoComplete string) htmfunc.AttributeRenderer {
    return attribute.Attribute("autocomplete", autoComplete)
}

// Blocking creates the blocking attribute - Whether the element is [potentially render-blocking]
//
// It can be applied to the following elements:
//   - [[link]]
//   - [[script]]
//   - [[style]]
//
// Value constraints: [Unordered set of unique space-separated tokens]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-blocking
// [script]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-blocking
// [style]: https://html.spec.whatwg.org/dev/semantics.html#attr-style-blocking
// [potentially render-blocking]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#potentially-render-blocking
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Blocking(blocking string) htmfunc.AttributeRenderer {
    return attribute.Attribute("blocking", blocking)
}

// Cite creates the cite attribute - Link to the source of the quotation or more information about the edit
//
// It can be applied to the following elements:
//   - [[blockquote]]
//   - [[del]]
//   - [[ins]]
//   - [[q]]
//
// Value constraints: [Valid URL potentially surrounded by spaces]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [blockquote]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-blockquote-cite
// [del]: https://html.spec.whatwg.org/dev/edits.html#attr-mod-cite
// [ins]: https://html.spec.whatwg.org/dev/edits.html#attr-mod-cite
// [q]: https://html.spec.whatwg.org/dev/text-level-semantics.html#attr-q-cite
// [Valid URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Cite(cite string) htmfunc.AttributeRenderer {
    return attribute.Attribute("cite", cite)
}

// Class creates the class attribute - Classes to which the element belongs
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: [Set of space-separated tokens]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#classes
// [Set of space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#set-of-space-separated-tokens
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Class(class string) htmfunc.AttributeRenderer {
    return attribute.Attribute("class", class)
}

// Color creates the color attribute - Color to use when customizing a site's icon (for [rel]="mask-icon")
//
// It can be applied to the following elements:
//   - [[link]]
//
// Value constraints: CSS [<color>]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-color
// [rel]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [<color>]: https://html.spec.whatwg.org/dev/https://drafts.csswg.org/css-color/#typedef-color
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Color(color string) htmfunc.AttributeRenderer {
    return attribute.Attribute("color", color)
}

// Cols creates the cols attribute - Maximum number of characters per line
//
// It can be applied to the following elements:
//   - [[textarea]]
//
// Value constraints: [Valid non-negative integer] greater than zero
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [textarea]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-cols
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Cols(cols string) htmfunc.AttributeRenderer {
    return attribute.Attribute("cols", cols)
}

// ColSpan creates the colspan attribute - Number of columns that the cell is to span
//
// It can be applied to the following elements:
//   - [[td]]
//   - [[th]]
//
// Value constraints: [Valid non-negative integer] greater than zero
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [td]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-colspan
// [th]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-colspan
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ColSpan(colSpan string) htmfunc.AttributeRenderer {
    return attribute.Attribute("colspan", colSpan)
}

// Coords creates the coords attribute - Coordinates for the shape to be created in an [image map]
//
// It can be applied to the following elements:
//   - [[area]]
//
// Value constraints: [Valid list of floating-point numbers]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [area]: https://html.spec.whatwg.org/dev/image-maps.html#attr-area-coords
// [image map]: https://html.spec.whatwg.org/dev/image-maps.html#image-map
// [Valid list of floating-point numbers]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-list-of-floating-point-numbers
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Coords(coords string) htmfunc.AttributeRenderer {
    return attribute.Attribute("coords", coords)
}

// Data creates the data attribute - Address of the resource
//
// It can be applied to the following elements:
//   - [[object]]
//
// Value constraints: [Valid non-empty URL potentially surrounded by spaces]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [object]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-object-data
// [Valid non-empty URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Data(data string) htmfunc.AttributeRenderer {
    return attribute.Attribute("data", data)
}

// DateTime creates the datetime attribute - Date and (optionally) time of the change
//
// It can be applied to the following elements:
//   - [[del]]
//   - [[ins]]
//
// Value constraints: [Valid date string with optional time]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [del]: https://html.spec.whatwg.org/dev/edits.html#attr-mod-datetime
// [ins]: https://html.spec.whatwg.org/dev/edits.html#attr-mod-datetime
// [Valid date string with optional time]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-date-string-with-optional-time
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func DateTime(dateTime string) htmfunc.AttributeRenderer {
    return attribute.Attribute("datetime", dateTime)
}

// DateTime creates the datetime attribute - Machine-readable value
//
// It can be applied to the following elements:
//   - [[time]]
//
// Value constraints: [Valid month string],[valid date string],[valid yearless date string],[valid time string],[valid local date and time string],[valid time-zone offset string],[valid global date and time string],[valid week string],[valid non-negative integer], or[valid duration string]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [time]: https://html.spec.whatwg.org/dev/text-level-semantics.html#attr-time-datetime
// [Valid month string]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-month-string
// [valid date string]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-date-string
// [valid yearless date string]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-yearless-date-string
// [valid time string]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-time-string
// [valid local date and time string]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-local-date-and-time-string
// [valid time-zone offset string]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-time-zone-offset-string
// [valid global date and time string]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-global-date-and-time-string
// [valid week string]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-week-string
// [valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [valid duration string]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-duration-string
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func DateTime(dateTime string) htmfunc.AttributeRenderer {
    return attribute.Attribute("datetime", dateTime)
}

// Download creates the download attribute - Whether to download the resource instead of navigating to it, and its filename if so
//
// It can be applied to the following elements:
//   - [[a]]
//   - [[area]]
//
// Value constraints: Text
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-download
// [area]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-download
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Download(download string) htmfunc.AttributeRenderer {
    return attribute.Attribute("download", download)
}

// For creates the for attribute - Associate the label with form control
//
// It can be applied to the following elements:
//   - [[label]]
//
// Value constraints: [ID]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [label]: https://html.spec.whatwg.org/dev/forms.html#attr-label-for
// [ID]: https://html.spec.whatwg.org/dev/https://dom.spec.whatwg.org/#concept-id
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func For(forV string) htmfunc.AttributeRenderer {
    return attribute.Attribute("for", forV)
}

// For creates the for attribute - Specifies controls from which the output was calculated
//
// It can be applied to the following elements:
//   - [[output]]
//
// Value constraints: [Unordered set of unique space-separated tokens] consisting of IDs*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [output]: https://html.spec.whatwg.org/dev/form-elements.html#attr-output-for
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func For(forV string) htmfunc.AttributeRenderer {
    return attribute.Attribute("for", forV)
}

// Form creates the form attribute - Associates the element with a [form] element
//
// It can be applied to the following elements:
//   - [[button]]
//   - [[fieldset]]
//   - [[input]]
//   - [[object]]
//   - [[output]]
//   - [[select]]
//   - [[textarea]]
//   - [[form-associated custom elements]]
//
// Value constraints: [ID]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [fieldset]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [input]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [object]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [output]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [select]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [textarea]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [form-associated custom elements]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fae-form
// [form]: https://html.spec.whatwg.org/dev/forms.html#the-form-element
// [ID]: https://html.spec.whatwg.org/dev/https://dom.spec.whatwg.org/#concept-id
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Form(form string) htmfunc.AttributeRenderer {
    return attribute.Attribute("form", form)
}

// FormAction creates the formaction attribute - [URL] to use for [form submission]
//
// It can be applied to the following elements:
//   - [[button]]
//   - [[input]]
//
// Value constraints: [Valid non-empty URL potentially surrounded by spaces]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formaction
// [input]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formaction
// [URL]: https://html.spec.whatwg.org/dev/https://url.spec.whatwg.org/#concept-url
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [Valid non-empty URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func FormAction(formAction string) htmfunc.AttributeRenderer {
    return attribute.Attribute("formaction", formAction)
}

// FormTarget creates the formtarget attribute - [Navigable] for [form submission]
//
// It can be applied to the following elements:
//   - [[button]]
//   - [[input]]
//
// Value constraints: [Valid navigable target name or keyword]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formtarget
// [input]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formtarget
// [Navigable]: https://html.spec.whatwg.org/dev/document-sequences.html#navigable
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [Valid navigable target name or keyword]: https://html.spec.whatwg.org/dev/document-sequences.html#valid-navigable-target-name-or-keyword
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func FormTarget(formTarget string) htmfunc.AttributeRenderer {
    return attribute.Attribute("formtarget", formTarget)
}

// Headers creates the headers attribute - The header cells for this cell
//
// It can be applied to the following elements:
//   - [[td]]
//   - [[th]]
//
// Value constraints: [Unordered set of unique space-separated tokens] consisting of IDs*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [td]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-headers
// [th]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-headers
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Headers(headers string) htmfunc.AttributeRenderer {
    return attribute.Attribute("headers", headers)
}

// Height creates the height attribute - Vertical dimension
//
// It can be applied to the following elements:
//   - [[canvas]]
//   - [[embed]]
//   - [[iframe]]
//   - [[img]]
//   - [[input]]
//   - [[object]]
//   - [[source] (in [picture])]
//   - [[video]]
//
// Value constraints: [Valid non-negative integer]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [canvas]: https://html.spec.whatwg.org/dev/canvas.html#attr-canvas-height
// [embed]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [iframe]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [img]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [input]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [object]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [source]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [picture]: https://html.spec.whatwg.org/dev/embedded-content.html#the-picture-element
// [video]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Height(height string) htmfunc.AttributeRenderer {
    return attribute.Attribute("height", height)
}

// Hidden creates the hidden attribute - Whether the element is relevant
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: "[until-found]";"[hidden]";the empty string
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/interaction.html#attr-hidden
// [until-found]: https://html.spec.whatwg.org/dev/interaction.html#attr-hidden-until-found
// [hidden]: https://html.spec.whatwg.org/dev/interaction.html#attr-hidden
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Hidden(hidden string) htmfunc.AttributeRenderer {
    return attribute.Attribute("hidden", hidden)
}

// High creates the high attribute - Low limit of high range
//
// It can be applied to the following elements:
//   - [[meter]]
//
// Value constraints: [Valid floating-point number]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-high
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func High(high string) htmfunc.AttributeRenderer {
    return attribute.Attribute("high", high)
}

// HRef creates the href attribute - Address of the [hyperlink]
//
// It can be applied to the following elements:
//   - [[a]]
//   - [[area]]
//
// Value constraints: [Valid URL potentially surrounded by spaces]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-href
// [area]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-href
// [hyperlink]: https://html.spec.whatwg.org/dev/links.html#hyperlink
// [Valid URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func HRef(hRef string) htmfunc.AttributeRenderer {
    return attribute.Attribute("href", hRef)
}

// HRef creates the href attribute - Address of the [hyperlink]
//
// It can be applied to the following elements:
//   - [[link]]
//
// Value constraints: [Valid non-empty URL potentially surrounded by spaces]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-href
// [hyperlink]: https://html.spec.whatwg.org/dev/links.html#hyperlink
// [Valid non-empty URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func HRef(hRef string) htmfunc.AttributeRenderer {
    return attribute.Attribute("href", hRef)
}

// HRef creates the href attribute - [Document base URL]
//
// It can be applied to the following elements:
//   - [[base]]
//
// Value constraints: [Valid URL potentially surrounded by spaces]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [base]: https://html.spec.whatwg.org/dev/semantics.html#attr-base-href
// [Document base URL]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#document-base-url
// [Valid URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func HRef(hRef string) htmfunc.AttributeRenderer {
    return attribute.Attribute("href", hRef)
}

// HRefLang creates the hreflang attribute - Language of the linked resource
//
// It can be applied to the following elements:
//   - [[a]]
//   - [[link]]
//
// Value constraints: Valid BCP 47 language tag
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-hreflang
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-hreflang
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func HRefLang(hRefLang string) htmfunc.AttributeRenderer {
    return attribute.Attribute("hreflang", hRefLang)
}

// ImageSizes creates the imagesizes attribute - Image sizes for different page layouts (for [rel]="[preload]")
//
// It can be applied to the following elements:
//   - [[link]]
//
// Value constraints: [Valid source size list]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-imagesizes
// [rel]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [preload]: https://html.spec.whatwg.org/dev/links.html#link-type-preload
// [Valid source size list]: https://html.spec.whatwg.org/dev/images.html#valid-source-size-list
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ImageSizes(imageSizes string) htmfunc.AttributeRenderer {
    return attribute.Attribute("imagesizes", imageSizes)
}

// ImageSrcSet creates the imagesrcset attribute - Images to use in different situations, e.g., high-resolution displays, small monitors, etc. (for [rel]="[preload]")
//
// It can be applied to the following elements:
//   - [[link]]
//
// Value constraints: Comma-separated list of [image candidate strings]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-imagesrcset
// [rel]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [preload]: https://html.spec.whatwg.org/dev/links.html#link-type-preload
// [image candidate strings]: https://html.spec.whatwg.org/dev/images.html#image-candidate-string
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ImageSrcSet(imageSrcSet string) htmfunc.AttributeRenderer {
    return attribute.Attribute("imagesrcset", imageSrcSet)
}

// Is creates the is attribute - Creates a [customized built-in element]
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: [Valid custom element name] of a defined [customized built-in element]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/custom-elements.html#attr-is
// [customized built-in element]: https://html.spec.whatwg.org/dev/custom-elements.html#customized-built-in-element
// [Valid custom element name]: https://html.spec.whatwg.org/dev/custom-elements.html#valid-custom-element-name
// [customized built-in element]: https://html.spec.whatwg.org/dev/custom-elements.html#customized-built-in-element
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Is(is string) htmfunc.AttributeRenderer {
    return attribute.Attribute("is", is)
}

// ItemId creates the itemid attribute - [Global identifier] for a microdata item
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: [Valid URL potentially surrounded by spaces]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/microdata.html#attr-itemid
// [Global identifier]: https://html.spec.whatwg.org/dev/microdata.html#global-identifier
// [Valid URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ItemId(itemId string) htmfunc.AttributeRenderer {
    return attribute.Attribute("itemid", itemId)
}

// ItemProp creates the itemprop attribute - [Property names] of a microdata item
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: [Unordered set of unique space-separated tokens] consisting of [valid absolute URLs], [defined property names], or text*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/microdata.html#names:-the-itemprop-attribute
// [Property names]: https://html.spec.whatwg.org/dev/microdata.html#property-names
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [valid absolute URLs]: https://html.spec.whatwg.org/dev/https://url.spec.whatwg.org/#syntax-url-absolute
// [defined property names]: https://html.spec.whatwg.org/dev/microdata.html#defined-property-name
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ItemProp(itemProp string) htmfunc.AttributeRenderer {
    return attribute.Attribute("itemprop", itemProp)
}

// ItemRef creates the itemref attribute - [Referenced] elements
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: [Unordered set of unique space-separated tokens] consisting of IDs*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/microdata.html#attr-itemref
// [Referenced]: https://html.spec.whatwg.org/dev/dom.html#referenced
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ItemRef(itemRef string) htmfunc.AttributeRenderer {
    return attribute.Attribute("itemref", itemRef)
}

// ItemType creates the itemtype attribute - [Item types] of a microdata item
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: [Unordered set of unique space-separated tokens] consisting of [valid absolute URLs]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/microdata.html#attr-itemtype
// [Item types]: https://html.spec.whatwg.org/dev/microdata.html#item-types
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [valid absolute URLs]: https://html.spec.whatwg.org/dev/https://url.spec.whatwg.org/#syntax-url-absolute
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ItemType(itemType string) htmfunc.AttributeRenderer {
    return attribute.Attribute("itemtype", itemType)
}

// Lang creates the lang attribute - Language of the element
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: Valid BCP 47 language tag or the empty string
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#attr-lang
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Lang(lang string) htmfunc.AttributeRenderer {
    return attribute.Attribute("lang", lang)
}

// List creates the list attribute - List of autocomplete options
//
// It can be applied to the following elements:
//   - [[input]]
//
// Value constraints: [ID]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-list
// [ID]: https://html.spec.whatwg.org/dev/https://dom.spec.whatwg.org/#concept-id
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func List(list string) htmfunc.AttributeRenderer {
    return attribute.Attribute("list", list)
}

// Low creates the low attribute - High limit of low range
//
// It can be applied to the following elements:
//   - [[meter]]
//
// Value constraints: [Valid floating-point number]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-low
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Low(low string) htmfunc.AttributeRenderer {
    return attribute.Attribute("low", low)
}

// Max creates the max attribute - Maximum value
//
// It can be applied to the following elements:
//   - [[input]]
//
// Value constraints: Varies*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-max
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Max(max string) htmfunc.AttributeRenderer {
    return attribute.Attribute("max", max)
}

// Max creates the max attribute - Upper bound of range
//
// It can be applied to the following elements:
//   - [[meter]]
//   - [[progress]]
//
// Value constraints: [Valid floating-point number]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-max
// [progress]: https://html.spec.whatwg.org/dev/form-elements.html#attr-progress-max
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Max(max string) htmfunc.AttributeRenderer {
    return attribute.Attribute("max", max)
}

// MaxLength creates the maxlength attribute - Maximum [length] of value
//
// It can be applied to the following elements:
//   - [[input]]
//   - [[textarea]]
//
// Value constraints: [Valid non-negative integer]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-maxlength
// [textarea]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-maxlength
// [length]: https://html.spec.whatwg.org/dev/https://infra.spec.whatwg.org/#string-length
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func MaxLength(maxLength string) htmfunc.AttributeRenderer {
    return attribute.Attribute("maxlength", maxLength)
}

// Media creates the media attribute - Applicable media
//
// It can be applied to the following elements:
//   - [[link]]
//   - [[meta]]
//   - [[source]]
//   - [[style]]
//
// Value constraints: [Valid media query list]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-media
// [meta]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-media
// [source]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-source-media
// [style]: https://html.spec.whatwg.org/dev/semantics.html#attr-style-media
// [Valid media query list]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-media-query-list
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Media(media string) htmfunc.AttributeRenderer {
    return attribute.Attribute("media", media)
}

// Min creates the min attribute - Minimum value
//
// It can be applied to the following elements:
//   - [[input]]
//
// Value constraints: Varies*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-min
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Min(min string) htmfunc.AttributeRenderer {
    return attribute.Attribute("min", min)
}

// Min creates the min attribute - Lower bound of range
//
// It can be applied to the following elements:
//   - [[meter]]
//
// Value constraints: [Valid floating-point number]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-min
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Min(min string) htmfunc.AttributeRenderer {
    return attribute.Attribute("min", min)
}

// MinLength creates the minlength attribute - Minimum [length] of value
//
// It can be applied to the following elements:
//   - [[input]]
//   - [[textarea]]
//
// Value constraints: [Valid non-negative integer]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-minlength
// [textarea]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-minlength
// [length]: https://html.spec.whatwg.org/dev/https://infra.spec.whatwg.org/#string-length
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func MinLength(minLength string) htmfunc.AttributeRenderer {
    return attribute.Attribute("minlength", minLength)
}

// Name creates the name attribute - Name of [content navigable]
//
// It can be applied to the following elements:
//   - [[iframe]]
//   - [[object]]
//
// Value constraints: [Valid navigable target name or keyword]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-name
// [object]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-object-name
// [content navigable]: https://html.spec.whatwg.org/dev/document-sequences.html#content-navigable
// [Valid navigable target name or keyword]: https://html.spec.whatwg.org/dev/document-sequences.html#valid-navigable-target-name-or-keyword
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Name(name string) htmfunc.AttributeRenderer {
    return attribute.Attribute("name", name)
}

// Optimum creates the optimum attribute - Optimum value in gauge
//
// It can be applied to the following elements:
//   - [[meter]]
//
// Value constraints: [Valid floating-point number]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-optimum
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Optimum(optimum string) htmfunc.AttributeRenderer {
    return attribute.Attribute("optimum", optimum)
}

// Pattern creates the pattern attribute - Pattern to be matched by the form control's value
//
// It can be applied to the following elements:
//   - [[input]]
//
// Value constraints: Regular expression matching the JavaScript [Pattern] production
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-pattern
// [Pattern]: https://html.spec.whatwg.org/dev/https://tc39.es/ecma262/#prod-Pattern
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Pattern(pattern string) htmfunc.AttributeRenderer {
    return attribute.Attribute("pattern", pattern)
}

// Ping creates the ping attribute - [URLs] to ping
//
// It can be applied to the following elements:
//   - [[a]]
//   - [[area]]
//
// Value constraints: [Set of space-separated tokens] consisting of [valid non-empty URLs]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#ping
// [area]: https://html.spec.whatwg.org/dev/links.html#ping
// [URLs]: https://html.spec.whatwg.org/dev/https://url.spec.whatwg.org/#concept-url
// [Set of space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#set-of-space-separated-tokens
// [valid non-empty URLs]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Ping(ping string) htmfunc.AttributeRenderer {
    return attribute.Attribute("ping", ping)
}

// PopOver creates the popover attribute - Makes the element a [popover] element
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: "[auto]";"[manual]";
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/popover.html#attr-popover
// [popover]: https://html.spec.whatwg.org/dev/popover.html#attr-popover
// [auto]: https://html.spec.whatwg.org/dev/popover.html#attr-popover-auto
// [manual]: https://html.spec.whatwg.org/dev/popover.html#attr-popover-manual
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func PopOver(popOver string) htmfunc.AttributeRenderer {
    return attribute.Attribute("popover", popOver)
}

// PopOverTarget creates the popovertarget attribute - Targets a popover element to toggle, show, or hide
//
// It can be applied to the following elements:
//   - [[button]]
//   - [[input]]
//
// Value constraints: [ID]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/popover.html#attr-popovertarget
// [input]: https://html.spec.whatwg.org/dev/popover.html#attr-popovertarget
// [ID]: https://html.spec.whatwg.org/dev/https://dom.spec.whatwg.org/#concept-id
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func PopOverTarget(popOverTarget string) htmfunc.AttributeRenderer {
    return attribute.Attribute("popovertarget", popOverTarget)
}

// Poster creates the poster attribute - Poster frame to show prior to video playback
//
// It can be applied to the following elements:
//   - [[video]]
//
// Value constraints: [Valid non-empty URL potentially surrounded by spaces]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [video]: https://html.spec.whatwg.org/dev/media.html#attr-video-poster
// [Valid non-empty URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Poster(poster string) htmfunc.AttributeRenderer {
    return attribute.Attribute("poster", poster)
}

// ReferrerPolicy creates the referrerpolicy attribute - [Referrer policy] for [fetches] initiated by the element
//
// It can be applied to the following elements:
//   - [[a]]
//   - [[area]]
//   - [[iframe]]
//   - [[img]]
//   - [[link]]
//   - [[script]]
//
// Value constraints: [Referrer policy]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-referrerpolicy
// [area]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-referrerpolicy
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-referrerpolicy
// [img]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-referrerpolicy
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-referrerpolicy
// [script]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-referrerpolicy
// [Referrer policy]: https://html.spec.whatwg.org/dev/https://w3c.github.io/webappsec-referrer-policy/#referrer-policy
// [fetches]: https://html.spec.whatwg.org/dev/https://fetch.spec.whatwg.org/#concept-fetch
// [Referrer policy]: https://html.spec.whatwg.org/dev/https://w3c.github.io/webappsec-referrer-policy/#referrer-policy
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ReferrerPolicy(referrerPolicy string) htmfunc.AttributeRenderer {
    return attribute.Attribute("referrerpolicy", referrerPolicy)
}

// Rel creates the rel attribute - Relationship between the location in the document containing the [hyperlink] and the destination resource
//
// It can be applied to the following elements:
//   - [[a]]
//   - [[area]]
//
// Value constraints: [Unordered set of unique space-separated tokens]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-rel
// [area]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-rel
// [hyperlink]: https://html.spec.whatwg.org/dev/links.html#hyperlink
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Rel(rel string) htmfunc.AttributeRenderer {
    return attribute.Attribute("rel", rel)
}

// Rel creates the rel attribute - Relationship between the document containing the [hyperlink] and the destination resource
//
// It can be applied to the following elements:
//   - [[link]]
//
// Value constraints: [Unordered set of unique space-separated tokens]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [hyperlink]: https://html.spec.whatwg.org/dev/links.html#hyperlink
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Rel(rel string) htmfunc.AttributeRenderer {
    return attribute.Attribute("rel", rel)
}

// Rows creates the rows attribute - Number of lines to show
//
// It can be applied to the following elements:
//   - [[textarea]]
//
// Value constraints: [Valid non-negative integer] greater than zero
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [textarea]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-rows
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Rows(rows string) htmfunc.AttributeRenderer {
    return attribute.Attribute("rows", rows)
}

// RowSpan creates the rowspan attribute - Number of rows that the cell is to span
//
// It can be applied to the following elements:
//   - [[td]]
//   - [[th]]
//
// Value constraints: [Valid non-negative integer]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [td]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-rowspan
// [th]: https://html.spec.whatwg.org/dev/tables.html#attr-tdth-rowspan
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func RowSpan(rowSpan string) htmfunc.AttributeRenderer {
    return attribute.Attribute("rowspan", rowSpan)
}

// SandBox creates the sandbox attribute - Security rules for nested content
//
// It can be applied to the following elements:
//   - [[iframe]]
//
// Value constraints: [Unordered set of unique space-separated tokens], [ASCII case-insensitive], consisting of"allow-downloads""allow-forms""allow-modals""allow-orientation-lock""allow-pointer-lock""allow-popups""allow-popups-to-escape-sandbox""allow-presentation""allow-same-origin""allow-scripts""allow-top-navigation""allow-top-navigation-by-user-activation""allow-top-navigation-to-custom-protocols"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-sandbox
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [ASCII case-insensitive]: https://html.spec.whatwg.org/dev/https://infra.spec.whatwg.org/#ascii-case-insensitive
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func SandBox(sandBox string) htmfunc.AttributeRenderer {
    return attribute.Attribute("sandbox", sandBox)
}

// Size creates the size attribute - Size of the control
//
// It can be applied to the following elements:
//   - [[input]]
//   - [[select]]
//
// Value constraints: [Valid non-negative integer] greater than zero
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-size
// [select]: https://html.spec.whatwg.org/dev/form-elements.html#attr-select-size
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Size(size string) htmfunc.AttributeRenderer {
    return attribute.Attribute("size", size)
}

// Sizes creates the sizes attribute - Sizes of the icons (for [rel]="[icon]")
//
// It can be applied to the following elements:
//   - [[link]]
//
// Value constraints: [Unordered set of unique space-separated tokens], [ASCII case-insensitive], consisting of sizes*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-sizes
// [rel]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [icon]: https://html.spec.whatwg.org/dev/links.html#rel-icon
// [Unordered set of unique space-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens
// [ASCII case-insensitive]: https://html.spec.whatwg.org/dev/https://infra.spec.whatwg.org/#ascii-case-insensitive
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Sizes(sizes string) htmfunc.AttributeRenderer {
    return attribute.Attribute("sizes", sizes)
}

// Sizes creates the sizes attribute - Image sizes for different page layouts
//
// It can be applied to the following elements:
//   - [[img]]
//   - [[source]]
//
// Value constraints: [Valid source size list]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [img]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-sizes
// [source]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-source-sizes
// [Valid source size list]: https://html.spec.whatwg.org/dev/images.html#valid-source-size-list
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Sizes(sizes string) htmfunc.AttributeRenderer {
    return attribute.Attribute("sizes", sizes)
}

// Span creates the span attribute - Number of columns spanned by the element
//
// It can be applied to the following elements:
//   - [[col]]
//   - [[colgroup]]
//
// Value constraints: [Valid non-negative integer] greater than zero
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [col]: https://html.spec.whatwg.org/dev/tables.html#attr-col-span
// [colgroup]: https://html.spec.whatwg.org/dev/tables.html#attr-colgroup-span
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Span(span string) htmfunc.AttributeRenderer {
    return attribute.Attribute("span", span)
}

// SpellCheck creates the spellcheck attribute - Whether the element is to have its spelling and grammar checked
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: "[true]";"[false]";the empty string
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/interaction.html#attr-spellcheck
// [true]: https://html.spec.whatwg.org/dev/interaction.html#attr-spellcheck-true
// [false]: https://html.spec.whatwg.org/dev/interaction.html#attr-spellcheck-false
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func SpellCheck(spellCheck string) htmfunc.AttributeRenderer {
    return attribute.Attribute("spellcheck", spellCheck)
}

// Src creates the src attribute - Address of the resource
//
// It can be applied to the following elements:
//   - [[audio]]
//   - [[embed]]
//   - [[iframe]]
//   - [[img]]
//   - [[input]]
//   - [[script]]
//   - [[source] (in [video] or [audio])]
//   - [[track]]
//   - [[video]]
//
// Value constraints: [Valid non-empty URL potentially surrounded by spaces]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [audio]: https://html.spec.whatwg.org/dev/media.html#attr-media-src
// [embed]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-embed-src
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-src
// [img]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-src
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-src
// [script]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-src
// [source]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-source-src
// [video]: https://html.spec.whatwg.org/dev/media.html#the-video-element
// [audio]: https://html.spec.whatwg.org/dev/media.html#the-audio-element
// [track]: https://html.spec.whatwg.org/dev/media.html#attr-track-src
// [video]: https://html.spec.whatwg.org/dev/media.html#attr-media-src
// [Valid non-empty URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Src(src string) htmfunc.AttributeRenderer {
    return attribute.Attribute("src", src)
}

// SrcDoc creates the srcdoc attribute - A document to render in the [iframe]
//
// It can be applied to the following elements:
//   - [[iframe]]
//
// Value constraints: The source of [an iframesrcdoc document]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-srcdoc
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#the-iframe-element
// [an iframesrcdoc document]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#an-iframe-srcdoc-document
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func SrcDoc(srcDoc string) htmfunc.AttributeRenderer {
    return attribute.Attribute("srcdoc", srcDoc)
}

// SrcLang creates the srclang attribute - Language of the text track
//
// It can be applied to the following elements:
//   - [[track]]
//
// Value constraints: Valid BCP 47 language tag
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [track]: https://html.spec.whatwg.org/dev/media.html#attr-track-srclang
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func SrcLang(srcLang string) htmfunc.AttributeRenderer {
    return attribute.Attribute("srclang", srcLang)
}

// SrcSet creates the srcset attribute - Images to use in different situations, e.g., high-resolution displays, small monitors, etc.
//
// It can be applied to the following elements:
//   - [[img]]
//   - [[source]]
//
// Value constraints: Comma-separated list of [image candidate strings]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [img]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-srcset
// [source]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-source-srcset
// [image candidate strings]: https://html.spec.whatwg.org/dev/images.html#image-candidate-string
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func SrcSet(srcSet string) htmfunc.AttributeRenderer {
    return attribute.Attribute("srcset", srcSet)
}

// Start creates the start attribute - Starting value of the list
//
// It can be applied to the following elements:
//   - [[ol]]
//
// Value constraints: [Valid integer]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [ol]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-ol-start
// [Valid integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Start(start string) htmfunc.AttributeRenderer {
    return attribute.Attribute("start", start)
}

// Step creates the step attribute - Granularity to be matched by the form control's value
//
// It can be applied to the following elements:
//   - [[input]]
//
// Value constraints: [Valid floating-point number] greater than zero, or "any"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-step
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Step(step string) htmfunc.AttributeRenderer {
    return attribute.Attribute("step", step)
}

// Style creates the style attribute - Presentational and formatting instructions
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: CSS declarations*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#attr-style
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Style(style string) htmfunc.AttributeRenderer {
    return attribute.Attribute("style", style)
}

// TabIndex creates the tabindex attribute - Whether the element is focusable and sequentially focusable, andthe relative order of the element for the purposes of sequential focus navigation
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: [Valid integer]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/interaction.html#attr-tabindex
// [Valid integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func TabIndex(tabIndex string) htmfunc.AttributeRenderer {
    return attribute.Attribute("tabindex", tabIndex)
}

// Target creates the target attribute - [Navigable] for [hyperlink][navigation]
//
// It can be applied to the following elements:
//   - [[a]]
//   - [[area]]
//
// Value constraints: [Valid navigable target name or keyword]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-target
// [area]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-target
// [Navigable]: https://html.spec.whatwg.org/dev/document-sequences.html#navigable
// [hyperlink]: https://html.spec.whatwg.org/dev/links.html#hyperlink
// [navigation]: https://html.spec.whatwg.org/dev/browsing-the-web.html#navigate
// [Valid navigable target name or keyword]: https://html.spec.whatwg.org/dev/document-sequences.html#valid-navigable-target-name-or-keyword
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Target(target string) htmfunc.AttributeRenderer {
    return attribute.Attribute("target", target)
}

// Target creates the target attribute - Default [navigable] for [hyperlink][navigation] and [form submission]
//
// It can be applied to the following elements:
//   - [[base]]
//
// Value constraints: [Valid navigable target name or keyword]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [base]: https://html.spec.whatwg.org/dev/semantics.html#attr-base-target
// [navigable]: https://html.spec.whatwg.org/dev/document-sequences.html#navigable
// [hyperlink]: https://html.spec.whatwg.org/dev/links.html#hyperlink
// [navigation]: https://html.spec.whatwg.org/dev/browsing-the-web.html#navigate
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [Valid navigable target name or keyword]: https://html.spec.whatwg.org/dev/document-sequences.html#valid-navigable-target-name-or-keyword
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Target(target string) htmfunc.AttributeRenderer {
    return attribute.Attribute("target", target)
}

// Target creates the target attribute - [Navigable] for [form submission]
//
// It can be applied to the following elements:
//   - [[form]]
//
// Value constraints: [Valid navigable target name or keyword]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [form]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-target
// [Navigable]: https://html.spec.whatwg.org/dev/document-sequences.html#navigable
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [Valid navigable target name or keyword]: https://html.spec.whatwg.org/dev/document-sequences.html#valid-navigable-target-name-or-keyword
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Target(target string) htmfunc.AttributeRenderer {
    return attribute.Attribute("target", target)
}

// Type creates the type attribute - Hint for the type of the referenced resource
//
// It can be applied to the following elements:
//   - [[a]]
//   - [[link]]
//
// Value constraints: [Valid MIME type string]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-type
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-type
// [Valid MIME type string]: https://html.spec.whatwg.org/dev/https://mimesniff.spec.whatwg.org/#valid-mime-type
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Type(typeV string) htmfunc.AttributeRenderer {
    return attribute.Attribute("type", typeV)
}

// Type creates the type attribute - Type of embedded resource
//
// It can be applied to the following elements:
//   - [[embed]]
//   - [[object]]
//   - [[source]]
//
// Value constraints: [Valid MIME type string]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [embed]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-embed-type
// [object]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-object-type
// [source]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-source-type
// [Valid MIME type string]: https://html.spec.whatwg.org/dev/https://mimesniff.spec.whatwg.org/#valid-mime-type
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Type(typeV string) htmfunc.AttributeRenderer {
    return attribute.Attribute("type", typeV)
}

// Type creates the type attribute - Type of form control
//
// It can be applied to the following elements:
//   - [[input]]
//
// Value constraints: [input type keyword]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-type
// [input type keyword]: https://html.spec.whatwg.org/dev/input.html#attr-input-type
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Type(typeV string) htmfunc.AttributeRenderer {
    return attribute.Attribute("type", typeV)
}

// Type creates the type attribute - Type of script
//
// It can be applied to the following elements:
//   - [[script]]
//
// Value constraints: "module"; a [valid MIME type string] that is not a [JavaScript MIME type essence match]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [script]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-type
// [valid MIME type string]: https://html.spec.whatwg.org/dev/https://mimesniff.spec.whatwg.org/#valid-mime-type
// [JavaScript MIME type essence match]: https://html.spec.whatwg.org/dev/https://mimesniff.spec.whatwg.org/#javascript-mime-type-essence-match
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Type(typeV string) htmfunc.AttributeRenderer {
    return attribute.Attribute("type", typeV)
}

// UseMap creates the usemap attribute - Name of [image map] to use
//
// It can be applied to the following elements:
//   - [[img]]
//
// Value constraints: [Valid hash-name reference]*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [img]: https://html.spec.whatwg.org/dev/image-maps.html#attr-hyperlink-usemap
// [image map]: https://html.spec.whatwg.org/dev/image-maps.html#image-map
// [Valid hash-name reference]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-hash-name-reference
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func UseMap(useMap string) htmfunc.AttributeRenderer {
    return attribute.Attribute("usemap", useMap)
}

// Value creates the value attribute - Value of the form control
//
// It can be applied to the following elements:
//   - [[input]]
//
// Value constraints: Varies*
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-value
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Value(value string) htmfunc.AttributeRenderer {
    return attribute.Attribute("value", value)
}

// Value creates the value attribute - Ordinal value of the list item
//
// It can be applied to the following elements:
//   - [[li]]
//
// Value constraints: [Valid integer]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [li]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-li-value
// [Valid integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Value(value string) htmfunc.AttributeRenderer {
    return attribute.Attribute("value", value)
}

// Value creates the value attribute - Current value of the element
//
// It can be applied to the following elements:
//   - [[meter]]
//   - [[progress]]
//
// Value constraints: [Valid floating-point number]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-value
// [progress]: https://html.spec.whatwg.org/dev/form-elements.html#attr-progress-value
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Value(value string) htmfunc.AttributeRenderer {
    return attribute.Attribute("value", value)
}

// Width creates the width attribute - Horizontal dimension
//
// It can be applied to the following elements:
//   - [[canvas]]
//   - [[embed]]
//   - [[iframe]]
//   - [[img]]
//   - [[input]]
//   - [[object]]
//   - [[source] (in [picture])]
//   - [[video]]
//
// Value constraints: [Valid non-negative integer]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [canvas]: https://html.spec.whatwg.org/dev/canvas.html#attr-canvas-width
// [embed]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [iframe]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [img]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [input]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [object]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [source]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [picture]: https://html.spec.whatwg.org/dev/embedded-content.html#the-picture-element
// [video]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [Valid non-negative integer]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-non-negative-integer
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Width(width string) htmfunc.AttributeRenderer {
    return attribute.Attribute("width", width)
}

// WritingSuggestions creates the writingsuggestions attribute - Whether the element can offer writing suggestions or not.
//
// It can be applied to the following elements:
//   - [[HTML elements]]
//
// Value constraints: "[true]";"[false]";the empty string
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/interaction.html#attr-writingsuggestions
// [true]: https://html.spec.whatwg.org/dev/interaction.html#attr-writingsuggestions-true
// [false]: https://html.spec.whatwg.org/dev/interaction.html#attr-writingsuggestions-false
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func WritingSuggestions(writingSuggestions string) htmfunc.AttributeRenderer {
    return attribute.Attribute("writingsuggestions", writingSuggestions)
}
