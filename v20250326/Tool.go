package v20250326

type Tool struct {
	Description string          `json:"description"`
	Properties  ToolProperties  `json:"properties"`
	Required    []string        `json:"required"`
	Type        AnnotationsType `json:"type"`
}
