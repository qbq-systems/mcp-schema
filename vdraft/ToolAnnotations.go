package vdraft

type ToolAnnotations struct {
	Description string                    `json:"description"`
	Properties  ToolAnnotationsProperties `json:"properties"`
	Type        AnnotationsType           `json:"type"`
}
