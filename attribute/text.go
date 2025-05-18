// Generated file. DO NOT EDIT.

package attribute

import (
	"github.com/KrischanCS/htmfunc"
)

// Abbr creates the abbr attribute - Alternative label to use for the header cell when referencing the cell in other contexts
//
// It can be applied to the following elements:
//   - [th]
//
// The attribute abbr is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [th]: https://html.spec.whatwg.org/dev/tables.html#attr-th-abbr
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Abbr() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("abbr")
}

// AcceptCharset creates the accept-charset attribute - Character encodings to use for [form submission]
//
// It can be applied to the following elements:
//   - [form]
//
// The attribute accept-charset is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [form]: https://html.spec.whatwg.org/dev/forms.html#attr-form-accept-charset
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [ASCII case-insensitive]: https://html.spec.whatwg.org/dev/https://infra.spec.whatwg.org/#ascii-case-insensitive
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func AcceptCharset() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("accept-charset")
}

// Action creates the action attribute - [URL] to use for [form submission]
//
// It can be applied to the following elements:
//   - [form]
//
// The attribute action is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [form]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-action
// [URL]: https://html.spec.whatwg.org/dev/https://url.spec.whatwg.org/#concept-url
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [Valid non-empty URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Action() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("action")
}

// Allow creates the allow attribute - [Permissions policy] to be applied to the [iframe]'s contents
//
// It can be applied to the following elements:
//   - [iframe]
//
// The attribute allow is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-allow
// [Permissions policy]: https://html.spec.whatwg.org/dev/https://w3c.github.io/webappsec-feature-policy/#permissions-policy
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#the-iframe-element
// [Serialized permissions policy]: https://html.spec.whatwg.org/dev/https://w3c.github.io/webappsec-feature-policy/#serialized-permissions-policy
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Allow() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("allow")
}

// Alt creates the alt attribute - Replacement text for use when images are not available
//
// It can be applied to the following elements:
//   - [area]
//   - [img]
//   - [input]
//
// The attribute alt is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [area]: https://html.spec.whatwg.org/dev/image-maps.html#attr-area-alt
// [img]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-alt
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-alt
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Alt() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("alt")
}

// As creates the as attribute - [Potential destination] for a preload request (for [rel]="[preload]" and [rel]="[modulepreload]")
//
// It can be applied to the following elements:
//   - [link]
//
// The attribute as is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-as
// [Potential destination]: https://html.spec.whatwg.org/dev/https://fetch.spec.whatwg.org/#concept-potential-destination
// [rel]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [preload]: https://html.spec.whatwg.org/dev/links.html#link-type-preload
// [modulepreload]: https://html.spec.whatwg.org/dev/links.html#link-type-modulepreload
// [Potential destination]: https://html.spec.whatwg.org/dev/https://fetch.spec.whatwg.org/#concept-potential-destination
// [rel]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [preload]: https://html.spec.whatwg.org/dev/links.html#link-type-preload
// [script-like destination]: https://html.spec.whatwg.org/dev/https://fetch.spec.whatwg.org/#request-destination-script-like
// [modulepreload]: https://html.spec.whatwg.org/dev/links.html#link-type-modulepreload
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func As() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("as")
}

// AutoComplete_InputSelectTextarea creates the autocomplete attribute - Hint for form autofill feature
//
// It can be applied to the following elements:
//   - [input]
//   - [select]
//   - [textarea]
//
// The attribute autocomplete is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-autocomplete
// [select]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-autocomplete
// [textarea]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-autocomplete
// [Autofill field]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#autofill-field
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func AutoComplete_InputSelectTextarea() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("autocomplete")
}

// Cite creates the cite attribute - Link to the source of the quotation or more information about the edit
//
// It can be applied to the following elements:
//   - [blockquote]
//   - [del]
//   - [ins]
//   - [q]
//
// The attribute cite is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [blockquote]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-blockquote-cite
// [del]: https://html.spec.whatwg.org/dev/edits.html#attr-mod-cite
// [ins]: https://html.spec.whatwg.org/dev/edits.html#attr-mod-cite
// [q]: https://html.spec.whatwg.org/dev/text-level-semantics.html#attr-q-cite
// [Valid URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Cite() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("cite")
}

