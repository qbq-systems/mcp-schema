package v20250618

type ListResourcesResult struct {
	Description string                        `json:"description"`
	Properties  ListResourcesResultProperties `json:"properties"`
	Required    []string                      `json:"required"`
	Type        AnnotationsType               `json:"type"`
}
