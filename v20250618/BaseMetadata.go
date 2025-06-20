package v20250618

type BaseMetadata struct {
	Description string                 `json:"description"`
	Properties  BaseMetadataProperties `json:"properties"`
	Required    []string               `json:"required"`
	Type        AnnotationsType        `json:"type"`
}
