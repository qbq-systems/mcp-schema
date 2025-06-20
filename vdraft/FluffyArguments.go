package vdraft

type FluffyArguments struct {
	AdditionalProperties Default         `json:"additionalProperties"`
	Description          string          `json:"description"`
	Type                 AnnotationsType `json:"type"`
}
