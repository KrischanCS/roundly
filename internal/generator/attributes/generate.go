// Package attributes generates all attribute functions based on the HTML standard.
package attributes

import (
	"embed"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"text/template"

	"golang.org/x/net/html"

	"github.com/KrischanCS/roundly/internal/generator/standard"
)

//go:embed templates
var templates embed.FS

//nolint:gochecknoglobals
var attrTemplates = template.Must(template.ParseFS(templates, "templates/*.go.tmpl"))

func GenerateAttributes(indicesBody *html.Node, inputBody *html.Node) {
	slog.Info("Generating attributes...")

	attributes := createAttributeGroups(indicesBody, inputBody)

	eventHandlerAttributes := findEventHandlerAttributes(indicesBody)

	slog.Info("Generating files...")

	const attribute = "attribute"
	generateFile("attributesText.go.tmpl", attributes.Text, attribute+"Text.go")
	generateFile("attributesBool.go.tmpl", attributes.Bool, attribute+"Bool.go")
	generateFile("attributesEnum.go.tmpl", attributes.Enum, attribute+"Enum.go")
	generateFile("attributesCommaSeparated.go.tmpl", attributes.ListComma, attribute+"CommaSeperated.go")
	generateFile(
		"attributesCommaSeparatedFloats.go.tmpl",
		attributes.ListCommaFloat,
		"commaSeperatedFloat.go",
	)
	generateFile("attributesSpaceSeparated.go.tmpl", attributes.ListSpace, attribute+"SpaceSeperated.go")
	generateFile("attributesFloat.go.tmpl", attributes.Float, attribute+"Float.go")
	generateFile("attributesInt.go.tmpl", attributes.Int, attribute+"Int.go")
	generateFile("attributesUint.go.tmpl", attributes.Uint, attribute+"Uint.go")
	generateFile("attributeList.go.tmpl", nil, attribute+"List.go")
	generateFile("attributesText.go.tmpl", eventHandlerAttributes, attribute+"EventHandler.go")

	slog.Info("Generated attributes.")
}

func findEventHandlerAttributes(body *html.Node) []attribute {
	eventHandlersTable := standard.FindNodeWithId(body, "ix-event-handlers")
	if eventHandlersTable == nil {
		log.Fatal("Error finding event handlers table")
	}

	tBody := standard.FindTBody(eventHandlersTable)

	attrs := make([]attribute, 0, 256)

	for row := range tBody.ChildNodes() {
		attr := parseAttribute(row)
		attrs = append(attrs, attr)
	}

	return attrs
}

func generateFile(name string, attributes []attribute, fileName string) {
	filePath := filepath.Join("../..", "html", fileName)

	slog.Debug("Generating file...", "file", filePath, "templateName", name)

	//nolint:gosec
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Print("Error closing file: ", err)
		}
	}(file)

	err = attrTemplates.Lookup(name).Execute(file, attributes)
	if err != nil {
		log.Print("Error executing template: ", err)
	}

	slog.Info("Generated file.", "file", filepath.Base(filePath))
}
