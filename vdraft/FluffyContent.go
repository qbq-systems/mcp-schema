package vdraft

type FluffyContent struct {
	AdditionalProperties AdditionalProperties `json:"additionalProperties"`
	Description          string               `json:"description"`
	Type                 AnnotationsType      `json:"type"`
}
