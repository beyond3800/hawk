package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/beyond3800/hawk/internal/templates"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)
func toTitle(s string) string{
	caser := cases.Title(language.English)
	return caser.String(s)
}
func GenerateTemplate(name string, templateName string, path string) error {
	// Parse the template file in the templates directory
	type templateDatas struct{
		Name string
		ShortName string
	}
	
	templateContent := fmt.Sprintf("%s.tmpl", strings.ToLower(templateName))
	tmpl, err := template.ParseFS(
		templates.Files,
		templateContent,
	)

	if err != nil {
		return err
	}

	// Set the output file path
	templateName = path

	fileName := filepath.Join(
		templateName, 
		name+".go",)
	// fileName = strings.ToUpper(string(name[0]))+string(name[1:])
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	caser := cases.Title(language.English)
	name = caser.String(name)

	// Fill in the template and write the file
	shortName := name[:1]
	data :=  templateDatas{Name:name,ShortName: shortName}
	return tmpl.Execute(file, data)
}

func MakeMigrationTemplate(name string, templateName string, migrationName string, migrationDir string) error {

	type templateDatas struct{
		Name string
		MigrationName string
	}

	templateContent := fmt.Sprintf("%s.tmpl",strings.ToLower(templateName))
	tmpl, err := template.ParseFS(
		templates.Files,
		templateContent,
	)
	if err != nil {
		return err
	}

	// Set the output file path
	fileName := filepath.Join(
        migrationDir,
        migrationName+".go",
    )
	// fileName = strings.ToUpper(string(name[0]))+string(name[1:])
	nameArr := strings.Split(name, "_")
	if len(nameArr) >= 3 &&nameArr[0] == "create" &&nameArr[len(nameArr)-1] == "table" {
		name = strings.Join(
			nameArr[1:len(nameArr)-1],
			"_",
		)
	}

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	name = toTitle(name)
	// Fill in the template and write the file
	data :=  templateDatas{Name:name, MigrationName: migrationName}
	return tmpl.Execute(file, data)
}

func MakeMiddlewareTemplate(name string, templateName string) error {

	type templateData struct {
		Name string
	}

	templateContent := fmt.Sprintf("%s.tmpl", strings.ToLower(templateName))

	tmpl, err := template.ParseFS(
		templates.Files,
		templateContent,
	)
	if err != nil {
		return err
	}
	
	name = toTitle(name)

	fileName := filepath.Join(
		"app/Http/Middleware",
		name+".go",
	)

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	data := templateData{
		Name: name,
	}

	return tmpl.Execute(file, data)
}

func MakeTemplate(name string, templateName string, path string, data string) error{
	type templateDatas struct{
		Name string
	}
	
	templateContent := fmt.Sprintf("%s.tmpl", strings.ToLower(templateName))
	tmpl, err := template.ParseFS(
		templates.Files,
		templateContent,
	)
	if err != nil {
		return err
	}

	fileName := filepath.Join(
		path,
		name,
	)
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	datas :=  templateDatas{Name:data}
	return tmpl.Execute(file, datas)
}
