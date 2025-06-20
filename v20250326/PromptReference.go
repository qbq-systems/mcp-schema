package v20250326

type PromptReference struct {
	Description string                    `json:"description"`
	Properties  PromptReferenceProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
