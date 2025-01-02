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

// TODO parse https://html.spec.whatwg.org/dev/input.html#attr-input-type instead of hardcoding this
var inputTypes = []string{
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

	return classify(attrs)
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

func classify(attrs []*attribute) attributes {
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

		if attr.Value == "[Boolean attribute]" {
			attrsClassified.Bool = append(attrsClassified.Bool, *attr)
			continue
		}

		if enumPattern.MatchString(attr.Value) {
			// TODO should be exchanged for something like scanner to avoid the two pass/allocations here,
			//   but since this is not performance critical keep it as is for the moment.
			attr.Value = strings.ReplaceAll(attr.Value, "; ", ";")
			attr.Values = strings.Split(attr.Value, ";")
			attrsClassified.Enum = append(attrsClassified.Enum, *attr)
			continue
		}

		if strings.HasPrefix(attr.Value, "[input type keyword]") {
			attr.Values = inputTypes

			attrsClassified.InputType = append(attrsClassified.InputType, *attr)

			continue
		}

		if isCommaSeparatedList(attr.Value) {
			attrsClassified.ListComma = append(attrsClassified.ListComma, *attr)
			continue
		}

		if strings.HasPrefix(attr.Value, "[Valid list of floating-point numbers]") {
			attrsClassified.ListCommaFloat = append(attrsClassified.ListCommaFloat, *attr)
			continue
		}

		if isSpaceSeparatedList(attr.Value) {
			attrsClassified.ListSpace = append(attrsClassified.ListSpace, *attr)
			continue
		}

		if strings.HasPrefix(attr.Value, "[Valid floating-point number]") {
			attrsClassified.Float = append(attrsClassified.Float, *attr)
			continue
		}

		if strings.Contains(attr.Value, "[Valid integer]") {
			attrsClassified.Int = append(attrsClassified.Int, *attr)
			continue
		}

		if strings.HasPrefix(attr.Value, "[Valid non-negative integer]") {
			attrsClassified.Uint = append(attrsClassified.Uint, *attr)
			continue
		}

		attrsClassified.Text = append(attrsClassified.Text, *attr)
	}

	return attrsClassified
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

func findNodeWithId(node *html.Node, s string) *html.Node {
	for _, attr := range node.Attr {
		if attr.Key == "id" && attr.Val == s {
			return node
		}
	}

	for child := range node.ChildNodes() {
		if n := findNodeWithId(child, s); n != nil {
			return n
		}
	}

	if node.NextSibling != nil {
		return findNodeWithId(node.NextSibling, s)
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
			if node.Data == "a" {
				addAsLink(node, &links, &sb)
				continue
			}

			if node.Data == "li" {
				if node.PrevSibling != nil && node.PrevSibling.Data == "li" {
					sb.WriteByte(',')
				}
				sb.WriteByte(' ')
			}

			text, ls := extractText(node)
			sb.WriteString(text)
			links = append(links, ls...)
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

func addAsLink(node *html.Node, links *[]link, sb *strings.Builder) {
	l := link{}
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			l.Url = htmlStandardUrl + attr.Val
			break
		}
	}

	text, ls := extractText(node)

	l.Name = text
	*links = append(*links, l)

	sb.WriteString("[" + text + "]")
	*links = append(*links, ls...)
}

func compactWhitespace(sb *strings.Builder) {
	scanner := bufio.NewScanner(strings.NewReader(sb.String()))
	sb.Reset()

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
