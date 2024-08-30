package usecases

import (
	"context"
	"fmt"

	"migrania/internal/domain/entities"
	"migrania/internal/infrastructure/openai"
	"migrania/internal/infrastructure/output"
)

type Migrator struct {
	client     *openai.Client
	fileWriter *output.FileWriter
}

func NewMigrator(client *openai.Client, fileWriter *output.FileWriter) *Migrator {
	return &Migrator{client: client, fileWriter: fileWriter}
}

func (m *Migrator) Migrate(originPath, originTech, destinyTech string) error {
	sources := entities.ReadFilePathsFromPath(originPath)

	prompt := fmt.Sprintf(entities.PROMPT_MIGRACION, sources, originTech, destinyTech)

	responseContent, err := m.client.RequestImprovement(context.Background(), prompt)
	if err != nil {
		return err
	}

	return m.fileWriter.WriteImprovementFiles(responseContent)
}
