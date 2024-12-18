package model

// Enum for Type
type ParameterType string

const (
	FreeType      ParameterType = "free"
	ListType      ParameterType = "list"
	SelectType    ParameterType = "select"
	FormattedType ParameterType = "formatted"
)

// Structs to match the YAML structure
type Parameter struct {
	ID        string        `json:"id"`
	Desc      string        `json:"desc"`
	Regex     string        `json:"regex"`
	Type      ParameterType `json:"type"`
	Required  bool          `json:"required"`
	Options   []string      `json:"main.go,omitempty"`
	Condition string        `json:"condition,omitempty"`
}

type Template struct {
	Parameters  []Parameter  `json:"parameters"`
	StringRules []StringRule `json:"rules"`
	Rules       []Rule
	Folder      []Folder
}
