package el

import (
	"github.com/ch-schulz/htmfunc"
	"github.com/ch-schulz/htmfunc/attr"
)

// Comment creates a [comment element].
//
// [comment element]: https://html.spec.whatwg.org/#the-comment-element
func Comment(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("!--", attributes, children...)
}

// Area creates a [area element].
//
// [area element]: https://html.spec.whatwg.org/#the-area-element
func Area(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("area", attributes, children...)
}

// Audio creates a [audio element].
//
// [audio element]: https://html.spec.whatwg.org/#the-audio-element
func Audio(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("audio", attributes, children...)
}

// Button creates a [button element].
//
// [button element]: https://html.spec.whatwg.org/#the-button-element
func Button(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("button", attributes, children...)
}

// Canvas creates a [canvas element].
//
// [canvas element]: https://html.spec.whatwg.org/#the-canvas-element
func Canvas(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("canvas", attributes, children...)
}

// Caption creates a [caption element].
//
// [caption element]: https://html.spec.whatwg.org/#the-caption-element
func Caption(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("caption", attributes, children...)
}

// Col creates a [col element].
//
// [col element]: https://html.spec.whatwg.org/#the-col-element
func Col(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("col", attributes, children...)
}

// Colgroup creates a [colgroup element].
//
// [colgroup element]: https://html.spec.whatwg.org/#the-colgroup-element
func Colgroup(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("colgroup", attributes, children...)
}

// Datalist creates a [datalist element].
//
// [datalist element]: https://html.spec.whatwg.org/#the-datalist-element
func Datalist(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("datalist", attributes, children...)
}

// Del creates a [del element].
//
// [del element]: https://html.spec.whatwg.org/#the-del-element
func Del(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("del", attributes, children...)
}

// Details creates a [details element].
//
// [details element]: https://html.spec.whatwg.org/#the-details-element
func Details(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("details", attributes, children...)
}

// Dialog creates a [dialog element].
//
// [dialog element]: https://html.spec.whatwg.org/#the-dialog-element
func Dialog(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("dialog", attributes, children...)
}

// Embed creates a [embed element].
//
// [embed element]: https://html.spec.whatwg.org/#the-embed-element
func Embed(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("embed", attributes, children...)
}

// Fieldset creates a [fieldset element].
//
// [fieldset element]: https://html.spec.whatwg.org/#the-fieldset-element
func Fieldset(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("fieldset", attributes, children...)
}

// Form creates a [form element].
//
// [form element]: https://html.spec.whatwg.org/#the-form-element
func Form(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("form", attributes, children...)
}

// Iframe creates a [iframe element].
//
// [iframe element]: https://html.spec.whatwg.org/#the-iframe-element
func Iframe(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("iframe", attributes, children...)
}

// Img creates a [img element].
//
// [img element]: https://html.spec.whatwg.org/#the-img-element
func Img(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("img", attributes, children...)
}

// Input creates a [input element].
//
// [input element]: https://html.spec.whatwg.org/#the-input-element
func Input(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("input", attributes, children...)
}

// Ins creates a [ins element].
//
// [ins element]: https://html.spec.whatwg.org/#the-ins-element
func Ins(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("ins", attributes, children...)
}

// Label creates a [label element].
//
// [label element]: https://html.spec.whatwg.org/#the-label-element
func Label(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("label", attributes, children...)
}

// Legend creates a [legend element].
//
// [legend element]: https://html.spec.whatwg.org/#the-legend-element
func Legend(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("legend", attributes, children...)
}

// Map creates a [map element].
//
// [map element]: https://html.spec.whatwg.org/#the-map-element
func Map(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("map", attributes, children...)
}

// Meter creates a [meter element].
//
// [meter element]: https://html.spec.whatwg.org/#the-meter-element
func Meter(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("meter", attributes, children...)
}

// Noscript creates a [noscript element].
//
// [noscript element]: https://html.spec.whatwg.org/#the-noscript-element
func Noscript(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("noscript", attributes, children...)
}

// Object creates a [object element].
//
// [object element]: https://html.spec.whatwg.org/#the-object-element
func Object(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("object", attributes, children...)
}

// Optgroup creates a [optgroup element].
//
// [optgroup element]: https://html.spec.whatwg.org/#the-optgroup-element
func Optgroup(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("optgroup", attributes, children...)
}

// Option creates a [option element].
//
// [option element]: https://html.spec.whatwg.org/#the-option-element
func Option(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("option", attributes, children...)
}

// Output creates a [output element].
//
// [output element]: https://html.spec.whatwg.org/#the-output-element
func Output(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("output", attributes, children...)
}

// Param creates a [param element].
//
// [param element]: https://html.spec.whatwg.org/#the-param-element
func Param(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("param", attributes, children...)
}

// Picture creates a [picture element].
//
// [picture element]: https://html.spec.whatwg.org/#the-picture-element
func Picture(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("picture", attributes, children...)
}

// Progress creates a [progress element]
//
// [progress element]: https://html.spec.whatwg.org/#the-progress-element
func Progress(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("progress", attributes, children...)
}

// Script creates a [script element].
//
// [script element]: https://html.spec.whatwg.org/#the-script-element
func Script(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("script", attributes, children...)
}

// Select creates a [select element].
//
// [select element]: https://html.spec.whatwg.org/#the-select-element
func Select(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("select", attributes, children...)
}

// Source creates a [source element].
//
// [source element]: https://html.spec.whatwg.org/#the-source-element
func Source(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("source", attributes, children...)
}

// Summary creates a [summary element].
//
// [summary element]: https://html.spec.whatwg.org/#the-summary-element
func Summary(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("summary", attributes, children...)
}

// Svg creates a [svg element].
//
// [svg element]: https://html.spec.whatwg.org/#the-svg-element
func Svg(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("svg", attributes, children...)
}

// Table creates a [table element].
//
// [table element]: https://html.spec.whatwg.org/#the-table-element
func Table(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("table", attributes, children...)
}

// Tbody creates a [tbody element].
//
// [tbody element]: https://html.spec.whatwg.org/#the-tbody-element
func Tbody(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("tbody", attributes, children...)
}

// Td creates a [td element].
//
// [td element]: https://html.spec.whatwg.org/#the-td-element
func Td(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("td", attributes, children...)
}

// Template creates a [template element].
//
// [template element]: https://html.spec.whatwg.org/#the-template-element
func Template(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("template", attributes, children...)
}

// Textarea creates a [textarea element].
//
// [textarea element]: https://html.spec.whatwg.org/#the-textarea-element
func Textarea(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("textarea", attributes, children...)
}

// Tfoot creates a [tfoot element].
//
// [tfoot element]: https://html.spec.whatwg.org/#the-tfoot-element
func Tfoot(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("tfoot", attributes, children...)
}

// Th creates a [th element].
//
// [th element]: https://html.spec.whatwg.org/#the-th-element
func Th(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("th", attributes, children...)
}

// Thead creates a [thead element].
//
// [thead element]: https://html.spec.whatwg.org/#the-thead-element
func Thead(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("thead", attributes, children...)
}

// Tr creates a [tr element].
//
// [tr element]: https://html.spec.whatwg.org/#the-tr-element
func Tr(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("tr", attributes, children...)
}

// Track creates a [track element].
//
// [track element]: https://html.spec.whatwg.org/#the-track-element
func Track(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("track", attributes, children...)
}

// Video creates a [video element].
//
// [video element]: https://html.spec.whatwg.org/#the-video-element
func Video(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("video", attributes, children...)
}
