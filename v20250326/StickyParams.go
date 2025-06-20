package v20250326

type StickyParams struct {
	Properties StickyProperties `json:"properties"`
	Required   []string         `json:"required"`
	Type       AnnotationsType  `json:"type"`
}
