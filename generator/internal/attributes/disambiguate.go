package attributes

import (
	"slices"
	"strings"

	"github.com/KrischanCS/go-toolbox/set"
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

// disambiguateAttrs checks multiple attributes with the same funcName exist.
//
// If so, it renames them based on their parameter types.
func disambiguateAttrs(attrs attributes) {
	duplicates := findDuplicatedNames(attrs)

	appendSuffixToDuplicates(attrs.Bool, "True", duplicates)
	appendSuffixToDuplicates(attrs.InputType, "InputType", duplicates)
	appendSuffixToDuplicates(attrs.ListComma, "Strings", duplicates)
	appendSuffixToDuplicates(attrs.ListSpace, "Strings", duplicates)
	appendSuffixToDuplicates(attrs.ListCommaFloat, "Floats", duplicates)
	appendSuffixToDuplicates(attrs.Float, "Float", duplicates)
	appendSuffixToDuplicates(attrs.Int, "Int", duplicates)
	appendSuffixToDuplicates(attrs.Uint, "UInt", duplicates)
}

func appendSuffixToDuplicates(attrs []attribute, suffix string, duplicates set.Set[string]) {
	for i, attr := range attrs {
		if duplicates.Contains(attr.FuncName) {
			attrs[i].FuncName += suffix
		}
	}
}

func findDuplicatedNames(attrs attributes) set.Set[string] {
	nameAppearances := make(map[string]int, 512)

	addFuncNameCounts(&nameAppearances, attrs.Text)
	addFuncNameCounts(&nameAppearances, attrs.Bool)
	addFuncNameCounts(&nameAppearances, attrs.Enum)
	addFuncNameCounts(&nameAppearances, attrs.InputType)
	addFuncNameCounts(&nameAppearances, attrs.ListComma)
	addFuncNameCounts(&nameAppearances, attrs.ListCommaFloat)
	addFuncNameCounts(&nameAppearances, attrs.ListSpace)
	addFuncNameCounts(&nameAppearances, attrs.Float)
	addFuncNameCounts(&nameAppearances, attrs.Int)
	addFuncNameCounts(&nameAppearances, attrs.Uint)

	duplicates := set.WithCapacity[string](32)
	for name, count := range nameAppearances {
		if count > 1 {
			duplicates.Add(name)
		}
	}

	return duplicates
}

func addFuncNameCounts(nameAppearances *map[string]int, a []attribute) {
	for _, attr := range a {
		(*nameAppearances)[attr.FuncName]++
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

func findDefinedGroup(elements []string) (string, bool) {
	for _, group := range definedElementGroups {
		if !slices.Equal(elements, group.elements) {
			continue
		}

		return group.name, true
	}

	return "", false
}
