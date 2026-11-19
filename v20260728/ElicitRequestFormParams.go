package v20260728

type ElicitRequestFormParams struct {
	Description string                            `json:"description"`
	Properties  ElicitRequestFormParamsProperties `json:"properties"`
	Required    []string                          `json:"required"`
	Type        TypeElement                       `json:"type"`
}
