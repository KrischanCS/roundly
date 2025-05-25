// Package elements implements the generation of the HTML element functions
// based on the HTML standard.
package elements

import (
	"embed"
	"fmt"
	"iter"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/KrischanCS/go-toolbox/set"
	"golang.org/x/net/html"

	"github.com/KrischanCS/go-toolbox/iterator"

	standard2 "github.com/KrischanCS/htmfunc/internal/standard"
)

//go:embed templates
var templateFs embed.FS

//nolint:gochecknoglobals
var elementTemplate = template.Must(template.ParseFS(templateFs, "templates/*.go.tmpl"))

type Element struct {
	FuncName string

	IsVoid bool

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
	DocumentationLink standard2.Link
	// Links are all links in the description of this element
	Links []standard2.Link
}

func GenerateElements(standardIndicesPage *html.Node) {
	elementGroups := getAllElements(standardIndicesPage)

	for group, elements := range elementGroups {
		generateFile(group, elements)
	}
}

func generateFile(group string, elements []Element) {
	filePath := filepath.Join("..", "element", group+".go")

	//nolint:gosec // Files are written for everyone
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(fmt.Sprintf("Error opening file %s: %s", group+".go", err))
	}
	defer file.Close() //nolint:errcheck

	err = elementTemplate.Lookup("element.go.tmpl").Execute(file, elements)
	if err != nil {
		panic(fmt.Sprintf("Error executing template %s: %s", group, err))
	}
}

func getAllElements(standardIndicesPage *html.Node) map[string][]Element {
	elementsTable, ok := findElementsTable(standardIndicesPage)
	if !ok {
		panic("Could not find elements table")
	}

	elements := createElements(elementsTable)

	for i, element := range elements {
		// TODO link deduplication is probably neither here, nor in attributes complete.
		//  Everything rendered to doc comment must be considered, not only description
		//  as it is in attributes (Or only dedup of exact same as I did here).
		element.Links = standard2.EliminateExactSameLinks(element.Links)
		elements[i] = element
	}

	e := iterator.Of(elements...)

	return groupElements(e)
}

func groupElements(e iter.Seq[Element]) map[string][]Element {
	groups := make(map[string][]Element)
	iterator.Reduce(e, &groups, func(accumulator *map[string][]Element, value Element) {
		group, ok := (*accumulator)[value.SemanticGroup]
		if !ok {
			group = make([]Element, 0)
		}

		group = append(group, value)
		(*accumulator)[value.SemanticGroup] = group
	})

	return groups
}

func findElementsTable(page *html.Node) (*html.Node, bool) {
	// The elements table has no id, but a caption with the text
	// "List of elements", searching by that…
	const caption = "List of elements"

	return standard2.FindTableWithCaption(page, caption)
}

func createElements(table *html.Node) []Element {
	tBody := standard2.FindTBody(table)

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
	names, links := standard2.ExtractText(nameNode)

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

// TODO fetch this from the standard instead if hardcoding

//nolint:gochecknoglobals
var voidElementTags = set.Of(
	"area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param")

func elementFromRow(
	name string,
	documentationLink standard2.Link,
	descriptionNode *html.Node,
) Element {
	element := Element{
		Tag:               name,
		FuncName:          strings.ToUpper(string(name[0])) + name[1:],
		DocumentationLink: documentationLink,
	}

	element.IsVoid = voidElementTags.Contains(name)

	description, links := standard2.ExtractText(descriptionNode)
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

func extractTokens(node *html.Node) ([]string, []standard2.Link) {
	text, links := standard2.ExtractText(node)

	categories := strings.Split(text, ";")
	for i := range categories {
		categories[i] = strings.Trim(categories[i], "*")
	}

	return categories, links
}

func extractSemanticGroup(link standard2.Link) string {
	s := strings.Split(link.Url, "#")[0]
	s, _ = strings.CutPrefix(s, "https://html.spec.whatwg.org/dev/")

	group, ok := semanticGroupMapping[s]
	if !ok {
		panic("Unknown semantic group link: " + link.Url)
	}

	return group
}

//nolint:gochecknoglobals
var semanticGroupMapping = map[string]string{
	"text-level-semantics.html":                "textSemantics",
	"sections.html":                            "sections",
	"image-maps.html":                          "imageMaps",
	"media.html":                               "media",
	"semantics.html":                           "semantics",
	"grouping-content.html":                    "grouping",
	"form-elements.html":                       "formElements",
	"canvas.html":                              "canvas",
	"tables.html":                              "tables",
	"edits.html":                               "edits",
	"interactive-elements.html":                "interactive",
	"iframe-embed-object.html":                 "iframeEmbedObject",
	"forms.html":                               "forms",
	"embedded-content.html":                    "embedded",
	"input.html":                               "input",
	"https://w3c.github.io/mathml-core/":       "mathml",
	"scripting.html":                           "scripting",
	"https://svgwg.org/svg2-draft/struct.html": "svg",
}
