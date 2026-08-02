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

func MakeResourceTemplate(name string, templateName string, dir string) error {
	type templateDatas struct {
		Name        string
		AppName     string
		PluralName  string
		SmallLetter string
	}
	templateContent := fmt.Sprintf("%s.tmpl", strings.ToLower(templateName))
	tmpl, err := template.ParseFS(
		templates.Files,
		templateContent,
	)
	if err != nil {
		return err
	}
	path := dir
	fileName := filepath.Join(
		path,
		name+".go",
	)
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if err := env.Load(".env"); err != nil {
		return fmt.Errorf("failed to load environment: %w", err)
	}

	appName, ok := env.Get("APP_NAME")
	if !ok {
		return fmt.Errorf("APP_NAME not set in environment")
	}
	newName := name
	if strings.HasSuffix(name, "y") {
		// newName,_ := strings.CutSuffix("y",name)
		newName = strings.Replace(newName, "y", "ies", 1)
	} else {
		newName = newName + "s"
	}

	name = ToTitle(name)
	newName = ToTitle(newName)
	smallLetter := strings.ToLower(newName)
	datas := templateDatas{Name: name, AppName: appName, PluralName: newName, SmallLetter: smallLetter}
	return tmpl.Execute(file, datas)
}