package attributes

import (
	"github.com/KrischanCS/roundly/internal/generator/standard"
)

type attribute struct {
	Name        string
	FuncName    string
	ParamName   string
	Elements    []string
	Description string
	Value       string
	Values      []string
	Links       []standard.Link
}

type attributes struct {
	Text           []attribute
	Bool           []attribute
	Enum           []attribute
	ListComma      []attribute
	ListCommaFloat []attribute
	ListSpace      []attribute
	Float          []attribute
	Int            []attribute
	Uint           []attribute
}

func newAttributes() attributes {
	attrsClassified := attributes{
		Text:           make([]attribute, 0, 16),
		Bool:           make([]attribute, 0, 16),
		Enum:           make([]attribute, 0, 16),
		ListComma:      make([]attribute, 0, 16),
		ListCommaFloat: make([]attribute, 0, 16),
		ListSpace:      make([]attribute, 0, 16),
		Float:          make([]attribute, 0, 16),
		Int:            make([]attribute, 0, 16),
		Uint:           make([]attribute, 0, 16),
	}

	return attrsClassified
}
