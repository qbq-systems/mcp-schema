package v20251125

type Icon struct {
	Description string          `json:"description"`
	Properties  IconProperties  `json:"properties"`
	Required    []string        `json:"required"`
	Type        AnnotationsType `json:"type"`
}
