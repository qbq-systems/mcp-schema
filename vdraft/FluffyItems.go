package vdraft

type FluffyItems struct {
	Description string            `json:"description"`
	Properties  CunningProperties `json:"properties"`
	Required    []string          `json:"required"`
	Type        AnnotationsType   `json:"type"`
}
