package vdraft

type CallToolRequest struct {
	Description string                    `json:"description"`
	Properties  CallToolRequestProperties `json:"properties"`
	Required    []CallToolRequestRequired `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
