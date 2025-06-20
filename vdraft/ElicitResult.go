package vdraft

type ElicitResult struct {
	Description string                 `json:"description"`
	Properties  ElicitResultProperties `json:"properties"`
	Required    []string               `json:"required"`
	Type        AnnotationsType        `json:"type"`
}
