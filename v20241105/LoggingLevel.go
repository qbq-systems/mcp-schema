package v20241105

type LoggingLevel struct {
	Description string      `json:"description"`
	Enum        []string    `json:"enum"`
	Type        TypeElement `json:"type"`
}
