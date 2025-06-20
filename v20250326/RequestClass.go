package v20250326

type RequestClass struct {
	Properties RequestProperties `json:"properties"`
	Required   []string          `json:"required"`
	Type       AnnotationsType   `json:"type"`
}
