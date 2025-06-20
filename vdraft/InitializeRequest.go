package vdraft

type InitializeRequest struct {
	Description string                      `json:"description"`
	Properties  InitializeRequestProperties `json:"properties"`
	Required    []string                    `json:"required"`
	Type        AnnotationsType             `json:"type"`
}
