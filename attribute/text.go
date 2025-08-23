// Generated file. DO NOT EDIT.

// Package element provides functions for the elements defined in the HTML
// standard.
package attribute

import (
	"github.com/KrischanCS/roundly"
)

// AbbrAttribute creates the abbr attribute - Alternative label to use for the header cell when referencing the cell in other contexts
//
// It can be applied to the following elements:
//   - [th]
//
// Value constraints: [Text] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [th]: https://html.spec.whatwg.org/dev/tables.html#attr-th-abbr
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func AbbrAttribute(abbr string) roundly.Attribute {
	return roundly.WriteAttribute("abbr", abbr)
}

// AcceptCharset creates the accept-charset attribute - Character encodings to use for [form submission]
//
// It can be applied to the following elements:
//   - [form]
//
// Value constraints: [ASCII case-insensitive] match for "UTF-8"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [form]: https://html.spec.whatwg.org/dev/forms.html#attr-form-accept-charset
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [ASCII case-insensitive]: https://html.spec.whatwg.org/dev/https://infra.spec.whatwg.org/#ascii-case-insensitive
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func AcceptCharset(acceptCharset string) roundly.Attribute {
	return roundly.WriteAttribute("accept-charset", acceptCharset)
}

// Action creates the action attribute - [URL] to use for [form submission]
//
// It can be applied to the following elements:
//   - [form]
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
func Action(action string) roundly.Attribute {
	return roundly.WriteAttribute("action", action)
}

// Allow creates the allow attribute - [Permissions policy] to be applied to the [iframe]'s contents
//
// It can be applied to the following elements:
//   - [iframe]
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
func Allow(allow string) roundly.Attribute {
	return roundly.WriteAttribute("allow", allow)
}

// Alt creates the alt attribute - Replacement text for use when images are not available
//
// It can be applied to the following elements:
//   - [area] [img] [input]
//   - [img]
//   - [input]
//
// Value constraints: [Text] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [area]: https://html.spec.whatwg.org/dev/image-maps.html#attr-area-alt
// [img]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-alt
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-alt
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Alt(alt string) roundly.Attribute {
	return roundly.WriteAttribute("alt", alt)
}

// As creates the as attribute - [Potential destination] for a preload request (for [rel]="[preload]" and [rel]="[modulepreload]")
//
// It can be applied to the following elements:
//   - [link]
//
// Value constraints: [Potential destination], for [rel]="[preload]"; [script-like destination], for [rel]="[modulepreload]"
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
func As(as string) roundly.Attribute {
	return roundly.WriteAttribute("as", as)
}

// AutoComplete creates the autocomplete attribute - Hint for form autofill feature
//
// It can be applied to the following elements:
//   - [input] [select] [textarea]
//   - [select]
//   - [textarea]
//
// Value constraints: [Autofill field] name and related tokens (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-autocomplete
// [select]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-autocomplete
// [textarea]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-autocomplete
// [Autofill field]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#autofill-field
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func AutoComplete(autoComplete string) roundly.Attribute {
	return roundly.WriteAttribute("autocomplete", autoComplete)
}

// CiteAttribute creates the cite attribute - Link to the source of the quotation or more information about the edit
//
// It can be applied to the following elements:
//   - [blockquote] [del] [ins] [q]
//   - [del]
//   - [ins]
//   - [q]
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
func CiteAttribute(cite string) roundly.Attribute {
	return roundly.WriteAttribute("cite", cite)
}

// Color creates the color attribute - Color to use when customizing a site's icon (for [rel]="mask-icon")
//
// It can be applied to the following elements:
//   - [link]
//
// Value constraints: CSS [<color>]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-color
// [rel]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [<color>]: https://html.spec.whatwg.org/dev/https://drafts.csswg.org/css-color/#typedef-color
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Color(color string) roundly.Attribute {
	return roundly.WriteAttribute("color", color)
}

