package bootstrap

import (
	"os"
	"os/exec"
)

func RunApp(cmd string) error{
	build := exec.Command("go", "build", "-o", ".tmp/app.exe", ".")
	build.Stdout = os.Stdout
	build.Stderr = os.Stderr

	if err := build.Run(); err != nil {
		return err
	}

	run := exec.Command(".tmp/app.exe", cmd)
	run.Stdout = os.Stdout
	run.Stderr = os.Stderr
	run.Stdin = os.Stdin

	return run.Run()

}