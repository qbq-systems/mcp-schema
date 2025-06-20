package v20250326

type IndigoParams struct {
	Properties IndigoProperties `json:"properties"`
	Required   []string         `json:"required"`
	Type       AnnotationsType  `json:"type"`
}
