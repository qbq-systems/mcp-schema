package v20251125

type PromptMessage struct {
	Description string                  `json:"description"`
	Properties  PromptMessageProperties `json:"properties"`
	Required    []string                `json:"required"`
	Type        AnnotationsType         `json:"type"`
}
