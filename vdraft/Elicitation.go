package vdraft

type Elicitation struct {
	Description string                `json:"description"`
	Properties  ElicitationProperties `json:"properties"`
	Type        TypeElement           `json:"type"`
}
