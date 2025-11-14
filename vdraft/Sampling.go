package vdraft

type Sampling struct {
	Description string             `json:"description"`
	Properties  SamplingProperties `json:"properties"`
	Type        AnnotationsType    `json:"type"`
}
