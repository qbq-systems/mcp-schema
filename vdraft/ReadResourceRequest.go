package vdraft

type ReadResourceRequest struct {
	Description string                        `json:"description"`
	Properties  ReadResourceRequestProperties `json:"properties"`
	Required    []CallToolRequestRequired     `json:"required"`
	Type        AnnotationsType               `json:"type"`
}
