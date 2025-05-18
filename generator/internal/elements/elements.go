package elements

import (
	"fmt"
	"iter"
	"sort"
	"strings"

	"github.com/KrischanCS/go-toolbox/set"
	"golang.org/x/net/html"

	"github.com/KrischanCS/go-toolbox/iterator"

	"github.com/KrischanCS/htmfunc/generator/internal/standard"
)

type Element struct {
	Tag           string
	Description   string
	SemanticGroup string
	Categories    []string
	Parents       []string
	Children      []string
	Attributes    []string

	DocumentationLink standard.Link
	Links             []standard.Link
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

	e := iterator.Of(elements...)

	printUniqueSorted(e, "Categories", func(e Element) []string { return e.Categories })
	printUniqueSorted(e, "Parents", func(e Element) []string { return e.Parents })
	printUniqueSorted(e, "Children", func(e Element) []string { return e.Children })

	printSemanticGroups(e)

	return elements
}

func printSemanticGroups(e iter.Seq[Element]) {
	groups := make(map[string][]string)
	iterator.Reduce(e, &groups, func(accumulator *map[string][]string, value Element) {
		group, ok := (*accumulator)[value.SemanticGroup]
		if !ok {
			group = make([]string, 0, 8)
		}

		group = append(group, value.Tag)
		(*accumulator)[value.SemanticGroup] = group
	})

	for group, tags := range groups {
		sort.Strings(tags)

		fmt.Println(group)
		for _, tag := range tags {
			fmt.Printf("  %s\n", tag)
		}
	}
}

func printUniqueSorted(e iter.Seq[Element], name string, mapper func(e Element) []string) {
	categorySet := set.WithCapacity[string](64)
	for categories := range iterator.Map(e, mapper) {
		for _, c := range categories {
			categorySet.Add(c)
		}
	}

	c := categorySet.Values()
	sort.Strings(c)
	fmt.Println(name + ":")
	for _, category := range c {
		fmt.Printf("  %s\n", category)
	}

	fmt.Println()
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

func extractSemanticGroup(link standard.Link) string {
	s := strings.Split(link.Url, "#")[0]

	switch s {
	default:
		panic("Unkown semantic groupt link: " + link.Url)
	case "https://html.spec.whatwg.org/dev/text-level-semantics.html":
		return "text"
	case "https://html.spec.whatwg.org/dev/sections.html":
		return "sections"
	case "https://html.spec.whatwg.org/dev/image-maps.html":
		return "imageMaps"
	case "https://html.spec.whatwg.org/dev/media.html":
		return "media"
	case "https://html.spec.whatwg.org/dev/semantics.html":
		return "semantics"
	case "https://html.spec.whatwg.org/dev/grouping-content.html":
		return "grouping"
	case "https://html.spec.whatwg.org/dev/form-elements.html":
		return "formElements"
	case "https://html.spec.whatwg.org/dev/canvas.html":
		return "canvas"
	case "https://html.spec.whatwg.org/dev/tables.html":
		return "tables"
	case "https://html.spec.whatwg.org/dev/edits.html":
		return "edits"
	case "https://html.spec.whatwg.org/dev/interactive-elements.html":
		return "interactive"
	case "https://html.spec.whatwg.org/dev/iframe-embed-object.html":
		return "iframeEmbedObject"
	case "https://html.spec.whatwg.org/dev/forms.html":
		return "forms"
	case "https://html.spec.whatwg.org/dev/embedded-content.html":
		return "embedded"
	case "https://html.spec.whatwg.org/dev/input.html":
		return "input"
	case "https://html.spec.whatwg.org/dev/https://w3c.github.io/mathml-core/":
		return "mathml"
	case "https://html.spec.whatwg.org/dev/scripting.html":
		return "scripting"
	case "https://html.spec.whatwg.org/dev/https://svgwg.org/svg2-draft/struct.html":
		return "svg"
	}
}
