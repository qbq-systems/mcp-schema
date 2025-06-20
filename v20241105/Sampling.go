package v20241105

type Sampling struct {
	AdditionalProperties bool                      `json:"additionalProperties"`
	Properties           AdditionalPropertiesClass `json:"properties"`
	Type                 AnnotationsType           `json:"type"`
	Description          *string                   `json:"description,omitempty"`
}
