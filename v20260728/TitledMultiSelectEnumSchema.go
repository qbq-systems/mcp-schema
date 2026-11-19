package v20260728

type TitledMultiSelectEnumSchema struct {
	Description string                                `json:"description"`
	Properties  TitledMultiSelectEnumSchemaProperties `json:"properties"`
	Required    []string                              `json:"required"`
	Type        TypeElement                           `json:"type"`
}
