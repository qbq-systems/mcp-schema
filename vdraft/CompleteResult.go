package vdraft

type CompleteResult struct {
	Description string                   `json:"description"`
	Properties  CompleteResultProperties `json:"properties"`
	Required    []string                 `json:"required"`
	Type        TypeElement              `json:"type"`
}
