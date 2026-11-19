package v20260728

type PaginatedRequestParams struct {
	Description string                           `json:"description"`
	Properties  PaginatedRequestParamsProperties `json:"properties"`
	Required    []string                         `json:"required"`
	Type        TypeElement                      `json:"type"`
}
