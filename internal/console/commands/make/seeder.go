package Make

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/beyond3800/hawk/internal/env"
	"github.com/beyond3800/hawk/internal/lib"
	templatecreator "github.com/beyond3800/hawk/internal/templateCreator"
	"github.com/spf13/cobra"
)

func databaseSeeder(dir string, seederName string) error{
	databaseSeederFile := "databaseSeeder.go"
	fmt.Println("Seeder created successfully")
	fileName := filepath.Join(dir,databaseSeederFile)

	if !lib.FileExist(dir, databaseSeederFile) {

		content := `package seeders

type DatabaseSeeder struct{}

func (DatabaseSeeder) Run() error {
	return nil
}
`

		if err := os.WriteFile(fileName, []byte(content), 0644); err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func addNewSeeder( fileName string, seederName string) error{
	// Read the existing file.
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	text := string(content)

	runner := fmt.Sprintf(
		"\tif err := %sSeeder{}.Run(); err != nil {\n\t\treturn err\n\t}\n\n",
		seederName,
	)


	// Insert before "return nil".
	idx := strings.LastIndex(text, "\treturn nil")
	if idx == -1 {
		return fmt.Errorf("could not find return nil")
	}

	text = text[:idx] + runner + text[idx:]
	return os.WriteFile(
		fileName,
		[]byte(text),
		0644,
	)
}

func CreateSeeder(name string, dir string){
	_,err := os.Stat(dir)
	if err !=nil{
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Println("Unable to create seeder directory:", err)
				return
			}
		}
	}
	if lib.FileExist(dir,name){
		fmt.Println(`This file already exist in ` + dir)
		return
	}
	seederName := templatecreator.ToTitle(name)
	env.Load(".env")
	appName,_ := env.Get("APP_NAME")
	if err := templatecreator.MakeTemplate(name+".go", "seeder", dir, seederName, appName); err != nil {
		fmt.Println(err)
	}
	if err := databaseSeeder(dir, seederName) ;err != nil {
		fmt.Println(err)
	}

}


var seederCmd = &cobra.Command{
	Use:   "seeder",
	Short: "Create a new seeder",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := "database/seeders"
		name := args[0]
		CreateSeeder(name, dir)
	},
}


func SeederCommand() *cobra.Command {
	return seederCmd
}