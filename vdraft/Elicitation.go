package vdraft

type Elicitation struct {
	AdditionalProperties bool                      `json:"additionalProperties"`
	Description          *string                   `json:"description,omitempty"`
	Properties           AdditionalPropertiesClass `json:"properties"`
	Type                 AnnotationsType           `json:"type"`
}
