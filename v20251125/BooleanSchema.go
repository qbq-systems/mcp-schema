package v20251125

type BooleanSchema struct {
	Properties BooleanSchemaProperties `json:"properties"`
	Required   []string                `json:"required"`
	Type       AnnotationsType         `json:"type"`
}
