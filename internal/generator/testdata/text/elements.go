package html

import (
	"github.com/KrischanCS/roundly"

	. "github.com/KrischanCS/roundly/html"
)

// DelText creates a [element.Del] element without any attributes and
// text as the only child.
func DelText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Del(Attributes(attributes...), Text(text))
}

// InsText creates a [element.Ins] element without any attributes and text as the only child.
func InsText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Ins(Attributes(attributes...), Text(text))
}

// ButtonText creates a [element.Button] element without any attributes and text as the only child.
func ButtonText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Button(Attributes(attributes...), Text(text))
}

// LegendText creates a [element.Legend] element without any attributes and text as the only child.
func LegendText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Legend(Attributes(attributes...), Text(text))
}

// OptionText creates a [element.Option] element without any attributes and text as the only child.
func OptionText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Option(Attributes(attributes...), Text(text))
}

// OutputText creates a [element.Output] element without any attributes and text as the only child.
func OutputText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Output(Attributes(attributes...), Text(text))
}

// ProgressText creates a [element.Progress] element without any attributes and text as the only
// child.
func ProgressText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Progress(Attributes(attributes...), Text(text))
}

// TextareaText creates a [element.Textarea] element without any attributes and text as the only
// child.
func TextareaText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Textarea(Attributes(attributes...), Text(text))
}

// LabelText creates a [element.Label] element without any attributes and text as the only child.
func LabelText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Label(Attributes(attributes...), Text(text))
}

// BlockquoteText creates a [element.Blockquote] element without any attributes and text as the only
// child.
func BlockquoteText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Blockquote(Attributes(attributes...), Text(text))
}

// DdText creates a [element.Dd] element without any attributes and text as the only child.
func DdText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Dd(Attributes(attributes...), Text(text))
}

// DtText creates a [element.Dt] element without any attributes and text as the only child.
func DtText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Dt(Attributes(attributes...), Text(text))
}

// FigcaptionText creates a [element.Figcaption] element without any attributes and text as the only
// child.
func FigcaptionText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Figcaption(Attributes(attributes...), Text(text))
}

// LiText creates a [element.Li] element without any attributes and text as the only child.
func LiText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Li(Attributes(attributes...), Text(text))
}

// PText creates a [element.P] element without any attributes and text as the only child.
func PText(text string, attributes ...roundly.Attribute) roundly.Element {
	return P(Attributes(attributes...), Text(text))
}

// PreText creates a [element.Pre] element without any attributes and text as the only child.
func PreText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Pre(Attributes(attributes...), Text(text))
}

// UlText creates a [element.Ul] element without any attributes and text as the only child.
func UlText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Ul(Attributes(attributes...), Text(text))
}

// SummaryText creates a [element.Summary] element without any attributes and text as the only
// child.
func SummaryText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Summary(Attributes(attributes...), Text(text))
}

// NoscriptText creates a [element.Noscript] element without any attributes and text as the only
// child.
func NoscriptText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Noscript(Attributes(attributes...), Text(text))
}

// H1Text creates a [element.H1] element without any attributes and text as the only child.
func H1Text(text string, attributes ...roundly.Attribute) roundly.Element {
	return H1(Attributes(attributes...), Text(text))
}

// H2Text creates a [element.H2] element without any attributes and text as the only child.
func H2Text(text string, attributes ...roundly.Attribute) roundly.Element {
	return H2(Attributes(attributes...), Text(text))
}

// H3Text creates a [element.H3] element without any attributes and text as the only child.
func H3Text(text string, attributes ...roundly.Attribute) roundly.Element {
	return H3(Attributes(attributes...), Text(text))
}

// H4Text creates a [element.H4] element without any attributes and text as the only child.
func H4Text(text string, attributes ...roundly.Attribute) roundly.Element {
	return H4(Attributes(attributes...), Text(text))
}

// H5Text creates a [element.H5] element without any attributes and text as the only child.
func H5Text(text string, attributes ...roundly.Attribute) roundly.Element {
	return H5(Attributes(attributes...), Text(text))
}

// H6Text creates a [element.H6] element without any attributes and text as the only child.
func H6Text(text string, attributes ...roundly.Attribute) roundly.Element {
	return H6(Attributes(attributes...), Text(text))
}

