package vdraft

type ResourceContentsClass struct {
	Description string                         `json:"description"`
	Properties  BlobResourceContentsProperties `json:"properties"`
	Required    []FormatElement                `json:"required"`
	Type        AnnotationsType                `json:"type"`
}
