package v20251125

type ReadResourceResult struct {
	Description string                       `json:"description"`
	Properties  ReadResourceResultProperties `json:"properties"`
	Required    []string                     `json:"required"`
	Type        AnnotationsType              `json:"type"`
}
