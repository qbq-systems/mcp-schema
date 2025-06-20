package v20250326

type ListResourceTemplatesResult struct {
	Description string                                `json:"description"`
	Properties  ListResourceTemplatesResultProperties `json:"properties"`
	Required    []string                              `json:"required"`
	Type        AnnotationsType                       `json:"type"`
}
