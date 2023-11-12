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

// A creates a [a element].
//
// [a element]: https://html.spec.whatwg.org/#the-a-element
func A(href string, attributes attr.Ls, text string) htmfunc.Component {
	attributes = append([]htmfunc.Attribute{attr.HRef(href)}, attributes...)
	return htmfunc.Element("a", attributes, Text(text))
}

// Abbr creates a [abbr element].
//
// [abbr element]: https://html.spec.whatwg.org/#the-abbr-element
func Abbr(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("abbr", attributes, children...)
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

// B creates a [b element].
//
// [b element]: https://html.spec.whatwg.org/#the-b-element
func B(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("b", attributes, children...)
}

// Bdi creates a [bdi element].
//
// [bdi element]: https://html.spec.whatwg.org/#the-bdi-element
func Bdi(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("bdi", attributes, children...)
}

// Bdo creates a [bdo element].
//
// [bdo element]: https://html.spec.whatwg.org/#the-bdo-element
func Bdo(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("bdo", attributes, children...)
}

// Blockquote creates a [blockquote element].
//
// [blockquote element]: https://html.spec.whatwg.org/#the-blockquote-element
func Blockquote(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("blockquote", attributes, children...)
}

// Br creates a [br element].
//
// [br element]: https://html.spec.whatwg.org/#the-br-element
func Br(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("br", attributes, children...)
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

// Cite creates a [cite element].
//
// [cite element]: https://html.spec.whatwg.org/#the-cite-element
func Cite(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("cite", attributes, children...)
}

// Code creates a [code element].
//
// [code element]: https://html.spec.whatwg.org/#the-code-element
func Code(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("code", attributes, children...)
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

// Data creates a [data element].
//
// [data element]: https://html.spec.whatwg.org/#the-data-element
func Data(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("data", attributes, children...)
}

// Datalist creates a [datalist element].
//
// [datalist element]: https://html.spec.whatwg.org/#the-datalist-element
func Datalist(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("datalist", attributes, children...)
}

// Dd creates a [dd element].
//
// [dd element]: https://html.spec.whatwg.org/#the-dd-element
func Dd(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("dd", attributes, children...)
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

// Dfn creates a [dfn element]
//
// [dfn element]: https://html.spec.whatwg.org/#the-dfn-element
func Dfn(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("dfn", attributes, children...)
}

// Dialog creates a [dialog element].
//
// [dialog element]: https://html.spec.whatwg.org/#the-dialog-element
func Dialog(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("dialog", attributes, children...)
}

// Div creates a [div element].
//
// [div element]: https://html.spec.whatwg.org/#the-div-element
func Div(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("div", attributes, children...)
}

// Dl creates a [dl element].
//
// [dl element]: https://html.spec.whatwg.org/#the-dl-element
func Dl(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("dl", attributes, children...)
}

// Dt creates a [dt element].
//
// [dt element]: https://html.spec.whatwg.org/#the-dt-element
func Dt(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("dt", attributes, children...)
}

// Em creates a [em element].
//
// [em element]: https://html.spec.whatwg.org/#the-em-element
func Em(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("em", attributes, children...)
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

// Figcaption creates a [figcaption element].
//
// [figcaption element]: https://html.spec.whatwg.org/#the-figcaption-element
func Figcaption(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("figcaption", attributes, children...)
}

// Figure creates a [figure element]
//
// [figure element]: https://html.spec.whatwg.org/#the-figu-element
func Figure(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("figure", attributes, children...)
}

// Form creates a [form element].
//
// [form element]: https://html.spec.whatwg.org/#the-form-element
func Form(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("form", attributes, children...)
}

// Hr creates a [hr element].
//
// [hr element]: https://html.spec.whatwg.org/#the-hr-element
func Hr(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("hr", attributes, children...)
}

// I creates a [i element].
//
// [i element]: https://html.spec.whatwg.org/#the-i-element
func I(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("i", attributes, children...)
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

// Kbd creates a [kbd element].
//
// [kbd element]: https://html.spec.whatwg.org/#the-kbd-element
func Kbd(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("kbd", attributes, children...)
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

// Li creates a [li element].
//
// [li element]: https://html.spec.whatwg.org/#the-li-element
func Li(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("li", attributes, children...)
}

// Main creates a [main element]
//
// [main element]: https://html.spec.whatwg.org/#the-main-element
func Main(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("main", attributes, children...)
}

// Map creates a [map element].
//
// [map element]: https://html.spec.whatwg.org/#the-map-element
func Map(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("map", attributes, children...)
}

// Mark creates a [mark element].
//
// [mark element]: https://html.spec.whatwg.org/#the-mark-element
func Mark(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("mark", attributes, children...)
}

// Menu creates a [menu element].
//
// [menu element]: https://html.spec.whatwg.org/#the-menu-element
func Menu(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("menu", attributes, children...)
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

// Ol creates a [ol element].
//
// [ol element]: https://html.spec.whatwg.org/#the-ol-element
func Ol(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("ol", attributes, children...)
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

// P creates a [p element].
//
// [p element]: https://html.spec.whatwg.org/#the-p-element
func P(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("p", attributes, children...)
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

// Pre creates a [pre element].
//
// [pre element]: https://html.spec.whatwg.org/#the-pre-element
func Pre(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("pre", attributes, children...)
}

// Progress creates a [progress element]
//
// [progress element]: https://html.spec.whatwg.org/#the-progress-element
func Progress(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("progress", attributes, children...)
}

// Q creates a [q element].
//
// [q element]: https://html.spec.whatwg.org/#the-q-element
func Q(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("q", attributes, children...)
}

// Rp creates a [rp element].
//
// [rp element]: https://html.spec.whatwg.org/#the-rp-element
func Rp(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("rp", attributes, children...)
}

// Rt creates a [rt element].
//
// [rt element]: https://html.spec.whatwg.org/#the-rt-element
func Rt(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("rt", attributes, children...)
}

// Ruby creates a [ruby element].
//
// [ruby element]: https://html.spec.whatwg.org/#the-ruby-element
func Ruby(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("ruby", attributes, children...)
}

// S creates a [s element].
//
// [s element]: https://html.spec.whatwg.org/#the-s-element
func S(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("s", attributes, children...)
}

// Samp creates a [samp element].
//
// [samp element]: https://html.spec.whatwg.org/#the-samp-element
func Samp(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("samp", attributes, children...)
}

// Script creates a [script element].
//
// [script element]: https://html.spec.whatwg.org/#the-script-element
func Script(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("script", attributes, children...)
}

// Search creates a [search element].
//
// [search element]: https://html.spec.whatwg.org/#the-search-element
func Search(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("search", attributes, children...)
}

// Select creates a [select element].
//
// [select element]: https://html.spec.whatwg.org/#the-select-element
func Select(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("select", attributes, children...)
}

// Small creates a [small element].
//
// [small element]: https://html.spec.whatwg.org/#the-small-element
func Small(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("small", attributes, children...)
}

// Source creates a [source element].
//
// [source element]: https://html.spec.whatwg.org/#the-source-element
func Source(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("source", attributes, children...)
}

// Span creates a [span element].
//
// [span element]: https://html.spec.whatwg.org/#the-span-element
func Span(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("span", attributes, children...)
}

// Strong creates a [strong element].
//
// [strong element]: https://html.spec.whatwg.org/#the-strong-element
func Strong(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("strong", attributes, children...)
}

// Sub creates a [sub element].
//
// [sub element]: https://html.spec.whatwg.org/#the-sub-element
func Sub(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("sub", attributes, children...)
}

// Summary creates a [summary element].
//
// [summary element]: https://html.spec.whatwg.org/#the-summary-element
func Summary(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("summary", attributes, children...)
}

// Sup creates a [sup element].
//
// [sup element]: https://html.spec.whatwg.org/#the-sup-element
func Sup(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("sup", attributes, children...)
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

// Time creates a [time element].
//
// [time element]: https://html.spec.whatwg.org/#the-time-element
func Time(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("time", attributes, children...)
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

// U creates a [u element].
//
// [u element]: https://html.spec.whatwg.org/#the-u-element
func U(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("u", attributes, children...)
}

// Ul creates a [ul element].
//
// [ul element]: https://html.spec.whatwg.org/#the-ul-element
func Ul(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("ul", attributes, children...)
}

// Var creates a [var element].
//
// [var element]: https://html.spec.whatwg.org/#the-var-element
func Var(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("var", attributes, children...)
}

// Video creates a [video element].
//
// [video element]: https://html.spec.whatwg.org/#the-video-element
func Video(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("video", attributes, children...)
}

// Wbr creates a [wbr element].
//
// [wbr element]: https://html.spec.whatwg.org/#the-wbr-element
func Wbr(attributes attr.Ls, children ...htmfunc.Component) htmfunc.Component {
	return htmfunc.Element("wbr", attributes, children...)
}