// Command creates the command attribute - Indicates to the targeted element which action to take.
//
// It can be applied to the following elements:
//   - [button]
//
// Value constraints: "[toggle-popover]";"[show-popover]";"[hide-popover]";"[close]";"[request-close]";"[show-modal]";a [custom command keyword]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-command
// [toggle-popover]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-command-toggle-popover
// [show-popover]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-command-show-popover
// [hide-popover]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-command-hide-popover
// [close]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-command-close
// [request-close]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-command-request-close
// [show-modal]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-command-show-modal
// [custom command keyword]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-command-custom
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Command(command string) roundly.Attribute {
	return roundly.WriteAttribute("command", command)
}

// Commandfor creates the commandfor attribute - Targets another element to be invoked.
//
// It can be applied to the following elements:
//   - [button]
//
// Value constraints: [ID] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-commandfor
// [ID]: https://html.spec.whatwg.org/dev/https://dom.spec.whatwg.org/#concept-id
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Commandfor(commandfor string) roundly.Attribute {
	return roundly.WriteAttribute("commandfor", commandfor)
}

// Content creates the content attribute - Value of the element
//
// It can be applied to the following elements:
//   - [meta]
//
// Value constraints: [Text] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meta]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-content
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Content(content string) roundly.Attribute {
	return roundly.WriteAttribute("content", content)
}

// DataAttribute creates the data attribute - Address of the resource
//
// It can be applied to the following elements:
//   - [object]
//
// Value constraints: [Valid non-empty URL potentially surrounded by spaces]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [object]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-object-data
// [Valid non-empty URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func DataAttribute(data string) roundly.Attribute {
	return roundly.WriteAttribute("data", data)
}

// DateTime creates the datetime attribute - It's semantics varies depending on the element it is applied to.
//
// It can be applied to the following elements:
//   - [del] [ins]: Date and (optionally) time of the change
//   - [time]: Machine-readable value
//
// Value constraints: 
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [del]: https://html.spec.whatwg.org/dev/edits.html#attr-mod-datetime
// [ins]: https://html.spec.whatwg.org/dev/edits.html#attr-mod-datetime
// [Valid date string with optional time]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-date-string-with-optional-time
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
func DateTime(dateTime string) roundly.Attribute {
	return roundly.WriteAttribute("datetime", dateTime)
}

// DirName creates the dirname attribute - Name of form control to use for sending the element's [directionality] in [form submission]
//
// It can be applied to the following elements:
//   - [input] [textarea]
//   - [textarea]
//
// Value constraints: [Text] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-dirname
// [textarea]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fe-dirname
// [directionality]: https://html.spec.whatwg.org/dev/dom.html#the-directionality
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func DirName(dirName string) roundly.Attribute {
	return roundly.WriteAttribute("dirname", dirName)
}

// Download creates the download attribute - Whether to download the resource instead of navigating to it, and its filename if so
//
// It can be applied to the following elements:
//   - [a] [area]
//   - [area]
//
// Value constraints: Text
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-download
// [area]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-download
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Download(download string) roundly.Attribute {
	return roundly.WriteAttribute("download", download)
}

// For creates the for attribute - Associate the label with form control
//
// It can be applied to the following elements:
//   - [label]
//
// Value constraints: [ID] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [label]: https://html.spec.whatwg.org/dev/forms.html#attr-label-for
// [ID]: https://html.spec.whatwg.org/dev/https://dom.spec.whatwg.org/#concept-id
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func For(forV string) roundly.Attribute {
	return roundly.WriteAttribute("for", forV)
}

// FormAttribute creates the form attribute - Associates the element with a [form] element
//
// It can be applied to the following elements:
//   - [button] [fieldset] [input] [object] [output] [select] [textarea] [form-associated custom elements]
//   - [fieldset]
//   - [input]
//   - [object]
//   - [output]
//   - [select]
//   - [textarea]
//   - [form-associated custom elements]
//
// Value constraints: [ID] (Additional rules apply, see elements documentation)
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
func FormAttribute(form string) roundly.Attribute {
	return roundly.WriteAttribute("form", form)
}

// FormAction creates the formaction attribute - [URL] to use for [form submission]
//
// It can be applied to the following elements:
//   - [button] [input]
//   - [input]
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
func FormAction(formAction string) roundly.Attribute {
	return roundly.WriteAttribute("formaction", formAction)
}

