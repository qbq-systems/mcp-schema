package vdraft

type GetTaskPayloadResult struct {
	AdditionalProperties AdditionalPropertiesClass `json:"additionalProperties"`
	Description          *string                   `json:"description,omitempty"`
	Type                 AnnotationsType           `json:"type"`
	Properties           *PurpleProperties         `json:"properties,omitempty"`
}
