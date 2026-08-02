package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	Make "github.com/beyond3800/hawk/internal/console/commands/make"
	"github.com/beyond3800/hawk/internal/lib"
	"github.com/spf13/cobra"
)

func createProject(projectName string) error {

	// create folders
	dirs := []string{
		projectName,

		// Creating the root-level directories
		filepath.Join(projectName, "app"),
		filepath.Join(projectName, "bootstrap"),
		filepath.Join(projectName, "config"),
		filepath.Join(projectName, "database"),
		filepath.Join(projectName, "routes"),	
		filepath.Join(projectName, "bootstrap"),
		filepath.Join(projectName, "routes"),
		filepath.Join(projectName, "config"),
		filepath.Join(projectName, "internal"),
		filepath.Join(projectName, "console"),

		
		// Creating the app directory and its subdirectories
		filepath.Join(projectName, "app/Models"),

		filepath.Join(projectName, "app/Http/Controllers"),
		filepath.Join(projectName, "app/Http/Middleware"),
		filepath.Join(projectName, "app/Http/Repository"),
		filepath.Join(projectName, "app/Http/Services"),

		// Creating the database directory and its subdirectories
		filepath.Join(projectName, "database/migrations"),
		filepath.Join(projectName, "database/seeders"),
		filepath.Join(projectName, "database/factory"),
		
		filepath.Join(projectName, "internal/console"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	if err := lib.MakeTemplate(".env","env",projectName+"/",projectName); err != nil{
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
	if err := lib.MakeTemplate(".air.toml","air",projectName+"/",""); err != nil{
		return err
	}
	if err := lib.MakeTemplate("databaseSeeder.go","databaseSeeder",projectName+"/database/seeders/",projectName); err != nil{
		return err
	}

	if err := lib.MakeTemplate("execute.go","execute",projectName+"/internal",""); err != nil{
		return err
	}
	if err := lib.MakeTemplate("commands.go","commands",projectName+"/internal",""); err != nil{
		return err
	}
	Make.CreateMigration("create_users_table",projectName+"/database/migrations")
	Make.CreateModel("user",projectName+"/app/Models")
	Make.CreateController("user",projectName+"/app/Http/Controllers")

	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = projectName
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return err
	}

	cmd = exec.Command("go", "get","github.com/beyond3800/hawk@latest")
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

		projectName := args[0]
		if projectName == ""{
			fmt.Println("The project needs a name")
			return
		}

		if err := createProject(projectName); err != nil {
			fmt.Println("Error:", err)
			return
		}
		installAir()
		fmt.Println("Project created successfully:", projectName)
	},
}


func NewProjectCommand () *cobra.Command{
	return newCmd
}


