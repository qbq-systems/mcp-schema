package v20250326

type BlobResourceContents struct {
	Properties BlobResourceContentsProperties `json:"properties"`
	Required   []string                       `json:"required"`
	Type       AnnotationsType                `json:"type"`
}
