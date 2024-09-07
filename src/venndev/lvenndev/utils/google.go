package utils

import (
	"os"
)

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
