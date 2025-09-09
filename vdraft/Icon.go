package vdraft

type Icon struct {
	Description string          `json:"description"`
	Properties  IconProperties  `json:"properties"`
	Required    []string        `json:"required"`
	Type        AnnotationsType `json:"type"`
}
