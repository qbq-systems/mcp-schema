package vdraft

type InitializeRequestParams struct {
	Description string                            `json:"description"`
	Properties  InitializeRequestParamsProperties `json:"properties"`
	Required    []string                          `json:"required"`
	Type        AnnotationsType                   `json:"type"`
}
