package vdraft

type CompleteRequest struct {
	Description string                    `json:"description"`
	Properties  CompleteRequestProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
