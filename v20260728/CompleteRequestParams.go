package v20260728

type CompleteRequestParams struct {
	Description string                          `json:"description"`
	Properties  CompleteRequestParamsProperties `json:"properties"`
	Required    []string                        `json:"required"`
	Type        TypeElement                     `json:"type"`
}
