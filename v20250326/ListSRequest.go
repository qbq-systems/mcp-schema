package v20250326

type ListSRequest struct {
	Description string                       `json:"description"`
	Properties  ListPromptsRequestProperties `json:"properties"`
	Required    []string                     `json:"required"`
	Type        AnnotationsType              `json:"type"`
}
