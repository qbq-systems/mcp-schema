package vdraft

type ResourceContents struct {
	Properties BlobResourceContentsProperties `json:"properties"`
	Required   []string                       `json:"required"`
	Type       AnnotationsType                `json:"type"`
}
