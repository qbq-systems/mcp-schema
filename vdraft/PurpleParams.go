package vdraft

type PurpleParams struct {
	Properties StickyProperties `json:"properties"`
	Required   []string         `json:"required"`
	Type       AnnotationsType  `json:"type"`
}
