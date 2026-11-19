package v20260728

type PromptArgument struct {
	Description string                   `json:"description"`
	Properties  PromptArgumentProperties `json:"properties"`
	Required    []string                 `json:"required"`
	Type        TypeElement              `json:"type"`
}
