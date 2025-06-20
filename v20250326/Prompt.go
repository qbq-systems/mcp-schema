package v20250326

type Prompt struct {
	Description string           `json:"description"`
	Properties  PromptProperties `json:"properties"`
	Required    []string         `json:"required"`
	Type        AnnotationsType  `json:"type"`
}
