package v20260728

type ElicitResult struct {
	Description string                 `json:"description"`
	Properties  ElicitResultProperties `json:"properties"`
	Required    []string               `json:"required"`
	Type        TypeElement            `json:"type"`
}
