package vdraft

type BooleanSchema struct {
	Properties BooleanSchemaProperties `json:"properties"`
	Required   []string                `json:"required"`
	Type       AnnotationsType         `json:"type"`
}