// FormTarget creates the formtarget attribute - [Navigable] for [form submission]
//
// It can be applied to the following elements:
//   - [button] [input]
//   - [input]
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
func FormTarget(formTarget string) roundly.Attribute {
	return roundly.WriteAttribute("formtarget", formTarget)
}

// HRef creates the href attribute - It's semantics varies depending on the element it is applied to.
//
// It can be applied to the following elements:
//   - [a] [area]: Address of the [hyperlink]
//   - [link]: Address of the [hyperlink]
//   - [base]: Document base URL
//
// Value constraints: 
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-href
// [area]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-href
// [hyperlink]: https://html.spec.whatwg.org/dev/links.html#hyperlink
// [Valid URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-url-potentially-surrounded-by-spaces
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-href
// [Valid non-empty URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces
// [base]: https://html.spec.whatwg.org/dev/semantics.html#attr-base-href
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func HRef(hRef string) roundly.Attribute {
	return roundly.WriteAttribute("href", hRef)
}

// HRefLang creates the hreflang attribute - Language of the linked resource
//
// It can be applied to the following elements:
//   - [a] [link]
//   - [link]
//
// Value constraints: Valid BCP 47 language tag
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-hreflang
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-hreflang
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func HRefLang(hRefLang string) roundly.Attribute {
	return roundly.WriteAttribute("hreflang", hRefLang)
}

// Id creates the id attribute - The element's [ID]
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Value constraints: [Text] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#the-id-attribute
// [ID]: https://html.spec.whatwg.org/dev/https://dom.spec.whatwg.org/#concept-id
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Id(id string) roundly.Attribute {
	return roundly.WriteAttribute("id", id)
}

// ImageSizes creates the imagesizes attribute - Image sizes for different page layouts (for [rel]="[preload]")
//
// It can be applied to the following elements:
//   - [link]
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
func ImageSizes(imageSizes string) roundly.Attribute {
	return roundly.WriteAttribute("imagesizes", imageSizes)
}

// Integrity creates the integrity attribute - Integrity metadata used in Subresource Integrity checks [[SRI]]
//
// It can be applied to the following elements:
//   - [link] [script]
//   - [script]
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
func Integrity(integrity string) roundly.Attribute {
	return roundly.WriteAttribute("integrity", integrity)
}

// Is creates the is attribute - Creates a [customized built-in element]
//
// It can be applied to the following elements:
//   - [HTML elements]
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
func Is(is string) roundly.Attribute {
	return roundly.WriteAttribute("is", is)
}

// ItemId creates the itemid attribute - [Global identifier] for a microdata item
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Value constraints: [Valid URL potentially surrounded by spaces]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/microdata.html#attr-itemid
// [Global identifier]: https://html.spec.whatwg.org/dev/microdata.html#global-identifier
// [Valid URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ItemId(itemId string) roundly.Attribute {
	return roundly.WriteAttribute("itemid", itemId)
}

// LabelAttribute creates the label attribute - User-visible label
//
// It can be applied to the following elements:
//   - [optgroup] [option] [track]
//   - [option]
//   - [track]
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
func LabelAttribute(label string) roundly.Attribute {
	return roundly.WriteAttribute("label", label)
}

// Lang creates the lang attribute - Language of the element
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Value constraints: Valid BCP 47 language tag or the empty string
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#attr-lang
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Lang(lang string) roundly.Attribute {
	return roundly.WriteAttribute("lang", lang)
}

// List creates the list attribute - List of autocomplete options
//
// It can be applied to the following elements:
//   - [input]
//
// Value constraints: [ID] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-list
// [ID]: https://html.spec.whatwg.org/dev/https://dom.spec.whatwg.org/#concept-id
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func List(list string) roundly.Attribute {
	return roundly.WriteAttribute("list", list)
}

// Max creates the max attribute - Maximum value
//
// It can be applied to the following elements:
//   - [input]
//
// Value constraints: Varies (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-max
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Max(max string) roundly.Attribute {
	return roundly.WriteAttribute("max", max)
}

