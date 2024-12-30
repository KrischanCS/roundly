package main

import (
	_ "embed"
	"flag"
	"log"
	"os"
	"text/template"
)

var reloadStandard = flag.Bool("reloadHtmlStandard", false, "reload html standard from the web instead of file system")

var textAttrTemplate = template.Must(template.ParseFiles("templates/attributesText.go.tmpl"))
var boolAttrTemplate = template.Must(template.ParseFiles("templates/attributesBool.go.tmpl"))
var enumAttrTemplate = template.Must(template.ParseFiles("templates/attributesEnum.go.tmpl"))
var inputTypeAttrTemplate = template.Must(template.ParseFiles("templates/attributesInputType.go.tmpl"))
var commaSeparatedAttrTemplate = template.Must(template.ParseFiles("templates/attributesCommaSeparated.go.tmpl"))
var commaSeparatedFloatAttrTemplate = template.Must(template.ParseFiles("templates/attributesCommaSeparatedFloats.go." +
	"" + "tmpl"))
var spaceSeparatedAttrTemplate = template.Must(template.ParseFiles("templates/attributesSpaceSeparated.go.tmpl"))
var floatAttrTemplate = template.Must(template.ParseFiles("templates/attributesFloat.go.tmpl"))
var intAttrTemplate = template.Must(template.ParseFiles("templates/attributesInt.go.tmpl"))
var uintAttrTemplate = template.Must(template.ParseFiles("templates/attributesUint.go.tmpl"))

func main() {
	flag.Parse()
	body := loadIndicesFromStandard()

	attributes := findAttributes(body)

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
}

func generateFile(tmpl *template.Template, attributes []attribute, fileName string) {
	file, err := os.OpenFile("generated/"+fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
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
