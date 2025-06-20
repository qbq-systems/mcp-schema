package v20250618

type EnumSchema struct {
	Properties EnumSchemaProperties `json:"properties"`
	Required   []string             `json:"required"`
	Type       AnnotationsType      `json:"type"`
}
