package elements

import (
	"github.com/KrischanCS/roundly/internal/generator/standard"
)

type Element struct {
	FuncName string

	IsVoid bool

	// Tag is the name of the element, e.g. "div", "span", "h1", etc.
	Tag string
	// Description is a short description of the element.
	Description string
	// SemanticGroup defines to which broad group the element belongs.
	SemanticGroup string
	// Categories are the categories the element belongs to.
	Categories []string
	// Parents are all possible elements of this element
	Parents []string
	// Children are all possible elements of this element
	Children []string
	// Attributes are all allowed attributes of this element
	Attributes []string

	// DocumentationLink is the link to the documentation of this in the html standard
	DocumentationLink standard.Link
	// Links are all links in the description of this element
	Links []standard.Link
}
