package main

import (
	"fmt"
	"os"

	"migrania/internal/application/usecases"
	"migrania/internal/infrastructure/input"
	"migrania/internal/infrastructure/openai"
	"migrania/internal/infrastructure/output"
)

func main() {
	openaiAPIKey := os.Getenv("OPENAI_API_KEY")

	client := openai.NewOpenAIClient(openaiAPIKey)
	prompter := input.NewPrompter()

	fmt.Println("🤖 Migracion asistida por ChatGPT")

	action := prompter.Prompt("\nSeleccione la operación (1: Migrar, 2: Mejorar Código, 3: Agregar Pruebas de Cobertura):")
	originPath := prompter.Prompt("\nIngrese ruta de fuentes a procesar:")
	destinyTech := prompter.Prompt("\nIngrese tecnología de destino (si aplica):")
	resultPath := prompter.Prompt("\nIngrese carpeta de destino de solucion:")
	fileWriter := output.NewFileWriter(resultPath)

	switch action {
	case "1":
		originTech := prompter.Prompt("\nIngrese tecnología de origen:")
		migrator := usecases.NewMigrator(client, fileWriter)
		err := migrator.Migrate(originPath, originTech, destinyTech)
		if err != nil {
			fmt.Println("Error durante la migración:", err)
		}
	case "2":
		codeImprover := usecases.NewCodeImprover(client, fileWriter)
		err := codeImprover.ImproveCode(originPath, destinyTech)
		if err != nil {
			fmt.Println("Error durante la mejora del código:", err)
		}
	case "3":
		codeImprover := usecases.NewCodeImprover(client, fileWriter)
		err := codeImprover.AddTestCoverage(originPath, destinyTech)
		if err != nil {
			fmt.Println("Error al agregar pruebas de cobertura:", err)
		}
	default:
		fmt.Println("Operación no válida")
	}
}
