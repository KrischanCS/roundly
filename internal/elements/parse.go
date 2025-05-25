package elements

import (
	"log"
	"log/slog"
	"strings"

	"github.com/KrischanCS/go-toolbox/set"
	"golang.org/x/net/html"

	"github.com/KrischanCS/htmfunc/internal/standard"
)

func parseAllElements(standardIndicesPage *html.Node) []Element {
	slog.Info("Creating element groups...")

	elementsTable, ok := findElementsTable(standardIndicesPage)
	if !ok {
		log.Panic("Could not find elements table")
	}

	elements := createElements(elementsTable)

	for i, element := range elements {
		// TODO link deduplication is probably neither here, nor in attributes complete.
		//  Everything rendered to doc comment must be considered, not only description
		//  as it is in attributes (Or only dedup of exact same as I did here).
		element.Links = standard.EliminateExactSameLinks(element.Links)
		elements[i] = element
	}

	slog.Info("Generated element groups.")

	return elements
}

func findElementsTable(page *html.Node) (*html.Node, bool) {
	// The elements table has no id, but a caption with the text
	// "List of elements", searching by that...
	const caption = "List of elements"

	return standard.FindTableWithCaption(page, caption)
}

func createElements(table *html.Node) []Element {
	slog.Debug("Searching tbody element of element index table...")

	tBody := standard.FindTBody(table)

	slog.Debug("Found tbody of element index table.")

	elements := make([]Element, 0)

	slog.Debug("Parsing table rows to elements...")

	for node := range tBody.ChildNodes() {
		if node.Type != html.ElementNode {
			panic("Expect only element nodes")
		}

		if node.Data != "tr" {
			panic("Expect only rows in tbody")
		}

		elements = append(elements, elementsFromRow(node)...)
	}

	slog.Debug("Parsed table nodes to elements.", "elementCount", len(elements))

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

	// TODO put link deduplication here

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

// TODO fetch this from the standard instead if hardcoding

//nolint:gochecknoglobals
var voidElementTags = set.Of(
	"area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param")

func elementFromRow(
	name string,
	documentationLink standard.Link,
	descriptionNode *html.Node,
) Element {
	element := Element{
		Tag:               name,
		FuncName:          strings.ToUpper(string(name[0])) + name[1:],
		DocumentationLink: documentationLink,
	}

	element.IsVoid = voidElementTags.Contains(name)

	description, links := standard.ExtractText(descriptionNode)
	element.Description = strings.Trim(description, "[]")
	element.Links = links

	element.SemanticGroup = extractSemanticGroup(documentationLink)

	categoriesNode := descriptionNode.NextSibling
	element.Categories, links = extractTokens(categoriesNode)
	element.Links = append(element.Links, links...)

	parentsNode := categoriesNode.NextSibling
	element.Parents, links = extractTokens(parentsNode)
	element.Links = append(element.Links, links...)

	childrenNode := parentsNode.NextSibling
	element.Children, links = extractTokens(childrenNode)
	element.Links = append(element.Links, links...)

	attributesNode := childrenNode.NextSibling
	element.Attributes, links = extractTokens(attributesNode)
	element.Links = append(element.Links, links...)

	return element
}

func extractTokens(node *html.Node) ([]string, []standard.Link) {
	text, links := standard.ExtractText(node)

	categories := strings.Split(text, ";")
	for i := range categories {
		categories[i] = strings.Trim(categories[i], "*")
	}

	return categories, links
}
