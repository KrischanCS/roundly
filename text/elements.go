package text

import (
	"fmt"

	"github.com/KrischanCS/htmfunc"
	. "github.com/KrischanCS/htmfunc/element"
)

// DelT creates an [element.Del] without any attributes and text as the only child.
func DelT(text string) htmfunc.Element {
	return Del(nil, Text(text))
}

// DelTf creates an [element.Del] without any attributes and the formatted template as its only
// child.
func DelTf(template string, args ...any) htmfunc.Element {
	return Del(nil, Text(fmt.Sprintf(template, args...)))
}

// InsT creates an [element.Ins] without any attributes and text as the only child.
func InsT(text string) htmfunc.Element {
	return Ins(nil, Text(text))
}

// InsTf creates an [element.Ins] without any attributes and the formatted template as its only
// child.
func InsTf(template string, args ...any) htmfunc.Element {
	return Ins(nil, Text(fmt.Sprintf(template, args...)))
}

// ButtonT creates an [element.Button] without any attributes and text as the only child.
func ButtonT(text string) htmfunc.Element {
	return Button(nil, Text(text))
}

// ButtonTf creates an [element.Button] without any attributes and the formatted template as its
// only child.
func ButtonTf(template string, args ...any) htmfunc.Element {
	return Button(nil, Text(fmt.Sprintf(template, args...)))
}

// LegendT creates an [element.Legend] without any attributes and text as the only child.
func LegendT(text string) htmfunc.Element {
	return Legend(nil, Text(text))
}

// LegendTf creates an [element.Legend] without any attributes and the formatted template as its
// only child.
func LegendTf(template string, args ...any) htmfunc.Element {
	return Legend(nil, Text(fmt.Sprintf(template, args...)))
}

// OptionT creates an [element.Option] without any attributes and text as the only child.
func OptionT(text string) htmfunc.Element {
	return Option(nil, Text(text))
}

// OptionTf creates an [element.Option] without any attributes and the formatted template as its
// only child.
func OptionTf(template string, args ...any) htmfunc.Element {
	return Option(nil, Text(fmt.Sprintf(template, args...)))
}

// OutputT creates an [element.Output] without any attributes and text as the only child.
func OutputT(text string) htmfunc.Element {
	return Output(nil, Text(text))
}

// OutputTf creates an [element.Output] without any attributes and the formatted template as its
// only child.
func OutputTf(template string, args ...any) htmfunc.Element {
	return Output(nil, Text(fmt.Sprintf(template, args...)))
}

// ProgressT creates an [element.Progress] without any attributes and text as the only child.
func ProgressT(text string) htmfunc.Element {
	return Progress(nil, Text(text))
}

// ProgressTf creates an [element.Progress] without any attributes and the formatted template as its
// only child.
func ProgressTf(template string, args ...any) htmfunc.Element {
	return Progress(nil, Text(fmt.Sprintf(template, args...)))
}

// TextareaT creates an [element.Textarea] without any attributes and text as the only child.
func TextareaT(text string) htmfunc.Element {
	return Textarea(nil, Text(text))
}

// TextareaTf creates an [element.Textarea] without any attributes and the formatted template as its
// only child.
func TextareaTf(template string, args ...any) htmfunc.Element {
	return Textarea(nil, Text(fmt.Sprintf(template, args...)))
}

// LabelT creates an [element.Label] without any attributes and text as the only child.
func LabelT(text string) htmfunc.Element {
	return Label(nil, Text(text))
}

// LabelTf creates an [element.Label] without any attributes and the formatted template as its only
// child.
func LabelTf(template string, args ...any) htmfunc.Element {
	return Label(nil, Text(fmt.Sprintf(template, args...)))
}

// BlockquoteT creates an [element.Blockquote] without any attributes and text as the only child.
func BlockquoteT(text string) htmfunc.Element {
	return Blockquote(nil, Text(text))
}

