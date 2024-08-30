package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sashabaranov/go-openai" // Necesitas instalar go-openai: `go get github.com/sashabaranov/go-openai`
)

const EOF = "#############################################"

const PROMPT_MIGRACION = `
    Migra los siguientes codigos fuentes : \n%s\n los cuales estan en %s, debes migrar esto a %s,
    quiero que tu respuesta este en el siguiento formato
    
    @@@@NOMBRE ARCHIVO
    CONTENIDO CODIGO MIGRADO
    #############################################

    No agregue espacios en blanco despues de la ultima linea de codigo, ya que esto puede generar errores en la migracion
    No coloques comentarios en el codigo, ya que esto puede generar errores en la migracion
    No menciones sugerencias ni acciones dentro del codigo, si quieres realizarlo comentalo en formato de comentario
    . \n
`

func main() {

	openaiAPIKey := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(openaiAPIKey)

	fmt.Println("🤖 Migracion asistida por ChatGPT")

	originPath := prompt("\nIngrese ruta de fuentes a migrar :")
	originTech := prompt("\nIngrese tecnologia de Origen  :")
	destinyTech := prompt("\nIngrese tecnologia de Destino :")
	sources := readFilePathsFromPath(originPath)

	prompt := fmt.Sprintf(PROMPT_MIGRACION, sources, originTech, destinyTech)

	requestIA := prompt

	// Contexto del asistente
	messages := []openai.ChatCompletionMessage{
		{Role: "system", Content: "Eres un developer senior."},
		{Role: "user", Content: requestIA},
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT4o20240513,
			Messages: messages,
		},
	)

	if err != nil {
		fmt.Println("Error al llamar a la API de OpenAI:", err)
		return
	}

	responseContent := resp.Choices[0].Message.Content

	// Iterar la lectura del responseContent linea a linea
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
			if line != EOF {
				contentFile += line + "\n"
			} else {
				createSource(nameFile, contentFile)
				isNameFile = true
				nameFile = ""
				contentFile = ""
			}
		}
	}
}

func prompt(label string) string {
	fmt.Print(label)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func createSource(filename, content string) {
	fmt.Printf("Archivo: %s\n", filename)
	fmt.Println(content)

	// verificar si carpeta existe, sino se crea una en la ruta especificada
	if _, err := os.Stat("./output"); os.IsNotExist(err) {
		fmt.Println("creando carpeta output")
		os.Mkdir("./output", os.ModePerm)
	}
	fmt.Println("Los fuentes seran dejados en la carpeta 'output'")

	// verificar si archivo existe, sino se crea un archivo nuevo
	path := "./output/" + filename
	if _, err := os.Stat(path); err == nil {
		fmt.Println("El archivo ya existe")
		os.Exit(-1)
	} else {
		fmt.Println("creando archivo  =>" + filename)
		file, err := os.Create(path)
		if err != nil {
			fmt.Println("Error creando archivo:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println("Error escribiendo en archivo:", err)
		}
	}
}

func readFilePathsFromPath(path string) string {
	var contentSourcesFile string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			contentSourcesFile += "\n" + info.Name() + readContentFromPath(path) + "\n"
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error leyendo archivos:", err)
	}

	return contentSourcesFile
}

func readContentFromPath(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error leyendo archivo:", err)
		return ""
	}
	return string(content)
}
