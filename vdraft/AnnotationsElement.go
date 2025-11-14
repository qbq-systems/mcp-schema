package vdraft

type AnnotationsElement struct {
	Ref         string  `json:"$ref"`
	Description *string `json:"description,omitempty"`
}
