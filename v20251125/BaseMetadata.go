package v20251125

type BaseMetadata struct {
	Description string                 `json:"description"`
	Properties  BaseMetadataProperties `json:"properties"`
	Required    []string               `json:"required"`
	Type        AnnotationsType        `json:"type"`
}
