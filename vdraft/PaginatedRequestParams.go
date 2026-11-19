package vdraft

type PaginatedRequestParams struct {
	Description string                           `json:"description"`
	Properties  PaginatedRequestParamsProperties `json:"properties"`
	Required    []string                         `json:"required"`
	Type        TypeElement                      `json:"type"`
}
