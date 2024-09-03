package output

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FileWriter struct {
	outputDir string
}

func NewFileWriter(outputDir string) *FileWriter {
	return &FileWriter{outputDir: outputDir}
}

func (fw *FileWriter) WriteImprovementFiles(responseContent string) error {
	isNameFile := true
	nameFile := ""
	contentFile := ""

	scanner := bufio.NewScanner(strings.NewReader(responseContent))
	for scanner.Scan() {
		line := scanner.Text()
		if isNameFile {
			if strings.HasPrefix(line, "@@@@") {
				nameFile = strings.Replace(line, "@@@@", "", 1)
				isNameFile = false
			}
		} else {
			if line != "#############################################" {
				contentFile += line + "\n"
			} else {
				err := fw.createSource(nameFile, contentFile)
				if err != nil {
					return err
				}
				isNameFile = true
				nameFile = ""
				contentFile = ""
			}
		}
	}
	return nil
}

func (fw *FileWriter) createSource(filename, content string) error {
	fmt.Printf("Archivo: %s\n", filename)
	fmt.Println(content)

	if _, err := os.Stat(fw.outputDir); os.IsNotExist(err) {
		fmt.Println("Creando carpeta output")
		os.Mkdir(fw.outputDir, os.ModePerm)
	}
	fmt.Println("Los fuentes serán dejados en la carpeta ", fw.outputDir)

	path := fw.outputDir + "/" + filename
	if _, err := os.Stat(path); err == nil {
		fmt.Println("El archivo ya existe")
		os.Exit(-1)
	} else {
		fmt.Println("Creando archivo  =>" + filename)
		file, err := os.Create(path)
		if err != nil {
			fmt.Println("Error creando archivo:", err)
			return err
		}
		defer file.Close()

		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println("Error escribiendo en archivo:", err)
			return err
		}
	}
	return nil
}
