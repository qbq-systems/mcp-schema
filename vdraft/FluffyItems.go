package vdraft

type FluffyItems struct {
	Description string          `json:"description"`
	Properties  Properties5     `json:"properties"`
	Required    []string        `json:"required"`
	Type        AnnotationsType `json:"type"`
}
