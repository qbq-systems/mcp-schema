package vdraft

type ElicitRequestFormParams struct {
	Description string                            `json:"description"`
	Properties  ElicitRequestFormParamsProperties `json:"properties"`
	Required    []string                          `json:"required"`
	Type        AnnotationsType                   `json:"type"`
}
