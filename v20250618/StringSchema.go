package v20250618

type StringSchema struct {
	Properties StringSchemaProperties `json:"properties"`
	Required   []string               `json:"required"`
	Type       AnnotationsType        `json:"type"`
}
