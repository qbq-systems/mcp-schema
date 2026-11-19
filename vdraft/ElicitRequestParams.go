package vdraft

type ElicitRequestParams struct {
	Description string                            `json:"description"`
	Properties  ElicitRequestFormParamsProperties `json:"properties"`
	Required    []string                          `json:"required"`
	Type        TypeElement                       `json:"type"`
}
