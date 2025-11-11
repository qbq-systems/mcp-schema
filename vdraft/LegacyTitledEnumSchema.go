package vdraft

type LegacyTitledEnumSchema struct {
	Description string                           `json:"description"`
	Properties  LegacyTitledEnumSchemaProperties `json:"properties"`
	Required    []string                         `json:"required"`
	Type        AnnotationsType                  `json:"type"`
}
