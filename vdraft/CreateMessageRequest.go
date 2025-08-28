package vdraft

type CreateMessageRequest struct {
	Description string                         `json:"description"`
	Properties  CreateMessageRequestProperties `json:"properties"`
	Required    []CallToolRequestRequired      `json:"required"`
	Type        AnnotationsType                `json:"type"`
}
