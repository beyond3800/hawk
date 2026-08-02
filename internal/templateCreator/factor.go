package templatecreator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/beyond3800/hawk/internal/env"
	"github.com/beyond3800/hawk/internal/templates"
)

func Factory(name string, templateName string, factoryName string, factoryDir string) error {

	type templateDatas struct{
		Name         string
		AppName      string
		DatabaseName string
		TableName    string
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
        factoryDir,
        factoryName+".go",
    )
	if err := env.Load(".env"); err != nil{
		return fmt.Errorf("failed to load env file: %w", err)
	}
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	name = ToTitle(name)
	dbName := strings.ToLower(Pluralize(name)) 
	appName,_ := env.Get("APP_NAME")
	// Fill in the template and write the file
	data :=  templateDatas{Name:name,TableName: strings.ToLower(name), DatabaseName: dbName, AppName: appName}
	return tmpl.Execute(file, data)
}