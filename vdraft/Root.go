package vdraft

type Root struct {
	Description string          `json:"description"`
	Properties  RootProperties  `json:"properties"`
	Required    []FormatElement `json:"required"`
	Type        AnnotationsType `json:"type"`
}
