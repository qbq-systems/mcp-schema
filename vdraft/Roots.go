package vdraft

type Roots struct {
	Description string          `json:"description"`
	Properties  RootsProperties `json:"properties"`
	Type        AnnotationsType `json:"type"`
}
