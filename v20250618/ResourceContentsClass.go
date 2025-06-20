package v20250618

type ResourceContentsClass struct {
	Description string                         `json:"description"`
	Properties  BlobResourceContentsProperties `json:"properties"`
	Required    []Required                     `json:"required"`
	Type        AnnotationsType                `json:"type"`
}
