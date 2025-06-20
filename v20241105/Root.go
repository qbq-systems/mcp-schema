package v20241105

type Root struct {
	Description string          `json:"description"`
	Properties  RootProperties  `json:"properties"`
	Required    []Required      `json:"required"`
	Type        AnnotationsType `json:"type"`
}
