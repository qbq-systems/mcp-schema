package v20260728

type BooleanSchema struct {
	Properties BooleanSchemaProperties `json:"properties"`
	Required   []string                `json:"required"`
	Type       TypeElement             `json:"type"`
}
