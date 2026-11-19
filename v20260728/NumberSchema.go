package v20260728

type NumberSchema struct {
	Properties NumberSchemaProperties `json:"properties"`
	Required   []string               `json:"required"`
	Type       TypeElement            `json:"type"`
}