// Media creates the media attribute - Applicable media
//
// It can be applied to the following elements:
//   - [link] [meta] [source] [style]
//   - [meta]
//   - [source]
//   - [style]
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
func Media(media string) roundly.Attribute {
	return roundly.WriteAttribute("media", media)
}

// Min creates the min attribute - Minimum value
//
// It can be applied to the following elements:
//   - [input]
//
// Value constraints: Varies (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-min
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Min(min string) roundly.Attribute {
	return roundly.WriteAttribute("min", min)
}

// Name creates the name attribute - It's semantics varies depending on the element it is applied to.
//
// It can be applied to the following elements:
//   - [button] [fieldset] [input] [output] [select] [textarea] [form-associated custom elements]: Name of the element to use for [form submission] and in the [form.elements] API
//   - [details]: Name of group of mutually-exclusive [details] elements
//   - [form]: Name of form to use in the [document.forms] API
//   - [iframe] [object]: Name of [content navigable]
//   - [map]: Name of [image map] to [reference] from the [usemap] attribute
//   - [meta]: Metadata name
//   - [slot]: Name of shadow tree slot
//
// Value constraints: 
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
// [details]: https://html.spec.whatwg.org/dev/interactive-elements.html#attr-details-name
// [details]: https://html.spec.whatwg.org/dev/interactive-elements.html#the-details-element
// [form]: https://html.spec.whatwg.org/dev/forms.html#attr-form-name
// [document.forms]: https://html.spec.whatwg.org/dev/dom.html#dom-document-forms
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-name
// [object]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-object-name
// [content navigable]: https://html.spec.whatwg.org/dev/document-sequences.html#content-navigable
// [Valid navigable target name or keyword]: https://html.spec.whatwg.org/dev/document-sequences.html#valid-navigable-target-name-or-keyword
// [map]: https://html.spec.whatwg.org/dev/image-maps.html#attr-map-name
// [image map]: https://html.spec.whatwg.org/dev/image-maps.html#image-map
// [reference]: https://html.spec.whatwg.org/dev/dom.html#referenced
// [usemap]: https://html.spec.whatwg.org/dev/image-maps.html#attr-hyperlink-usemap
// [meta]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-name
// [slot]: https://html.spec.whatwg.org/dev/scripting.html#attr-slot-name
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Name(name string) roundly.Attribute {
	return roundly.WriteAttribute("name", name)
}

// Nonce creates the nonce attribute - Cryptographic nonce used in Content Security Policy checks [[CSP]]
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Value constraints: [Text]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#attr-nonce
// [[CSP]]: https://html.spec.whatwg.org/dev/references.html#refsCSP
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Nonce(nonce string) roundly.Attribute {
	return roundly.WriteAttribute("nonce", nonce)
}

// Pattern creates the pattern attribute - Pattern to be matched by the form control's value
//
// It can be applied to the following elements:
//   - [input]
//
// Value constraints: Regular expression matching the JavaScript [Pattern] production
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-pattern
// [Pattern]: https://html.spec.whatwg.org/dev/https://tc39.es/ecma262/#prod-Pattern
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Pattern(pattern string) roundly.Attribute {
	return roundly.WriteAttribute("pattern", pattern)
}

// PlaceHolder creates the placeholder attribute - User-visible label to be placed within the form control
//
// It can be applied to the following elements:
//   - [input] [textarea]
//   - [textarea]
//
// Value constraints: [Text] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-placeholder
// [textarea]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-placeholder
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func PlaceHolder(placeHolder string) roundly.Attribute {
	return roundly.WriteAttribute("placeholder", placeHolder)
}

// PopOverTarget creates the popovertarget attribute - Targets a popover element to toggle, show, or hide
//
// It can be applied to the following elements:
//   - [button] [input]
//   - [input]
//
// Value constraints: [ID] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/popover.html#attr-popovertarget
// [input]: https://html.spec.whatwg.org/dev/popover.html#attr-popovertarget
// [ID]: https://html.spec.whatwg.org/dev/https://dom.spec.whatwg.org/#concept-id
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func PopOverTarget(popOverTarget string) roundly.Attribute {
	return roundly.WriteAttribute("popovertarget", popOverTarget)
}

