package v20250618

type PutSchema struct {
	Description string                `json:"description"`
	Properties  InputSchemaProperties `json:"properties"`
	Required    []string              `json:"required"`
	Type        AnnotationsType       `json:"type"`
}
