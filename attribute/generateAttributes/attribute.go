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
	Text  []attribute
	Bool  []attribute
	Enum  []attribute
	Other []attribute
}

type link struct {
	Name string
	Url  string
}

func findAttributes(body *html.Node) attributes {
	attributesTable := findNodeWithId(body, "attributes-1")
	if attributesTable == nil {
		log.Fatal("Error finding attributes table")
	}

	tBody := findTBody(attributesTable)

	attrs := attributes{
		Text:  make([]attribute, 0, 16),
		Bool:  make([]attribute, 0, 16),
		Other: make([]attribute, 0, 16),
	}

	for row := range tBody.ChildNodes() {
		attr := parseAttribute(row)

		classifyAndAdd(&attrs, attr)
	}

	return attrs
}

var enumPattern = regexp.MustCompile(`^".*?"(;".*?")*$`)

func classifyAndAdd(attrs *attributes, attr attribute) {
	if strings.HasPrefix(attr.Value, "[Text]") && !strings.Contains(attr.Value, ";") {
		attrs.Text = append(attrs.Text, attr)
		return
	}

	if attr.Value == "[Boolean attribute]" {
		attrs.Bool = append(attrs.Bool, attr)
		return
	}

	if enumPattern.MatchString(attr.Value) {
		// TODO should be exchanged for something like scanner to avoid the two pass/allocations here,
		//   but since this is not performance critical keep it as is for the moment.
		attr.Value = strings.ReplaceAll(attr.Value, "; ", ";")
		attr.Values = strings.Split(attr.Value, ";")
		attrs.Enum = append(attrs.Enum, attr)
		return
	}

	attrs.Other = append(attrs.Other, attr)
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
			if node.Type == html.ElementNode && node.Data == "a" {
				addAsLink(node, &links, &sb)
				continue
			}

			text, ls := extractText(node)
			sb.WriteString(text)
			links = append(links, ls...)
		default:
			log.Panic("Unexpected node type: ", node.Type)
		}
	}

	compactWhitespace(&sb)

	return sb.String(), links
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
