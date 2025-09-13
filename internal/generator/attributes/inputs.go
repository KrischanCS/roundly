package attributes

import (
	"log/slog"

	"golang.org/x/net/html"

	standard2 "github.com/KrischanCS/roundly/internal/generator/standard"
)

func parseInputTypes(inputBody *html.Node) []string {
	slog.Info("Parsing input types...")

	tBody := findInputTypesTable(inputBody)

	inputTypes := make([]string, 0, 16)

	for row := range tBody.ChildNodes() {
		text, _ := standard2.ExtractText(row.FirstChild)
		inputTypes = append(inputTypes, text)
	}

	slog.Info("Parsed input types.")

	return inputTypes
}

func findInputTypesTable(inputBody *html.Node) *html.Node {
	inputsTable := standard2.FindNodeWithId(inputBody, "attr-input-type-keywords")
	if inputsTable == nil {
		panic("Unable to find inputs table")
	}

	tBody := standard2.FindTBody(inputsTable)

	return tBody
}
