package vdraft

type Request struct {
	Description string                     `json:"description"`
	Properties  ListRootsRequestProperties `json:"properties"`
	Required    []CallToolRequestRequired  `json:"required"`
	Type        AnnotationsType            `json:"type"`
}
