package v20251125

type EmbeddedResource struct {
	Description string                     `json:"description"`
	Properties  EmbeddedResourceProperties `json:"properties"`
	Required    []string                   `json:"required"`
	Type        AnnotationsType            `json:"type"`
}
