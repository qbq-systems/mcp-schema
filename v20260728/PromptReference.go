package v20260728

type PromptReference struct {
	Description string                    `json:"description"`
	Properties  PromptReferenceProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        TypeElement               `json:"type"`
}
