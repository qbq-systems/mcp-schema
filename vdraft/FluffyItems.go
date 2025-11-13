package vdraft

type FluffyItems struct {
	Description string              `json:"description"`
	Properties  AmbitiousProperties `json:"properties"`
	Required    []string            `json:"required"`
	Type        AnnotationsType     `json:"type"`
}
