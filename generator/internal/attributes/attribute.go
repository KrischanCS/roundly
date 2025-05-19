package attributes

import (
	"log"
	"regexp"
	"strings"

	"github.com/KrischanCS/go-toolbox/set"
	"golang.org/x/net/html"

	"github.com/KrischanCS/htmfunc/generator/internal/standard"
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
	InputType      []attribute
	ListComma      []attribute
	ListCommaFloat []attribute
	ListSpace      []attribute
	Float          []attribute
	Int            []attribute
	Uint           []attribute
}

// TODO parse https://html.spec.whatwg.org/dev/input.html#attr-input-type instead of hardcoding this.
var inputTypes = []string{ //nolint:gochecknoglobals
	"hidden",
	"text",
	"search",
	"tel",
	"url",
	"email",
	"password",
	"date",
	"month",
	"week",
	"time",
	"datetime",
	"number",
	"range",
	"color",
	"checkbox",
	"radio",
	"file",
	"submit",
	"image",
	"reset",
	"button",
}

func findAttributes(body *html.Node) attributes {
	attributesTable := standard.FindNodeWithId(body, "attributes-1")
	if attributesTable == nil {
		log.Fatal("Error finding attributes table")
	}

	tBody := standard.FindTBody(attributesTable)

	attrsByName := make(map[string][]*attribute)
	attrs := make([]*attribute, 0, 256) //nolint:mnd

	for row := range tBody.ChildNodes() {
		attr := parseAttribute(row)
		addByName(attrsByName, &attr)
		attrs = append(attrs, &attr)
	}

	attrsClassified := classifyAttributes(attrs)
	attrsClassified.Enum = DecomposeEnums(attrsClassified.Enum)

	// TODO adapt this to handle attributesType
	//disambiguateAttrs(attrsClassified)

	return attrsClassified
}

func DecomposeEnums(enum []attribute) []attribute {
	decomposed := make([]attribute, 0, len(enum)*3)

	for _, attr := range enum {
		for _, value := range attr.Values {
			value = strings.Trim(value, `[]"`)

			if value == "the empty string" {
				value = "empty"
			}

			for i := strings.IndexByte(value, '-'); i != -1; i = strings.IndexByte(value, '-') {
				value = strings.Replace(value, "-", "", 1)
				value = value[:i] + strings.ToUpper(value[i:i+1]) + value[i+1:]
			}

			for i := strings.IndexByte(value, '/'); i != -1; i = strings.IndexByte(value, '/') {
				value = strings.Replace(value, "/", "", 1)
				value = value[:i] + strings.ToUpper(value[i:i+1]) + value[i+1:]
			}

			funcName := attr.FuncName + strings.ToUpper(value[:1]) + value[1:]

			funcName = handleOrderedListTypeAttributes(attr, value, funcName)

			newAttr := attr
			newAttr.FuncName = funcName
			newAttr.ParamName = ""
			newAttr.Value = value

			decomposed = append(decomposed, newAttr)
		}
	}

	return disambiguateEnumAttrs(decomposed)

}

// handleOrderedListTypeAttributes checks for an annoying special case… The
// single-character values of the type attribute of ol elements are
// case-sensitive, thus they cannot be converted simply to camel case as i and I
// as well as a and A would be converted to the same name.
//
// Instead, the generated function names chosen here are more explicit but
// longer:
//   - 1 -> Numbered
//   - a -> RomanLower
//   - b -> RomanUpper
//   - i -> AlphaLower
//   - I -> AlphaUpper
func handleOrderedListTypeAttributes(attr attribute, value string, funcName string) (string, bool) {
	if attr.Name != "type" {
		return funcName, false
	}

	switch value {
	default:
		return funcName, false
	case "1":
		return "Numbered", true
	case "i":
		return "RomanLower", true
	case "I":
		return "RomanUpper", true
	case "a":
		return "AlphaLower", true
	case "A":
		return "AlphaUpper", true
	}
}

func disambiguateEnumAttrs(enumAttrs []attribute) []attribute {
	duplicates := make(map[string]set.Set[attribute])
	duplicateIndices := set.WithCapacity[int](16)

	for i, this := range enumAttrs {

		for j, other := range enumAttrs {
			if this.Name != other.Name {
				continue
			}


		}
	}
}

func findEventHandlerAttributes(body *html.Node) []attribute {
	eventHandlersTable := standard.FindNodeWithId(body, "ix-event-handlers")
	if eventHandlersTable == nil {
		log.Fatal("Error finding event handlers table")
	}

	tBody := standard.FindTBody(eventHandlersTable)

	attrs := make([]attribute, 0, 256) //nolint:mnd

	for row := range tBody.ChildNodes() {
		attr := parseAttribute(row)
		attrs = append(attrs, attr)
	}

	return attrs
}

