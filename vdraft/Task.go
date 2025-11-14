package vdraft

type Task struct {
	Description string          `json:"description"`
	Properties  TaskProperties  `json:"properties"`
	Required    []string        `json:"required"`
	Type        AnnotationsType `json:"type"`
}
