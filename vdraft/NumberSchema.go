package vdraft

type NumberSchema struct {
	Properties NumberSchemaProperties `json:"properties"`
	Required   []string               `json:"required"`
	Type       AnnotationsType        `json:"type"`
}
