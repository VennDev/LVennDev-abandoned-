package utils

import (
	"os/exec"
)

func CheckVSCode() bool {
	_, err := exec.LookPath("code")
	if err != nil {
		return false
	}
	return true
}
