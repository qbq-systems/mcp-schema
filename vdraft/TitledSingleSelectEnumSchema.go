package vdraft

type TitledSingleSelectEnumSchema struct {
	Description string                                 `json:"description"`
	Properties  TitledSingleSelectEnumSchemaProperties `json:"properties"`
	Required    []string                               `json:"required"`
	Type        AnnotationsType                        `json:"type"`
}
