package vdraft

type CallToolRequestParams struct {
	Description string                          `json:"description"`
	Properties  CallToolRequestParamsProperties `json:"properties"`
	Required    []string                        `json:"required"`
	Type        AnnotationsType                 `json:"type"`
}
