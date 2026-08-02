package templatecreator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/beyond3800/hawk/internal/templates"
)

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