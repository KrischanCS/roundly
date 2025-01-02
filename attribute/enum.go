//nolint:nolintlint,lll
package attribute

import (
	"github.com/ch-schulz/htmfunc"
)

// AutoCapitalize creates the autocapitalize attribute - Recommended autocapitalization behavior (for supported input methods)
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Can hold one of the following values:
//   - "[on]"
//   - "[off]"
//   - "[none]"
//   - "[sentences]"
//   - "[words]"
//   - "[characters]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/interaction.html#attr-autocapitalize
// [on]: https://html.spec.whatwg.org/dev/interaction.html#attr-autocapitalize-on
// [off]: https://html.spec.whatwg.org/dev/interaction.html#attr-autocapitalize-off
// [none]: https://html.spec.whatwg.org/dev/interaction.html#attr-autocapitalize-none
// [sentences]: https://html.spec.whatwg.org/dev/interaction.html#attr-autocapitalize-sentences
// [words]: https://html.spec.whatwg.org/dev/interaction.html#attr-autocapitalize-words
// [characters]: https://html.spec.whatwg.org/dev/interaction.html#attr-autocapitalize-characters
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func AutoCapitalize(autoCapitalize string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("autocapitalize", autoCapitalize)
}

// AutoComplete_Form creates the autocomplete attribute - Default setting for autofill feature for controls in the form
//
// It can be applied to the following elements:
//   - [form]
//
// Can hold one of the following values:
//   - "on"
//   - "off"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [form]: https://html.spec.whatwg.org/dev/forms.html#attr-form-autocomplete
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func AutoComplete_Form(autoComplete string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("autocomplete", autoComplete)
}

// AutoCorrect creates the autocorrect attribute - Recommended autocorrection behavior (for supported input methods)
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Can hold one of the following values:
//   - "[on]"
//   - "[off]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/interaction.html#attr-autocorrect
// [on]: https://html.spec.whatwg.org/dev/interaction.html#attr-autocorrect-on
// [off]: https://html.spec.whatwg.org/dev/interaction.html#attr-autocorrect-off
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func AutoCorrect(autoCorrect string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("autocorrect", autoCorrect)
}

// CharSet creates the charset attribute - [Character encoding declaration]
//
// It can be applied to the following elements:
//   - [meta]
//
// Can hold one of the following values:
//   - "utf-8"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meta]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-charset
// [Character encoding declaration]: https://html.spec.whatwg.org/dev/semantics.html#character-encoding-declaration
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func CharSet(charSet string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("charset", charSet)
}

// ColorSpace creates the colorspace attribute - The color space of the serialized color
//
// It can be applied to the following elements:
//   - [input]
//
// Can hold one of the following values:
//   - "[limited-srgb]"
//   - "[display-p3]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-colorspace
// [limited-srgb]: https://html.spec.whatwg.org/dev/input.html#attr-input-colorspace-limited-srgb
// [display-p3]: https://html.spec.whatwg.org/dev/input.html#attr-input-colorspace-display-p3
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ColorSpace(colorSpace string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("colorspace", colorSpace)
}

// ContentEditable creates the contenteditable attribute - Whether the element is editable
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Can hold one of the following values:
//   - "true"
//   - "plaintext-only"
//   - "false"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/interaction.html#attr-contenteditable
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ContentEditable(contentEditable string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("contenteditable", contentEditable)
}

// CrossOrigin creates the crossorigin attribute - How the element handles crossorigin requests
//
// It can be applied to the following elements:
//   - [audio]
//   - [img]
//   - [link]
//   - [script]
//   - [video]
//
// Can hold one of the following values:
//   - "[anonymous]"
//   - "[use-credentials]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [audio]: https://html.spec.whatwg.org/dev/media.html#attr-media-crossorigin
// [img]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-crossorigin
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-crossorigin
// [script]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-crossorigin
// [video]: https://html.spec.whatwg.org/dev/media.html#attr-media-crossorigin
// [anonymous]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#attr-crossorigin-anonymous-keyword
// [use-credentials]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#attr-crossorigin-use-credentials-keyword
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func CrossOrigin(crossOrigin string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("crossorigin", crossOrigin)
}

// Decoding creates the decoding attribute - Decoding hint to use when processing this image for presentation
//
// It can be applied to the following elements:
//   - [img]
//
// Can hold one of the following values:
//   - "sync"
//   - "async"
//   - "auto"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [img]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-decoding
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Decoding(decoding string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("decoding", decoding)
}

