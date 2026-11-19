package v20251125

type StringSchema struct {
	Properties StringSchemaProperties `json:"properties"`
	Required   []string               `json:"required"`
	Type       AnnotationsType        `json:"type"`
}
