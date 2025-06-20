package vdraft

type StickyParams struct {
	Properties IndigoProperties `json:"properties"`
	Required   []string         `json:"required"`
	Type       AnnotationsType  `json:"type"`
}
