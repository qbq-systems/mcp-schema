package v20250326

type IndecentParams struct {
	Properties IndecentProperties `json:"properties"`
	Required   []string           `json:"required"`
	Type       AnnotationsType    `json:"type"`
}
