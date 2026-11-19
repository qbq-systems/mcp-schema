package vdraft

type ResourceParams struct {
	Description string                          `json:"description"`
	Properties  ResourceRequestParamsProperties `json:"properties"`
	Required    []string                        `json:"required"`
	Type        TypeElement                     `json:"type"`
}