// BlockquoteTf creates an [element.Blockquote] without any attributes and the formatted template as
// its only child.
func BlockquoteTf(template string, args ...any) htmfunc.Element {
	return Blockquote(nil, Text(fmt.Sprintf(template, args...)))
}

// DdT creates an [element.Dd] without any attributes and text as the only child.
func DdT(text string) htmfunc.Element {
	return Dd(nil, Text(text))
}

// DdTf creates an [element.Dd] without any attributes and the formatted template as its only child.
func DdTf(template string, args ...any) htmfunc.Element {
	return Dd(nil, Text(fmt.Sprintf(template, args...)))
}

// DtT creates an [element.Dt] without any attributes and text as the only child.
func DtT(text string) htmfunc.Element {
	return Dt(nil, Text(text))
}

// DtTf creates an [element.Dt] without any attributes and the formatted template as its only child.
func DtTf(template string, args ...any) htmfunc.Element {
	return Dt(nil, Text(fmt.Sprintf(template, args...)))
}

// FigcaptionT creates an [element.Figcaption] without any attributes and text as the only child.
func FigcaptionT(text string) htmfunc.Element {
	return Figcaption(nil, Text(text))
}

// FigcaptionTf creates an [element.Figcaption] without any attributes and the formatted template as
// its only child.
func FigcaptionTf(template string, args ...any) htmfunc.Element {
	return Figcaption(nil, Text(fmt.Sprintf(template, args...)))
}

// LiT creates an [element.Li] without any attributes and text as the only child.
func LiT(text string) htmfunc.Element {
	return Li(nil, Text(text))
}

// LiTf creates an [element.Li] without any attributes and the formatted template as its only child.
func LiTf(template string, args ...any) htmfunc.Element {
	return Li(nil, Text(fmt.Sprintf(template, args...)))
}

// PT creates an [element.P] without any attributes and text as the only child.
func PT(text string) htmfunc.Element {
	return P(nil, Text(text))
}

// PTf creates an [element.P] without any attributes and the formatted template as its only child.
func PTf(template string, args ...any) htmfunc.Element {
	return P(nil, Text(fmt.Sprintf(template, args...)))
}

// PreT creates an [element.Pre] without any attributes and text as the only child.
func PreT(text string) htmfunc.Element {
	return Pre(nil, Text(text))
}

// PreTf creates an [element.Pre] without any attributes and the formatted template as its only
// child.
func PreTf(template string, args ...any) htmfunc.Element {
	return Pre(nil, Text(fmt.Sprintf(template, args...)))
}

// UlT creates an [element.Ul] without any attributes and text as the only child.
func UlT(text string) htmfunc.Element {
	return Ul(nil, Text(text))
}

// UlTf creates an [element.Ul] without any attributes and the formatted template as its only child.
func UlTf(template string, args ...any) htmfunc.Element {
	return Ul(nil, Text(fmt.Sprintf(template, args...)))
}

// SummaryT creates an [element.Summary] without any attributes and text as the only child.
func SummaryT(text string) htmfunc.Element {
	return Summary(nil, Text(text))
}

// SummaryTf creates an [element.Summary] without any attributes and the formatted template as its
// only child.
func SummaryTf(template string, args ...any) htmfunc.Element {
	return Summary(nil, Text(fmt.Sprintf(template, args...)))
}

// NoscriptT creates an [element.Noscript] without any attributes and text as the only child.
func NoscriptT(text string) htmfunc.Element {
	return Noscript(nil, Text(text))
}

// NoscriptTf creates an [element.Noscript] without any attributes and the formatted template as its
// only child.
func NoscriptTf(template string, args ...any) htmfunc.Element {
	return Noscript(nil, Text(fmt.Sprintf(template, args...)))
}

// H1T creates an [element.H1] without any attributes and text as the only child.
func H1T(text string) htmfunc.Element {
	return H1(nil, Text(text))
}

// H1Tf creates an [element.H1] without any attributes and the formatted template as its only child.
func H1Tf(template string, args ...any) htmfunc.Element {
	return H1(nil, Text(fmt.Sprintf(template, args...)))
}

