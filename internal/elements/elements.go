// Package elements implements the generation of the HTML element functions
// based on the HTML standard.
package elements

import (
	"log/slog"

	"golang.org/x/net/html"
)

func GenerateElements(standardIndicesPage *html.Node) {
	slog.Info("Generating elements...")

	elements := parseAllElements(standardIndicesPage)

	elementGroups := groupElements(elements)

	generateFileForEachGroup(elementGroups)

	slog.Info("Generated elements.")
}