// TitleText creates a [element.Title] element without any attributes and text as the only child.
func TitleText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Title(Attributes(attributes...), Text(text))
}

// CaptionText creates a [element.Caption] element without any attributes and text as the only
// child.
func CaptionText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Caption(Attributes(attributes...), Text(text))
}

// TdText creates a [element.Td] element without any attributes and text as the only child.
func TdText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Td(Attributes(attributes...), Text(text))
}

// ThText creates a [element.Th] element without any attributes and text as the only child.
func ThText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Th(Attributes(attributes...), Text(text))
}

// AText creates a [element.A] element without any attributes and text as the only child.
func AText(text string, attributes ...roundly.Attribute) roundly.Element {
	return A(Attributes(attributes...), Text(text))
}

// AbbrText creates a [element.Abbr] element without any attributes and text as the only child.
func AbbrText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Abbr(Attributes(attributes...), Text(text))
}

// BText creates a [element.B] element without any attributes and text as the only child.
func BText(text string, attributes ...roundly.Attribute) roundly.Element {
	return B(Attributes(attributes...), Text(text))
}

// CiteText creates a [element.Cite] element without any attributes and text as the only child.
func CiteText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Cite(Attributes(attributes...), Text(text))
}

// CodeText creates a [element.Code] element without any attributes and text as the only child.
func CodeText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Code(Attributes(attributes...), Text(text))
}

// DataText creates a [element.Data] element without any attributes and text as the only child.
func DataText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Data(Attributes(attributes...), Text(text))
}

// DfnText creates a [element.Dfn] element without any attributes and text as the only child.
func DfnText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Dfn(Attributes(attributes...), Text(text))
}

// EmText creates a [element.Em] element without any attributes and text as the only child.
func EmText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Em(Attributes(attributes...), Text(text))
}

// IText creates a [element.I] element without any attributes and text as the only child.
func IText(text string, attributes ...roundly.Attribute) roundly.Element {
	return I(Attributes(attributes...), Text(text))
}

// KbdText creates a [element.Kbd] element without any attributes and text as the only child.
func KbdText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Kbd(Attributes(attributes...), Text(text))
}

// MarkText creates a [element.Mark] element without any attributes and text as the only child.
func MarkText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Mark(Attributes(attributes...), Text(text))
}

// QText creates a [element.Q] element without any attributes and text as the only child.
func QText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Q(Attributes(attributes...), Text(text))
}

// RpText creates a [element.Rp] element without any attributes and text as the only child.
func RpText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Rp(Attributes(attributes...), Text(text))
}

// RtText creates a [element.Rt] element without any attributes and text as the only child.
func RtText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Rt(Attributes(attributes...), Text(text))
}

// SText creates a [element.S] element without any attributes and text as the only child.
func SText(text string, attributes ...roundly.Attribute) roundly.Element {
	return S(Attributes(attributes...), Text(text))
}

// SampText creates a [element.Samp] element without any attributes and text as the only child.
func SampText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Samp(Attributes(attributes...), Text(text))
}

// SmallText creates a [element.Small] element without any attributes and text as the only child.
func SmallText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Small(Attributes(attributes...), Text(text))
}

// SpanText creates a [element.Span] element without any attributes and text as the only child.
func SpanText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Span(Attributes(attributes...), Text(text))
}

// StrongText creates a [element.Strong] element without any attributes and text as the only child.
func StrongText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Strong(Attributes(attributes...), Text(text))
}

// SubText creates a [element.Sub] element without any attributes and text as the only child.
func SubText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Sub(Attributes(attributes...), Text(text))
}

// SupText creates a [element.Sup] element without any attributes and text as the only child.
func SupText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Sup(Attributes(attributes...), Text(text))
}

// UText creates a [element.U] element without any attributes and text as the only child.
func UText(text string, attributes ...roundly.Attribute) roundly.Element {
	return U(Attributes(attributes...), Text(text))
}

// VarText creates a [element.Var] element without any attributes and text as the only child.
func VarText(text string, attributes ...roundly.Attribute) roundly.Element {
	return Var(Attributes(attributes...), Text(text))
}
