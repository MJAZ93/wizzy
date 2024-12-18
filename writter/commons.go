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
		return false, fmt.Errorf("directory does not exist: %s", dir)
	}

	filePath := filepath.Join(dir, file)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false, nil
	}

	return true, nil
}
