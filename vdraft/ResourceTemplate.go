package vdraft

type ResourceTemplate struct {
	Description string                     `json:"description"`
	Properties  ResourceTemplateProperties `json:"properties"`
	Required    []string                   `json:"required"`
	Type        AnnotationsType            `json:"type"`
}
