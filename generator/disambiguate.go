package main

import (
	"slices"
	"strings"
)

type namedElementGroups struct {
	elements []string
	name     string
}

//nolint:gochecknoglobals
var definedElementGroups = []namedElementGroups{
	{
		elements: []string{
			"[button]",
			"[input]",
			"[optgroup]",
			"[option]",
			"[select]",
			"[textarea]",
			"[form-associated custom elements]",
		},
		name: "Inputs",
	},
	{
		elements: []string{
			"[button]",
			"[fieldset]",
			"[input]",
			"[output]",
			"[select]",
			"[textarea]",
			"[form-associated custom elements]",
		},
		name: "InputsOutputs",
	},
}

// disambiguateAttrs checks multiple attributes with the same name exist.
//
// If so, it renames them based on the elements they can be used on.
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

// renameByElements implements the renaming rules.
//   - Attributes which can be applied to all elements ("[HTML elements]") are not renamed
//   - If the element set equals a defined group, the group name is used
//   - If it can be applied to form-associated custom elements, the name is suffixed with "CustomFormElements"
//   - If it nonthin of that applies, the name is suffixed with all the applicable element names
func renameByElements(attr *attribute) {
	if len(attr.Elements) == 1 && attr.Elements[0] == "[HTML elements]" {
		return
	}

	attr.FuncName += "_"

	for _, elementName := range attr.Elements {
		if elementName == "[HTML elements]" {
			continue
		}

		if groupName, ok := findDefinedGroup(attr.Elements); ok {
			attr.FuncName += groupName
			continue
		}

		if elementName == "[form-associated custom elements]" {
			attr.FuncName += "CustomFormElements"
			continue
		}

		runes := []rune(strings.Trim(elementName, "[]"))
		attr.FuncName += strings.ToUpper(string(runes[0])) + string(runes[1:])
	}
}

func findDefinedGroup(elements []string) (string, bool) {
	for _, group := range definedElementGroups {
		if !slices.Equal(elements, group.elements) {
			continue
		}

		return group.name, true
	}

	return "", false
}
