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
	Categorys     []string
	Parents       []string
	Children      []string

	Links []standard.Link
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
	names, _ := standard.ExtractText(nameNode)

	elements := make([]Element, 0, 1)
	for s := range strings.SplitSeq(names, ", ") {
		if s == "[autonomous custom elements]" {
			continue
		}

		elements = append(elements, elementFromRow(s, nameNode.NextSibling))
	}

	return elements
}

func elementFromRow(name string, descriptionNode *html.Node) Element {
	name = strings.Trim(name, "[]")

	if strings.Contains(name, " ") {
		// Spaces are not allowed in element names. Known special cases are:
		//   - "SVG svg" -> "svg"
		//   - "MathML math" -> "math"
		switch name {
		case "SVG svg":
			name = "svg"
		case "MathML math":
			name = "math"
		default:
			panic("Unexpected element name with space: " + name)
		}
	}

	element := Element{
		Tag: name,
	}

	description, links := standard.ExtractText(descriptionNode)
	element.Description = strings.Trim(description, "[]")
	element.Links = links

	return element
}
