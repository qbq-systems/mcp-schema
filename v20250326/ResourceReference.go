package v20250326

type ResourceReference struct {
	Description string                      `json:"description"`
	Properties  ResourceReferenceProperties `json:"properties"`
	Required    []string                    `json:"required"`
	Type        AnnotationsType             `json:"type"`
}
