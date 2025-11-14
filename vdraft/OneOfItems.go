package vdraft

type OneOfItems struct {
	Properties Properties3     `json:"properties"`
	Required   []string        `json:"required"`
	Type       AnnotationsType `json:"type"`
}
