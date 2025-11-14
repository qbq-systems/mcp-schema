package vdraft

type Context struct {
	Description string           `json:"description"`
	Properties  FluffyProperties `json:"properties"`
	Type        AnnotationsType  `json:"type"`
}
