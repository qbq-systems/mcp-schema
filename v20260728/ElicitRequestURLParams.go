package v20260728

type ElicitRequestURLParams struct {
	Description string                           `json:"description"`
	Properties  ElicitRequestURLParamsProperties `json:"properties"`
	Required    []string                         `json:"required"`
	Type        TypeElement                      `json:"type"`
}
