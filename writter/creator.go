package writter

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"wizzy/core/model"
)

type block struct {
	kind      string
	condition string
	lines     []string
}

func createFileWithContent(dir, fileName, content string) (string, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return "", fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	filePath := filepath.Join(dir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return "", fmt.Errorf("failed to write content to file %s: %w", filePath, err)
	}

	return filePath, nil
}

func generateFileContent(templatePath string, params []model.Param) (string, error) {
	templateBytes, err := os.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template file %s: %w", templatePath, err)
	}
	templateContent := string(templateBytes)

	paramMap := make(map[string]string)
	for _, p := range params {
		paramMap[p.Id] = p.Value
	}

	return processTemplate(templateContent, paramMap)
}

func processTemplate(templateContent string, paramMap map[string]string) (string, error) {
	var output strings.Builder

	// Regular expressions
	variablesRegex := regexp.MustCompile(`\{\{(\w+)\}\}`)
	inlineIfRegex := regexp.MustCompile(`{%if\((.*?)\)%}(.*?){%endif%}`)
	blockIfRegex := regexp.MustCompile(`^{%if\((.*?)\)%}$`)
	blockForRegex := regexp.MustCompile(`^{%for\s*\((.*?)\s+as\s+(.*?)\)\s*%}$`)
	blockEndRegex := regexp.MustCompile(`^{%end(for|if)%}$`)

	// Stack to handle nested blocks
	var blockStack []block

	// Split template content into lines
	lines := strings.Split(templateContent, "\n")

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Handle block-level {%if(condition)%}
		if matches := blockIfRegex.FindStringSubmatch(trimmedLine); matches != nil {
			condition := matches[1]
			blockStack = append(blockStack, block{
				kind:      "if",
				condition: condition,
				lines:     []string{},
			})
			continue
		}

		// Handle block-level {%for (collection as item)%}
		if matches := blockForRegex.FindStringSubmatch(trimmedLine); matches != nil {
			collection := matches[1]
			item := matches[2]
			blockStack = append(blockStack, block{
				kind:      "for",
				condition: fmt.Sprintf("%s as %s", collection, item),
				lines:     []string{},
			})
			continue
		}

		// Handle {%endif%} or {%endfor%}
		if blockEndRegex.MatchString(trimmedLine) {
			if len(blockStack) == 0 {
				return "", fmt.Errorf("unmatched end tag: %s", trimmedLine)
			}

			// Pop the last block
			lastBlock := blockStack[len(blockStack)-1]
			blockStack = blockStack[:len(blockStack)-1]

			// Process the block
			var blockResult string
			var err error
			switch lastBlock.kind {
			case "if":
				blockResult, err = processIfBlock(lastBlock, paramMap)
			case "for":
				blockResult, err = processForBlock(lastBlock, paramMap)
			}
			if err != nil {
				return "", err
			}

			// Add the block result to the parent block or output
			if len(blockStack) > 0 {
				blockStack[len(blockStack)-1].lines = append(blockStack[len(blockStack)-1].lines, blockResult)
			} else {
				output.WriteString(blockResult)
			}
			continue
		}

		// Handle content inside blocks
		if len(blockStack) > 0 {
			blockStack[len(blockStack)-1].lines = append(blockStack[len(blockStack)-1].lines, line)
		} else {
			// Handle inline {%if%} conditions
			line = inlineIfRegex.ReplaceAllStringFunc(line, func(match string) string {
				matches := inlineIfRegex.FindStringSubmatch(match)
				if len(matches) != 3 {
					return match // Leave it as is if the pattern doesn't match properly
				}
				condition := matches[1]
				content := matches[2]

				// Evaluate the condition
				result, err := evaluateCondition(condition, paramMap)
				if err != nil || !result {
					return "" // Omit content if condition fails
				}
				return content // Return content if condition passes
			})

			// Replace variables in the processed line
			processedLine := variablesRegex.ReplaceAllStringFunc(line, func(match string) string {
				varName := match[2 : len(match)-2]
				if val, ok := paramMap[varName]; ok {
					return val
				}
				return match // Leave as is if not found
			})

			output.WriteString(processedLine + "\n")
		}
	}

	return output.String(), nil
}

// evaluateCondition evaluates a condition string like "param==value"
func evaluateCondition(condition string, paramMap map[string]string) (bool, error) {
	conditionParts := strings.Split(condition, "==")
	if len(conditionParts) != 2 {
		return false, fmt.Errorf("invalid condition format: %s", condition)
	}
	paramName := strings.TrimSpace(conditionParts[0])
	expectedValue := strings.TrimSpace(conditionParts[1])

	actualValue, ok := paramMap[paramName]
	if !ok {
		return false, nil // Parameter not found, condition is false
	}

	return actualValue == expectedValue, nil
}

// processIfBlock processes an if block and returns the result.
func processIfBlock(b block, paramMap map[string]string) (string, error) {
	result, err := evaluateCondition(b.condition, paramMap)
	if err != nil {
		return "", err
	}

	if result {
		// Process the lines inside the if block
		blockContent := strings.Join(b.lines, "\n")
		return processTemplate(blockContent, paramMap)
	}
	// Condition is false; return empty string
	return "", nil
}

// processForBlock processes a for block and returns the result.
func processForBlock(b block, paramMap map[string]string) (string, error) {
	// Parse the condition: "collection as item"
	parts := strings.Split(b.condition, " as ")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid for loop format: %s", b.condition)
	}
	collectionName := strings.TrimSpace(parts[0])
	itemName := strings.TrimSpace(parts[1])

	// Get the collection from paramMap
	collectionValue, ok := paramMap[collectionName]
	if !ok {
		return "", fmt.Errorf("collection '%s' not found in params", collectionName)
	}

	// Assume the collection is a comma-separated string
	items := strings.Split(collectionValue, ",")

	var result strings.Builder
	for _, item := range items {
		// Trim spaces
		item = strings.TrimSpace(item)

		// Set the item in paramMap temporarily
		paramMap[itemName] = item

		// Process the lines inside the for block
		blockContent := strings.Join(b.lines, "\n")
		processedBlock, err := processTemplate(blockContent, paramMap)
		if err != nil {
			return "", err
		}
		result.WriteString(processedBlock)

		// Remove the temporary item from paramMap
		delete(paramMap, itemName)
	}

	return result.String(), nil
}
