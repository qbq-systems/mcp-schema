package v20241105

type InitializeResult struct {
	Description string                     `json:"description"`
	Properties  InitializeResultProperties `json:"properties"`
	Required    []string                   `json:"required"`
	Type        AnnotationsType            `json:"type"`
}
