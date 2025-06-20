package vdraft

type ReadResourceResult struct {
	Description string                       `json:"description"`
	Properties  ReadResourceResultProperties `json:"properties"`
	Required    []string                     `json:"required"`
	Type        AnnotationsType              `json:"type"`
}
