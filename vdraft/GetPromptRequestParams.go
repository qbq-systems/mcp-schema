package vdraft

type GetPromptRequestParams struct {
	Description string                           `json:"description"`
	Properties  GetPromptRequestParamsProperties `json:"properties"`
	Required    []string                         `json:"required"`
	Type        AnnotationsType                  `json:"type"`
}
