package vdraft

type EnumSchema struct {
	Properties EnumSchemaProperties `json:"properties"`
	Required   []string             `json:"required"`
	Type       AnnotationsType      `json:"type"`
}
