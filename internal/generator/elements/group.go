package elements

import (
	"log/slog"
	"strings"

	"github.com/KrischanCS/go-toolbox/iterator"

	"github.com/KrischanCS/roundly/internal/generator/standard"
)

func groupElements(elements []Element) map[string][]Element {
	slog.Info("assigning elements to semantic groups...")

	e := iterator.Of(elements...)

	groups := make(map[string][]Element)
	iterator.Reduce(e, &groups, func(accumulator *map[string][]Element, value Element) {
		group, ok := (*accumulator)[value.SemanticGroup]
		if !ok {
			group = make([]Element, 0)
		}

		group = append(group, value)
		(*accumulator)[value.SemanticGroup] = group
	})

	slog.Info("assigned elements to semantic groups.")

	return groups
}

func extractSemanticGroup(link standard.Link) string {
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
