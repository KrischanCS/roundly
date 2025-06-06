package attributes

import (
	"log"
	"log/slog"

	"golang.org/x/net/html"

	"github.com/KrischanCS/roundly/internal/standard"
)

func parseAttributes(body *html.Node) []*attribute {
	slog.Info("Parsing attributes...")

	tBody := findAttributesTableBody(body)
	attrs := createAttrbibutesFromRows(tBody)

	slog.Info("Parsed attributes.", "attributesCount", len(attrs))

	return attrs
}

func findAttributesTableBody(body *html.Node) *html.Node {
	slog.Debug("Finding attributes table...")

	attributesTable := standard.FindNodeWithId(body, "attributes-1")
	if attributesTable == nil {
		log.Fatal("Error finding attributes table")
	}

	tBody := standard.FindTBody(attributesTable)

	slog.Debug("Found attributes table.")

	return tBody
}

func createAttrbibutesFromRows(tBody *html.Node) []*attribute {
	slog.Debug("Creating attributes from rows...")

	attrs := make([]*attribute, 0, 256)

	for row := range tBody.ChildNodes() {
		attr := parseAttribute(row)
		attrs = append(attrs, &attr)
	}

	slog.Debug("Created attributes from rows.")

	return attrs
}
