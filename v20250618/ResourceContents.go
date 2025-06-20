package v20250618

type ResourceContents struct {
	Properties BlobResourceContentsProperties `json:"properties"`
	Required   []string                       `json:"required"`
	Type       AnnotationsType                `json:"type"`
}
