package v20250326

type CompleteResult struct {
	Description string                   `json:"description"`
	Properties  CompleteResultProperties `json:"properties"`
	Required    []string                 `json:"required"`
	Type        AnnotationsType          `json:"type"`
}
