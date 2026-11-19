package vdraft

type InputResponseRequestParams struct {
	Properties InputResponseRequestParamsProperties `json:"properties"`
	Required   []string                             `json:"required"`
	Type       TypeElement                          `json:"type"`
}
