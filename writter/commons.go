package writter

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"wizzy/core/model"
)

func replacePlaceholders(template string, params []model.Param) string {
	paramMap := make(map[string]string)
	for _, p := range params {
		paramMap[p.Id] = p.Value
	}

	placeholderRegex := regexp.MustCompile(`\{\{(\w+)\}\}`)

	result := placeholderRegex.ReplaceAllStringFunc(template, func(match string) string {
		paramName := strings.Trim(match, "{}")
		if value, exists := paramMap[paramName]; exists {
			return value
		}

		return match
	})

	return result
}

func fileExistsInDir(dir string, file string) (bool, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if mkErr := os.MkdirAll(dir, 0755); mkErr != nil {
			return false, fmt.Errorf("failed to create directory: %s, error: %w", dir, mkErr)
		}
		return false, nil // Directory created, file does not exist yet
	} else if err != nil {
		return false, fmt.Errorf("failed to check directory: %s, error: %w", dir, err)
	}

	filePath := filepath.Join(dir, file)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false, nil // File does not exist
	} else if err != nil {
		return false, fmt.Errorf("failed to check file: %s, error: %w", filePath, err)
	}

	return true, nil // File exists
}
