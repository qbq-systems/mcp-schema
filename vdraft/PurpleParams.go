package vdraft

type PurpleParams struct {
	Properties MagentaProperties `json:"properties"`
	Required   []string          `json:"required"`
	Type       AnnotationsType   `json:"type"`
}
