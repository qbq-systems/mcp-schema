package v20260728

type TitledSingleSelectEnumSchema struct {
	Description string                                 `json:"description"`
	Properties  TitledSingleSelectEnumSchemaProperties `json:"properties"`
	Required    []string                               `json:"required"`
	Type        TypeElement                            `json:"type"`
}
