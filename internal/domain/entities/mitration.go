package entities

import (
	"os"
)

const EOF = "#############################################"

func readContentFromPath(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(content)
}
