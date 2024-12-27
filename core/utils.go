package core

type Action int

const (
	Navigate Action = iota
	Execute
	NONE
)

func (s Action) String() string {
	switch s {
	case Navigate:
		return "Navigate"
	case Execute:
		return "Execute"
	default:
		return "Unknown"
	}
}

func ActionFromString(act string) Action {
	switch act {
	case "Navigate":
		return Navigate
	case "Execute":
		return Execute
	default:
		return NONE
	}
}
