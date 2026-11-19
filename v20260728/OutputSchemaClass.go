package v20260728

type OutputSchemaClass struct {
	AdditionalProperties AdditionalPropertiesClass `json:"additionalProperties"`
	Description          string                    `json:"description"`
	Type                 TypeElement               `json:"type"`
	Properties           *PurpleProperties         `json:"properties,omitempty"`
}
