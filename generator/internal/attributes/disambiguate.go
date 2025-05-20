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
func disambiguateAttrs(attrs *attributes) {
	attrs.Text = mergeAttributes(attrs.Text)
	attrs.Bool = mergeAttributes(attrs.Bool)
	attrs.InputType = mergeAttributes(attrs.InputType)
	attrs.ListComma = mergeAttributes(attrs.ListComma)
	attrs.ListSpace = mergeAttributes(attrs.ListSpace)
	attrs.ListCommaFloat = mergeAttributes(attrs.ListCommaFloat)
	attrs.Float = mergeAttributes(attrs.Float)
	attrs.Int = mergeAttributes(attrs.Int)
	attrs.Uint = mergeAttributes(attrs.Uint)

	duplicates := findDuplicatedNames(*attrs)

	appendSuffixToDuplicates(attrs.Bool, "True", duplicates)
	appendSuffixToDuplicates(attrs.InputType, "InputType", duplicates)
	appendSuffixToDuplicates(attrs.ListComma, "Strings", duplicates)
	appendSuffixToDuplicates(attrs.ListSpace, "Strings", duplicates)
	appendSuffixToDuplicates(attrs.ListCommaFloat, "Floats", duplicates)
	appendSuffixToDuplicates(attrs.Float, "Float", duplicates)
	appendSuffixToDuplicates(attrs.Int, "Int", duplicates)
	appendSuffixToDuplicates(attrs.Uint, "UInt", duplicates)
}

func mergeAttributes(attrs []attribute) []attribute {
	merged := make([]attribute, 0, len(attrs))
	visitedDuplicates := set.WithCapacity[int](len(attrs))

	for i, attr := range attrs {
		if visitedDuplicates.Contains(i) {
			continue
		}

		duplicates := make([]attribute, 0)

		for j, other := range attrs[i+1:] {
			if attr.FuncName != other.FuncName || visitedDuplicates.Contains(i+j+1) {
				continue
			}

			visitedDuplicates.Add(i+j+1)
			duplicates = append(duplicates, other)
		}

		if len(duplicates) == 0 {
			merged = append(merged, attr)
			continue
		}

		newAttr := joinDuplicateAttrsOfSameType(attr, duplicates)

		merged = append(merged, newAttr)
	}

	return merged
}

func joinDuplicateAttrsOfSameType(attr attribute, duplicates []attribute) attribute {
	newAttr := attribute{
		Name:        attr.Name,
		FuncName:    attr.FuncName,
		ParamName:   attr.ParamName,
		Elements:    make([]string, 0, len(duplicates)+1),
		Description: "It's semantics varies depending on the element it is applied to.",
		Values:      attr.Values,
		Links:       attr.Links,
	}

	newAttr.Elements = joinElementsAndDescription(newAttr.Elements, attr)

	for _, duplicate := range duplicates {
		newAttr.Elements = joinElementsAndDescription(newAttr.Elements, duplicate)

		for _, value := range duplicate.Values {
			if !slices.Contains(newAttr.Values, value) {
				newAttr.Elements = append(newAttr.Values, value)
			}
		}

		for _, link := range duplicate.Links {
			if !slices.Contains(newAttr.Links, link) {
				newAttr.Links = append(newAttr.Links, link)
			}
		}
	}

	return newAttr
}

func joinElementsAndDescription(elements []string, attr attribute) []string {
	return append(elements, strings.Join(attr.Elements, " ")+": "+attr.Description)
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
