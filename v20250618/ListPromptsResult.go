package v20250618

type ListPromptsResult struct {
	Description string                      `json:"description"`
	Properties  ListPromptsResultProperties `json:"properties"`
	Required    []string                    `json:"required"`
	Type        AnnotationsType             `json:"type"`
}
