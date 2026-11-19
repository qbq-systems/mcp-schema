package vdraft

type FluffyContent struct {
	AdditionalProperties ContentAdditionalProperties `json:"additionalProperties"`
	Description          string                      `json:"description"`
	Type                 TypeElement                 `json:"type"`
}
