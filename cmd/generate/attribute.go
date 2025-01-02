package main

import (
	"bufio"
	"bytes"
	"golang.org/x/net/html"
	"log"
	"regexp"
	"strings"
)

type attribute struct {
	Name        string
	FuncName    string
	ParamName   string
	Elements    []string
	Description string
	Value       string
	Values      []string
	Links       []link
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

type link struct {
	Name string
	Url  string
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
	attributesTable := findNodeWithId(body, "attributes-1")
	if attributesTable == nil {
		log.Fatal("Error finding attributes table")
	}

	tBody := findTBody(attributesTable)

	attrsByName := make(map[string][]*attribute)
	attrs := make([]*attribute, 0, 256)

	for row := range tBody.ChildNodes() {
		attr := parseAttribute(row)
		addByName(attrsByName, &attr)
		attrs = append(attrs, &attr)
	}

	disambiguateAttrs(attrsByName)

	return classifyAttributes(attrs)
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

func findNodeWithId(node *html.Node, id string) *html.Node {
	for _, attr := range node.Attr {
		if attr.Key == "id" && attr.Val == id {
			return node
		}
	}

	for child := range node.ChildNodes() {
		if n := findNodeWithId(child, id); n != nil {
			return n
		}
	}

	if node.NextSibling != nil {
		return findNodeWithId(node.NextSibling, id)
	}

	return nil
}

func findTBody(table *html.Node) *html.Node {
	for current := table.FirstChild; current != nil; current = current.NextSibling {
		if current.Data == "tbody" {
			return current
		}
	}

	log.Fatal("Error finding tbody")

	return nil
}

func parseAttribute(row *html.Node) attribute {
	data := row.FirstChild

	var attr attribute

	text, links := extractText(data)
	setNames(&attr, text)
	attr.Links = links

	data = data.NextSibling

	elements, links := extractElements(data)
	attr.Elements = elements
	attr.Links = append(attr.Links, links...)

	data = data.NextSibling

	text, links = extractText(data)
	attr.Description = text
	attr.Links = append(attr.Links, links...)

	data = data.NextSibling

	text, links = extractText(data)
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

func extractText(data *html.Node) (string, []link) {
	var sb strings.Builder
	links := make([]link, 0, 3)

	for node := data.FirstChild; node != nil; node = node.NextSibling {
		switch node.Type {
		case html.CommentNode:
			continue
		case html.TextNode:
			if isAllWhitespace(node.Data) {
				continue
			}

			sb.WriteString(node.Data)
		case html.ElementNode:
			extractTextFromElement(node, &links, &sb)
		default:
			log.Panic("Unexpected node type: ", node.Type)
		}
	}

	compactWhitespace(&sb)

	s := sb.String()

	// Special case currently needed for "popover" which has a trailing ;
	s = strings.Trim(s, ";")

	return s, links
}

func extractTextFromElement(node *html.Node, links *[]link, sb *strings.Builder) {
	if node.Data == "a" {
		addAsLink(node, links, sb)
		return
	}

	if node.Data == "li" {
		if node.PrevSibling != nil && node.PrevSibling.Data == "li" {
			sb.WriteByte(',')
		}

		sb.WriteByte(' ')
	}

	text, ls := extractText(node)

	sb.WriteString(text)

	*links = append(*links, ls...)
}

func addAsLink(node *html.Node, links *[]link, sb *strings.Builder) {
	link := link{}

	for _, attr := range node.Attr {
		if attr.Key == "href" {
			link.Url = htmlStandardUrl + attr.Val
			break
		}
	}

	text, ls := extractText(node)

	link.Name = text
	*links = append(*links, link)

	sb.WriteString("[" + text + "]")

	*links = append(*links, ls...)
}

func compactWhitespace(sb *strings.Builder) {
	scanner := bufio.NewScanner(strings.NewReader(sb.String()))
	sb.Reset()

	//nolint:wsl
	for scanner.Scan() {
		bs := scanner.Bytes()
		bs = bytes.Trim(bs, " \t")

		lastWasWhitespace := false
		for _, b := range bs {
			if b != ' ' && b != '\t' {
				lastWasWhitespace = false
				sb.WriteByte(b)
				continue
			}

			if lastWasWhitespace {
				continue
			}

			lastWasWhitespace = true
			sb.WriteByte(' ')
		}
	}
}

var allWhitespace = regexp.MustCompile(`^[ \t\n\r]*$`)

func isAllWhitespace(data string) bool {
	return allWhitespace.MatchString(data)
}

func extractElements(data *html.Node) ([]string, []link) {
	text, links := extractText(data)

	return strings.Split(text, ";"), links
}
