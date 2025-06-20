package v20241105

type FluffyParams struct {
	Properties FluffyProperties `json:"properties"`
	Required   []string         `json:"required"`
	Type       AnnotationsType  `json:"type"`
}
