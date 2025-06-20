package v20241105

type StickyParams struct {
	Properties StickyProperties `json:"properties"`
	Required   []string         `json:"required"`
	Type       AnnotationsType  `json:"type"`
}
