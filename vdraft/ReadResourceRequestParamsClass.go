package vdraft

type ReadResourceRequestParamsClass struct {
	Description string                              `json:"description"`
	Properties  ReadResourceRequestParamsProperties `json:"properties"`
	Required    []FormatElement                     `json:"required"`
	Type        AnnotationsType                     `json:"type"`
}
