package v20241105

type MagentaParams struct {
	Properties MischievousProperties `json:"properties"`
	Required   []string              `json:"required"`
	Type       AnnotationsType       `json:"type"`
}
