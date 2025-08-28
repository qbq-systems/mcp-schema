package vdraft

type ListSRequest struct {
	Description string                       `json:"description"`
	Properties  ListPromptsRequestProperties `json:"properties"`
	Required    []CallToolRequestRequired    `json:"required"`
	Type        AnnotationsType              `json:"type"`
}
