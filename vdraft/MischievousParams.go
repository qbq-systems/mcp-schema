package vdraft

type MischievousParams struct {
	Properties Properties3     `json:"properties"`
	Required   []string        `json:"required"`
	Type       AnnotationsType `json:"type"`
}
