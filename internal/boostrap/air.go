package bootstrap

import (
	"fmt"
	"os"
	"os/exec"
)

func Air() {
	_, err := exec.LookPath("air")
	if err != nil {
		fmt.Println("Air is not installed.")
		fmt.Println("Install it with:")
		fmt.Println("go install github.com/air-verse/air@latest")
		return
	}

	command := exec.Command("air")

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	err = command.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Running at " + os.Getenv("APP_PORT"))
}