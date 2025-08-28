package vdraft

type ElicitRequest struct {
	Description string                    `json:"description"`
	Properties  ElicitRequestProperties   `json:"properties"`
	Required    []CallToolRequestRequired `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
