package v20250326

type Implementation struct {
	Description string                   `json:"description"`
	Properties  ImplementationProperties `json:"properties"`
	Required    []string                 `json:"required"`
	Type        AnnotationsType          `json:"type"`
}
