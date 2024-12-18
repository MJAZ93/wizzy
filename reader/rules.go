package reader

import (
	"fmt"
	"strings"
	"wizzy/core/model"
)

func getRules(rules []model.StringRule, filePath string) ([]model.Rule, error) {
	var objRules []model.Rule
	for _, rule := range rules {
		parts := strings.Split(rule.Rule, "->")
		if len(parts) != 2 {
			return objRules, fmt.Errorf("error reading rule, expected a -> b, got: %s", rule)
		}

		origin := strings.TrimSpace(parts[0])
		destination := strings.TrimSpace(parts[1])

		if origin == "template" {
			objRules = append(objRules, model.Rule{
				Rule:        model.TemplateRule,
				Origin:      "N/A",
				Destination: destination,
				Path:        filePath,
				Condition:   rule.Condition,
			})
		} else {
			objRules = append(objRules, model.Rule{
				Rule:        model.FileRule,
				Origin:      origin,
				Destination: destination,
				Path:        filePath,
				Condition:   rule.Condition,
			})
		}
	}

	return objRules, nil
}
