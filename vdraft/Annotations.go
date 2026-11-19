package vdraft

type Annotations struct {
	Description string                `json:"description"`
	Properties  AnnotationsProperties `json:"properties"`
	Type        TypeElement           `json:"type"`
}
