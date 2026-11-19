package v20251125

type PaginatedRequestParams struct {
	Description string                           `json:"description"`
	Properties  PaginatedRequestParamsProperties `json:"properties"`
	Type        AnnotationsType                  `json:"type"`
}
