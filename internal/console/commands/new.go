package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/beyond3800/hawk/lib"
	"github.com/spf13/cobra"
)


func createProject(projectName string) error {

	// create folders
	dirs := []string{
		projectName,

		filepath.Join(projectName, "app/Models"),

		filepath.Join(projectName, "app/Http/Controllers"),
		filepath.Join(projectName, "app/Http/Middleware"),
		filepath.Join(projectName, "app/Http/Repository"),
		filepath.Join(projectName, "app/Http/Services"),

		filepath.Join(projectName, "database/migrations"),
		filepath.Join(projectName, "database/seeders"),
		filepath.Join(projectName, "database/factory"),
		
		filepath.Join(projectName, "routes"),
		filepath.Join(projectName, "config"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	if err := lib.MakeTemplate("main.go","main",projectName+"/", projectName); err != nil{
		return err
	}
	if err := lib.MakeTemplate(".env","env",projectName+"/",""); err != nil{
		return err
	}
	if err := lib.MakeTemplate("web.go","web",projectName+"/routes/",""); err != nil{
		return err
	}
	createMigration("create_users_table",projectName+"/database/migrations")
	createModel("user",projectName+"/app/Models")
	createController("user",projectName+"/app/Http/Controllers")

	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = projectName
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("go", "get","github.com/beyond3800/hawk")
	cmd.Dir = projectName
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err = cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = projectName
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


