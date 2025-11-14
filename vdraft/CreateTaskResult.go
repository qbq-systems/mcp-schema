package vdraft

type CreateTaskResult struct {
	Description string                     `json:"description"`
	Properties  CreateTaskResultProperties `json:"properties"`
	Required    []string                   `json:"required"`
	Type        AnnotationsType            `json:"type"`
}
