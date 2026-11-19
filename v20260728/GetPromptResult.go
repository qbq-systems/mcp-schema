package v20260728

type GetPromptResult struct {
	Description string                    `json:"description"`
	Properties  GetPromptResultProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        TypeElement               `json:"type"`
}
