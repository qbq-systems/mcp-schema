package v20251125

type CompleteRequestParams struct {
	Description string                          `json:"description"`
	Properties  CompleteRequestParamsProperties `json:"properties"`
	Required    []string                        `json:"required"`
	Type        AnnotationsType                 `json:"type"`
}
