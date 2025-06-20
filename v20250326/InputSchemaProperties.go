package v20250326

type InputSchemaProperties struct {
	Properties PropertiesProperties `json:"properties"`
	Required   StopSequences        `json:"required"`
	Type       TypeClass            `json:"type"`
}