// Color creates the color attribute - Color to use when customizing a site's icon (for [rel]="mask-icon")
//
// It can be applied to the following elements:
//   - [link]
//
// The attribute color is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-color
// [rel]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [<color>]: https://html.spec.whatwg.org/dev/https://drafts.csswg.org/css-color/#typedef-color
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Color() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("color")
}

// Command creates the command attribute - Indicates to the targeted element which action to take.
//
// It can be applied to the following elements:
//   - [button]
//
// The attribute command is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-command
// [toggle-popover]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-command-toggle-popover
// [show-popover]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-command-show-popover
// [hide-popover]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-command-hide-popover
// [close]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-command-close
// [show-modal]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-command-show-modal
// [custom command keyword]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-command-custom
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Command() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("command")
}

// Commandfor creates the commandfor attribute - Targets another element to be invoked.
//
// It can be applied to the following elements:
//   - [button]
//
// The attribute commandfor is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-commandfor
// [ID]: https://html.spec.whatwg.org/dev/https://dom.spec.whatwg.org/#concept-id
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Commandfor() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("commandfor")
}

// Content creates the content attribute - Value of the element
//
// It can be applied to the following elements:
//   - [meta]
//
// The attribute content is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meta]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-content
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Content() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("content")
}

// Data creates the data attribute - Address of the resource
//
// It can be applied to the following elements:
//   - [object]
//
// The attribute data is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [object]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-object-data
// [Valid non-empty URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Data() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("data")
}

// DateTime_DelIns creates the datetime attribute - Date and (optionally) time of the change
//
// It can be applied to the following elements:
//   - [del]
//   - [ins]
//
// The attribute datetime is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [del]: https://html.spec.whatwg.org/dev/edits.html#attr-mod-datetime
// [ins]: https://html.spec.whatwg.org/dev/edits.html#attr-mod-datetime
// [Valid date string with optional time]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-date-string-with-optional-time
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func DateTime_DelIns() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("datetime")
}

// DateTime_Time creates the datetime attribute - Machine-readable value
//
// It can be applied to the following elements:
//   - [time]
//
// The attribute datetime is a [Boolean attribute].
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
func DateTime_Time() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("datetime")
}

// DirName creates the dirname attribute - Name of form control to use for sending the element's [directionality] in [form submission]
//
// It can be applied to the following elements:
//   - [input]
//   - [textarea]
//
// The attribute dirname is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-dirname
// [textarea]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-dirname
// [directionality]: https://html.spec.whatwg.org/dev/dom.html#the-directionality
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func DirName() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("dirname")
}

// Download creates the download attribute - Whether to download the resource instead of navigating to it, and its filename if so
//
// It can be applied to the following elements:
//   - [a]
//   - [area]
//
// The attribute download is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-download
// [area]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-download
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Download() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("download")
}

// For_Label creates the for attribute - Associate the label with form control
//
// It can be applied to the following elements:
//   - [label]
//
// The attribute for is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [label]: https://html.spec.whatwg.org/dev/forms.html#attr-label-for
// [ID]: https://html.spec.whatwg.org/dev/https://dom.spec.whatwg.org/#concept-id
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func For_Label() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("for")
}

// Form creates the form attribute - Associates the element with a [form] element
//
// It can be applied to the following elements:
//   - [button]
//   - [fieldset]
//   - [input]
//   - [object]
//   - [output]
//   - [select]
//   - [textarea]
//   - [form-associated custom elements]
//
// The attribute form is a [Boolean attribute].
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
func Form() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("form")
}

// FormAction creates the formaction attribute - [URL] to use for [form submission]
//
// It can be applied to the following elements:
//   - [button]
//   - [input]
//
// The attribute formaction is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formaction
// [input]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formaction
// [URL]: https://html.spec.whatwg.org/dev/https://url.spec.whatwg.org/#concept-url
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [Valid non-empty URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func FormAction() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("formaction")
}