// H2T creates an [element.H2] without any attributes and text as the only child.
func H2T(text string) htmfunc.Element {
	return H2(nil, Text(text))
}

// H2Tf creates an [element.H2] without any attributes and the formatted template as its only child.
func H2Tf(template string, args ...any) htmfunc.Element {
	return H2(nil, Text(fmt.Sprintf(template, args...)))
}

// H3T creates an [element.H3] without any attributes and text as the only child.
func H3T(text string) htmfunc.Element {
	return H3(nil, Text(text))
}

// H3Tf creates an [element.H3] without any attributes and the formatted template as its only child.
func H3Tf(template string, args ...any) htmfunc.Element {
	return H3(nil, Text(fmt.Sprintf(template, args...)))
}

// H4T creates an [element.H4] without any attributes and text as the only child.
func H4T(text string) htmfunc.Element {
	return H4(nil, Text(text))
}

// H4Tf creates an [element.H4] without any attributes and the formatted template as its only child.
func H4Tf(template string, args ...any) htmfunc.Element {
	return H4(nil, Text(fmt.Sprintf(template, args...)))
}

// H5T creates an [element.H5] without any attributes and text as the only child.
func H5T(text string) htmfunc.Element {
	return H5(nil, Text(text))
}

// H5Tf creates an [element.H5] without any attributes and the formatted template as its only child.
func H5Tf(template string, args ...any) htmfunc.Element {
	return H5(nil, Text(fmt.Sprintf(template, args...)))
}

// H6T creates an [element.H6] without any attributes and text as the only child.
func H6T(text string) htmfunc.Element {
	return H6(nil, Text(text))
}

// H6Tf creates an [element.H6] without any attributes and the formatted template as its only child.
func H6Tf(template string, args ...any) htmfunc.Element {
	return H6(nil, Text(fmt.Sprintf(template, args...)))
}

// TitleT creates an [element.Title] without any attributes and text as the only child.
func TitleT(text string) htmfunc.Element {
	return Title(nil, Text(text))
}

// TitleTf creates an [element.Title] without any attributes and the formatted template as its only
// child.
func TitleTf(template string, args ...any) htmfunc.Element {
	return Title(nil, Text(fmt.Sprintf(template, args...)))
}

// CaptionT creates an [element.Caption] without any attributes and text as the only child.
func CaptionT(text string) htmfunc.Element {
	return Caption(nil, Text(text))
}

// CaptionTf creates an [element.Caption] without any attributes and the formatted template as its
// only child.
func CaptionTf(template string, args ...any) htmfunc.Element {
	return Caption(nil, Text(fmt.Sprintf(template, args...)))
}

// TdT creates an [element.Td] without any attributes and text as the only child.
func TdT(text string) htmfunc.Element {
	return Td(nil, Text(text))
}

// TdTf creates an [element.Td] without any attributes and the formatted template as its only child.
func TdTf(template string, args ...any) htmfunc.Element {
	return Td(nil, Text(fmt.Sprintf(template, args...)))
}

// ThT creates an [element.Th] without any attributes and text as the only child.
func ThT(text string) htmfunc.Element {
	return Th(nil, Text(text))
}

// ThTf creates an [element.Th] without any attributes and the formatted template as its only child.
func ThTf(template string, args ...any) htmfunc.Element {
	return Th(nil, Text(fmt.Sprintf(template, args...)))
}

// AT creates an [element.A] without any attributes and text as the only child.
func AT(text string) htmfunc.Element {
	return A(nil, Text(text))
}

// ATf creates an [element.A] without any attributes and the formatted template as its only child.
func ATf(template string, args ...any) htmfunc.Element {
	return A(nil, Text(fmt.Sprintf(template, args...)))
}

// AbbrT creates an [element.Abbr] without any attributes and text as the only child.
func AbbrT(text string) htmfunc.Element {
	return Abbr(nil, Text(text))
}

