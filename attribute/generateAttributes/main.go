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

func main() {
	flag.Parse()
	body := loadIndicesFromStandard()

	attributes := findAttributes(body)

	generateFile(textAttrTemplate, attributes.Text, "text.go")
	generateFile(boolAttrTemplate, attributes.Bool, "bool.go")
	generateFile(enumAttrTemplate, attributes.Enum, "enum.go")
	generateFile(textAttrTemplate, attributes.Other, "other.go")
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