// FormTarget creates the formtarget attribute - [Navigable] for [form submission]
//
// It can be applied to the following elements:
//   - [button]
//   - [input]
//
// The attribute formtarget is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formtarget
// [input]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formtarget
// [Navigable]: https://html.spec.whatwg.org/dev/document-sequences.html#navigable
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [Valid navigable target name or keyword]: https://html.spec.whatwg.org/dev/document-sequences.html#valid-navigable-target-name-or-keyword
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func FormTarget() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("formtarget")
}

// HRef_AArea creates the href attribute - Address of the [hyperlink]
//
// It can be applied to the following elements:
//   - [a]
//   - [area]
//
// The attribute href is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-href
// [area]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-href
// [hyperlink]: https://html.spec.whatwg.org/dev/links.html#hyperlink
// [Valid URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func HRef_AArea() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("href")
}

// HRef_Link creates the href attribute - Address of the [hyperlink]
//
// It can be applied to the following elements:
//   - [link]
//
// The attribute href is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-href
// [hyperlink]: https://html.spec.whatwg.org/dev/links.html#hyperlink
// [Valid non-empty URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func HRef_Link() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("href")
}

// HRef_Base creates the href attribute - [Document base URL]
//
// It can be applied to the following elements:
//   - [base]
//
// The attribute href is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [base]: https://html.spec.whatwg.org/dev/semantics.html#attr-base-href
// [Document base URL]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#document-base-url
// [Valid URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func HRef_Base() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("href")
}

// HRefLang creates the hreflang attribute - Language of the linked resource
//
// It can be applied to the following elements:
//   - [a]
//   - [link]
//
// The attribute hreflang is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-hreflang
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-hreflang
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func HRefLang() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("hreflang")
}

// Id creates the id attribute - The element's [ID]
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// The attribute id is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#the-id-attribute
// [ID]: https://html.spec.whatwg.org/dev/https://dom.spec.whatwg.org/#concept-id
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Id() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("id")
}

// ImageSizes creates the imagesizes attribute - Image sizes for different page layouts (for [rel]="[preload]")
//
// It can be applied to the following elements:
//   - [link]
//
// The attribute imagesizes is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-imagesizes
// [rel]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [preload]: https://html.spec.whatwg.org/dev/links.html#link-type-preload
// [Valid source size list]: https://html.spec.whatwg.org/dev/images.html#valid-source-size-list
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ImageSizes() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("imagesizes")
}

// Integrity creates the integrity attribute - Integrity metadata used in Subresource Integrity checks [[SRI]]
//
// It can be applied to the following elements:
//   - [link]
//   - [script]
//
// The attribute integrity is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-integrity
// [script]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-integrity
// [[SRI]]: https://html.spec.whatwg.org/dev/references.html#refsSRI
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Integrity() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("integrity")
}

// Is creates the is attribute - Creates a [customized built-in element]
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// The attribute is is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/custom-elements.html#attr-is
// [customized built-in element]: https://html.spec.whatwg.org/dev/custom-elements.html#customized-built-in-element
// [Valid custom element name]: https://html.spec.whatwg.org/dev/custom-elements.html#valid-custom-element-name
// [customized built-in element]: https://html.spec.whatwg.org/dev/custom-elements.html#customized-built-in-element
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Is() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("is")
}

// ItemId creates the itemid attribute - [Global identifier] for a microdata item
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// The attribute itemid is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/microdata.html#attr-itemid
// [Global identifier]: https://html.spec.whatwg.org/dev/microdata.html#global-identifier
// [Valid URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ItemId() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("itemid")
}

// Label creates the label attribute - User-visible label
//
// It can be applied to the following elements:
//   - [optgroup]
//   - [option]
//   - [track]
//
// The attribute label is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [optgroup]: https://html.spec.whatwg.org/dev/form-elements.html#attr-optgroup-label
// [option]: https://html.spec.whatwg.org/dev/form-elements.html#attr-option-label
// [track]: https://html.spec.whatwg.org/dev/media.html#attr-track-label
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Label() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("label")
}

// Lang creates the lang attribute - Language of the element
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// The attribute lang is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#attr-lang
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Lang() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("lang")
}

// List creates the list attribute - List of autocomplete options
//
// It can be applied to the following elements:
//   - [input]
//
// The attribute list is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-list
// [ID]: https://html.spec.whatwg.org/dev/https://dom.spec.whatwg.org/#concept-id
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func List() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("list")
}