// AbbrTf creates an [element.Abbr] without any attributes and the formatted template as its only
// child.
func AbbrTf(template string, args ...any) htmfunc.Element {
	return Abbr(nil, Text(fmt.Sprintf(template, args...)))
}

// BT creates an [element.B] without any attributes and text as the only child.
func BT(text string) htmfunc.Element {
	return B(nil, Text(text))
}

// BTf creates an [element.B] without any attributes and the formatted template as its only child.
func BTf(template string, args ...any) htmfunc.Element {
	return B(nil, Text(fmt.Sprintf(template, args...)))
}

// CiteT creates an [element.Cite] without any attributes and text as the only child.
func CiteT(text string) htmfunc.Element {
	return Cite(nil, Text(text))
}

// CiteTf creates an [element.Cite] without any attributes and the formatted template as its only
// child.
func CiteTf(template string, args ...any) htmfunc.Element {
	return Cite(nil, Text(fmt.Sprintf(template, args...)))
}

// CodeT creates an [element.Code] without any attributes and text as the only child.
func CodeT(text string) htmfunc.Element {
	return Code(nil, Text(text))
}

// CodeTf creates an [element.Code] without any attributes and the formatted template as its only
// child.
func CodeTf(template string, args ...any) htmfunc.Element {
	return Code(nil, Text(fmt.Sprintf(template, args...)))
}

// DataT creates an [element.Data] without any attributes and text as the only child.
func DataT(text string) htmfunc.Element {
	return Data(nil, Text(text))
}

// DataTf creates an [element.Data] without any attributes and the formatted template as its only
// child.
func DataTf(template string, args ...any) htmfunc.Element {
	return Data(nil, Text(fmt.Sprintf(template, args...)))
}

// DfnT creates an [element.Dfn] without any attributes and text as the only child.
func DfnT(text string) htmfunc.Element {
	return Dfn(nil, Text(text))
}

// DfnTf creates an [element.Dfn] without any attributes and the formatted template as its only
// child.
func DfnTf(template string, args ...any) htmfunc.Element {
	return Dfn(nil, Text(fmt.Sprintf(template, args...)))
}

// EmT creates an [element.Em] without any attributes and text as the only child.
func EmT(text string) htmfunc.Element {
	return Em(nil, Text(text))
}

// EmTf creates an [element.Em] without any attributes and the formatted template as its only child.
func EmTf(template string, args ...any) htmfunc.Element {
	return Em(nil, Text(fmt.Sprintf(template, args...)))
}

// IT creates an [element.I] without any attributes and text as the only child.
func IT(text string) htmfunc.Element {
	return I(nil, Text(text))
}

// ITf creates an [element.I] without any attributes and the formatted template as its only child.
func ITf(template string, args ...any) htmfunc.Element {
	return I(nil, Text(fmt.Sprintf(template, args...)))
}

// KbdT creates an [element.Kbd] without any attributes and text as the only child.
func KbdT(text string) htmfunc.Element {
	return Kbd(nil, Text(text))
}

// KbdTf creates an [element.Kbd] without any attributes and the formatted template as its only
// child.
func KbdTf(template string, args ...any) htmfunc.Element {
	return Kbd(nil, Text(fmt.Sprintf(template, args...)))
}

// MarkT creates an [element.Mark] without any attributes and text as the only child.
func MarkT(text string) htmfunc.Element {
	return Mark(nil, Text(text))
}

// MarkTf creates an [element.Mark] without any attributes and the formatted template as its only
// child.
func MarkTf(template string, args ...any) htmfunc.Element {
	return Mark(nil, Text(fmt.Sprintf(template, args...)))
}

// QT creates an [element.Q] without any attributes and text as the only child.
func QT(text string) htmfunc.Element {
	return Q(nil, Text(text))
}

// QTf creates an [element.Q] without any attributes and the formatted template as its only child.
func QTf(template string, args ...any) htmfunc.Element {
	return Q(nil, Text(fmt.Sprintf(template, args...)))
}

// RpT creates an [element.Rp] without any attributes and text as the only child.
func RpT(text string) htmfunc.Element {
	return Rp(nil, Text(text))
}

