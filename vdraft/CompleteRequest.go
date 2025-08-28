package vdraft

type CompleteRequest struct {
	Description string                    `json:"description"`
	Properties  CompleteRequestProperties `json:"properties"`
	Required    []CallToolRequestRequired `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
