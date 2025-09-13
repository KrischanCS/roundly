package elements

import (
	"embed"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates
var templateFs embed.FS

//nolint:gochecknoglobals
var elementTemplate = template.Must(template.ParseFS(templateFs, "templates/*.go.tmpl"))

func generateFileForEachGroup(elementGroups map[string][]Element) {
	slog.Info("Generating files for groups...", "groupCount", len(elementGroups))

	for group, elements := range elementGroups {
		generateFile(group, elements)
	}

	slog.Info("Generated files for all groups.")
}

func generateFile(group string, elements []Element) {
	filePath := filepath.Join("../..", "html/element"+strings.ToUpper(string(group[0]))+group[1:]+".go")

	slog.Debug("Creating file...", "file", filePath)

	//nolint:gosec // Files are written for everyone
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(fmt.Sprintf("Error opening file %s: %s", group+".go", err))
	}
	defer file.Close() //nolint:errcheck

	err = elementTemplate.Lookup("element.go.tmpl").Execute(file, elements)
	if err != nil {
		panic(fmt.Sprintf("Error executing template %s: %s", group, err))
	}

	slog.Info("Created file.", "file", filePath)
}
