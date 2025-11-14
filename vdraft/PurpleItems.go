package vdraft

type PurpleItems struct {
	Description string          `json:"description"`
	Properties  Properties2     `json:"properties"`
	Required    []string        `json:"required"`
	Type        AnnotationsType `json:"type"`
}
