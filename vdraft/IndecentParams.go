package vdraft

type IndecentParams struct {
	Properties AmbitiousProperties `json:"properties"`
	Required   []string            `json:"required"`
	Type       AnnotationsType     `json:"type"`
}
