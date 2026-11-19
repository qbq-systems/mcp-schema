package v20260728

type ResourceTemplateReference struct {
	Description string                              `json:"description"`
	Properties  ResourceTemplateReferenceProperties `json:"properties"`
	Required    []string                            `json:"required"`
	Type        TypeElement                         `json:"type"`
}
