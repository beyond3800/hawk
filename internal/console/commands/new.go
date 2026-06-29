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

		filepath.Join(projectName, "bootstrap"),

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
	if err := lib.MakeTemplate(".env","env",projectName+"/",""); err != nil{
		return err
	}
	if err := lib.MakeTemplate("main.go","main",projectName+"/",projectName); err != nil{
		return err
	}
	if err := lib.MakeTemplate("web.go","web",projectName+"/routes/",""); err != nil{
		return err
	}
	if err := lib.MakeTemplate("app.go","app",projectName+"/bootstrap/",projectName); err != nil{
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
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = projectName
	if err := cmd.Run(); err != nil {
		return err
	}
		return nil
}
func installAir() {
    if _, err := exec.LookPath("air"); err == nil {
        fmt.Println("Air already installed.")
        return
    }

    fmt.Println("Installing Air...")

    cmd := exec.Command(
        "go",
        "install",
        "github.com/air-verse/air@latest",
    )

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("Failed to install Air:", err)
        return
    }

    fmt.Println("Air installed successfully.")
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
		installAir()
		fmt.Println("Project created successfully:", project)
	},
}


func NewProjectCommand () *cobra.Command{
	return newCmd
}


