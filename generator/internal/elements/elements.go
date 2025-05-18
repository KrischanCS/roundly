// Package elements implements the generation of the HTML element functions
// based on the HTML standard.
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
	// Tag is the name of the element, e.g. "div", "span", "h1", etc.
	Tag string
	// Description is a short description of the element.
	Description string
	// SemanticGroup defines to which broad group the element belongs.
	SemanticGroup string
	// Categories are the categories the element belongs to.
	Categories []string
	// Parents are all possible elements of this element
	Parents []string
	// Children are all possible elements of this element
	Children []string
	// Attributes are all allowed attributes of this element
	Attributes []string

	// DocumentationLink is the link to the documentation of this in the html standard
	DocumentationLink standard.Link
	// Links are all links in the description of this element
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

	group, ok := semanticGroupMapping[s]
	if !ok {
		panic("Unknown semantic group link: " + link.Url)
	}

	return group
}

//nolint:gochecknoglobals
var semanticGroupMapping = map[string]string{
	"https://html.spec.whatwg.org/dev/text-level-semantics.html":                "text",
	"https://html.spec.whatwg.org/dev/sections.html":                            "sections",
	"https://html.spec.whatwg.org/dev/image-maps.html":                          "imageMaps",
	"https://html.spec.whatwg.org/dev/media.html":                               "media",
	"https://html.spec.whatwg.org/dev/semantics.html":                           "semantics",
	"https://html.spec.whatwg.org/dev/grouping-content.html":                    "grouping",
	"https://html.spec.whatwg.org/dev/form-elements.html":                       "formElements",
	"https://html.spec.whatwg.org/dev/canvas.html":                              "canvas",
	"https://html.spec.whatwg.org/dev/tables.html":                              "tables",
	"https://html.spec.whatwg.org/dev/edits.html":                               "edits",
	"https://html.spec.whatwg.org/dev/interactive-elements.html":                "interactive",
	"https://html.spec.whatwg.org/dev/iframe-embed-object.html":                 "iframeEmbedObject",
	"https://html.spec.whatwg.org/dev/forms.html":                               "forms",
	"https://html.spec.whatwg.org/dev/embedded-content.html":                    "embedded",
	"https://html.spec.whatwg.org/dev/input.html":                               "input",
	"https://html.spec.whatwg.org/dev/https://w3c.github.io/mathml-core/":       "mathml",
	"https://html.spec.whatwg.org/dev/scripting.html":                           "scripting",
	"https://html.spec.whatwg.org/dev/https://svgwg.org/svg2-draft/struct.html": "svg",
}