// Max_Input creates the max attribute - Maximum value
//
// It can be applied to the following elements:
//   - [input]
//
// The attribute max is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-max
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Max_Input() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("max")
}

// Media creates the media attribute - Applicable media
//
// It can be applied to the following elements:
//   - [link]
//   - [meta]
//   - [source]
//   - [style]
//
// The attribute media is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-media
// [meta]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-media
// [source]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-source-media
// [style]: https://html.spec.whatwg.org/dev/semantics.html#attr-style-media
// [Valid media query list]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-media-query-list
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Media() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("media")
}

// Min_Input creates the min attribute - Minimum value
//
// It can be applied to the following elements:
//   - [input]
//
// The attribute min is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-min
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Min_Input() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("min")
}

// Name_InputsOutputs creates the name attribute - Name of the element to use for [form submission] and in the [form.elements] API
//
// It can be applied to the following elements:
//   - [button]
//   - [fieldset]
//   - [input]
//   - [output]
//   - [select]
//   - [textarea]
//   - [form-associated custom elements]
//
// The attribute name is a [Boolean attribute].
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
func Name_InputsOutputs() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("name")
}

// Name_Details creates the name attribute - Name of group of mutually-exclusive [details] elements
//
// It can be applied to the following elements:
//   - [details]
//
// The attribute name is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [details]: https://html.spec.whatwg.org/dev/interactive-elements.html#attr-details-name
// [details]: https://html.spec.whatwg.org/dev/interactive-elements.html#the-details-element
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Name_Details() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("name")
}

// Name_Form creates the name attribute - Name of form to use in the [document.forms] API
//
// It can be applied to the following elements:
//   - [form]
//
// The attribute name is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [form]: https://html.spec.whatwg.org/dev/forms.html#attr-form-name
// [document.forms]: https://html.spec.whatwg.org/dev/dom.html#dom-document-forms
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Name_Form() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("name")
}

// Name_IframeObject creates the name attribute - Name of [content navigable]
//
// It can be applied to the following elements:
//   - [iframe]
//   - [object]
//
// The attribute name is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-name
// [object]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-object-name
// [content navigable]: https://html.spec.whatwg.org/dev/document-sequences.html#content-navigable
// [Valid navigable target name or keyword]: https://html.spec.whatwg.org/dev/document-sequences.html#valid-navigable-target-name-or-keyword
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Name_IframeObject() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("name")
}

// Name_Map creates the name attribute - Name of [image map] to [reference] from the [usemap] attribute
//
// It can be applied to the following elements:
//   - [map]
//
// The attribute name is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [map]: https://html.spec.whatwg.org/dev/image-maps.html#attr-map-name
// [image map]: https://html.spec.whatwg.org/dev/image-maps.html#image-map
// [reference]: https://html.spec.whatwg.org/dev/dom.html#referenced
// [usemap]: https://html.spec.whatwg.org/dev/image-maps.html#attr-hyperlink-usemap
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Name_Map() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("name")
}

// Name_Meta creates the name attribute - Metadata name
//
// It can be applied to the following elements:
//   - [meta]
//
// The attribute name is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meta]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-name
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Name_Meta() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("name")
}

// Name_Slot creates the name attribute - Name of shadow tree slot
//
// It can be applied to the following elements:
//   - [slot]
//
// The attribute name is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [slot]: https://html.spec.whatwg.org/dev/scripting.html#attr-slot-name
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Name_Slot() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("name")
}

// Nonce creates the nonce attribute - Cryptographic nonce used in Content Security Policy checks [[CSP]]
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// The attribute nonce is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#attr-nonce
// [[CSP]]: https://html.spec.whatwg.org/dev/references.html#refsCSP
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Nonce() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("nonce")
}

// Pattern creates the pattern attribute - Pattern to be matched by the form control's value
//
// It can be applied to the following elements:
//   - [input]
//
// The attribute pattern is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-pattern
// [Pattern]: https://html.spec.whatwg.org/dev/https://tc39.es/ecma262/#prod-Pattern
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Pattern() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("pattern")
}

