package vdraft

type FluffyParams struct {
	Properties TentacledProperties `json:"properties"`
	Required   []string            `json:"required"`
	Type       AnnotationsType     `json:"type"`
}
