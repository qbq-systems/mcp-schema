package vdraft

type ContentClass struct {
	AdditionalProperties AdditionalProperties `json:"additionalProperties"`
	Description          string               `json:"description"`
	Type                 AnnotationsType      `json:"type"`
}
