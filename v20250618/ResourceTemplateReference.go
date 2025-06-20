package v20250618

type ResourceTemplateReference struct {
	Description string                              `json:"description"`
	Properties  ResourceTemplateReferenceProperties `json:"properties"`
	Required    []string                            `json:"required"`
	Type        AnnotationsType                     `json:"type"`
}