// Dir creates the dir attribute - [The text directionality] of the element
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Can hold one of the following values:
//   - "[ltr]"
//   - "[rtl]"
//   - "[auto]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#attr-dir
// [The text directionality]: https://html.spec.whatwg.org/dev/dom.html#the-directionality
// [ltr]: https://html.spec.whatwg.org/dev/dom.html#attr-dir-ltr
// [rtl]: https://html.spec.whatwg.org/dev/dom.html#attr-dir-rtl
// [auto]: https://html.spec.whatwg.org/dev/dom.html#attr-dir-auto
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Dir(dir string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("dir", dir)
}

// Dir_Bdo creates the dir attribute - [The text directionality] of the element
//
// It can be applied to the following elements:
//   - [bdo]
//
// Can hold one of the following values:
//   - "[ltr]"
//   - "[rtl]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [bdo]: https://html.spec.whatwg.org/dev/text-level-semantics.html#the-bdo-element
// [The text directionality]: https://html.spec.whatwg.org/dev/dom.html#the-directionality
// [ltr]: https://html.spec.whatwg.org/dev/dom.html#attr-dir-ltr
// [rtl]: https://html.spec.whatwg.org/dev/dom.html#attr-dir-rtl
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Dir_Bdo(dir string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("dir", dir)
}

// Draggable creates the draggable attribute - Whether the element is draggable
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Can hold one of the following values:
//   - "true"
//   - "false"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dnd.html#attr-draggable
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Draggable(draggable string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("draggable", draggable)
}

// EncType creates the enctype attribute - Entry list encoding type to use for [form submission]
//
// It can be applied to the following elements:
//   - [form]
//
// Can hold one of the following values:
//   - "[application/x-www-form-urlencoded]"
//   - "[multipart/form-data]"
//   - "[text/plain]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [form]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-enctype
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [application/x-www-form-urlencoded]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-enctype-urlencoded
// [multipart/form-data]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-enctype-formdata
// [text/plain]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-enctype-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func EncType(encType string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("enctype", encType)
}

// EnterKeyHint creates the enterkeyhint attribute - Hint for selecting an enter key action
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Can hold one of the following values:
//   - "[enter]"
//   - "[done]"
//   - "[go]"
//   - "[next]"
//   - "[previous]"
//   - "[search]"
//   - "[send]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/interaction.html#attr-enterkeyhint
// [enter]: https://html.spec.whatwg.org/dev/interaction.html#attr-enterkeyhint-keyword-enter
// [done]: https://html.spec.whatwg.org/dev/interaction.html#attr-enterkeyhint-keyword-done
// [go]: https://html.spec.whatwg.org/dev/interaction.html#attr-enterkeyhint-keyword-go
// [next]: https://html.spec.whatwg.org/dev/interaction.html#attr-enterkeyhint-keyword-next
// [previous]: https://html.spec.whatwg.org/dev/interaction.html#attr-enterkeyhint-keyword-previous
// [search]: https://html.spec.whatwg.org/dev/interaction.html#attr-enterkeyhint-keyword-search
// [send]: https://html.spec.whatwg.org/dev/interaction.html#attr-enterkeyhint-keyword-send
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func EnterKeyHint(enterKeyHint string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("enterkeyhint", enterKeyHint)
}

// FetchPriority creates the fetchpriority attribute - Sets the [priority] for [fetches] initiated by the element
//
// It can be applied to the following elements:
//   - [img]
//   - [link]
//   - [script]
//
// Can hold one of the following values:
//   - "[auto]"
//   - "[high]"
//   - "[low]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [img]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-fetchpriority
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-fetchpriority
// [script]: https://html.spec.whatwg.org/dev/scripting.html#attr-script-fetchpriority
// [priority]: https://html.spec.whatwg.org/dev/https://fetch.spec.whatwg.org/#request-priority
// [fetches]: https://html.spec.whatwg.org/dev/https://fetch.spec.whatwg.org/#concept-fetch
// [auto]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#attr-fetchpriority-auto
// [high]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#attr-fetchpriority-high
// [low]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#attr-fetchpriority-low
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func FetchPriority(fetchPriority string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("fetchpriority", fetchPriority)
}

