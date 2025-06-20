package v20241105

type ImageContent struct {
	Description string                 `json:"description"`
	Properties  ImageContentProperties `json:"properties"`
	Required    []string               `json:"required"`
	Type        AnnotationsType        `json:"type"`
}
