package elements

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"

	"github.com/KrischanCS/htmfunc/generator/internal/standard"
)

type Element struct {
	Tag           string
	Description   string
	SemanticGroup string
	Categories    []string
	Parents       []string
	Children      []string

	Links             []standard.Link
	DocumentationLink standard.Link
}

func GenerateElements(standardIndicesPage *html.Node) {
	getAllElements(standardIndicesPage)
}

func getAllElements(standardIndicesPage *html.Node) []Element {
	elementsTable, ok := findElementsTable(standardIndicesPage)
	if !ok {
		panic("Could not find elements table")
	}

	elements := createElements(elementsTable)

	for _, element := range elements {
		fmt.Println(element)
	}

	return elements
}

func findElementsTable(page *html.Node) (*html.Node, bool) {
	// The elements table has no id, but a caption with the text
	// "List of elements", searching by that…

	const caption = "List of elements"

	return standard.FindTableWithCaption(page, caption)
}

func createElements(table *html.Node) []Element {
	tBody := standard.FindTBody(table)

	elements := make([]Element, 0)
	for node := range tBody.ChildNodes() {
		if node.Type != html.ElementNode {
			panic("Expect only element nodes")
		}

		if node.Data != "tr" {
			panic("Expect only rows in tbody")
		}

		elements = append(elements, elementsFromRow(node)...)
	}

	return elements
}

func elementsFromRow(node *html.Node) []Element {
	nameNode := node.FirstChild
	if nameNode.Type != html.ElementNode || nameNode.Data != "td" && nameNode.Data != "th" {
		panic("Expect first child to be a td, was: " + nameNode.Data)
	}

	// There exist one case with multiple elements in a single ro
	// (h1, h2, h3, h4, h5, h6)
	names, links := standard.ExtractText(nameNode)

	elements := make([]Element, 0, 1)
	for name := range strings.SplitSeq(names, ", ") {
		name, ok := normalizeName(name)
		if !ok {
			continue
		}

		mainLink := links[0]
		mainLink.Name = "(More)"

		elements = append(elements, elementFromRow(name, mainLink, nameNode.NextSibling))
	}

	return elements
}

func normalizeName(name string) (string, bool) {
	name = strings.Trim(name, "[]")

	if !strings.Contains(name, " ") {
		return name, true
	}

	switch name {
	case "autonomous custom elements": // Not an element
		return "", false
	case "SVG svg":
		return "svg", true
	case "MathML math":
		return "math", true
	default:
		panic("Unexpected element name with space: " + name)
	}
}

func elementFromRow(name string, documentationLink standard.Link, descriptionNode *html.Node) Element {
	element := Element{
		Tag:               name,
		DocumentationLink: documentationLink,
	}

	description, links := standard.ExtractText(descriptionNode)
	element.Description = strings.Trim(description, "[]")
	element.Links = links

	return element
}