// FormEnctype creates the formenctype attribute - Entry list encoding type to use for [form submission]
//
// It can be applied to the following elements:
//   - [button]
//   - [input]
//
// Can hold one of the following values:
//   - "[application/x-www-form-urlencoded]"
//   - "[multipart/form-data]"
//   - "[text/plain]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formenctype
// [input]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formenctype
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [application/x-www-form-urlencoded]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-enctype-urlencoded
// [multipart/form-data]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-enctype-formdata
// [text/plain]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-enctype-text
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func FormEnctype(formEnctype string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("formenctype", formEnctype)
}

// FormMethod creates the formmethod attribute - Variant to use for [form submission]
//
// It can be applied to the following elements:
//   - [button]
//   - [input]
//
// Can hold one of the following values:
//   - "GET"
//   - "POST"
//   - "dialog"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formmethod
// [input]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-formmethod
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func FormMethod(formMethod string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("formmethod", formMethod)
}

// Hidden creates the hidden attribute - Whether the element is relevant
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Can hold one of the following values:
//   - "[until-found]"
//   - "[hidden]"
//   - the empty string
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/interaction.html#attr-hidden
// [until-found]: https://html.spec.whatwg.org/dev/interaction.html#attr-hidden-until-found
// [hidden]: https://html.spec.whatwg.org/dev/interaction.html#attr-hidden
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Hidden(hidden string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("hidden", hidden)
}

// HttpEquiv creates the http-equiv attribute - Pragma directive
//
// It can be applied to the following elements:
//   - [meta]
//
// Can hold one of the following values:
//   - "[content-type]"
//   - "[default-style]"
//   - "[refresh]"
//   - "[x-ua-compatible]"
//   - "[content-security-policy]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meta]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-http-equiv
// [content-type]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-http-equiv-keyword-content-type
// [default-style]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-http-equiv-keyword-default-style
// [refresh]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-http-equiv-keyword-refresh
// [x-ua-compatible]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-http-equiv-keyword-x-ua-compatible
// [content-security-policy]: https://html.spec.whatwg.org/dev/semantics.html#attr-meta-http-equiv-keyword-content-security-policy
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func HttpEquiv(httpEquiv string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("http-equiv", httpEquiv)
}

// InputMode creates the inputmode attribute - Hint for selecting an input modality
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Can hold one of the following values:
//   - "[none]"
//   - "[text]"
//   - "[tel]"
//   - "[email]"
//   - "[url]"
//   - "[numeric]"
//   - "[decimal]"
//   - "[search]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/interaction.html#attr-inputmode
// [none]: https://html.spec.whatwg.org/dev/interaction.html#attr-inputmode-keyword-none
// [text]: https://html.spec.whatwg.org/dev/interaction.html#attr-inputmode-keyword-text
// [tel]: https://html.spec.whatwg.org/dev/interaction.html#attr-inputmode-keyword-tel
// [email]: https://html.spec.whatwg.org/dev/interaction.html#attr-inputmode-keyword-email
// [url]: https://html.spec.whatwg.org/dev/interaction.html#attr-inputmode-keyword-url
// [numeric]: https://html.spec.whatwg.org/dev/interaction.html#attr-inputmode-keyword-numeric
// [decimal]: https://html.spec.whatwg.org/dev/interaction.html#attr-inputmode-keyword-decimal
// [search]: https://html.spec.whatwg.org/dev/interaction.html#attr-inputmode-keyword-search
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func InputMode(inputMode string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("inputmode", inputMode)
}

// Kind creates the kind attribute - The type of text track
//
// It can be applied to the following elements:
//   - [track]
//
// Can hold one of the following values:
//   - "[subtitles]"
//   - "[captions]"
//   - "[descriptions]"
//   - "[chapters]"
//   - "[metadata]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [track]: https://html.spec.whatwg.org/dev/media.html#attr-track-kind
// [subtitles]: https://html.spec.whatwg.org/dev/media.html#attr-track-kind-keyword-subtitles
// [captions]: https://html.spec.whatwg.org/dev/media.html#attr-track-kind-keyword-captions
// [descriptions]: https://html.spec.whatwg.org/dev/media.html#attr-track-kind-keyword-descriptions
// [chapters]: https://html.spec.whatwg.org/dev/media.html#attr-track-kind-keyword-chapters
// [metadata]: https://html.spec.whatwg.org/dev/media.html#attr-track-kind-keyword-metadata
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Kind(kind string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("kind", kind)
}