func addByName(attrsByName map[string][]*attribute, attr *attribute) {
	attrsWithName, ok := attrsByName[attr.Name]
	if !ok {
		attrsWithName = make([]*attribute, 0, 1)
	}

	attrsWithName = append(attrsWithName, attr)
	attrsByName[attr.Name] = attrsWithName
}

var enumPattern = regexp.MustCompile(`^".*?"(;".*?")*(;the empty string)?$`)

//nolint:mnd
func classifyAttributes(attrs []*attribute) attributes {
	attrsClassified := attributes{
		Text:           make([]attribute, 0, 16),
		Bool:           make([]attribute, 0, 16),
		Enum:           make([]attribute, 0, 16),
		InputType:      make([]attribute, 0, 16),
		ListComma:      make([]attribute, 0, 16),
		ListCommaFloat: make([]attribute, 0, 16),
		ListSpace:      make([]attribute, 0, 16),
		Float:          make([]attribute, 0, 16),
		Int:            make([]attribute, 0, 16),
		Uint:           make([]attribute, 0, 16),
	}

	for _, attr := range attrs {
		classifyAttribute(&attrsClassified, attr)
	}

	return attrsClassified
}

func classifyAttribute(attrs *attributes, attr *attribute) {
	if attr.Value == "[Boolean attribute]" {
		attrs.Bool = append(attrs.Bool, *attr)
		return
	}

	if enumPattern.MatchString(attr.Value) {
		// TODO should be exchanged for something like scanner to avoid the two pass/allocations here,
		//   but since this is not performance critical keep it as is for the moment.
		attr.Value = strings.ReplaceAll(attr.Value, "; ", ";")
		attr.Values = strings.Split(attr.Value, ";")
		attrs.Enum = append(attrs.Enum, *attr)

		return
	}

	if strings.HasPrefix(attr.Value, "[input type keyword]") {
		attr.Values = inputTypes

		attrs.InputType = append(attrs.InputType, *attr)

		return
	}

	if isCommaSeparatedList(attr.Value) {
		attrs.ListComma = append(attrs.ListComma, *attr)
		return
	}

	if strings.HasPrefix(attr.Value, "[Valid list of floating-point numbers]") {
		attrs.ListCommaFloat = append(attrs.ListCommaFloat, *attr)
		return
	}

	if isSpaceSeparatedList(attr.Value) {
		attrs.ListSpace = append(attrs.ListSpace, *attr)
		return
	}

	if strings.HasPrefix(attr.Value, "[Valid floating-point number]") {
		attrs.Float = append(attrs.Float, *attr)
		return
	}

	if strings.Contains(attr.Value, "[Valid integer]") {
		attrs.Int = append(attrs.Int, *attr)
		return
	}

	if strings.HasPrefix(attr.Value, "[Valid non-negative integer]") {
		attrs.Uint = append(attrs.Uint, *attr)
		return
	}

	attrs.Text = append(attrs.Text, *attr)
}

func isCommaSeparatedList(value string) bool {
	if strings.HasPrefix(value, "[Set of comma-separated tokens]") {
		return true
	}

	return strings.HasPrefix(value, "Comma-separated list of")
}

func isSpaceSeparatedList(value string) bool {
	if strings.HasPrefix(value, "[Ordered set of unique space-separated tokens]") {
		return true
	}

	if strings.HasPrefix(value, "[Unordered set of unique space-separated tokens]") {
		return true
	}

	return strings.HasPrefix(value, "[Set of space-separated tokens]")
}

func parseAttribute(row *html.Node) attribute {
	data := row.FirstChild

	var attr attribute

	text, links := standard.ExtractText(data)
	setNames(&attr, text)
	attr.Links = links

	data = data.NextSibling

	elements, links := extractElements(data)
	attr.Elements = elements
	attr.Links = append(attr.Links, links...)

	data = data.NextSibling

	text, links = standard.ExtractText(data)
	attr.Description = text
	attr.Links = append(attr.Links, links...)

	data = data.NextSibling

	text, links = standard.ExtractText(data)
	text = strings.ReplaceAll(text, "*", " (Additional rules apply, see elements documentation)")

	attr.Value = text
	attr.Links = append(attr.Links, links...)

	return attr
}

func setNames(attr *attribute, attributeName string) {
	attr.Name = attributeName

	if mapping, ok := mappings[attributeName]; ok {
		attr.FuncName = mapping.FuncName
		attr.ParamName = mapping.paramName

		return
	}

	attr.ParamName = attributeName
	attr.FuncName = strings.ToUpper(attr.ParamName[0:1]) + attr.ParamName[1:]
}

func extractElements(data *html.Node) ([]string, []standard.Link) {
	text, links := standard.ExtractText(data)

	return strings.Split(text, ";"), links
}
