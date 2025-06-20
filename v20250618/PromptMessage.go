package v20250618

type PromptMessage struct {
	Description string                  `json:"description"`
	Properties  PromptMessageProperties `json:"properties"`
	Required    []string                `json:"required"`
	Type        AnnotationsType         `json:"type"`
}
