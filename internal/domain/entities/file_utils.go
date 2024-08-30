package entities

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadFilePathsFromPath(path string) string {
	var contentSourcesFile string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			contentSourcesFile += "\n" + info.Name() + ReadContentFromPath(path) + "\n"
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error leyendo archivos:", err)
	}

	return contentSourcesFile
}

func ReadContentFromPath(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error leyendo archivo:", err)
		return ""
	}
	return string(content)
}
