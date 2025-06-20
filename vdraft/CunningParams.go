package vdraft

type CunningParams struct {
	Properties MischievousProperties `json:"properties"`
	Required   []string              `json:"required"`
	Type       AnnotationsType       `json:"type"`
}
