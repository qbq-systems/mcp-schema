package v20241105

type TentacledParams struct {
	Properties TentacledProperties `json:"properties"`
	Required   []string            `json:"required"`
	Type       AnnotationsType     `json:"type"`
}
