package vdraft

type TentacledParams struct {
	Properties StickyProperties `json:"properties"`
	Required   []string         `json:"required"`
	Type       AnnotationsType  `json:"type"`
}
