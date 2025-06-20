package v20250618

type Content struct {
	Description string                 `json:"description"`
	Properties  AudioContentProperties `json:"properties"`
	Required    []string               `json:"required"`
	Type        AnnotationsType        `json:"type"`
}
