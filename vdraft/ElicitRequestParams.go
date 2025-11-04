package vdraft

type ElicitRequestParams struct {
	Description string                        `json:"description"`
	Properties  ElicitRequestParamsProperties `json:"properties"`
	Required    []string                      `json:"required"`
	Type        AnnotationsType               `json:"type"`
}
