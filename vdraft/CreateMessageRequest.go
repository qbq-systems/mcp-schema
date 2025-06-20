package vdraft

type CreateMessageRequest struct {
	Description string                         `json:"description"`
	Properties  CreateMessageRequestProperties `json:"properties"`
	Required    []string                       `json:"required"`
	Type        AnnotationsType                `json:"type"`
}
