package vdraft

type UntitledMultiSelectEnumSchema struct {
	Description string                                  `json:"description"`
	Properties  UntitledMultiSelectEnumSchemaProperties `json:"properties"`
	Required    []string                                `json:"required"`
	Type        AnnotationsType                         `json:"type"`
}
