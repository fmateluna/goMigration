package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"migrania/internal/application/usecases"
	"migrania/internal/domain/entities"
	"migrania/internal/infrastructure/input"
	"migrania/internal/infrastructure/openai"
	"migrania/internal/infrastructure/output"
)

func loadPrompts() (map[int]string, map[int]string, error) {
	file, err := os.Open("ia.context.json")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var data entities.JSONData

	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, nil, err
	}

	promptMap := make(map[int]string)
	taskMap := make(map[int]string)
	for i, t := range data.Tasks {
		promptMap[i+1] = t.Prompt
		taskMap[i+1] = t.Task
	}

	return promptMap, taskMap, nil
}

func main() {
	openaiAPIKey := os.Getenv("OPENAI_API_KEY")

	client := openai.NewOpenAIClient(openaiAPIKey)
	prompter := input.NewPrompter()

	fmt.Println("🤖 Migracion asistida por ChatGPT")

	prompts, tasks, err := loadPrompts()
	if err != nil {
		fmt.Println("Error al cargar los prompts:", err)
		return
	}

	// Present options to user
	fmt.Println("Seleccione la operación:")
	for index, task := range tasks {
		fmt.Printf("%d) %s\n", index, task)
	}

	actionStr := prompter.Prompt("\nIngrese el número de la operación:")
	action, err := strconv.Atoi(strings.TrimSpace(actionStr))
	if err != nil {
		fmt.Println("Entrada no válida")
		return
	}

	originPath := prompter.Prompt("\nIngrese ruta de fuentes a procesar:")
	destinyTech := prompter.Prompt("\nIngrese tecnología de destino (si aplica):")
	resultPath := prompter.Prompt("\nIngrese carpeta de destino de solución:")
	fileWriter := output.NewFileWriter(resultPath)

	prompt, exists := prompts[action]
	if !exists {
		fmt.Println("Operación no válida")
		return
	}

	// Format prompt with user inputs
	prompt = fmt.Sprintf(prompt, originPath, destinyTech)

	fmt.Printf("Ejecutando tarea: %s\n", tasks[action])

	taskDev := usecases.NewTaskDev(client, fileWriter)
	err = taskDev.TaskDev(prompt)
	if err != nil {
		fmt.Println("Error durante la tarea:", err)
	}
}