// Loading creates the loading attribute - Used when determining loading deferral
//
// It can be applied to the following elements:
//   - [iframe]
//   - [img]
//
// Can hold one of the following values:
//   - "[lazy]"
//   - "[eager]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [iframe]: https://html.spec.whatwg.org/dev/iframe-embed-object.html#attr-iframe-loading
// [img]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-loading
// [lazy]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#attr-loading-lazy
// [eager]: https://html.spec.whatwg.org/dev/urls-and-fetching.html#attr-loading-eager
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Loading(loading string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("loading", loading)
}

// Method creates the method attribute - Variant to use for [form submission]
//
// It can be applied to the following elements:
//   - [form]
//
// Can hold one of the following values:
//   - "[GET]"
//   - "[POST]"
//   - "[dialog]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [form]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-method
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [GET]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-method-get-keyword
// [POST]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-method-post-keyword
// [dialog]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#attr-fs-method-dialog-keyword
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Method(method string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("method", method)
}

// PopOver creates the popover attribute - Makes the element a [popover] element
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Can hold one of the following values:
//   - "[auto]"
//   - "[manual]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/popover.html#attr-popover
// [popover]: https://html.spec.whatwg.org/dev/popover.html#attr-popover
// [auto]: https://html.spec.whatwg.org/dev/popover.html#attr-popover-auto
// [manual]: https://html.spec.whatwg.org/dev/popover.html#attr-popover-manual
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func PopOver(popOver string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("popover", popOver)
}

// PopOverTargetAction creates the popovertargetaction attribute - Indicates whether a targeted popover element is to be toggled, shown, or hidden
//
// It can be applied to the following elements:
//   - [button]
//   - [input]
//
// Can hold one of the following values:
//   - "[toggle]"
//   - "[show]"
//   - "[hide]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/popover.html#attr-popovertargetaction
// [input]: https://html.spec.whatwg.org/dev/popover.html#attr-popovertargetaction
// [toggle]: https://html.spec.whatwg.org/dev/popover.html#attr-popovertargetaction-toggle
// [show]: https://html.spec.whatwg.org/dev/popover.html#attr-popovertargetaction-show
// [hide]: https://html.spec.whatwg.org/dev/popover.html#attr-popovertargetaction-hide
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func PopOverTargetAction(popOverTargetAction string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("popovertargetaction", popOverTargetAction)
}

// PreLoad creates the preload attribute - Hints how much buffering the [media resource] will likely need
//
// It can be applied to the following elements:
//   - [audio]
//   - [video]
//
// Can hold one of the following values:
//   - "[none]"
//   - "[metadata]"
//   - "[auto]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [audio]: https://html.spec.whatwg.org/dev/media.html#attr-media-preload
// [video]: https://html.spec.whatwg.org/dev/media.html#attr-media-preload
// [media resource]: https://html.spec.whatwg.org/dev/media.html#media-resource
// [none]: https://html.spec.whatwg.org/dev/media.html#attr-media-preload-none
// [metadata]: https://html.spec.whatwg.org/dev/media.html#attr-media-preload-metadata
// [auto]: https://html.spec.whatwg.org/dev/media.html#attr-media-preload-auto
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func PreLoad(preLoad string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("preload", preLoad)
}

// Scope creates the scope attribute - Specifies which cells the header cell applies to
//
// It can be applied to the following elements:
//   - [th]
//
// Can hold one of the following values:
//   - "[row]"
//   - "[col]"
//   - "[rowgroup]"
//   - "[colgroup]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [th]: https://html.spec.whatwg.org/dev/tables.html#attr-th-scope
// [row]: https://html.spec.whatwg.org/dev/tables.html#attr-th-scope-row
// [col]: https://html.spec.whatwg.org/dev/tables.html#attr-th-scope-col
// [rowgroup]: https://html.spec.whatwg.org/dev/tables.html#attr-th-scope-rowgroup
// [colgroup]: https://html.spec.whatwg.org/dev/tables.html#attr-th-scope-colgroup
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Scope(scope string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("scope", scope)
}