// Poster creates the poster attribute - Poster frame to show prior to video playback
//
// It can be applied to the following elements:
//   - [video]
//
// Value constraints: [Valid non-empty URL potentially surrounded by spaces]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [video]: https://html.spec.whatwg.org/dev/media.html#attr-video-poster
// [Valid non-empty URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Poster(poster string) roundly.Attribute {
	return roundly.WriteAttribute("poster", poster)
}

// ReferrerPolicy creates the referrerpolicy attribute - [Referrer policy] for [fetches] initiated by the element
//
// It can be applied to the following elements:
//   - [a] [area] [iframe] [img] [link] [script]
//   - [area]
//   - [iframe]
//   - [img]
//   - [link]
//   - [script]
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
func ReferrerPolicy(referrerPolicy string) roundly.Attribute {
	return roundly.WriteAttribute("referrerpolicy", referrerPolicy)
}

// Sizes creates the sizes attribute - Image sizes for different page layouts
//
// It can be applied to the following elements:
//   - [img] [source]
//   - [source]
//
// Value constraints: [Valid source size list]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [img]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-sizes
// [source]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-source-sizes
// [Valid source size list]: https://html.spec.whatwg.org/dev/images.html#valid-source-size-list
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Sizes(sizes string) roundly.Attribute {
	return roundly.WriteAttribute("sizes", sizes)
}

// SlotAttribute creates the slot attribute - The element's desired slot
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Value constraints: [Text]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#the-id-attribute
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func SlotAttribute(slot string) roundly.Attribute {
	return roundly.WriteAttribute("slot", slot)
}

// Src creates the src attribute - Address of the resource
//
// It can be applied to the following elements:
//   - [audio] [embed] [iframe] [img] [input] [script] [source] (in [video] or [audio (1)]) [track] [video (1)]
//   - [embed]
//   - [iframe]
//   - [img]
//   - [input]
//   - [script]
//   - [source] (in [video] or [audio (1)])
//   - [track]
//   - [video (1)]
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
// [audio (1)]: https://html.spec.whatwg.org/dev/media.html#the-audio-element
// [track]: https://html.spec.whatwg.org/dev/media.html#attr-track-src
// [video (1)]: https://html.spec.whatwg.org/dev/media.html#attr-media-src
// [Valid non-empty URL potentially surrounded by spaces]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Src(src string) roundly.Attribute {
	return roundly.WriteAttribute("src", src)
}

// SrcDoc creates the srcdoc attribute - A document to render in the [iframe]
//
// It can be applied to the following elements:
//   - [iframe]
//
// Value constraints: The source of [an iframesrcdoc document] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-srcdoc
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#the-iframe-element
// [an iframesrcdoc document]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#an-iframe-srcdoc-document
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func SrcDoc(srcDoc string) roundly.Attribute {
	return roundly.WriteAttribute("srcdoc", srcDoc)
}

// SrcLang creates the srclang attribute - Language of the text track
//
// It can be applied to the following elements:
//   - [track]
//
// Value constraints: Valid BCP 47 language tag
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [track]: https://html.spec.whatwg.org/dev/media.html#attr-track-srclang
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func SrcLang(srcLang string) roundly.Attribute {
	return roundly.WriteAttribute("srclang", srcLang)
}

