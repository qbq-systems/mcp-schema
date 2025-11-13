package vdraft

type OneOfItems struct {
	Properties IndecentProperties `json:"properties"`
	Required   []string           `json:"required"`
	Type       AnnotationsType    `json:"type"`
}
