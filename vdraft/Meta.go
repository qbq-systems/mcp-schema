package vdraft

type Meta struct {
	AdditionalProperties AdditionalPropertiesClass `json:"additionalProperties"`
	Description          string                    `json:"description"`
	Type                 AnnotationsType           `json:"type"`
	Properties           *PurpleProperties         `json:"properties,omitempty"`
}
