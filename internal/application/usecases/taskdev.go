package usecases

import (
	"context"
	"fmt"
	"migrania/internal/infrastructure/openai"
	"migrania/internal/infrastructure/output"
)

type TaskDev struct {
	client     *openai.Client
	fileWriter *output.FileWriter
}

func NewTaskDev(client *openai.Client, fileWriter *output.FileWriter) *TaskDev {
	return &TaskDev{client: client, fileWriter: fileWriter}
}

func (td *TaskDev) TaskDev(prompt string) error {
	responseContent, err := td.client.RequestImprovement(context.Background(), prompt)

	fmt.Println(prompt)

	if err != nil {
		return err
	}

	return td.fileWriter.WriteImprovementFiles(responseContent)
}
