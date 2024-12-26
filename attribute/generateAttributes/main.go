package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type attribute struct {
	name        string
	elements    []string
	description string
	value       string
	links       []link
}

type link struct {
	name string
	url  string
}

const (
	attributeFuncTemplate = `
// %s creates the %s attribute - %s
// 
// It can be applied to the following elements:
%s
//
// Value constraints: %s
//
// Source: [The HTML Standard]
//
%s
// [The HTML Standard]: https://html.spec.whatwg.org/multipage/
func %s(%s string) htmfunc.AttributeRenderer {
	return attr.Attribute("%s", %s)
}

`
	elementListTemplate = "//   - %s"
	linkTemplate        = "// [%s]: %s"
)

func main() {
	body := loadIndicesFromStandard()

	attributes := findAttributes(body)

	file, err := os.OpenFile("generated/attributes.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, attr := range attributes {
		var linkSb strings.Builder

		for i, l := range attr.links {
			if i != 0 {
				linkSb.WriteString("\n")
			}
			linkSb.WriteString(fmt.Sprintf(linkTemplate, l.name, l.url))
		}

		var elementSb strings.Builder
		for i, e := range attr.elements {
			if i != 0 {
				elementSb.WriteString("\n")
			}
			elementSb.WriteString(fmt.Sprintf(elementListTemplate, e))
		}

		title := strings.Title(attr.name)
		fmt.Fprintf(file, attributeFuncTemplate,
			title,
			attr.name,
			attr.description,
			elementSb.String(),
			attr.value,
			linkSb.String(),
			title,
			strings.ReplaceAll(attr.name, "type", "typeValue"),
			attr.name,
			strings.ReplaceAll(attr.name, "type", "typeValue"))
	}
}

const htmlStandardUrl = `https://html.spec.whatwg.org/multipage/`

func loadIndicesFromStandard() *html.Node {
	const indicesUrl = htmlStandardUrl + "indices.html"

	response, err := http.Get(indicesUrl)
	if err != nil {
		log.Fatal("Error loading indices from standard: ", indicesUrl)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatal("Unexpected status loading indices from standard: ", response.StatusCode)
	}

	body, err := html.Parse(response.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", response.StatusCode)
	}

	return body
}

func findAttributes(body *html.Node) []attribute {
	attributesTable := findNodeWithId(body, "attributes-1")
	if attributesTable == nil {
		log.Fatal("Error finding attributes table")
	}

	tBody := findTBody(attributesTable)

	attributes := make([]attribute, 0, 64)
	for row := range tBody.ChildNodes() {
		attr := parseAttribute(row)
		attributes = append(attributes, attr)
	}

	return attributes
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
	attr.name = text
	attr.links = links

	data = data.NextSibling

	elements, links := extractElements(data)
	attr.elements = elements
	attr.links = append(attr.links, links...)

	data = data.NextSibling

	text, links = extractText(data)
	attr.description = text
	attr.links = append(attr.links, links...)

	data = data.NextSibling

	text, links = extractText(data)
	attr.value = text
	attr.links = append(attr.links, links...)

	return attr
}

func extractText(data *html.Node) (string, []link) {
	var sb strings.Builder
	links := make([]link, 0, 3)

	for node := data.FirstChild; node != nil; node = node.NextSibling {
		if node.Type == html.CommentNode {
			continue
		}

		if node.Type == html.TextNode && !isAllWhitespace(node.Data) {
			sb.WriteString(node.Data)
			continue
		}

		if node.Type == html.ElementNode && node.Data == "a" {
			l := link{}
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					l.url = htmlStandardUrl + attr.Val
					break
				}
			}

			text, ls := extractText(node)

			l.name = text
			links = append(links, l)

			sb.WriteString("[" + text + "]")
			links = append(links, ls...)
			continue
		}

		text, ls := extractText(node)
		sb.WriteString(text)
		links = append(links, ls...)
	}

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

	return sb.String(), links
}

var allWhitespace = regexp.MustCompile(`^[ \t\n\r]*$`)

func isAllWhitespace(data string) bool {
	return allWhitespace.MatchString(data)
}

func extractElements(data *html.Node) ([]string, []link) {
	text, links := extractText(data)

	return strings.Split(text, ";"), links
}
