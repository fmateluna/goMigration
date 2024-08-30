package usecases

import (
	"context"
	"fmt"

	"migrania/internal/domain/entities"
	"migrania/internal/infrastructure/openai"
	"migrania/internal/infrastructure/output"
)

type CodeImprover struct {
	client     *openai.Client
	fileWriter *output.FileWriter
}

func NewCodeImprover(client *openai.Client, fileWriter *output.FileWriter) *CodeImprover {
	return &CodeImprover{client: client, fileWriter: fileWriter}
}

func (ci *CodeImprover) MigrateCode(originPath, destinyTech string) error {
	sources := entities.ReadFilePathsFromPath(originPath)

	prompt := fmt.Sprintf(entities.PROMPT_MIGRACION, sources, destinyTech)

	responseContent, err := ci.client.RequestImprovement(context.Background(), prompt)
	if err != nil {
		return err
	}

	return ci.fileWriter.WriteImprovementFiles(responseContent)
}

func (ci *CodeImprover) ImproveCode(originPath, destinyTech string) error {
	sources := entities.ReadFilePathsFromPath(originPath)

	prompt := fmt.Sprintf(entities.PROMPT_IMPROVE_CODE, sources, destinyTech)

	responseContent, err := ci.client.RequestImprovement(context.Background(), prompt)
	if err != nil {
		return err
	}

	return ci.fileWriter.WriteImprovementFiles(responseContent)
}

func (ci *CodeImprover) AddTestCoverage(originPath, destinyTech string) error {
	sources := entities.ReadFilePathsFromPath(originPath)

	prompt := fmt.Sprintf(entities.PROMPT_ADD_TEST_COVERAGE, sources, destinyTech)

	responseContent, err := ci.client.RequestImprovement(context.Background(), prompt)
	if err != nil {
		return err
	}

	return ci.fileWriter.WriteImprovementFiles(responseContent)
}
