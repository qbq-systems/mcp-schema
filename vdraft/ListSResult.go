package vdraft

type ListSResult struct {
	Description string                      `json:"description"`
	Properties  ListPromptsResultProperties `json:"properties"`
	Required    []string                    `json:"required"`
	Type        TypeElement                 `json:"type"`
}
