package attributes

import (
	"log/slog"
	"strings"

	"golang.org/x/net/html"

	"github.com/KrischanCS/htmfunc/internal/standard"
)

func createAcctributeGroups(body *html.Node) attributes {
	slog.Info("Creating attribute groups...")

	attrs := parseAttributes(body)

	attrsClassified := classifyAttributes(attrs)

	attrsClassified.Enum = DecomposeEnums(attrsClassified.Enum)

	disambiguateAttrs(&attrsClassified)

	slog.Info("Created attribute groups.")

	return attrsClassified
}

func parseAttribute(row *html.Node) attribute {
	data := row.FirstChild

	var attr attribute

	text, links := standard.ExtractText(data)
	setNames(&attr, text)
	attr.Links = links

	data = data.NextSibling

	elements, links := extractElements(data)
	attr.Elements = elements
	attr.Links = append(attr.Links, links...)

	data = data.NextSibling

	text, links = standard.ExtractText(data)
	attr.Description = text
	attr.Links = append(attr.Links, links...)

	data = data.NextSibling

	text, links = standard.ExtractText(data)
	text = strings.ReplaceAll(text, "*", " (Additional rules apply, see elements documentation)")

	attr.Value = text
	attr.Links = append(attr.Links, links...)

	return attr
}

func setNames(attr *attribute, attributeName string) {
	attr.Name = attributeName

	if mapping, ok := mappings[attributeName]; ok {
		attr.FuncName = mapping.FuncName
		attr.ParamName = mapping.paramName

		return
	}

	attr.ParamName = attributeName
	attr.FuncName = strings.ToUpper(attr.ParamName[0:1]) + attr.ParamName[1:]
}

func extractElements(data *html.Node) ([]string, []standard.Link) {
	text, links := standard.ExtractText(data)

	return strings.Split(text, ";"), links
}
