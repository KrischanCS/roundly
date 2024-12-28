package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strings"
	"text/template"

	"golang.org/x/net/html"
)

type attribute struct {
	Name        string
	FuncName    string
	ParamName   string
	Elements    []string
	Description string
	Value       string
	Links       []link
}

type link struct {
	Name string
	Url  string
}

var attrTemplate = template.Must(template.ParseFiles("attributes.go.tmpl"))

func main() {
	body := loadIndicesFromStandard()

	attributes := findAttributes(body)

	file, err := os.OpenFile("generated/attributes.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Print("Error closing file: ", err)
		}
	}(file)

	err = attrTemplate.Execute(file, attributes)
	if err != nil {
		log.Print("Error executing template: ", err)
	}
}

const htmlStandardUrl = `https://html.spec.whatwg.org/multipage/`

func loadIndicesFromStandard() *html.Node {
	const indicesUrl = htmlStandardUrl + "indices.html"

	response, err := http.Get(indicesUrl)
	if err != nil {
		log.Fatal("Error loading indices from standard: ", indicesUrl)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Print("Error closing body: ", err)
		}
	}(response.Body)

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
	attr.Value = text
	attr.Links = append(attr.Links, links...)

	return attr
}

func setNames(attr *attribute, attributeName string) {
	attr.Name = attributeName

	runes := []rune(attributeName)
	for len(runes) > 0 {
		i := slices.Index(runes, '-')
		if i == -1 {
			attr.ParamName += string(runes)
			break
		}

		attr.ParamName += string(runes[:i])
		attr.ParamName += strings.ToUpper(string(runes[i+1]))
		runes = runes[i+2:]
	}

	attr.FuncName = strings.ToUpper(attr.ParamName[0:1]) + attr.ParamName[1:]
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
					l.Url = htmlStandardUrl + attr.Val
					break
				}
			}

			text, ls := extractText(node)

			l.Name = text
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
