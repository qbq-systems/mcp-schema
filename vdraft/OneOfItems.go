package vdraft

type OneOfItems struct {
	Properties HilariousProperties `json:"properties"`
	Required   []string            `json:"required"`
	Type       AnnotationsType     `json:"type"`
}
