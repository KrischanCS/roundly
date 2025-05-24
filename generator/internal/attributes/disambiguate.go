package attributes

import (
	"slices"
	"strings"

	"github.com/KrischanCS/go-toolbox/set"
)

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

		duplicates := findAttrDuplicates(attrs, i, attr, visitedDuplicates)

		if len(duplicates) == 0 {
			merged = append(merged, attr)
			continue
		}

		newAttr := joinDuplicateAttrsOfSameType(attr, duplicates)

		merged = append(merged, newAttr)
	}

	return merged
}

func findAttrDuplicates(
	attrs []attribute,
	i int,
	attr attribute,
	visitedDuplicates set.Set[int],
) []attribute {
	duplicates := make([]attribute, 0)

	for j, other := range attrs[i+1:] {
		if attr.FuncName != other.FuncName || visitedDuplicates.Contains(i+j+1) {
			continue
		}

		visitedDuplicates.Add(i + j + 1)

		duplicates = append(duplicates, other)
	}

	return duplicates
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
		mergeAttrValues(duplicate, newAttr)
		mergeAttrLinks(duplicate, newAttr)
	}

	return newAttr
}

// TODO code duplication of mergeAttrValues and mergeAttrLinks

func mergeAttrValues(duplicate attribute, newAttr attribute) {
	for _, value := range duplicate.Values {
		if !slices.Contains(newAttr.Values, value) {
			newAttr.Values = append(newAttr.Values, value)
		}
	}
}

func mergeAttrLinks(duplicate attribute, newAttr attribute) {
	for _, link := range duplicate.Links {
		if !slices.Contains(newAttr.Links, link) {
			newAttr.Links = append(newAttr.Links, link)
		}
	}
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

	duplicates := set.WithCapacity[string](32) //nolint:mnd

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
