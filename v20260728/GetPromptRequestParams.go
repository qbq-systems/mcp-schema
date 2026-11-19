package v20260728

type GetPromptRequestParams struct {
	Description string                           `json:"description"`
	Properties  GetPromptRequestParamsProperties `json:"properties"`
	Required    []string                         `json:"required"`
	Type        TypeElement                      `json:"type"`
}
