package vdraft

type ElicitRequestURLParams struct {
	Description string                           `json:"description"`
	Properties  ElicitRequestURLParamsProperties `json:"properties"`
	Required    []string                         `json:"required"`
	Type        AnnotationsType                  `json:"type"`
}
