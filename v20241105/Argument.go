package v20241105

type Argument struct {
	Description string             `json:"description"`
	Properties  ArgumentProperties `json:"properties"`
	Required    []string           `json:"required"`
	Type        AnnotationsType    `json:"type"`
}
