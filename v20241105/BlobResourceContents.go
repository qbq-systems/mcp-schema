package v20241105

type BlobResourceContents struct {
	Properties BlobResourceContentsProperties `json:"properties"`
	Required   []string                       `json:"required"`
	Type       AnnotationsType                `json:"type"`
}
