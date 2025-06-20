package vdraft

type Implementation struct {
	Description string                   `json:"description"`
	Properties  ImplementationProperties `json:"properties"`
	Required    []string                 `json:"required"`
	Type        AnnotationsType          `json:"type"`
}
