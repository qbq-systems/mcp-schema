package v20250326

type EmbeddedResource struct {
	Description string                     `json:"description"`
	Properties  EmbeddedResourceProperties `json:"properties"`
	Required    []string                   `json:"required"`
	Type        AnnotationsType            `json:"type"`
}
