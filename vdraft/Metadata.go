package vdraft

type Metadata struct {
	AdditionalProperties bool                      `json:"additionalProperties"`
	Properties           AdditionalPropertiesClass `json:"properties"`
	Type                 AnnotationsType           `json:"type"`
	Description          *string                   `json:"description,omitempty"`
}
