package vdraft

type PurpleParams struct {
	Properties FluffyProperties `json:"properties"`
	Required   []string         `json:"required"`
	Type       AnnotationsType  `json:"type"`
}
