package elements

import (
	"fmt"

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
}

func GenerateElements(standardIndicesPage *html.Node) {
	getAllElements(standardIndicesPage)
}

func getAllElements(standardIndicesPage *html.Node) []Element {
	elementsTable, ok := findElementsTable(standardIndicesPage)
	if !ok {
		panic("Could not find elements table")
	}

	text, _ := standard.ExtractText(elementsTable)
	fmt.Println("Found elements table:\n",  text)

	return nil
}

func findElementsTable(page *html.Node) (*html.Node, bool) {
	// The elements table has no id, but a caption with the text
	// "List of elements", searching by that…

	const caption = "List of elements"

	return standard.FindTableWithCaption(page, caption)
}
