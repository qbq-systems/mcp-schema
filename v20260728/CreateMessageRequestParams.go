package v20260728

type CreateMessageRequestParams struct {
	Description string                               `json:"description"`
	Properties  CreateMessageRequestParamsProperties `json:"properties"`
	Required    []string                             `json:"required"`
	Type        TypeElement                          `json:"type"`
}
