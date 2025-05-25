package attributes

import (
	"log/slog"

	"golang.org/x/net/html"

	"github.com/KrischanCS/htmfunc/internal/standard"
)

func parseInputTypes(inputBody *html.Node) []string {
	slog.Info("Parsing input types...")

	tBody := findInputTypesTable(inputBody)

	inputTypes := make([]string, 0, 16)

	for row := range tBody.ChildNodes() {
		text, _ := standard.ExtractText(row.FirstChild)
		inputTypes = append(inputTypes, text)
	}

	slog.Info("Parsed input types.")

	return inputTypes
}

func findInputTypesTable(inputBody *html.Node) *html.Node {
	inputsTable := standard.FindNodeWithId(inputBody, "attr-input-type-keywords")
	if inputsTable == nil {
		panic("Unable to find inputs table")
	}

	tBody := standard.FindTBody(inputsTable)

	return tBody
}
