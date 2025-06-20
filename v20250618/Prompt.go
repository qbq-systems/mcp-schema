package v20250618

type Prompt struct {
	Description string           `json:"description"`
	Properties  PromptProperties `json:"properties"`
	Required    []string         `json:"required"`
	Type        AnnotationsType  `json:"type"`
}
