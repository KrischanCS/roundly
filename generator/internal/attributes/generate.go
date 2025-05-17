package attributes

import (
	"log"
	"os"
	"text/template"

	"golang.org/x/net/html"
)

//nolint:gochecknoglobals
var (
	textAttrTemplate = template.Must(
		template.ParseFiles("./internal/attributes/templates/attributesText.go.tmpl"))
	boolAttrTemplate = template.Must(
		template.ParseFiles("./internal/attributes/templates/attributesBool.go.tmpl"))
	enumAttrTemplate = template.Must(
		template.ParseFiles("./internal/attributes/templates/attributesEnum.go.tmpl"))
	inputTypeAttrTemplate = template.Must(
		template.ParseFiles("./internal/attributes/templates/attributesInputType.go.tmpl"))
	commaSeparatedAttrTemplate = template.Must(
		template.ParseFiles("./internal/attributes/templates/attributesCommaSeparated.go.tmpl"))
	commaSeparatedFloatAttrTemplate = template.Must(
		template.ParseFiles("./internal/attributes/templates/attributesCommaSeparatedFloats.go." + "" + "tmpl"))
	spaceSeparatedAttrTemplate = template.Must(
		template.ParseFiles("./internal/attributes/templates/attributesSpaceSeparated.go.tmpl"))
	floatAttrTemplate = template.Must(
		template.ParseFiles("./internal/attributes/templates/attributesFloat.go.tmpl"))
	intAttrTemplate = template.Must(
		template.ParseFiles("./internal/attributes/templates/attributesInt.go.tmpl"))
	uintAttrTemplate = template.Must(
		template.ParseFiles("./internal/attributes/templates/attributesUint.go.tmpl"))
	attrsTemplate = template.Must(
		template.ParseFiles("./internal/attributes/templates/attributeList.go.tmpl"))
)

func GenerateAttributes(body *html.Node) {
	attributes := findAttributes(body)
	eventHandlerAttributes := findEventHandlerAttributes(body)

	generateFile(textAttrTemplate, attributes.Text, "text.go")
	generateFile(boolAttrTemplate, attributes.Bool, "bool.go")
	generateFile(enumAttrTemplate, attributes.Enum, "enum.go")
	generateFile(inputTypeAttrTemplate, attributes.InputType, "inputType.go")
	generateFile(commaSeparatedAttrTemplate, attributes.ListComma, "commaSeperated.go")
	generateFile(commaSeparatedFloatAttrTemplate, attributes.ListCommaFloat, "commaSeperatedFloat.go")
	generateFile(spaceSeparatedAttrTemplate, attributes.ListSpace, "spaceSeperated.go")
	generateFile(floatAttrTemplate, attributes.Float, "float.go")
	generateFile(intAttrTemplate, attributes.Int, "int.go")
	generateFile(uintAttrTemplate, attributes.Uint, "uint.go")
	generateFile(attrsTemplate, nil, "attributeList.go")
	generateFile(textAttrTemplate, eventHandlerAttributes, "eventHandler.go")
}

func generateFile(tmpl *template.Template, attributes []attribute, fileName string) {
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
	err = tmpl.Execute(file, attributes)
	if err != nil {
		log.Print("Error executing template: ", err)
	}
}
