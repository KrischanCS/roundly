package main

import (
	"slices"
	"strings"
)

type namedElementGroups struct {
	elements []string
	name     string
}

var definedElementGroups = []namedElementGroups{
	{
		elements: []string{"[button]", "[input]", "[optgroup]", "[option]", "[select]", "[textarea]", "[form-associated custom elements]"},
		name:     "Inputs",
	},
	{
		elements: []string{"[button]", "[fieldset]", "[input]", "[output]", "[select]", "[textarea]", "[form-associated custom elements]"},
		name:     "InputsOutputs",
	},
}

func disambiguateAttrs(attrsByName map[string][]*attribute) {
	for _, attrs := range attrsByName {
		if len(attrs) == 1 {
			continue
		}

		for _, attr := range attrs {
			renameByElements(attr)
		}
	}
}

func renameByElements(attr *attribute) {
	if len(attr.Elements) == 1 && attr.Elements[0] == "[HTML elements]" {
		return
	}

	attr.FuncName += "_"

	for _, elementName := range attr.Elements {
		if elementName == "[HTML elements]" {
			continue
		}

		for _, group := range definedElementGroups {
			if !slices.Equal(attr.Elements, group.elements) {
				continue
			}
			attr.FuncName += group.name
			return
		}

		if elementName == "[form-associated custom elements]" {
			attr.FuncName += "CustomFormElements"
			continue
		}

		runes := []rune(strings.Trim(elementName, "[]"))
		attr.FuncName += strings.ToUpper(string(runes[0])) + string(runes[1:])
	}
}