// PlaceHolder creates the placeholder attribute - User-visible label to be placed within the form control
//
// It can be applied to the following elements:
//   - [input]
//   - [textarea]
//
// The attribute placeholder is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-placeholder
// [textarea]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-placeholder
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func PlaceHolder() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("placeholder")
}

// PopOverTarget creates the popovertarget attribute - Targets a popover element to toggle, show, or hide
//
// It can be applied to the following elements:
//   - [button]
//   - [input]
//
// The attribute popovertarget is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/popover.html#attr-popovertarget
// [input]: https://html.spec.whatwg.org/dev/popover.html#attr-popovertarget
// [ID]: https://html.spec.whatwg.org/dev/https://dom.spec.whatwg.org/#concept-id
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func PopOverTarget() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("popovertarget")
}

// Poster creates the poster attribute - Poster frame to show prior to video playback
//
// It can be applied to the following elements:
//   - [video]
//
// The attribute poster is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [video]: https://html.spec.whatwg.org/dev/media.html#attr-video-poster
// [Valid non-empty URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Poster() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("poster")
}

// ReferrerPolicy creates the referrerpolicy attribute - [Referrer policy] for [fetches] initiated by the element
//
// It can be applied to the following elements:
//   - [a]
//   - [area]
//   - [iframe]
//   - [img]
//   - [link]
//   - [script]
//
// The attribute referrerpolicy is a [Boolean attribute].
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
func ReferrerPolicy() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("referrerpolicy")
}

// Sizes_ImgSource creates the sizes attribute - Image sizes for different page layouts
//
// It can be applied to the following elements:
//   - [img]
//   - [source]
//
// The attribute sizes is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [img]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-sizes
// [source]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-source-sizes
// [Valid source size list]: https://html.spec.whatwg.org/dev/images.html#valid-source-size-list
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Sizes_ImgSource() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("sizes")
}

// Slot creates the slot attribute - The element's desired slot
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// The attribute slot is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#the-id-attribute
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Slot() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("slot")
}

// Src creates the src attribute - Address of the resource
//
// It can be applied to the following elements:
//   - [audio]
//   - [embed]
//   - [iframe]
//   - [img]
//   - [input]
//   - [script]
//   - [source] (in [video] or [audio (1)])
//   - [track]
//   - [video (1)]
//
// The attribute src is a [Boolean attribute].
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
// [audio (1)]: https://html.spec.whatwg.org/dev/media.html#the-audio-element
// [track]: https://html.spec.whatwg.org/dev/media.html#attr-track-src
// [video (1)]: https://html.spec.whatwg.org/dev/media.html#attr-media-src
// [Valid non-empty URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Src() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("src")
}

// SrcDoc creates the srcdoc attribute - A document to render in the [iframe]
//
// It can be applied to the following elements:
//   - [iframe]
//
// The attribute srcdoc is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-srcdoc
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#the-iframe-element
// [an iframesrcdoc document]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#an-iframe-srcdoc-document
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func SrcDoc() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("srcdoc")
}

// SrcLang creates the srclang attribute - Language of the text track
//
// It can be applied to the following elements:
//   - [track]
//
// The attribute srclang is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [track]: https://html.spec.whatwg.org/dev/media.html#attr-track-srclang
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func SrcLang() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("srclang")
}

// Style creates the style attribute - Presentational and formatting instructions
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// The attribute style is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#attr-style
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Style() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("style")
}

// Target_AArea creates the target attribute - [Navigable] for [hyperlink][navigation]
//
// It can be applied to the following elements:
//   - [a]
//   - [area]
//
// The attribute target is a [Boolean attribute].
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
func Target_AArea() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("target")
}

// Target_Base creates the target attribute - Default [navigable] for [hyperlink][navigation] and [form submission]
//
// It can be applied to the following elements:
//   - [base]
//
// The attribute target is a [Boolean attribute].
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
func Target_Base() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("target")
}

// Target_Form creates the target attribute - [Navigable] for [form submission]
//
// It can be applied to the following elements:
//   - [form]
//
// The attribute target is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [form]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-target
// [Navigable]: https://html.spec.whatwg.org/dev/document-sequences.html#navigable
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [Valid navigable target name or keyword]: https://html.spec.whatwg.org/dev/document-sequences.html#valid-navigable-target-name-or-keyword
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Target_Form() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("target")
}

