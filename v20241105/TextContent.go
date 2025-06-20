package v20241105

type TextContent struct {
	Description string                `json:"description"`
	Properties  TextContentProperties `json:"properties"`
	Required    []string              `json:"required"`
	Type        AnnotationsType       `json:"type"`
}
