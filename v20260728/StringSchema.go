package v20260728

type StringSchema struct {
	Properties StringSchemaProperties `json:"properties"`
	Required   []string               `json:"required"`
	Type       TypeElement            `json:"type"`
}
