package v20241105

type InputSchema struct {
	Description string                `json:"description"`
	Properties  InputSchemaProperties `json:"properties"`
	Required    []string              `json:"required"`
	Type        AnnotationsType       `json:"type"`
}
