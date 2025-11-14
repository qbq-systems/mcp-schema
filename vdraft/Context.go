package vdraft

type Context struct {
	Description string              `json:"description"`
	Properties  AmbitiousProperties `json:"properties"`
	Type        AnnotationsType     `json:"type"`
}
