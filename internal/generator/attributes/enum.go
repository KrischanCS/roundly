package attributes

import (
	"log/slog"
	"sort"
	"strings"

	"github.com/KrischanCS/go-toolbox/iterator"
	"github.com/KrischanCS/go-toolbox/iterator/reducer"
	"github.com/KrischanCS/go-toolbox/set"

	"github.com/KrischanCS/roundly/internal/generator/standard"
)

func DecomposeEnums(enum []attribute) []attribute {
	slog.Debug("Decomposing enum attributes...")

	decomposed := make([]attribute, 0, len(enum)*3)

	for _, attr := range enum {
		for _, value := range attr.Values {
			value = strings.Trim(value, `[]"`)

			if value == "the empty string" {
				value = ""
			}

			newAttr := attr
			newAttr.FuncName = createEnumAttrFuncName(attr, value)
			newAttr.ParamName = ""
			newAttr.Value = value

			decomposed = append(decomposed, newAttr)
		}
	}

	attrs := disambiguateEnumAttrs(decomposed)

	sort.Slice(attrs, func(i, j int) bool {
		return attrs[i].FuncName < attrs[j].FuncName
	})

	slog.Info("Decomposed enum attributes.")

	return attrs
}

func createEnumAttrFuncName(attr attribute, value string) string {
	funcNameSuffix := value
	if value == "" {
		funcNameSuffix = "Empty"
	}

	funcName, ok := handleOrderedListTypeAttributes(attr, funcNameSuffix, attr.FuncName)
	if ok {
		return funcName
	}

	funcNameSuffix = replaceAndUppercaseNext(funcNameSuffix, '-')
	funcNameSuffix = replaceAndUppercaseNext(funcNameSuffix, '/')

	return attr.FuncName + strings.ToUpper(funcNameSuffix[:1]) + funcNameSuffix[1:]
}

// handleOrderedListTypeAttributes checks for an annoying special case... The
// single-character values of the type attribute of ol elements are
// case-sensitive, thus they cannot be converted simply to camel case as i and I
// as well as a and A would be converted to the same name.
//
// Instead, the generated function names chosen here are more explicit but
// longer:
//   - 1 -> Numeric
//   - a -> RomanLower
//   - b -> Roman
//   - i -> AlphaLower
//   - I -> Alpha
func handleOrderedListTypeAttributes(attr attribute, value string, funcName string) (string, bool) {
	if attr.Name != "type" {
		return funcName, false
	}

	switch value {
	default:
		return funcName, false
	case "1":
		return funcName + "Numeric", true
	case "i":
		return funcName + "RomanLower", true
	case "I":
		return funcName + "Roman", true
	case "a":
		return funcName + "AlphaLower", true
	case "A":
		return funcName + "Alpha", true
	}
}

func replaceAndUppercaseNext(funcNameSuffix string, char byte) string {
	for i := index(funcNameSuffix, char); i != -1; i = strings.IndexByte(funcNameSuffix, char) {
		funcNameSuffix = strings.Replace(funcNameSuffix, string(char), "", 1)
		funcNameSuffix = uppercaseAt(funcNameSuffix, i)
	}

	return funcNameSuffix
}

func index(text string, b byte) int {
	return strings.IndexByte(text, b)
}

func uppercaseAt(funcNameSuffix string, i int) string {
	return funcNameSuffix[:i] + strings.ToUpper(funcNameSuffix[i:i+1]) + funcNameSuffix[i+1:]
}

func disambiguateEnumAttrs(enumAttrs []attribute) []attribute {
	attrsByName := make(map[string][]attribute)

	iterator.Reduce(
		iterator.Of(enumAttrs...),
		&attrsByName,
		reducer.GroupBy(func(attr attribute) string { return attr.FuncName }),
	)

	disambiguated := make([]attribute, 0, len(attrsByName))

	for _, attrs := range attrsByName {
		if len(attrs) == 1 {
			disambiguated = append(disambiguated, attrs[0])
			continue
		}

		attrs[0].Elements = mergeElements(attrs)
		attrs[0].Values = mergeValues(attrs)
		attrs[0].Links = mergeLinks(attrs)

		disambiguated = append(disambiguated, attrs[0])
	}

	return disambiguated
}

func mergeElements(disambiguated []attribute) []string {
	elements := set.Of(disambiguated[0].Elements...)

	for _, a := range disambiguated[1:] {
		elements.Add(a.Elements...)
	}

	values := elements.Values()
	sort.Strings(values)

	return values
}

func mergeValues(disambiguated []attribute) []string {
	values := set.Of(disambiguated[0].Values...)

	for _, a := range disambiguated[1:] {
		values.Add(a.Values...)
	}

	v := values.Values()
	sort.Strings(v)

	return v
}

func mergeLinks(disambiguated []attribute) []standard.Link {
	links := set.Of(disambiguated[0].Links...)

	for _, a := range disambiguated[1:] {
		links.Add(a.Links...)
	}

	v := links.Values()
	sort.Slice(v, func(i, j int) bool {
		return v[i].Name < v[j].Name
	})

	return v
}
