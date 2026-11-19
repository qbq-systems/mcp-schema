package vdraft

type ModelHint struct {
	Description string              `json:"description"`
	Properties  ModelHintProperties `json:"properties"`
	Type        TypeElement         `json:"type"`
}
