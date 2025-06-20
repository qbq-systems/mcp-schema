package v20241105

type InputSchemaProperties struct {
	Properties PropertiesProperties `json:"properties"`
	Required   StopSequences        `json:"required"`
	Type       Method               `json:"type"`
}
