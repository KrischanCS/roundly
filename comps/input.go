package comps

import (
	"github.com/KrischanCS/roundly"

	. "github.com/KrischanCS/roundly/html"
)

func InputForm(
	actionTarget string,
	attrs roundly.Attribute,
	inputs ...roundly.Element,
) roundly.Element {
	return Form(
		Attributes(
			Action(actionTarget),
			If(attrs != nil, attrs),
			StyleAttribute(`display:grid; grid-template-columns: fit-content(1%) 1fr`),
		),
		Group(inputs...),
	)
}

// InputAttributes are the basic attributes for inputs. Label, name and id should always be
// provided, at least one of Value or Placeholder is recommended.
type InputAttributes struct {
	Label string
	Name  string
	Id    string

	Value       string
	Placeholder string
}

// ImageAttributes are the attributes provided to an IptImage. All values should be filled
type ImageAttributes struct {
	Label string
	Id    string

	Src    string
	Alt    string
	Width  uint
	Height uint
}

// IptRadioButtonOption represents the data for an element of a group of radio buttons
type IptRadioButtonOption struct {
	Label string
	Id    string

	Value string

	OtherAttributes []roundly.Attribute
}

func input(inputType func() roundly.Attribute, attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return Group(
		LabelText(attributes.Label, For(attributes.Id)),
		Input(
			Attributes(
				inputType(),
				Name(attributes.Name),
				Id(attributes.Id),
				Value(attributes.Value),
				PlaceHolder(attributes.Placeholder),
				Attributes(otherAttrs...),
			),
		),
	)
}

func inputLeft(inputType func() roundly.Attribute, attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return Group(
		Input(
			Attributes(
				inputType(),
				Name(attributes.Name),
				Id(attributes.Id),
				Value(attributes.Value),
				PlaceHolder(attributes.Placeholder),
				Attributes(otherAttrs...),
			),
		),
		LabelText(attributes.Label, For(attributes.Id)),
	)
}

// IptText creates an input of type text with a label.
func IptText(attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return input(TypeText, attributes, otherAttrs...)
}

// IptPassword creates an input of type password with a label.
func IptPassword(attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return input(TypePassword, attributes, otherAttrs...)
}

// IptSubmit creates a submit button.
func IptSubmit(value string, otherAttrs ...roundly.Attribute) roundly.Element {
	return Input(Attributes(
		TypeSubmit(),
		Value(value),
		Attributes(otherAttrs...),
	))
}

// IptReset create a reset button.
func IptReset(value string, otherAttrs ...roundly.Attribute) roundly.Element {
	return Input(Attributes(
		TypeReset(),
		Value(value),
		Attributes(otherAttrs...),
	))
}

// IptRadioGroup creates a group of radio buttons.
func IptRadioGroup(name string, options ...IptRadioButtonOption) roundly.Element {
	return Range(options, func(_ int, attributes IptRadioButtonOption) roundly.Element {
		return inputLeft(
			TypeRadio,
			InputAttributes{
				Label: attributes.Label,
				Name:  name,
				Id:    attributes.Id,
				Value: attributes.Value,
			},
			attributes.OtherAttributes...,
		)
	})
}

// IptCheckBox creates an input with type checkbox with a label
func IptCheckBox(attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return inputLeft(TypeCheckbox, attributes, otherAttrs...)
}

// IptButton creates a button with the given onclick action.
func IptButton(value string, onClick string, otherAttrs ...roundly.Attribute) roundly.Element {
	return Input(Attributes(
		TypeButton(),
		Value(value),
		OnClick(onClick),
		Attributes(otherAttrs...),
	))
}

// IptColor creates an input of type color with a label.
func IptColor(attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return input(TypeColor, attributes, otherAttrs...)
}

// IptDate creates an input of type date with a label.
func IptDate(attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return input(TypeDate, attributes, otherAttrs...)
}

// IptDateTimeLocal creates an input if type datetime-local with a label.
func IptDateTimeLocal(attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return input(TypeDatetimeLocal, attributes, otherAttrs...)
}

// IptEmail create an input of type email with a label.
func IptEmail(attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return input(TypeEmail, attributes, otherAttrs...)
}

// IptImage creates an input of type image with a label.
func IptImage(attributes ImageAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return Div(
		StyleAttribute(`display:flex; flex-direction:column;`),
		Label(Id(attributes.Id), Text(attributes.Label)),
		Input(
			Attributes(
				TypeImage(),
				Id(attributes.Id),
				Src(attributes.Src),
				Alt(attributes.Alt),
				Width(attributes.Width),
				Height(attributes.Height),
				Attributes(otherAttrs...),
			),
		),
	)
}

// IptFile create an input of type file with a label.
func IptFile(attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return input(TypeFile, attributes, otherAttrs...)
}

// IptHidden creates a input of type hidden with a label.
func IptHidden(name, value string, otherAttrs ...roundly.Attribute) roundly.Element {
	return Input(
		Attributes(
			TypeHidden(),
			Name(name),
			Value(value),
			Attributes(otherAttrs...),
		),
	)
}

// IptMonth creates an input of type month with a label.
//
// Deprecated: the 'month' input-type is not supported in Firefox & Safari, wouldn't recommend
// using it.
func IptMonth(attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return input(TypeMonth, attributes, otherAttrs...)
}

// IptNumber creates an input of type number with a label.
func IptNumber(attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return input(TypeNumber, attributes, otherAttrs...)
}

// IptRange creates an input of type range with a label.
func IptRange(attributes InputAttributes, min, max string, otherAttrs ...roundly.Attribute) roundly.Element {
	return input(TypeRange, attributes,
		Attributes(
			Min(min),
			Max(max),
			Attributes(otherAttrs...),
		),
	)
}

// IptSearch creates an input of type search with a label.
func IptSearch(attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return input(TypeSearch, attributes, otherAttrs...)
}

// IptTel creates an input of type tel with a label.
func IptTel(attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return input(TypeTel, attributes, otherAttrs...)
}

// IptTime creates an input of type time with a label.
func IptTime(attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return input(TypeTime, attributes, otherAttrs...)
}

// IptUrl creates an input of type url with a label.
func IptUrl(attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return input(TypeUrl, attributes, otherAttrs...)
}

// IptWeek creates an input of type week with a label.
func IptWeek(attributes InputAttributes, otherAttrs ...roundly.Attribute) roundly.Element {
	return input(TypeWeek, attributes, otherAttrs...)
}
