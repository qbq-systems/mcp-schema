package v20260728

type CallToolRequestParams struct {
	Description string                          `json:"description"`
	Properties  CallToolRequestParamsProperties `json:"properties"`
	Required    []string                        `json:"required"`
	Type        TypeElement                     `json:"type"`
}
