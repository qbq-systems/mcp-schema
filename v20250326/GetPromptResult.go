package v20250326

type GetPromptResult struct {
	Description string                    `json:"description"`
	Properties  GetPromptResultProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
