package vdraft

type CreateMessageRequestParams struct {
	Description string                               `json:"description"`
	Properties  CreateMessageRequestParamsProperties `json:"properties"`
	Required    []string                             `json:"required"`
	Type        AnnotationsType                      `json:"type"`
}
