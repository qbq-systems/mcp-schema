package vdraft

type Elicitation struct {
	Description string                `json:"description"`
	Properties  ElicitationProperties `json:"properties"`
	Type        AnnotationsType       `json:"type"`
}
