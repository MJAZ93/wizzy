package model

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
