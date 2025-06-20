package vdraft

type Resource struct {
	Description string             `json:"description"`
	Properties  ResourceProperties `json:"properties"`
	Required    []string           `json:"required"`
	Type        AnnotationsType    `json:"type"`
}
