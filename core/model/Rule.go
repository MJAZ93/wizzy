package model

import (
	"strings"
)

type ExecutionRule string

const (
	FileRule     ExecutionRule = "file"
	TemplateRule ExecutionRule = "template"
)

type Rule struct {
	Rule        ExecutionRule
	Origin      string
	Destination string
	Path        string
	Condition   string
}

type StringRule struct {
	Rule      string
	Condition string
}

func (rule Rule) MatchesRule(param []Param) bool {

	if rule.Condition == "" {
		return true
	}

	parts := strings.Split(rule.Condition, "==")
	if len(parts) == 2 {
		condition := parts[0]
		value := parts[1]

		for _, p := range param {
			if p.Id == condition && p.Value == value {
				return true
			}
		}
	}

	return false
}