// Title creates the title attribute - Advisory information for the element
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// The attribute title is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#attr-title
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Title() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("title")
}

// Title_AbbrDfn creates the title attribute - Full term or expansion of abbreviation
//
// It can be applied to the following elements:
//   - [abbr]
//   - [dfn]
//
// The attribute title is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [abbr]: https://html.spec.whatwg.org/dev/text-level-semantics.html#attr-abbr-title
// [dfn]: https://html.spec.whatwg.org/dev/text-level-semantics.html#attr-dfn-title
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Title_AbbrDfn() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("title")
}

// Title_Input creates the title attribute - Description of pattern (when used with [pattern] attribute)
//
// It can be applied to the following elements:
//   - [input]
//
// The attribute title is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-title
// [pattern]: https://html.spec.whatwg.org/dev/input.html#attr-input-pattern
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Title_Input() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("title")
}

// Title_Link creates the title attribute - Title of the link
//
// It can be applied to the following elements:
//   - [link]
//
// The attribute title is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-title
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Title_Link() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("title")
}

// Title_LinkStyle creates the title attribute - [CSS style sheet set name]
//
// It can be applied to the following elements:
//   - [link]
//   - [style]
//
// The attribute title is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-title
// [style]: https://html.spec.whatwg.org/dev/semantics.html#attr-style-title
// [CSS style sheet set name]: https://html.spec.whatwg.org/dev/https://drafts.csswg.org/cssom/#css-style-sheet-set-name
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Title_LinkStyle() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("title")
}

// Type_ALink creates the type attribute - Hint for the type of the referenced resource
//
// It can be applied to the following elements:
//   - [a]
//   - [link]
//
// The attribute type is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-type
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-type
// [Valid MIME type string]: https://html.spec.whatwg.org/dev/https://mimesniff.spec.whatwg.org/#valid-mime-type
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Type_ALink() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("type")
}

// Type_EmbedObjectSource creates the type attribute - Type of embedded resource
//
// It can be applied to the following elements:
//   - [embed]
//   - [object]
//   - [source]
//
// The attribute type is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [embed]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-embed-type
// [object]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-object-type
// [source]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-source-type
// [Valid MIME type string]: https://html.spec.whatwg.org/dev/https://mimesniff.spec.whatwg.org/#valid-mime-type
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Type_EmbedObjectSource() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("type")
}

// Type_Script creates the type attribute - Type of script
//
// It can be applied to the following elements:
//   - [script]
//
// The attribute type is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [script]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-type
// [valid MIME type string]: https://html.spec.whatwg.org/dev/https://mimesniff.spec.whatwg.org/#valid-mime-type
// [JavaScript MIME type essence match]: https://html.spec.whatwg.org/dev/https://mimesniff.spec.whatwg.org/#javascript-mime-type-essence-match
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Type_Script() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("type")
}

// UseMap creates the usemap attribute - Name of [image map] to use
//
// It can be applied to the following elements:
//   - [img]
//
// The attribute usemap is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [img]: https://html.spec.whatwg.org/dev/image-maps.html#attr-hyperlink-usemap
// [image map]: https://html.spec.whatwg.org/dev/image-maps.html#image-map
// [Valid hash-name reference]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-hash-name-reference
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func UseMap() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("usemap")
}

// Value_ButtonOption creates the value attribute - Value to be used for [form submission]
//
// It can be applied to the following elements:
//   - [button]
//   - [option]
//
// The attribute value is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-value
// [option]: https://html.spec.whatwg.org/dev/form-elements.html#attr-option-value
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Value_ButtonOption() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("value")
}

// Value_Data creates the value attribute - Machine-readable value
//
// It can be applied to the following elements:
//   - [data]
//
// The attribute value is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [data]: https://html.spec.whatwg.org/dev/text-level-semantics.html#attr-data-value
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Value_Data() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("value")
}

// Value_Input creates the value attribute - Value of the form control
//
// It can be applied to the following elements:
//   - [input]
//
// The attribute value is a [Boolean attribute].
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-value
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Value_Input() htmfunc.Attribute {
	return htmfunc.WriteBoolAttribute("value")
}
