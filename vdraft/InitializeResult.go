package vdraft

type InitializeResult struct {
	Description string                     `json:"description"`
	Properties  InitializeResultProperties `json:"properties"`
	Required    []string                   `json:"required"`
	Type        AnnotationsType            `json:"type"`
}
