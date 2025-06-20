package v20250618

type NumberSchema struct {
	Properties NumberSchemaProperties `json:"properties"`
	Required   []string               `json:"required"`
	Type       AnnotationsType        `json:"type"`
}
