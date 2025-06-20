package vdraft

type GetPromptRequest struct {
	Description string                     `json:"description"`
	Properties  GetPromptRequestProperties `json:"properties"`
	Required    []string                   `json:"required"`
	Type        AnnotationsType            `json:"type"`
}
