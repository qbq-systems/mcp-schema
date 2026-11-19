package v20251125

type SetLevelRequestParams struct {
	Description string                          `json:"description"`
	Properties  SetLevelRequestParamsProperties `json:"properties"`
	Required    []string                        `json:"required"`
	Type        AnnotationsType                 `json:"type"`
}
