package v20250618

type PromptArgument struct {
	Description string                   `json:"description"`
	Properties  PromptArgumentProperties `json:"properties"`
	Required    []string                 `json:"required"`
	Type        AnnotationsType          `json:"type"`
}
