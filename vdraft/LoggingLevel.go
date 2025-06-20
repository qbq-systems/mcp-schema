package vdraft

type LoggingLevel struct {
	Description string      `json:"description"`
	Enum        []string    `json:"enum"`
	Type        TypeElement `json:"type"`
}
