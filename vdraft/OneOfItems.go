package vdraft

type OneOfItems struct {
	Properties Properties2     `json:"properties"`
	Required   []string        `json:"required"`
	Type       AnnotationsType `json:"type"`
}
