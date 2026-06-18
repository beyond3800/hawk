package lib

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func GenerateTemplate(name string, templateName string, path string) error {
	// Parse the template file in the templates directory
	type templateDatas struct{
		Name string
		ShortName string
	}
	
	templateContent := fmt.Sprintf("./templates/%s.tmpl",strings.ToLower(templateName))
	tmpl, err := template.ParseFiles(templateContent)
	if err != nil {
		return err
	}

	// Set the output file path
	templateName = "app/Http/"+strings.Title(templateName)
	fileName := fmt.Sprintf("%s/%s.go", templateName, name)
	// fileName = strings.ToUpper(string(name[0]))+string(name[1:])
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Fill in the template and write the file
	shortName := name[:1]
	data :=  templateDatas{Name:strings.Title(name),ShortName: shortName}
	return tmpl.Execute(file, data)
}

func MakeMigrationTemplate(name string, templateName string, migrationName string) error {

	type templateDatas struct{
		Name string
		MigrationName string
	}
	fmt.Println(name,templateName)
	templateContent := fmt.Sprintf("./templates/%s.tmpl",strings.ToLower(templateName))
	tmpl, err := template.ParseFiles(templateContent)
	if err != nil {
		return err
	}

	// Set the output file path
	
	fileName := "database/migrations/" + migrationName + ".go"
	// fileName = strings.ToUpper(string(name[0]))+string(name[1:])
	nameArr := strings.Split(name, "_")
	if len(nameArr) == 3{
		// if nameArr[1] != "create" || nameArr[1] != "table"{
			name = nameArr[1]
		// }
	}
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Fill in the template and write the file
	data :=  templateDatas{Name:strings.Title(name), MigrationName: migrationName}
	return tmpl.Execute(file, data)
}

func MakeMiddlewareTemplate(name string, templateName string) error {

	type templateDatas struct{
		Name string
	}
	
	// if !strings.HasSuffix(name,"middleware"){
	// 	fileName := fmt.Sprintf("./middleware/%s.go", name)
	// 	return
	// }
	templateContent := fmt.Sprintf("./templates/%s.tmpl",strings.ToLower(templateName))
	tmpl, err := template.ParseFiles(templateContent)
	if err != nil {
		return err
	}

	name = strings.Title(name)
	fileName := fmt.Sprintf("./middleware/%s.go", name)
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	data :=  templateDatas{Name:strings.Title(name)}
	return tmpl.Execute(file, data)
	
}