// ShadowRootMode creates the shadowrootmode attribute - Enables streaming declarative shadow roots
//
// It can be applied to the following elements:
//   - [template]
//
// Can hold one of the following values:
//   - "open"
//   - "closed"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [template]: https://html.spec.whatwg.org/dev/scripting.html#attr-template-shadowrootmode
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ShadowRootMode(shadowRootMode string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("shadowrootmode", shadowRootMode)
}

// Shape creates the shape attribute - The kind of shape to be created in an [image map]
//
// It can be applied to the following elements:
//   - [area]
//
// Can hold one of the following values:
//   - "[circle]"
//   - "[default]"
//   - "[poly]"
//   - "[rect]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [area]: https://html.spec.whatwg.org/dev/image-maps.html#attr-area-shape
// [image map]: https://html.spec.whatwg.org/dev/image-maps.html#image-map
// [circle]: https://html.spec.whatwg.org/dev/image-maps.html#attr-area-shape-keyword-circle
// [default]: https://html.spec.whatwg.org/dev/image-maps.html#attr-area-shape-keyword-default
// [poly]: https://html.spec.whatwg.org/dev/image-maps.html#attr-area-shape-keyword-poly
// [rect]: https://html.spec.whatwg.org/dev/image-maps.html#attr-area-shape-keyword-rect
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Shape(shape string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("shape", shape)
}

// SpellCheck creates the spellcheck attribute - Whether the element is to have its spelling and grammar checked
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Can hold one of the following values:
//   - "[true]"
//   - "[false]"
//   - the empty string
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/interaction.html#attr-spellcheck
// [true]: https://html.spec.whatwg.org/dev/interaction.html#attr-spellcheck-true
// [false]: https://html.spec.whatwg.org/dev/interaction.html#attr-spellcheck-false
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func SpellCheck(spellCheck string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("spellcheck", spellCheck)
}

// Translate creates the translate attribute - Whether the element is to be translated when the page is localized
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Can hold one of the following values:
//   - "yes"
//   - "no"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/dom.html#attr-translate
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Translate(translate string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("translate", translate)
}

// Type_Button creates the type attribute - Type of button
//
// It can be applied to the following elements:
//   - [button]
//
// Can hold one of the following values:
//   - "[submit]"
//   - "[reset]"
//   - "[button]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [button]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-type
// [submit]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-type-submit
// [reset]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-type-reset
// [button]: https://html.spec.whatwg.org/dev/form-elements.html#attr-button-type-button
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Type_Button(typeV string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("type", typeV)
}

// Type_Ol creates the type attribute - Kind of list marker
//
// It can be applied to the following elements:
//   - [ol]
//
// Can hold one of the following values:
//   - "[1]"
//   - "[a]"
//   - "[A]"
//   - "[i]"
//   - "[I]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [ol]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-ol-type
// [1]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-ol-type-keyword-decimal
// [a]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-ol-type-keyword-lower-alpha
// [A]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-ol-type-keyword-upper-alpha
// [i]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-ol-type-keyword-lower-roman
// [I]: https://html.spec.whatwg.org/dev/grouping-content.html#attr-ol-type-keyword-upper-roman
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Type_Ol(typeV string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("type", typeV)
}

// Wrap creates the wrap attribute - How the value of the form control is to be wrapped for [form submission]
//
// It can be applied to the following elements:
//   - [textarea]
//
// Can hold one of the following values:
//   - "[soft]"
//   - "[hard]"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [textarea]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-wrap
// [form submission]: https://html.spec.whatwg.org/dev/form-control-infrastructure.html#form-submission-2
// [soft]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-wrap-soft
// [hard]: https://html.spec.whatwg.org/dev/form-elements.html#attr-textarea-wrap-hard
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Wrap(wrap string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("wrap", wrap)
}

// WritingSuggestions creates the writingsuggestions attribute - Whether the element can offer writing suggestions or not.
//
// It can be applied to the following elements:
//   - [HTML elements]
//
// Can hold one of the following values:
//   - "[true]"
//   - "[false]"
//   - the empty string
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [HTML elements]: https://html.spec.whatwg.org/dev/interaction.html#attr-writingsuggestions
// [true]: https://html.spec.whatwg.org/dev/interaction.html#attr-writingsuggestions-true
// [false]: https://html.spec.whatwg.org/dev/interaction.html#attr-writingsuggestions-false
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func WritingSuggestions(writingSuggestions string) htmfunc.Attribute {
	return htmfunc.WriteAttribute("writingsuggestions", writingSuggestions)
}
