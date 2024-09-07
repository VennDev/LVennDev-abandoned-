package utils

import (
	"os/exec"
	"path/filepath"
)

func RunCommand(command string) {
	absPath, err := filepath.Abs(command)
	if err != nil {
		panic(err)
	}
	exec.Command("cmd", "/C", absPath).Start()
}
