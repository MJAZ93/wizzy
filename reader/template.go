package reader

import (
	"encoding/json"
	"fmt"
	"os"
	"wizzy/core/model"
)

func readTemplate(filePath string) (model.Template, error) {
	fileName := filePath + "template.json"
	file, err := os.Open(fileName)
	if err != nil {
		return model.Template{}, fmt.Errorf("failed to open file %s: %w", fileName, err)
	}
	defer file.Close()

	var template model.Template
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&template); err != nil {
		return template, fmt.Errorf("failed to decode JSON in %s: %w", fileName, err)
	}

	for _, param := range template.Parameters {
		if param.Type != model.FreeType && param.Type != model.ListType && param.Type != model.FormattedType && param.Type != model.SelectType {
			return model.Template{}, fmt.Errorf("failed to parse file: '%s': type '%s' from param '%s' is not allowed", fileName, param.Type, param.ID)
		}
	}

	foldersObjects, err := listFolders(filePath)
	if err != nil {
		return model.Template{}, fmt.Errorf("failed to list folders in %s: %w", filePath, err)
	}
	template.Folder = foldersObjects

	rules, err := getRules(template.StringRules, filePath)
	if err != nil {
		return model.Template{}, fmt.Errorf("failed to parse rule in %s: %w", filePath, err)
	}
	template.Rules = rules

	return template, nil
}
