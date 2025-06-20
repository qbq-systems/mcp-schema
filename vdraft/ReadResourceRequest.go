package vdraft

type ReadResourceRequest struct {
	Description string                        `json:"description"`
	Properties  ReadResourceRequestProperties `json:"properties"`
	Required    []string                      `json:"required"`
	Type        AnnotationsType               `json:"type"`
}