// RpTf creates an [element.Rp] without any attributes and the formatted template as its only child.
func RpTf(template string, args ...any) htmfunc.Element {
	return Rp(nil, Text(fmt.Sprintf(template, args...)))
}

// RtT creates an [element.Rt] without any attributes and text as the only child.
func RtT(text string) htmfunc.Element {
	return Rt(nil, Text(text))
}

// RtTf creates an [element.Rt] without any attributes and the formatted template as its only child.
func RtTf(template string, args ...any) htmfunc.Element {
	return Rt(nil, Text(fmt.Sprintf(template, args...)))
}

// ST creates an [element.S] without any attributes and text as the only child.
func ST(text string) htmfunc.Element {
	return S(nil, Text(text))
}

// STf creates an [element.S] without any attributes and the formatted template as its only child.
func STf(template string, args ...any) htmfunc.Element {
	return S(nil, Text(fmt.Sprintf(template, args...)))
}

// SampT creates an [element.Samp] without any attributes and text as the only child.
func SampT(text string) htmfunc.Element {
	return Samp(nil, Text(text))
}

// SampTf creates an [element.Samp] without any attributes and the formatted template as its only
// child.
func SampTf(template string, args ...any) htmfunc.Element {
	return Samp(nil, Text(fmt.Sprintf(template, args...)))
}

// SmallT creates an [element.Small] without any attributes and text as the only child.
func SmallT(text string) htmfunc.Element {
	return Small(nil, Text(text))
}

// SmallTf creates an [element.Small] without any attributes and the formatted template as its only
// child.
func SmallTf(template string, args ...any) htmfunc.Element {
	return Small(nil, Text(fmt.Sprintf(template, args...)))
}

// SpanT creates an [element.Span] without any attributes and text as the only child.
func SpanT(text string) htmfunc.Element {
	return Span(nil, Text(text))
}

// SpanTf creates an [element.Span] without any attributes and the formatted template as its only
// child.
func SpanTf(template string, args ...any) htmfunc.Element {
	return Span(nil, Text(fmt.Sprintf(template, args...)))
}

// StrongT creates an [element.Strong] without any attributes and text as the only child.
func StrongT(text string) htmfunc.Element {
	return Strong(nil, Text(text))
}

// StrongTf creates an [element.Strong] without any attributes and the formatted template as its
// only child.
func StrongTf(template string, args ...any) htmfunc.Element {
	return Strong(nil, Text(fmt.Sprintf(template, args...)))
}

// SubT creates an [element.Sub] without any attributes and text as the only child.
func SubT(text string) htmfunc.Element {
	return Sub(nil, Text(text))
}

// SubTf creates an [element.Sub] without any attributes and the formatted template as its only
// child.
func SubTf(template string, args ...any) htmfunc.Element {
	return Sub(nil, Text(fmt.Sprintf(template, args...)))
}

// SupT creates an [element.Sup] without any attributes and text as the only child.
func SupT(text string) htmfunc.Element {
	return Sup(nil, Text(text))
}

// SupTf creates an [element.Sup] without any attributes and the formatted template as its only
// child.
func SupTf(template string, args ...any) htmfunc.Element {
	return Sup(nil, Text(fmt.Sprintf(template, args...)))
}

// UT creates an [element.U] without any attributes and text as the only child.
func UT(text string) htmfunc.Element {
	return U(nil, Text(text))
}

// UTf creates an [element.U] without any attributes and the formatted template as its only child.
func UTf(template string, args ...any) htmfunc.Element {
	return U(nil, Text(fmt.Sprintf(template, args...)))
}

// VarT creates an [element.Var] without any attributes and text as the only child.
func VarT(text string) htmfunc.Element {
	return Var(nil, Text(text))
}

// VarTf creates an [element.Var] without any attributes and the formatted template as its only
// child.
func VarTf(template string, args ...any) htmfunc.Element {
	return Var(nil, Text(fmt.Sprintf(template, args...)))
}
