package vdraft

type URLElicitationRequiredError struct {
	Description string                                `json:"description"`
	Properties  URLElicitationRequiredErrorProperties `json:"properties"`
	Required    []string                              `json:"required"`
	Type        AnnotationsType                       `json:"type"`
}
