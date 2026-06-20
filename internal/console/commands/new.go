package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/beyond3800/hawk/lib"
	"github.com/spf13/cobra"
)


func createProject(name string) error {

	// create folders
	dirs := []string{
		name,

		filepath.Join(name, "app/Models"),

		filepath.Join(name, "app/Http/Controllers"),
		filepath.Join(name, "app/Http/Middleware"),
		filepath.Join(name, "app/Http/Repository"),
		filepath.Join(name, "app/Http/Services"),

		filepath.Join(name, "database/migrations"),
		filepath.Join(name, "database/seeders"),
		filepath.Join(name, "database/factory"),
		
		filepath.Join(name, "routes"),
		filepath.Join(name, "config"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	if err := lib.MakeTemplate("main","main","/", name); err != nil{
		return err
	}
	if err := lib.MakeTemplate(".env","env","",""); err != nil{
		return err
	}
	if err := lib.MakeTemplate("web","web","routes/",""); err != nil{
		return err
	}
	createMigration("create_users_table","database/migrations")
	createModel("user","app/Models")

	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = name
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

var newCmd = &cobra.Command{
	Use:   "new [projectName]",
	Short: "Create a new Hawk project",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		project := args[0]
		if project == ""{
			fmt.Println("The project needs a name")
			return
		}

		if err := createProject(project); err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Project created:", project)
	},
}


func NewProjectCommand () *cobra.Command{
	return newCmd
}


