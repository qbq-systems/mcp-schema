package v20250326

type Sampling struct {
	AdditionalProperties bool                      `json:"additionalProperties"`
	Properties           AdditionalPropertiesClass `json:"properties"`
	Type                 AnnotationsType           `json:"type"`
	Description          *string                   `json:"description,omitempty"`
}
