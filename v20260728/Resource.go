package v20260728

type Resource struct {
	Description string             `json:"description"`
	Properties  ResourceProperties `json:"properties"`
	Required    []string           `json:"required"`
	Type        TypeElement        `json:"type"`
}
