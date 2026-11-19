package v20251125

type Argument struct {
	Description string             `json:"description"`
	Properties  ArgumentProperties `json:"properties"`
	Required    []string           `json:"required"`
	Type        AnnotationsType    `json:"type"`
}
