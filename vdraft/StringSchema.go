package vdraft

type StringSchema struct {
	Properties StringSchemaProperties `json:"properties"`
	Required   []string               `json:"required"`
	Type       AnnotationsType        `json:"type"`
}
