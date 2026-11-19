package vdraft

type ReadResourceRequestParams struct {
	Description string                               `json:"description"`
	Properties  InputResponseRequestParamsProperties `json:"properties"`
	Required    []string                             `json:"required"`
	Type        TypeElement                          `json:"type"`
}
