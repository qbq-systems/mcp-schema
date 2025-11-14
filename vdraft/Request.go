package vdraft

type Request struct {
	Description string                      `json:"description"`
	Properties  CancelTaskRequestProperties `json:"properties"`
	Required    []CallToolRequestRequired   `json:"required"`
	Type        AnnotationsType             `json:"type"`
}
