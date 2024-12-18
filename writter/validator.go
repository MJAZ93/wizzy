package writter

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"wizzy/core/model"
)

func canCreateFileInDir(dir string) error {
	testFilePath := filepath.Join(dir, ".testfile")
	file, err := os.Create(testFilePath)
	if err != nil {
		return fmt.Errorf("cannot create a file in directory %s: %w", dir, err)
	}
	defer os.Remove(testFilePath)
	defer file.Close()
	return nil
}

func checkTemplateParams(params []model.Param, templatePath string) ([]string, error) {
	templateBytes, err := os.ReadFile(templatePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read template file %s: %w", templatePath, err)
	}
	templateContent := string(templateBytes)

	// Convert params slice to a map for easy lookup
	paramMap := make(map[string]bool)
	for _, p := range params {
		paramMap[p.Id] = true
	}

	// Extract variables from the template
	usedParams, err := extractTemplateVariables(templateContent)
	if err != nil {
		return nil, fmt.Errorf("failed to extract variables from template: %w", err)
	}

	// Check for missing parameters
	var missingParams []string
	for param := range usedParams {
		if _, ok := paramMap[param]; !ok {
			missingParams = append(missingParams, param)
		}
	}

	for _, param := range params {
		if param.Required {
			for _, missingParam := range missingParams {
				if param.Id == missingParam {
					return missingParams, errors.New("the required param '" + param.Id + "' is missing")
				}
			}
		}
	}

	return missingParams, nil
}

func extractTemplateVariables(templateContent string) (map[string]bool, error) {
	var usedParams = make(map[string]bool)
	scanner := bufio.NewScanner(strings.NewReader(templateContent))

	// Regular expressions to find variables and parameters in conditions/loops
	variableRegex := regexp.MustCompile(`\{\{(\w+)\}\}`)
	conditionRegex := regexp.MustCompile(`\{%\s*(if|for)\s*(.*?)\s*%\}`)

	for scanner.Scan() {
		line := scanner.Text()

		// Find variables in the line
		variables := variableRegex.FindAllStringSubmatch(line, -1)
		for _, match := range variables {
			varName := match[1]
			usedParams[varName] = true
		}

		// Find conditions and loops
		conditions := conditionRegex.FindAllStringSubmatch(line, -1)
		for _, match := range conditions {
			conditionContent := match[2]

			// Extract variables from conditions
			conditionVariables := extractVariablesFromCondition(conditionContent)
			for _, varName := range conditionVariables {
				usedParams[varName] = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading template content: %w", err)
	}

	return usedParams, nil
}

func extractVariablesFromCondition(condition string) []string {
	var variables []string

	// Remove operators and split by non-word characters
	cleanedCondition := strings.ReplaceAll(condition, "==", " ")
	cleanedCondition = strings.ReplaceAll(cleanedCondition, "!=", " ")
	cleanedCondition = strings.ReplaceAll(cleanedCondition, ">", " ")
	cleanedCondition = strings.ReplaceAll(cleanedCondition, "<", " ")
	cleanedCondition = strings.ReplaceAll(cleanedCondition, ">=", " ")
	cleanedCondition = strings.ReplaceAll(cleanedCondition, "<=", " ")
	cleanedCondition = strings.ReplaceAll(cleanedCondition, "(", " ")
	cleanedCondition = strings.ReplaceAll(cleanedCondition, ")", " ")
	cleanedCondition = strings.ReplaceAll(cleanedCondition, ",", " ")
	cleanedCondition = strings.ReplaceAll(cleanedCondition, ".", " ")

	// Split by whitespace
	parts := strings.Fields(cleanedCondition)

	// Assume variables are words that are not numbers or keywords
	keywords := map[string]bool{
		"if":    true,
		"for":   true,
		"as":    true,
		"in":    true,
		"and":   true,
		"or":    true,
		"not":   true,
		"true":  true,
		"false": true,
	}

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		if _, isKeyword := keywords[part]; isKeyword {
			continue
		}
		if isNumber(part) {
			continue
		}
		// It's a variable
		variables = append(variables, part)
	}

	return variables
}

// isNumber checks if a string represents a number.
func isNumber(s string) bool {
	_, err := fmt.Sscan(s, new(float64))
	return err == nil
}
