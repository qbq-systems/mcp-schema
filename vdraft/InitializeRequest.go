package vdraft

type InitializeRequest struct {
	Description string                      `json:"description"`
	Properties  InitializeRequestProperties `json:"properties"`
	Required    []CallToolRequestRequired   `json:"required"`
	Type        AnnotationsType             `json:"type"`
}
