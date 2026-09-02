package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	templatecreator "github.com/beyond3800/hawk/internal/templateCreator"
	"github.com/beyond3800/hawk/internal/templates"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)
func toTitle(s string) string{
	caser := cases.Title(language.English)
	return caser.String(s)
}

func GenerateTemplate(name, templateName, path string) error {
	// Parse the template file in the templates directory
	type templateDatas struct{
		Name      string
		ShortName string
		TableName string
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
	tableName := templatecreator.Pluralize(name)
	// Fill in the template and write the file
	
	shortName := strings.ToLower(name[:1])
	data :=  templateDatas{Name:name, ShortName:shortName, TableName:strings.ToLower(tableName)}
	return tmpl.Execute(file, data)
}

func MakeMigrationTemplate( name, templateName, migrationName, migrationDir string) error {

	type templateDatas struct{
		Name string
		MigrationName string
		TableName string
	}
	name = templatecreator.Pluralize(name)
	templateContent := fmt.Sprintf("%s.tmpl",strings.ToLower(templateName))
	tmpl, err := template.ParseFS(
		templates.Files,
		templateContent,
	)
	if err != nil {
		return err
	}
	timestamp := time.Now().Format("20060102150405")
	
	// fileName = strings.ToUpper(string(name[0]))+string(name[1:])
	nameArr := strings.Split(name, "_")
	if len(nameArr) >= 3 &&nameArr[0] == "create" &&nameArr[len(nameArr)-1] == "table" {
		name = strings.Join(
			nameArr[1:len(nameArr)-1],
			"_",
		)
	}else{
		// strings.CutPrefix(name,"_create")
		migrationName = fmt.Sprintf("%v_%v_%v","create",name,"table")
	}
	
	// Set the output file path
	migrationName = fmt.Sprintf("%v_%v",timestamp,migrationName)
	fileName := filepath.Join(
        migrationDir,
        migrationName+".go",
    )
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	name = toTitle(name)
	// Fill in the template and write the file
	data :=  templateDatas{Name:name, MigrationName: migrationName, TableName: strings.ToLower(name)}
	return tmpl.Execute(file, data)
}

func MakeMiddlewareTemplate(name, templateName string) error {

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

func MakeTemplate(name, templateName, path, data string) error{
	type templateDatas struct{
		Name string
	}
	
	templateContent := fmt.Sprintf("%s.tmpl", templateName)
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


