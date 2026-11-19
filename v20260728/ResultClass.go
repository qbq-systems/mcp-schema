package v20260728

type ResultClass struct {
	AdditionalProperties *AdditionalPropertiesClass `json:"additionalProperties,omitempty"`
	Description          string                     `json:"description"`
	Properties           IndecentProperties         `json:"properties"`
	Required             []string                   `json:"required"`
	Type                 TypeElement                `json:"type"`
}
