package utils

import (
	"os"
	"os/exec"
)

func CheckVSCode() bool {
	_, err := exec.LookPath("code")
	return err == nil
}

func CheckGoogleHasDownloaded() bool {
	chromePaths := []string{
		"C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe",
		"C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe",
	}

	for _, path := range chromePaths {
		if _, err := os.Stat(path); err == nil {
			return true
		}
	}

	return false
}

func CheckComposer() bool {
	_, err := exec.LookPath("composer")
	return err == nil
}