// StyleAttribute creates the style attribute - Presentational and formatting instructions
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Value constraints: CSS declarations (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#attr-style
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func StyleAttribute(style string) roundly.Attribute {
	return roundly.WriteAttribute("style", style)
}

// Target creates the target attribute - It's semantics varies depending on the element it is applied to.
//
// It can be applied to the following elements:
//   - [a] [area]: [Navigable] for [hyperlink][navigation]
//   - [base]: Default [navigable] for [hyperlink][navigation] and [form submission]
//   - [form]: [Navigable] for [form submission]
//
// Value constraints: 
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-target
// [area]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-target
// [Navigable]: https://html.spec.whatwg.org/dev/document-sequences.html#navigable
// [hyperlink]: https://html.spec.whatwg.org/dev/links.html#hyperlink
// [navigation]: https://html.spec.whatwg.org/dev/browsing-the-web.html#navigate
// [Valid navigable target name or keyword]: https://html.spec.whatwg.org/dev/document-sequences.html#valid-navigable-target-name-or-keyword
// [base]: https://html.spec.whatwg.org/dev/semantics.html#attr-base-target
// [navigable]: https://html.spec.whatwg.org/dev/document-sequences.html#navigable
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [form]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-target
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Target(target string) roundly.Attribute {
	return roundly.WriteAttribute("target", target)
}

// TitleAttribute creates the title attribute - It's semantics varies depending on the element it is applied to.
//
// It can be applied to the following elements:
//   - [HTML elements]: Advisory information for the element
//   - [abbr] [dfn]: Full term or expansion of abbreviation
//   - [input]: Description of pattern (when used with [pattern] attribute)
//   - [link]: Title of the link
//   - [link] [style]: [CSS style sheet set name]
//
// Value constraints: 
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#attr-title
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [abbr]: https://html.spec.whatwg.org/dev/text-level-semantics.html#attr-abbr-title
// [dfn]: https://html.spec.whatwg.org/dev/text-level-semantics.html#attr-dfn-title
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-title
// [pattern]: https://html.spec.whatwg.org/dev/input.html#attr-input-pattern
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-title
// [style]: https://html.spec.whatwg.org/dev/semantics.html#attr-style-title
// [CSS style sheet set name]: https://html.spec.whatwg.org/dev/https://drafts.csswg.org/cssom/#css-style-sheet-set-name
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func TitleAttribute(title string) roundly.Attribute {
	return roundly.WriteAttribute("title", title)
}

// Type creates the type attribute - It's semantics varies depending on the element it is applied to.
//
// It can be applied to the following elements:
//   - [a] [link]: Hint for the type of the referenced resource
//   - [embed] [object] [source]: Type of embedded resource
//   - [script]: Type of script
//
// Value constraints: 
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [a]: https://html.spec.whatwg.org/dev/links.html#attr-hyperlink-type
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-type
// [Valid MIME type string]: https://html.spec.whatwg.org/dev/https://mimesniff.spec.whatwg.org/#valid-mime-type
// [embed]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-embed-type
// [object]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-object-type
// [source]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-source-type
// [script]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-type
// [valid MIME type string]: https://html.spec.whatwg.org/dev/https://mimesniff.spec.whatwg.org/#valid-mime-type
// [JavaScript MIME type essence match]: https://html.spec.whatwg.org/dev/https://mimesniff.spec.whatwg.org/#javascript-mime-type-essence-match
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Type(typeV string) roundly.Attribute {
	return roundly.WriteAttribute("type", typeV)
}

// UseMap creates the usemap attribute - Name of [image map] to use
//
// It can be applied to the following elements:
//   - [img]
//
// Value constraints: [Valid hash-name reference] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [img]: https://html.spec.whatwg.org/dev/image-maps.html#attr-hyperlink-usemap
// [image map]: https://html.spec.whatwg.org/dev/image-maps.html#image-map
// [Valid hash-name reference]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-hash-name-reference
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func UseMap(useMap string) roundly.Attribute {
	return roundly.WriteAttribute("usemap", useMap)
}

// Value creates the value attribute - It's semantics varies depending on the element it is applied to.
//
// It can be applied to the following elements:
//   - [button] [option]: Value to be used for [form submission]
//   - [data]: Machine-readable value
//   - [input]: Value of the form control
//
// Value constraints: 
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-value
// [option]: https://html.spec.whatwg.org/dev/form-elements.html#attr-option-value
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [Text]: https://html.spec.whatwg.org/dev/dom.html#attribute-text
// [data]: https://html.spec.whatwg.org/dev/text-level-semantics.html#attr-data-value
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-value
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Value(value string) roundly.Attribute {
	return roundly.WriteAttribute("value", value)
}
