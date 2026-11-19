package v20260728

type ResourceContentsClass struct {
	Description string                         `json:"description"`
	Properties  BlobResourceContentsProperties `json:"properties"`
	Required    []FormatElement                `json:"required"`
	Type        TypeElement                    `json:"type"`
}
