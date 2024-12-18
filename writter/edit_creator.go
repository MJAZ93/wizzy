package writter

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"wizzy/core/model"
)

// Function to extract blocks between @@ and -@@ markers
func extractBlocksWithMarkers(templateContent string) ([]struct {
	RegexPattern string
	Content      string
}, error) {
	var blocks []struct {
		RegexPattern string
		Content      string
	}

	// Regular expression to find blocks between @@ and -@@
	blockRegex := regexp.MustCompile(`@@\s*(.*?)\n([\s\S]*?)\n-@@`)
	matches := blockRegex.FindAllStringSubmatch(templateContent, -1)

	if matches == nil {
		return nil, fmt.Errorf("no blocks found between @@ and -@@ markers")
	}

	for _, match := range matches {
		if len(match) > 2 {
			regexPattern := strings.TrimSpace(match[1])
			content := match[2]
			blocks = append(blocks, struct {
				RegexPattern string
				Content      string
			}{
				RegexPattern: regexPattern,
				Content:      content,
			})
		}
	}

	return blocks, nil
}

// Function to insert processed blocks into the destination file
func insertProcessedBlocks(templatePath, destFilePath string, params []model.Param) error {
	// Read the template file
	templateBytes, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template file %s: %w", templatePath, err)
	}
	templateContent := string(templateBytes)

	// Extract blocks between @@ and -@@ markers
	blocks, err := extractBlocksWithMarkers(templateContent)
	if err != nil {
		return err
	}

	// Convert params to map
	paramMap := make(map[string]string)
	for _, p := range params {
		paramMap[p.Id] = p.Value
	}

	// Read the destination file
	destFileBytes, err := os.ReadFile(destFilePath)
	if err != nil {
		return fmt.Errorf("failed to read destination file %s: %w", destFilePath, err)
	}
	destContent := string(destFileBytes)

	// Process each block and insert into destination file
	for _, block := range blocks {
		// Process the code template
		processedContent, err := processTemplate(block.Content, paramMap)
		if err != nil {
			return fmt.Errorf("failed to process block for regex '%s': %w", block.RegexPattern, err)
		}

		// Insert the processed content into the destination content
		destContent, err = insertCodeAtRegex(destContent, processedContent, block.RegexPattern)
		if err != nil {
			return fmt.Errorf("failed to insert code at regex '%s': %w", block.RegexPattern, err)
		}
	}

	// Write the updated content back to the destination file
	err = os.WriteFile(destFilePath, []byte(destContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write updated content to destination file %s: %w", destFilePath, err)
	}

	return nil
}

// Function to insert code into destination content at the regex pattern
func insertCodeAtRegex(destContent, codeToInsert, regexPattern string) (string, error) {
	// Compile the regex pattern
	re, err := regexp.Compile(regexPattern)
	if err != nil {
		return "", fmt.Errorf("invalid regex pattern '%s': %w", regexPattern, err)
	}

	// Find the index in the destination content where the regex matches
	loc := re.FindStringIndex(destContent)
	if loc == nil {
		return "", fmt.Errorf("regex pattern '%s' not found in destination content", regexPattern)
	}

	// Decide where to insert the code
	insertionPoint := loc[1] // Insert after the matched regex
	// If you want to insert before the matched regex, use loc[0]

	// Ensure proper indentation
	// Get the line where the match ends
	lineStart := strings.LastIndex(destContent[:insertionPoint], "\n") + 1
	lineContent := destContent[lineStart:insertionPoint]
	indent := strings.Repeat(" ", len(lineContent)-len(strings.TrimLeft(lineContent, " ")))

	// Indent the code to insert
	codeLines := strings.Split(codeToInsert, "\n")
	for i, line := range codeLines {
		if strings.TrimSpace(line) != "" {
			codeLines[i] = indent + line
		}
	}
	indentedCode := strings.Join(codeLines, "\n")

	// Insert the code at the insertion point
	newContent := destContent[:insertionPoint] + "\n" + indentedCode + destContent[insertionPoint:]

	return newContent, nil
}
