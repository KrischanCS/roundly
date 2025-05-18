// Package attributes generates all attribute functions based on the HTML standard.
package attributes

import (
	"embed"
	"log"
	"os"
	"text/template"

	"golang.org/x/net/html"
)

//go:embed templates
var templates embed.FS

//nolint:gochecknoglobals
var attrTemplates = template.Must(template.ParseFS(templates, "templates/*.go.tmpl"))

func GenerateAttributes(body *html.Node) {
	attributes := findAttributes(body)
	eventHandlerAttributes := findEventHandlerAttributes(body)

	generateFile("attributesText.go.tmpl", attributes.Text, "text.go")
	generateFile("attributesBool.go.tmpl", attributes.Bool, "bool.go")
	generateFile("attributesEnum.go.tmpl", attributes.Enum, "enum.go")
	generateFile("attributesInputType.go.tmpl", attributes.InputType, "inputType.go")
	generateFile("attributesCommaSeparated.go.tmpl", attributes.ListComma, "commaSeperated.go")
	generateFile("attributesCommaSeparatedFloats.go.tmpl", attributes.ListCommaFloat, "commaSeperatedFloat.go")
	generateFile("attributesSpaceSeparated.go.tmpl", attributes.ListSpace, "spaceSeperated.go")
	generateFile("attributesFloat.go.tmpl", attributes.Float, "float.go")
	generateFile("attributesInt.go.tmpl", attributes.Int, "int.go")
	generateFile("attributesUint.go.tmpl", attributes.Uint, "uint.go")
	generateFile("attributeList.go.tmpl", nil, "attributeList.go")
	generateFile("attributesText.go.tmpl", eventHandlerAttributes, "eventHandler.go")
}

func generateFile(name string, attributes []attribute, fileName string) {
	//nolint:gosec
	file, err := os.OpenFile("../attribute/"+fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //nolint:mnd
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
}
