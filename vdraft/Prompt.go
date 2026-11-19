package vdraft

type Prompt struct {
	Description string           `json:"description"`
	Properties  PromptProperties `json:"properties"`
	Required    []string         `json:"required"`
	Type        TypeElement      `json:"type"`
}
